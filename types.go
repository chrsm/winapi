package winapi

import "unsafe"

type (
	HWND   unsafe.Pointer
	HANDLE uintptr
	DWORD  uint32
	LONG   uint32
	ATOM   uint16
	WORD   uint16
	WPARAM uintptr
	LPARAM uintptr

	Rect struct {
		Left, Top, Right, Bottom LONG
	}
)
