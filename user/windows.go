package user

import (
	"syscall"
	"unsafe"

	"github.com/chrsm/winapi"
)

var (
	pCreateWindowEx           = userapi.NewProc("CreateWindowExW")
	pDefWindowProc            = userapi.NewProc("DefWindowProcW")
	pDestroyWindow            = userapi.NewProc("DestroyWindow")
	pDispatchMessage          = userapi.NewProc("DispatchMessageW")
	pEnumWindows              = userapi.NewProc("EnumWindows")
	pGetActiveWindow          = userapi.NewProc("GetActiveWindow")
	pGetForegroundWindow      = userapi.NewProc("GetForegroundWindow")
	pGetMessage               = userapi.NewProc("GetMessageW")
	pGetWindowInfo            = userapi.NewProc("GetWindowInfo")
	pGetWindowThreadProcessId = userapi.NewProc("GetWindowThreadProcessId")
	pIsWindowVisible          = userapi.NewProc("IsWindowVisible")
	pPostQuitMessage          = userapi.NewProc("PostQuitMessage")
	pTranslateMessage         = userapi.NewProc("TranslateMessage")
	pRegisterClassEx          = userapi.NewProc("RegisterClassExW")
	pRegisterWindowMessage    = userapi.NewProc("RegisterWindowMessageA")
	pSetForegroundWindow      = userapi.NewProc("SetForegroundWindow")
	pShowWindow               = userapi.NewProc("ShowWindow")
)

func CreateWindow(
	class, title string,
	style wsExStyle,
	x, y, width, height int32,
	parent winapi.HWND,
	menu winapi.HMENU,
	instance winapi.HMODULE,
	lpParam winapi.LPVOID,
) winapi.HWND {
	ret, _, _ := pCreateWindowEx.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(class))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(0),
	)

	return winapi.HWND(ret)
}

func DestroyWindow(hwnd winapi.HWND) {
	pDestroyWindow.Call(uintptr(hwnd))
}

func DefWindowProc(hwnd winapi.HWND, msg uint32, wparam winapi.WPARAM, lparam winapi.LPARAM) uintptr {
	ret, _, _ := pDefWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam),
	)

	return ret
}

func EnumWindows(enumfn EnumWindowsCallback, v uintptr) bool {
	cb := syscall.NewCallback(enumfn)

	ret, _, _ := pEnumWindows.Call(
		cb,
		0,
	)

	return ret != 0
}

func GetMessage(msg *Message, hwnd winapi.HWND, filterMin, filterMax uint) bool {
	ret, _, _ := pGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(filterMin),
		uintptr(filterMax),
	)

	return ret != 0
}

func TranslateMessage(m *Message) {
	pTranslateMessage.Call(uintptr(unsafe.Pointer(m)))
}

func DispatchMessage(m *Message) {
	pDispatchMessage.Call(uintptr(unsafe.Pointer(m)))
}

func PostQuitMessage(code int32) {
	pPostQuitMessage.Call(uintptr(code))
}

// ShowWindow returns true if the window was previously visible.
// MS should probably have called this "SetWindowState" or some shit.
func ShowWindow(hwnd winapi.HWND, state swFlags) bool {
	ret, _, _ := pShowWindow.Call(
		uintptr(hwnd),
		uintptr(state),
	)

	return ret != 0
}

func GetWindowInfo(hwnd winapi.HWND) *WindowInfo {
	info := &WindowInfo{}
	info.CbSize = szWindowInfo

	ret, _, _ := pGetWindowInfo.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(info)),
	)

	if ret == 0 {
		return nil
	}

	return info
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

func SetForegroundWindow(window winapi.HWND) bool {
	ret, _, _ := pSetForegroundWindow.Call(uintptr(window))

	return ret != 0
}

func GetWindowThreadProcessId(window winapi.HWND) (winapi.HANDLE, int) {
	pid := 0

	ret, _, _ := pGetWindowThreadProcessId.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&pid)),
	)

	return winapi.HANDLE(ret), pid
}

func GetActiveWindow() winapi.HWND {
	ret, _, _ := pGetActiveWindow.Call()

	return winapi.HWND(ret)
}

func RegisterWindowMessage(name string) uint {
	ret, _, _ := pRegisterWindowMessage.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))))

	return uint(ret)
}

func RegisterClass(wc *WindowClass) uint16 {
	ret, _, _ := pRegisterClassEx.Call(uintptr(unsafe.Pointer(wc)))

	return uint16(ret)
}
