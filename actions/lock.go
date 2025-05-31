package actions

import "syscall"

var (
	user32              = syscall.NewLazyDLL("User32.dll")
	procLockWorkStation = user32.NewProc("LockWorkStation")
)

func LockScreenWindows() {
	procLockWorkStation.Call()
}
