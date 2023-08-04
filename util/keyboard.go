package util

import (
	"github.com/HountryLiu/pubg-mortar-assistant/screen"
	hook "github.com/robotn/gohook"
)

func HookKeyboard() {
	var (
		KeyOpen  uint16 // 对应键盘 "+"
		KeyClose uint16 // 对应键盘 "-"
	)

	pubg := screen.GetPubgInstance()

	if pubg.OS == "windows" {
		KeyOpen = 187
		KeyClose = 189
	} else {
		// linux macos
		KeyOpen = 24
		KeyClose = 27
	}
	hooks := hook.Start()
	defer hook.End()
	for ev := range hooks {
		//	监听键盘弹起
		if ev.Kind == hook.KeyUp {
			if ev.Rawcode == KeyOpen {
				screen.ScreenShot.Create()
			} else if ev.Rawcode == KeyClose {
				screen.ScreenShot.Destroy()
			}
		}
	}
}
