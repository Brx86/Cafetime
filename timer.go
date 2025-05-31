package main

import (
	"strconv"
	"time"

	"fyne.io/systray"
)

func NewTimer(m *systray.MenuItem, count int, f func()) {
	timerTitleCh := make(chan string, 1)
	ticker := time.NewTicker(time.Second)

	// 关闭 channel，恢复菜单标题
	defer m.SetTitle("启动定时器")
	defer SetTimerStatus(false)
	defer close(timerTitleCh)
	defer ticker.Stop()

	// 更新标题
	go func() {
		for t := range timerTitleCh {
			m.SetTitle(t)
		}
	}()

	// 开始计时
	for c := count; c > 0; c-- {
		timerTitleCh <- "剩余" + strconv.Itoa(c) + "s ，点击取消"
		select {
		case <-ticker.C:
		case <-stopTimerCh:
			return
		}
	}
	f()
}

func SetTimerStatus(status bool) {
	timerFlag = status
	if status {
		systray.SetIcon(IconOn)
	} else {
		systray.SetIcon(IconOff)
	}
}
