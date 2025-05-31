package actions

import (
	"os/exec"
	"runtime"
)

func OpenFileWindows(filePath string) {
	if runtime.GOOS != "windows" {
		println("此函数仅适用于 Windows 操作系统")
	}
	cmd := exec.Command("cmd", "/C", "start", "", filePath)
	cmd.Run()
}
