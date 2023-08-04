package main

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/HountryLiu/pubg-mortar-assistant/resource"
	"github.com/HountryLiu/pubg-mortar-assistant/screen"
	"github.com/HountryLiu/pubg-mortar-assistant/theme"
	"github.com/HountryLiu/pubg-mortar-assistant/util"
)

func init() {
	pubg := screen.GetPubgInstance()
	pubg.App = app.NewWithID("pubg-mortar-assistant")
	// 支持中文显示
	pubg.App.Settings().SetTheme(&theme.ShanGShouJianSongTheme{RefThemeApp: pubg.App})
	// 创建主窗口
	pubg.Win = pubg.App.NewWindow("绝地求生迫击炮助手")
	// 设置桌面图标
	pubg.Win.SetIcon(resource.DesktopIcon)
	// 默认选择主屏幕
	pubg.DisplayIndex = 0
	// 设置当前屏幕分辨率
	pubg.ScreenWidth, pubg.ScreenHeight = util.GetScreenSize()
	// 设置当前操作系统
	pubg.OS = runtime.GOOS
}

func main() {
	// 全局监听键盘，快捷键截图
	go util.HookKeyboard()

	pubg := screen.GetPubgInstance()
	textLabel := widget.NewLabel("欢迎使用绝地求生迫击炮助手! '+'键开始测量，'-'键关闭测量")
	pubg.Win.SetContent(textLabel)
	pubg.Win.Resize(fyne.NewSize(150, 50))
	pubg.Win.ShowAndRun()

}
