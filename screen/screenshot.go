package screen

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type screenShot struct{}

var ScreenShot = new(screenShot)

func (o *screenShot) Create(mode GameScreenshotMode) {
	pubg := GetPubgInstance()
	if pubg.WinScreenShotStatus {
		return
	}
	pubg.WinScreenShot = fyne.CurrentApp().NewWindow("截图")
	// 关闭截图padding
	pubg.WinScreenShot.SetPadded(false)
	// 设置全屏展示
	pubg.WinScreenShot.SetFullScreen(true)
	pubg.WinScreenShotStatus = true

	if pubg.GameMapCellLength <= 0 && mode != MeasureLengthGameCellScreen {
		text := "还未测量游戏地图一格长度，请同时按 %s 键进行测量"
		text = fmt.Sprintf(text, strings.Join(OpenMeasureLengthGameCellScreenHotKey, " 与 "))
		cnf := dialog.NewConfirm("错误", text, o.confirmCallback, pubg.Win)
		cnf.SetConfirmText("确定")
		cnf.SetDismissText("关闭")
		cnf.Show()
		return
	}
	if err := pubg.MakeScreenshot(); err != nil {
		cnf := dialog.NewConfirm("错误", err.Error(), o.confirmCallback, pubg.Win)
		cnf.SetConfirmText("确定")
		cnf.SetDismissText("关闭")
		cnf.Show()
		return
	}
	pubg.GameScreenshotMode = mode
	pubg.ViewPort = NewViewPort(pubg)

	pubg.WinScreenShot.SetContent(pubg.ViewPort)
	pubg.WinScreenShot.Show()
}

func (o *screenShot) Destroy() {
	pubg := GetPubgInstance()
	if !pubg.WinScreenShotStatus {
		return
	}
	pubg.GameScreenshotMode = NoOption
	pubg.WinScreenShotStatus = false
	pubg.WinScreenShot.Close()
}

func (o *screenShot) confirmCallback(response bool) {
	pubg := GetPubgInstance()
	pubg.WinScreenShotStatus = false
}
