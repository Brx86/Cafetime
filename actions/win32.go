package actions

import "syscall"

var (
	user32                        = syscall.NewLazyDLL("User32.dll")
	procLockWorkStation           = user32.NewProc("LockWorkStation")
	setProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")
)
