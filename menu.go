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
	Timer50,
	Timer30,
	Timer10,
	TimerS,
	StartUp,
	Quit *systray.MenuItem
}

func NewMenu() *Menu {
	m := &Menu{}
	m.Actions = systray.AddMenuItem("超时动作", "")
	m.ActionLck = m.Actions.AddSubMenuItemCheckbox("锁定屏幕", "", true)
	m.ActionRun = m.Actions.AddSubMenuItem("打开文件/程序", "")
	m.ActionWeb = m.Actions.AddSubMenuItem("打开网站", "")
	m.Timer50 = systray.AddMenuItem("定时 50 min", "")
	m.Timer30 = systray.AddMenuItem("定时 30 min", "")
	m.Timer10 = systray.AddMenuItem("定时 10 min", "")
	m.TimerS = systray.AddMenuItem("停止定时器", "")
	m.TimerS.Hide()
	systray.AddSeparator() // 分隔线
	m.StartUp = systray.AddMenuItemCheckbox("开机自启", "", false)
	m.Quit = systray.AddMenuItem("退出", "")
	return m
}

func OnClickMenu(m *Menu) {
	for {
		select {
		case <-m.ActionLck.ClickedCh: // 超时锁屏
			actionFunc = actions.LockScreenWindows
			UncheckAllExcept(m.ActionLck, m.ActionRun, m.ActionWeb)
		case <-m.ActionRun.ClickedCh: // 超时打开文件或程序
			if filePath := actions.SelectFileWindows(); filePath != "" {
				actionFunc = actions.OpenFileWindows(filePath)
				UncheckAllExcept(m.ActionRun, m.ActionLck, m.ActionWeb)
			}
		case <-m.ActionWeb.ClickedCh: // 超时打开网页
			if link := actions.SelectLnkWindows(); link != "" {
				actionFunc = actions.OpenFileWindows(link)
				UncheckAllExcept(m.ActionWeb, m.ActionLck, m.ActionRun)
			}
		case <-m.TimerS.ClickedCh: // 点击启动/取消定时器
			SetTimerStatus(m, false)
			stopTimerCh <- struct{}{}
		case <-m.Timer50.ClickedCh:
			go NewTimer(m, 50*secPm)
		case <-m.Timer30.ClickedCh:
			go NewTimer(m, 30*secPm)
		case <-m.Timer10.ClickedCh:
			go NewTimer(m, 10)
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

func UncheckAllExcept(clickedItem *systray.MenuItem, items ...*systray.MenuItem) {
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
