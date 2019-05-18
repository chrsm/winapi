// Package kernel implements Windows APIs for kernel32.dll in Go.
package kernel

import (
	"syscall"
	"unsafe"

	"github.com/chrsm/winapi"
)

var (
	kernapi = syscall.NewLazyDLL("kernel32.dll")

	pGetLastError    = kernapi.NewProc("GetLastError")
	pGetModuleHandle = kernapi.NewProc("GetModuleHandleW")
)

func GetLastError() winapi.DWORD {
	ret, _, _ := pGetLastError.Call()

	return winapi.DWORD(ret)
}

func GetModuleHandle(name string) winapi.HANDLE {
	var ret uintptr

	if name == "" {
		// get ourself
		ret, _, _ = pGetModuleHandle.Call(0)
	} else {
		ret, _, _ = pGetModuleHandle.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))))
	}

	return winapi.HANDLE(ret)
}
