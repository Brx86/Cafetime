package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/systray"
)

const (
	secPm = 60
	secPh = 60 * secPm
	secPd = 24 * secPh
)

// 新建定时器
func NewTimer(m *systray.MenuItem, timeout int, f func()) {
	timerTitleCh := make(chan int, 1)
	ticker := time.NewTicker(time.Second)

	// 关闭 channel，恢复菜单标题
	defer m.SetTitle("启动定时器")
	defer SetTimerStatus(false)
	defer close(timerTitleCh)
	defer ticker.Stop()

	// 更新标题
	go func() {
		for t := range timerTitleCh {
			m.SetTitle(`剩余 ` + SecToStr(t) + `，点击取消`)
		}
	}()

	// 开始计时
	for c := timeout; c > 0; c-- {
		timerTitleCh <- c
		select {
		case <-ticker.C:
		case <-stopTimerCh:
			return
		}
	}
	f()
}

// 从输入时间字符串计算得秒数
func StrToSec(timeoutStr string) (int, error) {
	// 判断输入有效
	if len(timeoutStr) < 2 {
		return 0, fmt.Errorf("无效的超时字符串: '%s', 长度至少为2 (例如 '1s')", timeoutStr)
	}
	suffix := strings.ToLower(timeoutStr[len(timeoutStr)-1:])
	count, err := strconv.Atoi(timeoutStr[:len(timeoutStr)-1])
	if err != nil {
		return 0, fmt.Errorf("无法将 '%s' 解析为数字: %w", timeoutStr, err)
	}
	if count < 0 {
		return 0, fmt.Errorf("超时时间不能为负数: %s", timeoutStr)
	}

	// 返回秒数
	switch suffix {
	case "s":
		return count, nil
	case "m":
		return count * secPm, nil
	case "h":
		return count * secPh, nil
	case "d":
		return count * secPd, nil
	default: // 默认为分钟
		return count * secPm, nil
	}
}

// 从秒数计算得剩余时间
func SecToStr(sec int) string {
	switch {
	case sec < secPm:
		return fmt.Sprintf("%d 秒", sec)
	case sec < secPh:
		minutes := sec / secPm
		rSeconds := sec % secPm
		return fmt.Sprintf("%d 分 %d 秒", minutes, rSeconds)
	case sec < secPd:
		hours := sec / secPh
		rMinutes := (sec % secPh) / secPm
		return fmt.Sprintf("%d 小时 %d 分钟", hours, rMinutes)
	default:
		days := sec / secPd
		rHours := (sec % secPd) / secPh
		rMinutes := (sec % secPh) / secPm
		return fmt.Sprintf("%d 天 %d 小时 %d 分钟", days, rHours, rMinutes)
	}
}

// 设置定时器状态和托盘图标
func SetTimerStatus(status bool) {
	timerFlag = status
	if status {
		systray.SetIcon(IconOn)
	} else {
		systray.SetIcon(IconOff)
	}
}
