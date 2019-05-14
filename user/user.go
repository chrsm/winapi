// Package user implements Windows APIs for user32.dll in Go.
// See documentation for more info:
// 	https://docs.microsoft.com/en-us/windows/desktop/api/winuser/
package user

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/chrsm/winapi"
)

var (
	userapi = syscall.NewLazyDLL("user32.dll")

	pGetMessage               = userapi.NewProc("GetMessageW")
	pEnumWindows              = userapi.NewProc("EnumWindows")
	pShowWindow               = userapi.NewProc("ShowWindow")
	pGetWindowInfo            = userapi.NewProc("GetWindowInfo")
	pRegisterHotkey           = userapi.NewProc("RegisterHotKey")
	pUnregisterHotkey         = userapi.NewProc("UnregisterHotKey")
	pIsWindowVisible          = userapi.NewProc("IsWindowVisible")
	pGetForegroundWindow      = userapi.NewProc("GetForegroundWindow")
	pGetWindowThreadProcessId = userapi.NewProc("GetWindowThreadProcessId")
)

type EnumWindowsCallback func(window winapi.HWND, v winapi.LPARAM) uintptr

func EnumWindows(enumfn EnumWindowsCallback, v uintptr) bool {
	cb := syscall.NewCallback(enumfn)

	ret, _, _ := pEnumWindows.Call(
		cb,
		0,
	)

	return ret != 0
}

// ShowWindow returns true if the window was previously visible.
// MS should probably have called this "SetWindowState" or some shit.
func ShowWindow(window winapi.HWND, state swFlags) bool {
	ret, _, _ := pShowWindow.Call(
		uintptr(window),
		uintptr(state),
	)

	return ret != 0
}

func GetWindowInfo(window winapi.HWND) (*WindowInfo, error) {
	info := &WindowInfo{}
	info.CbSize = szWindowInfo

	ret, _, _ := pGetWindowInfo.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(info)),
	)

	if ret == 0 {
		return nil, errors.New("...getLastError")
	}

	return info, nil
}

func RegisterHotkey(window winapi.HWND, id int, modifiers modKey, vk virtualKey) bool {
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

func IsWindowVisible(window winapi.HWND) bool {
	ret, _, _ := pIsWindowVisible.Call(
		uintptr(window),
	)

	return ret != 0
}

func GetForegroundWindow() winapi.HWND {
	ret, _, _ := pGetForegroundWindow.Call()

	return winapi.HWND(ret)
}

func GetMessage(window winapi.HWND, filterMin, filterMax uint) (*Message, bool) {
	msg := &Message{}
	ret, _, _ := pGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(window),
		uintptr(filterMin),
		uintptr(filterMax),
	)

	return msg, ret != 0
}

func GetWindowThreadProcessId(window winapi.HWND) (winapi.HANDLE, int) {
	pid := 0

	ret, _, _ := pGetWindowThreadProcessId.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&pid)),
	)

	return winapi.HANDLE(ret), pid
}
