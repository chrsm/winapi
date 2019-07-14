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
	pGetClassName             = userapi.NewProc("GetClassNameW")
	pGetDesktopWindow         = userapi.NewProc("GetDesktopWindow")
	pGetForegroundWindow      = userapi.NewProc("GetForegroundWindow")
	pGetMessage               = userapi.NewProc("GetMessageW")
	pGetParent                = userapi.NewProc("GetParent")
	pGetWindow                = userapi.NewProc("GetWindow")
	pGetWindowInfo            = userapi.NewProc("GetWindowInfo")
	pGetWindowLong            = userapi.NewProc("GetWindowLongW")
	pGetWindowText            = userapi.NewProc("GetWindowTextW")
	pGetWindowTextLength      = userapi.NewProc("GetWindowTextLengthW")
	pGetWindowThreadProcessId = userapi.NewProc("GetWindowThreadProcessId")
	pIsWindow                 = userapi.NewProc("IsWindow")
	pIsWindowVisible          = userapi.NewProc("IsWindowVisible")
	pPostQuitMessage          = userapi.NewProc("PostQuitMessage")
	pTranslateMessage         = userapi.NewProc("TranslateMessage")
	pRegisterClassEx          = userapi.NewProc("RegisterClassExW")
	pRegisterWindowMessage    = userapi.NewProc("RegisterWindowMessageA")
	pSendMessage              = userapi.NewProc("SendMessageW")
	pSetFocus                 = userapi.NewProc("SetFocus")
	pSetForegroundWindow      = userapi.NewProc("SetForegroundWindow")
	pSetWindowPos             = userapi.NewProc("SetWindowPos")
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

func GetParent(hwnd winapi.HWND) winapi.HWND {
	ret, _, _ := pGetParent.Call(uintptr(hwnd))

	return winapi.HWND(ret)
}

func GetWindow(hwnd winapi.HWND, cmd uint) winapi.HWND {
	ret, _, _ := pGetWindow.Call(uintptr(hwnd), uintptr(cmd))

	return winapi.HWND(ret)
}

func GetWindowLong(hwnd winapi.HWND, index int) uint32 {
	ret, _, _ := pGetWindowLong.Call(uintptr(hwnd), uintptr(index))

	return uint32(ret)
}

func GetWindowText(hwnd winapi.HWND, dst []uint16) string {
	ret, _, _ := pGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&dst[0])),
		uintptr(len(dst)), // max length
	)

	_ = ret

	return syscall.UTF16ToString(dst)
}

func GetWindowTextLength(hwnd winapi.HWND) int {
	ret, _, _ := pGetWindowTextLength.Call(uintptr(hwnd))

	return int(ret)
}

func IsWindow(hwnd winapi.HWND) bool {
	ret, _, _ := pIsWindow.Call(uintptr(hwnd))

	return ret != 0
}

func GetDesktopWindow() winapi.HWND {
	ret, _, _ := pGetDesktopWindow.Call()

	return winapi.HWND(ret)
}

func GetClassName(hwnd winapi.HWND, dst []uint16) string {
	ret, _, _ := pGetClassName.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&dst[0])),
		uintptr(len(dst)), // max length
	)

	_ = ret

	return syscall.UTF16ToString(dst)
}

func SetFocus(hwnd winapi.HWND) {
	pSetFocus.Call(uintptr(hwnd))
}

func SetWindowPos(hwnd winapi.HWND, hwndAfter winapi.HWND, x, y, width, height int, flags uint) bool {
	ret, _, _ := pSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hwndAfter),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)

	return ret != 0
}

func SendMessage(hwnd winapi.HWND, msg uint, wparam winapi.WPARAM, lparam winapi.LPARAM) {
	ret, _, _ := pSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam),
	)

	// @TODO(chrsm): there's a bunch of different ways to handle lresult,
	// should probably just hand it back to the caller..
	_ = ret
}
