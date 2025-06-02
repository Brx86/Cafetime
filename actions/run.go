package actions

import (
	"os/exec"
)

func OpenFileWindows(filePath string) {
	cmd := exec.Command("cmd", "/C", "start", "", filePath)
	cmd.Run()
}
