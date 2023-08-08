package screen

type GameScreenshotMode int

const (
	// 无操作
	NoOption GameScreenshotMode = iota
	// 测量游戏地图一格在当前显示器屏幕长度
	MeasureLengthGameCellScreen
	// 测量迫击炮发射距离
	MeasureLengthMortarShoot
)

var (
	// 快捷键监听
	OpenMeasureLengthGameCellScreenHotKey = []string{"alt", "-"}
	OpenMeasureLengthMortarShootHotKey    = []string{"alt", "="}
	CloseMeasureLengthViewHotKey          = []string{"alt", "delete"}
)
