package screen

import (
	"fyne.io/fyne/v2"
)

type screenShot struct{}

var ScreenShot = new(screenShot)

func (o *screenShot) Create() {
	pubg := GetPubgInstance()
	if pubg.WinScreenShotStatus {
		return
	}
	pubg.WinScreenShot = fyne.CurrentApp().NewWindow("截图")
	// 关闭截图padding
	pubg.WinScreenShot.SetPadded(false)
	// 设置全屏展示
	pubg.WinScreenShot.SetFullScreen(true)

	pubg.MakeScreenshot()
	pubg.ViewPort = NewViewPort(pubg)

	pubg.WinScreenShot.SetContent(pubg.ViewPort)
	pubg.WinScreenShotStatus = true
	pubg.WinScreenShot.Show()
}

func (o *screenShot) Destroy() {
	pubg := GetPubgInstance()
	if !pubg.WinScreenShotStatus {
		return
	}
	pubg.WinScreenShotStatus = false
	pubg.WinScreenShot.Close()
}
