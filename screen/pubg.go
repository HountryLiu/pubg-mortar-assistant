package screen

import (
	"image"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/golang/glog"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
)

type pubg struct {
	// 应用和窗口
	App fyne.App

	// 主窗口
	Win fyne.Window

	// 截图窗口
	WinScreenShot fyne.Window
	// 截图窗口开启状态
	WinScreenShotStatus bool

	// 截图信息
	Screenshot *image.RGBA
	CropRect   image.Rectangle

	// 联系我们界面
	ConnectUsDialog dialog.Dialog

	// 记录当前需要截取那个屏幕,默认情况下是0
	DisplayIndex int

	// 显示器屏幕宽
	ScreenWidth int

	// 显示器屏幕高
	ScreenHeight int

	//当前操作系统
	OS string

	// 预览窗口
	ViewPort *ViewPort

	// 游戏地图一格 在当前显示器屏幕长度
	GameMapCellLength float64
	// 游戏地图一格 在游戏中的比例，1比多少米
	GameMapCellRatio float64
}

var (
	p    *pubg
	once = &sync.Once{}
)

func GetPubgInstance() *pubg {
	once.Do(func() {
		p = &pubg{}
	})
	return p
}

// MakeScreenshot 开始截屏
func (o *pubg) MakeScreenshot() error {

	// 获取当前显示器左上角和右下角的位置信息 eg (0,0) (1920, 1080)
	bounds := screenshot.GetDisplayBounds(o.DisplayIndex)

	var err error
	// 根据指定的bounds信息截取屏幕
	o.Screenshot, err = screenshot.CaptureRect(bounds)
	if err != nil {
		glog.Errorf("CaptureRect failed.")
		return err
	}
	o.CropRect = o.Screenshot.Bounds()

	return nil
}

func (o *pubg) HookKeyboard() {
	var (
		KeyOpen  uint16 // 对应键盘 "+"
		KeyClose uint16 // 对应键盘 "-"
	)

	pubg := GetPubgInstance()

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
				ScreenShot.Create()
			} else if ev.Rawcode == KeyClose {
				ScreenShot.Destroy()
			}
		}
	}
}
