package main

import (
	_ "embed"
	"os"
	"os/signal"
	"syscall"

	"cafetime/actions"

	"fyne.io/systray"
)

var (
	actionFunc  = actions.LockScreenWindows
	stopTimerCh = make(chan struct{}, 1)
	//go:embed asset/cafeon.ico
	IconOn []byte
	//go:embed asset/cafeoff.ico
	IconOff []byte
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// 设置托盘
	systray.SetIcon(IconOff)
	systray.SetTooltip("CafeTime：曼城时光~")

	// 添加菜单项，处理点击事件
	m := NewMenu()
	go OnClickMenu(m)

	// 处理 SYSCALL 信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		println("收到系统退出信号")
		systray.Quit()
	}()
}

func onExit() {
	println("程序已退出")
}
