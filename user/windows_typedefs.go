package user

import (
	"unsafe"

	"github.com/chrsm/winapi"
)

var (
	szWindowInfo  = winapi.DWORD(unsafe.Sizeof(WindowInfo{}))
	szWindowClass = winapi.DWORD(unsafe.Sizeof(WindowClass{}))
)

type (
	swFlags       int
	wsExStyle     uint32
	windowMessage uint
	windowStyle   uint

	EnumWindowsCallback func(window winapi.HWND, v winapi.LPARAM) uintptr

	Message struct {
		HWnd     winapi.HWND
		Message  uint
		WParam   winapi.WPARAM
		LParam   winapi.LPARAM
		Time     winapi.DWORD
		Point    Point
		LPrivate winapi.DWORD
	}

	Point struct {
		X, Y winapi.LONG
	}

	WindowClass struct {
		CbSize      winapi.DWORD
		Style       uint32
		WndProc     uintptr
		ClsExtra    int32
		CbWndExtra  int32
		HInstance   winapi.HINSTANCE
		HIcon       winapi.HICON
		HCursor     winapi.HCURSOR
		HBackground winapi.HBACKGROUND
		MenuName    *uint16
		ClassName   *uint16
		HIconSmall  winapi.HICON
	}

	WindowInfo struct {
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
)

const (
	SwForceMinimize   = 11
	SwHide            = 0
	SwMaximize        = 3
	SwMinimize        = 6
	SwRestore         = 9
	SwShow            = 5
	SwShowDefault     = 6
	SwShowMaximized   = SwMaximize
	SwShowMinimized   = 2
	SwShowMinNoActive = 7
	SwShowNA          = 8
	SwShowNoActivate  = 4
	SwShowNormal      = 1

	WmNull   = 0x0000
	WmQuit   = 0x0012
	WmHotkey = 0x0312

	WsVisible = 0x10000000

	WsExAcceptFiles         = 0x00000010
	WsExAppWindow           = 0x00040000
	WsExClientEdge          = 0x00000200
	WsExComposited          = 0x02000000
	WsExContextHelp         = 0x00000400
	WsExControlParent       = 0x00010000
	WsExDlgModalFrame       = 0x00000001
	WsExLayered             = 0x00080000
	WsExLayoutRTL           = 0x00400000
	WsExLeft                = 0x00000000
	WsExLeftScrollBar       = 0x00004000
	WsExLTRReading          = 0x00000000
	WsExMDIChild            = 0x00000040
	WsExNoActivate          = 0x08000000
	WsExNoInheritLayout     = 0x00100000
	WsExNoParentNotify      = 0x00000004
	WsExNoRedirectionBitmap = 0x00200000
	WsExOverlappedWindow    = WsExWindowEdge | WsExClientEdge
	WsExPaletteWindow       = WsExWindowEdge | WsExToolWindow | WsExTopMost
	WsExRight               = 0x00001000
	WsExRightScrollbar      = 0x00000000
	WsExRTLReading          = 0x00002000
	WsExStaticEdge          = 0x00020000
	WsExToolWindow          = 0x00000080
	WsExTopMost             = 0x00000008
	WsExTransparent         = 0x00000020
	WsExWindowEdge          = 0x00000100
)
