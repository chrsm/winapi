package user

import "syscall"

var (
	userapi      = syscall.NewLazyDLL("user32.dll")
	pEnumWindows = userapi.NewProc("EnumWindows")
)

type EnumWindowsCallback func(uintptr, uintptr)

func EnumWindows(enumfn EnumWindowsCallback, v uintptr) bool {
	cb := syscall.NewCallback(enumfn)

	ret, _, _ := pEnumWindows.Call(
		cb,
		0,
	)

	return ret != 0
}
