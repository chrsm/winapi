package winapi

import "unsafe"

type (
	HANDLE    uintptr
	DWORD     uint32
	LONG      uint32
	ATOM      uint16
	WORD      uint16
	WPARAM    uintptr
	LPARAM    uintptr
	HWND      HANDLE
	HMENU     HANDLE
	HINSTANCE HANDLE
	HMODULE   HANDLE
	LPVOID    unsafe.Pointer

	HBACKGROUND HANDLE
	HCURSOR     HANDLE
	HICON       HANDLE

	Rect struct {
		Left, Top, Right, Bottom LONG
	}

	GUID struct {
		Data1 DWORD
		Data2 WORD
		Data3 WORD
		Data4 [8]byte
	}
)
