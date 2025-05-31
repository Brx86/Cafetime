package main

import (
	"cafetime/actions"

	"fyne.io/systray"
)

type Menu struct {
	Actions,
	ActionLck,
	ActionRun,
	ActionWeb,
	Timer,
	StartUp,
	Quit *systray.MenuItem
}

func NewMenu() *Menu {
	m := &Menu{}
	m.Actions = systray.AddMenuItem("超时动作", "")
	m.ActionLck = m.Actions.AddSubMenuItemCheckbox("锁定屏幕", "", true)
	m.ActionRun = m.Actions.AddSubMenuItem("打开文件/程序", "")
	m.ActionWeb = m.Actions.AddSubMenuItem("打开网站", "")
	m.Timer = systray.AddMenuItem("启动定时器", "")
	systray.AddSeparator() // 分隔线
	m.StartUp = systray.AddMenuItemCheckbox("开机自启", "", false)
	m.Quit = systray.AddMenuItem("退出", "")
	return m
}

func OnClickMenu(m *Menu) {
	for {
		select {
		case <-m.ActionLck.ClickedCh: // 超时锁屏
			uncheckAllExcept(m.ActionLck, m.ActionRun, m.ActionWeb)
			actionFunc = actions.LockScreenWindows
		case <-m.ActionRun.ClickedCh: // 超时打开文件或程序
			uncheckAllExcept(m.ActionRun, m.ActionLck, m.ActionWeb)
			actionFunc = func() { actions.OpenFileWindows("gohome.mp3") }
		case <-m.ActionWeb.ClickedCh: // 超时打开网页
			uncheckAllExcept(m.ActionWeb, m.ActionLck, m.ActionRun)
			actionFunc = func() { actions.OpenFileWindows("https://www.bilibili.com/video/BV1uh7pzGEiN/") }
		case <-m.Timer.ClickedCh: // 点击启动/取消定时器
			if timerFlag {
				SetTimerStatus(false)
				stopTimerCh <- struct{}{}
			} else {
				SetTimerStatus(true)
				go NewTimer(m.Timer, 5, actionFunc)
			}
		case <-m.StartUp.ClickedCh: // 设置开机自启
			if m.StartUp.Checked() {
				m.StartUp.Uncheck()
			} else {
				m.StartUp.Check()
			}
		case <-m.Quit.ClickedCh: // 退出
			systray.Quit()
			return
		}
	}
}

func uncheckAllExcept(clickedItem *systray.MenuItem, items ...*systray.MenuItem) {
	if !clickedItem.Checked() {
		clickedItem.Check()
		for _, item := range items {
			if item.Checked() {
				item.Uncheck()
				return
			}
		}
	}
}
