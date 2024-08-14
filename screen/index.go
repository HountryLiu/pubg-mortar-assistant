package screen

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/HountryLiu/pubg-mortar-assistant/util"
	"golang.org/x/image/colornames"
)

var _ fyne.Widget = (*Index)(nil)

// Index is a widget for displaying a index with themeable color.
//
// Since: 1.4
type Index struct {
	widget.BaseWidget
	L1 *canvas.Text
	L2 *canvas.Text
}

// NewIndex creates a new index.
//
// Since: 1.4
func NewIndex() *Index {
	s := &Index{}
	pubg := GetPubgInstance()
	pubg.IndexPage = s
	return s
}

// CreateRenderer returns a new renderer for the index.
//
// Implements: fyne.Widget
func (s *Index) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	bar := canvas.NewRectangle(theme.DisabledColor())

	mainFn := func(title string) (l *widget.Label) {
		l = widget.NewLabel(title)
		l.TextStyle.Bold = true
		l.Alignment = fyne.TextAlignTrailing
		return l
	}
	titleFn := func(title string) (l *widget.Label) {
		l = widget.NewLabel(title)
		l.TextStyle.Bold = true
		l.Alignment = fyne.TextAlignLeading
		return l
	}
	descFn := func(desc string) (l *widget.Label) {
		l = widget.NewLabel(desc)
		l.Alignment = fyne.TextAlignCenter
		return l
	}
	contentFn := func(shortcut string) (l *canvas.Text) {
		l = canvas.NewText(shortcut, colornames.Red)
		l.TextStyle.Italic = true
		l.TextStyle.Bold = true
		l.Alignment = fyne.TextAlignTrailing
		return l
	}
	pubg := GetPubgInstance()
	s.L1 = contentFn(strconv.Itoa(util.Round(pubg.GameMapCellLength)) + " 像素")
	s.L2 = contentFn(strconv.Itoa(pubg.GameMortarShootLength) + " 米")
	ctnTop := container.NewVBox(
		mainFn("欢迎使用绝地求生迫击炮助手! "),
		titleFn("全局快捷键："),
		container.NewGridWithColumns(2,
			descFn("测量游戏地图一格长度"), contentFn("'Alt' + '-'"),
			descFn("测量迫击炮的发射距离"), contentFn("'Alt' + '='"),
		),
		widget.NewLabel(""),
		titleFn("展示测量长度数据："),
		container.NewGridWithColumns(2,
			descFn("游戏地图一格长度："), s.L1,
			descFn("迫击炮的发射距离："), s.L2,
		),
	)
	ctnBottom := container.NewCenter(
		container.NewHBox(
			widget.NewHyperlink("博客", util.ParseURL("https://olo.ink/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("Github", util.ParseURL("https://github.com/HountryLiu/pubg-mortar-assistant")),
		),
	)
	ctn := container.NewBorder(ctnTop, ctnBottom, nil, nil)
	return &indexRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(ctn),
		bar:            bar,
		d:              s,
	}
}

// MinSize returns the minimal size of the index.
//
// Implements: fyne.Widget
func (s *Index) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return fyne.NewSize(500, 350)
}

var _ fyne.WidgetRenderer = (*indexRenderer)(nil)

type indexRenderer struct {
	fyne.WidgetRenderer
	bar *canvas.Rectangle
	d   *Index
}

func (r *indexRenderer) Layout(s fyne.Size) {
}

func (r *indexRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *indexRenderer) Refresh() {
	r.bar.FillColor = theme.DisabledColor()
	canvas.Refresh(r.d)
}
