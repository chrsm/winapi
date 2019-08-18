package user

import (
	"unsafe"

	"github.com/chrsm/winapi"
)

var (
	pRegisterHotkey   = userapi.NewProc("RegisterHotKey")
	pUnregisterHotkey = userapi.NewProc("UnregisterHotKey")
	pGetKeyboardState = userapi.NewProc("GetKeyboardState")
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

func GetKeyboardState() []byte {
	state := [256]byte{}

	ret, _, _ := pGetKeyboardState.Call(
		uintptr(unsafe.Pointer(&state)),
	)

	if ret == 0 {
		return nil
	}

	return state[:]
}
