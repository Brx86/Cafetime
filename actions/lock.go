package actions

// 调用 WIN32 锁屏
func LockScreenWindows() {
	procLockWorkStation.Call()
}

// 启用 DPI 感知 https://github.com/getlantern/systray/issues/285
func EnableDPIAwareness() {
	// 定义 DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2
	// 在 Windows 中，这是一个预定义的常量，值为 HANDLE(-4)
	const DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 = uintptr(^uintptr(3)) // -4 的补码表示
	// 调用 SetProcessDpiAwarenessContext
	setProcessDpiAwarenessContext.Call(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2)
}
