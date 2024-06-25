package gb28181

// Record 录像
type Record struct {
	DeviceID  string
	Name      string
	FilePath  string
	Address   string
	StartTime string
	EndTime   string
	Secrecy   int
	Type      string
}

// 预置位
type Preset struct {
	DeviceID   string // 设备ID
	PresetID   int    // 预置位编号
	PresetName string // 预置位名称
}
