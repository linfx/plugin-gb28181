package gb28181

import (
	"sync"
	"time"

	"github.com/ghettovoice/gosip/sip"
	"m7s.live/engine/v4/log"
)

// 设备
type Device struct {
	ID              string
	Name            string
	Manufacturer    string
	Model           string
	Owner           string
	RegisterTime    time.Time
	UpdateTime      time.Time
	LastKeepaliveAt time.Time
	Status          DeviceStatus
	SN              int
	Addr            sip.Address `json:"-" yaml:"-"`
	SipIP           string      //设备对应网卡的服务器ip
	MediaIP         string      //设备对应网卡的服务器ip
	NetAddr         string
	channelMap      sync.Map
	subscriber      struct {
		CallID  string
		Timeout time.Time
	}
	lastSyncTime time.Time
	GpsTime      time.Time //gps时间
	Longitude    string    //经度
	Latitude     string    //纬度
	*log.Logger  `json:"-" yaml:"-"`
}

// 设备位置
type DevicePosition struct {
	ID        string
	GpsTime   time.Time //gps时间
	Longitude string    //经度
	Latitude  string    //纬度
}

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
