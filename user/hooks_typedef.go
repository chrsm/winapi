package user

import "github.com/chrsm/winapi"

type Hook uintptr
type HookFn func(int, winapi.WPARAM, winapi.LPARAM) uintptr

type windowsHookType int

const (
	WindowsHookCBT              windowsHookType = 5
	WindowsHookShell                            = 10
	WindowsHookCallWndProc                      = 12
	WindowsHookKeyboardLowLevel                 = 13
)

type CallWndProc struct {
	LParam  uintptr
	WParam  uintptr
	Message uint
	HWND    uintptr
}

const (
	// see docs:
	// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644977(v=vs.85)
	HBCTActivateWindow = 5
	HBCTCreateWindow   = 3
	HBCTDestroyWindow  = 4
	HBCTMinMax         = 1
	HBCTMoveSize       = 0
	HBCTSetFocus       = 9
)

const (
	HShellWindowActivated = 4
	HShellWindowCreated   = 1
	HShellWindowDestroyed = 2
	HShellWindowReplaced  = 13
)

type KBDLLHOOKSTRUCT struct {
	VKCode      winapi.DWORD
	ScanCode    winapi.DWORD
	Flags       winapi.DWORD
	Time        winapi.DWORD
	DwExtraInfo winapi.ULONG_PTR
}
