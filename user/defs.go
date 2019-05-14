package user

import (
	"unsafe"

	"github.com/chrsm/winapi"
)

type Message struct {
	HWnd     winapi.HWND
	Message  uint
	WParam   winapi.WPARAM
	LParam   winapi.LPARAM
	Time     winapi.DWORD
	Point    Point
	LPrivate winapi.DWORD
}

type Point struct {
	X, Y winapi.LONG
}

var szWindowInfo = winapi.DWORD(unsafe.Sizeof(WindowInfo{}))

type swFlags int

const (
	SwForceMinimize   swFlags = 11
	SwHide                    = 0
	SwMaximize                = 3
	SwMinimize                = 6
	SwRestore                 = 9
	SwShow                    = 5
	SwShowDefault             = 6
	SwShowMaximized           = SwMaximize
	SwShowMinimized           = 2
	SwShowMinNoActive         = 7
	SwShowNA                  = 8
	SwShowNoActivate          = 4
	SwShowNormal              = 1
)

type WindowInfo struct {
	CbSize winapi.DWORD

	RcWindow winapi.Rect
	RcClient winapi.Rect

	DwStyle        winapi.DWORD
	DwExStyle      winapi.DWORD
	DwWindowStatus winapi.DWORD

	CxWindowBorders uint
	CyWindowBorders uint

	AtomWindowType winapi.ATOM

	WCreatorVersion winapi.WORD
}

type modKey uint

const (
	ModAlt      modKey = 0x0001
	ModControl         = 0x0002
	ModNoRepeat        = 0x4000
	ModShift           = 0x0004
	ModWin             = 0x0008
)

type virtualKey uint

const (
	VirtKeyShift    virtualKey = 0x10
	VirtKeyControl             = 0x11
	VirtKeyEscape              = 0x1B
	VirtKeyLeft                = 0x25
	VirtKeyUp                  = 0x26
	VirtKeyRight               = 0x27
	VirtKeyDown                = 0x28
	VirtKeyLeftWin             = 0x5B
	VirtKeyRightWin            = 0x5C
)

type windowMessage uint

const (
	WmNull   windowMessage = 0x0000
	WmQuit                 = 0x0012
	WmHotkey               = 0x0312
)

type windowStyle uint

const (
	WsVisible      windowStyle = 0x10000000
	WsExToolWindow             = 0x0000008
)
