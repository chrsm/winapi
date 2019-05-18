// Package user implements Windows APIs for user32.dll in Go.
// See documentation for more info:
// 	https://docs.microsoft.com/en-us/windows/desktop/api/winuser/
package user

import "syscall"

var userapi = syscall.NewLazyDLL("user32.dll")
