package user

import (
	"syscall"

	"github.com/chrsm/winapi"
)

var (
	pSetWindowsHookExA       = userapi.NewProc("SetWindowsHookExA")
	pCallNextHookEx          = userapi.NewProc("CallNextHookEx")
	pUnhookWindowsHookEx     = userapi.NewProc("UnhookWindowsHookEx")
	pRegisterShellHookWindow = userapi.NewProc("RegisterShellHookWindow")
)

func SetWindowsHookEx(typ windowsHookType, fn HookFn, hmod uintptr, thread winapi.DWORD) Hook {
	ret, _, _ := pSetWindowsHookExA.Call(
		uintptr(typ),
		uintptr(syscall.NewCallback(fn)),
		uintptr(hmod),
		uintptr(thread),
	)

	return Hook(ret)
}

func CallNextHook(hook Hook, ncode int, wparam winapi.WPARAM, lparam winapi.LPARAM) uintptr {
	ret, _, _ := pCallNextHookEx.Call(
		uintptr(hook),
		uintptr(ncode),
		uintptr(wparam),
		uintptr(lparam),
	)

	return ret
}

func UnhookWindowsHookEx(hook Hook) bool {
	ret, _, _ := pUnhookWindowsHookEx.Call(
		uintptr(hook),
	)

	return ret != 0
}

func RegisterShellHookWindow(window winapi.HWND) bool {
	ret, _, _ := pRegisterShellHookWindow.Call(uintptr(window))

	return ret != 0
}
