package screen

import (
	"fmt"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

type ViewPort struct {
	widget.BaseWidget

	// 应用全局变量指针.
	pubg *pubg

	// 鼠标是否开始拖动
	isDrag    bool
	dragStart fyne.Position

	// 鼠标拖动的直线
	DrawLine *canvas.Line

	// 截图图片
	Screenshot *canvas.Image

	// 迫击炮攻击距离显示
	MortarAttachDistanceView *canvas.Text
}

// NewViewPort 视窗，放置需要编辑的视图
func NewViewPort(pubg *pubg) (vp *ViewPort) {
	vp = &ViewPort{}
	vp.pubg = pubg
	vp.DrawLine = canvas.NewLine(colornames.Red)
	vp.DrawLine.StrokeWidth = 3
	vp.Screenshot = canvas.NewImageFromImage(pubg.Screenshot)
	vp.Screenshot.Resize(fyne.NewSize(float32(pubg.ScreenWidth), float32(pubg.ScreenHeight)))
	vp.MortarAttachDistanceView = canvas.NewText("", colornames.Red)
	vp.MortarAttachDistanceView.TextSize = 16 * float32(pubg.Win.Canvas().Scale())
	return
}

// DragEnd implements fyne.Draggable
func (vp *ViewPort) DragEnd() {
	vp.isDrag = false
}

// Dragged function.
func (vp *ViewPort) Dragged(e *fyne.DragEvent) {
	if !vp.isDrag {
		vp.dragStart = e.Position
		vp.isDrag = true
	}
	vp.DrawLine.Position1 = vp.dragStart
	vp.DrawLine.Position2 = e.Position
	x1 := vp.DrawLine.Position1.X
	x2 := vp.DrawLine.Position2.X
	y1 := vp.DrawLine.Position1.Y
	y2 := vp.DrawLine.Position2.Y
	drawLineDistince := math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))

	vp.MortarAttachDistanceView.Text = fmt.Sprintf(" %v 米", int(drawLineDistince*float64(vp.pubg.GameMapRatio)))
	vp.MortarAttachDistanceView.Move(e.Position)
	vp.MortarAttachDistanceView.Refresh()
	vp.DrawLine.Refresh()
}

// CreateRenderer returns a new renderer for the ViewPort.
//
// Implements: fyne.Widget
func (vp *ViewPort) CreateRenderer() fyne.WidgetRenderer {
	vp.ExtendBaseWidget(vp)
	bar := canvas.NewRectangle(theme.DisabledColor())

	ctn := container.NewWithoutLayout(vp.Screenshot, vp.DrawLine, vp.MortarAttachDistanceView)
	return &ViewPortRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(ctn),
		bar:            bar,
		d:              vp,
	}
}

// MinSize returns the minimal size of the ViewPort.
//
// Implements: fyne.Widget
func (vp *ViewPort) MinSize() fyne.Size {
	vp.ExtendBaseWidget(vp)
	return fyne.NewSize(float32(vp.pubg.ScreenWidth), float32(vp.pubg.ScreenHeight))
}

var _ fyne.WidgetRenderer = (*ViewPortRenderer)(nil)

type ViewPortRenderer struct {
	fyne.WidgetRenderer
	bar *canvas.Rectangle
	d   *ViewPort
}

func (r *ViewPortRenderer) Layout(s fyne.Size) {
	if !r.d.isDrag {
		r.d.DrawLine.Position1 = fyne.NewPos(0, 0)
		r.d.DrawLine.Position2 = fyne.NewPos(0, 0)
	}
}

func (r *ViewPortRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *ViewPortRenderer) Refresh() {
	r.bar.FillColor = theme.DisabledColor()
	canvas.Refresh(r.d)
}
