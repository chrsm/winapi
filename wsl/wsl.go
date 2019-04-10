// Package wsl implements Windows APIs for wslapi.dll in Go.
// See documentation for more info:
// 	https://docs.microsoft.com/en-us/windows/desktop/api/_wsl/
package wsl

import (
	"fmt"
	"syscall"
	"unsafe"
)

var wslapi = syscall.NewLazyDLL("wslapi.dll")

type wslFlags uint32

const (
	// No flags are being supplied
	FlagNone wslFlags = 0
	// Allow the distribution to interoperate with Windows processes,
	// ex. the user can run `cmd.exe` from within WSL
	FlagEnableInterop = 1
	// Add Windows %PATH% environment variables to WSL sessions
	FlagAppendNTPath = 2
	// Automount Windows drives inside of WSL sessions
	FlagEnableDriveMounting = 4

	ValidFlags   = FlagEnableInterop | FlagAppendNTPath | FlagEnableDriveMounting
	DefaultFlags = FlagEnableInterop | FlagAppendNTPath | FlagEnableDriveMounting
)

var (
	pWslConfigureDistribution        = wslapi.NewProc("WslConfigureDistribution")
	pWslGetDistributionConfiguration = wslapi.NewProc("WslGetDistributionConfiguration")
	pWslIsDistributionRegistered     = wslapi.NewProc("WslIsDistributionRegistered")
	pWslLaunch                       = wslapi.NewProc("WslLaunch")
	pWslLaunchInteractive            = wslapi.NewProc("WslLaunchInteractive")
	pWslRegisterDistribution         = wslapi.NewProc("WslRegisterDistribution")
	pWslUnregisterDistribution       = wslapi.NewProc("WslUnregisterDistribution")
)

// ConfigureDistribution calls WslConfigureDistribution.
// Returns true if the distribution was successfully configured.
func ConfigureDistribution(name string, defuid uint64, flags wslFlags) bool {
	ret, _, _ := pWslConfigureDistribution.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(defuid),
		uintptr(flags),
	)

	return ret == 0
}

// DistributionConfiguration describes a distribution's configuration
// based on when it was set up or last modified.
type DistributionConfiguration struct {
	WslVer                      uint64
	DefaultUID                  uint64
	Flags                       wslFlags
	DefaultEnvironmentVariables []string
}

func GetDistributionConfiguration(name string) (version, defUID uint64, flags wslFlags) {
	var (
		env     **uint16
		lenvars uint64
	)

	ret, _, _ := pWslGetDistributionConfiguration.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(&version)),
		uintptr(unsafe.Pointer(&defUID)),
		uintptr(unsafe.Pointer(&flags)),
		uintptr(unsafe.Pointer(&env)), // TODO: wtf **PSTR
		uintptr(unsafe.Pointer(&lenvars)),
	)

	// TODO: handle.
	_ = ret

	return
}

func IsDistributionRegistered(name string) bool {
	ret, _, _ := pWslIsDistributionRegistered.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)

	return ret == 1
}

// Hook up Std(x) by getting the Fd.
func Launch(name, command string, useCWD bool, stdin, stdout, stderr uintptr) (bool, uintptr) {
	var proc uintptr

	ret, _, _ := pWslLaunch.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(command))),
		uintptr(btou32(useCWD)),
		stdin,
		stdout,
		stderr,
		uintptr(unsafe.Pointer(&proc)),
	)

	return ret == 0, proc
}

func LaunchInteractive(name, command string, useCWD bool) (bool, uint32) {
	var exit uint32

	ret, ret2, lerr := pWslLaunchInteractive.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(command))),
		uintptr(btou32(useCWD)),
		uintptr(unsafe.Pointer(&exit)),
	)

	fmt.Printf("ret: %d\n", ret)
	fmt.Printf("ret2: %d\n", ret2)
	fmt.Printf("lerr: %#v\n", lerr)

	return ret == 0, exit
}

// RegisterDistribution registers a new distribution with WSL.
func RegisterDistribution(name, archivePath string) bool {
	ret, _, _ := pWslRegisterDistribution.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(archivePath))),
	)

	return ret == 0
}

// UnregisterDistribution removes ("unregisters") a distro from WSL.
func UnregisterDistribution(name string) bool {
	ret, _, _ := pWslUnregisterDistribution.Call(
		uintptr(syscall.StringToUTF16Ptr(name)),
	)

	return ret == 0
}

func btou32(b bool) uint32 {
	if b {
		return 1
	}

	return 0
}
