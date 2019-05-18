package user

import "github.com/chrsm/winapi"

var (
	pRegisterHotkey   = userapi.NewProc("RegisterHotKey")
	pUnregisterHotkey = userapi.NewProc("UnregisterHotKey")
)

func RegisterHotkey(window winapi.HWND, id int, modifiers ModKey, vk VirtualKey) bool {
	ret, _, _ := pRegisterHotkey.Call(
		uintptr(window),
		uintptr(id),
		uintptr(modifiers),
		uintptr(vk),
	)

	return ret != 0
}

func UnregisterHotkey(window winapi.HWND, id int) bool {
	ret, _, _ := pUnregisterHotkey.Call(
		uintptr(window),
		uintptr(id),
	)

	return ret != 0
}
