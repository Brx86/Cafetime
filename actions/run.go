package actions

import (
	"os/exec"

	"github.com/sqweek/dialog"
)

func OpenFileWindows(filePath string) func() {
	return func() {
		cmd := exec.Command("cmd", "/C", "start", "", filePath)
		cmd.Run()
	}
}

func SelectFileWindows() string {
	filePath, _ := dialog.File().Title("选择要打开的文件或程序").Filter("文件或程序", "*").Load()
	return filePath
}

func SelectLnkWindows() string {
	link, _ := dialog.File().Title("选择要打开的网页快捷方式").Filter("快捷方式", "lnk").Load()
	return link
}
