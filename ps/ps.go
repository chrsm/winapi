package ps

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	psapi = syscall.NewLazyDLL("psapi.dll")

	pEnumProcessModules = psapi.NewProc("EnumProcessModules")
	pGetModuleBaseName  = psapi.NewProc("GetModuleBaseNameW")

	szhmod uintptr

	userapi      = syscall.NewLazyDLL("user32.dll")
	pEnumWindows = userapi.NewProc("EnumWindows")
)

func EnumProcessModules(hproc uintptr, mod *uintptr, needed *uint32) bool {
	ret, _, _ := pEnumProcessModules.Call(
		uintptr(hproc),
		uintptr(unsafe.Pointer(mod)),
		uintptr(unsafe.Sizeof(szhmod)),
		uintptr(unsafe.Pointer(needed)),
	)

	return ret != 0
}

func GetModuleBaseName(hproc uintptr, mod uintptr) string {
	var name = make([]uint16, 32)

	ret, _, _ := pGetModuleBaseName.Call(
		uintptr(hproc),
		uintptr(mod),
		uintptr(unsafe.Pointer(&name[0])),
		uintptr(len(name)),
	)

	if ret != 0 {
		return windows.UTF16ToString(name[:ret])
	}

	return ""
}
