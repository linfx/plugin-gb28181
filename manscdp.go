package gb28181

import (
	"fmt"
	"strconv"
	"time"
)

func intTotime(t int64) time.Time {
	tstr := strconv.FormatInt(t, 10)
	if len(tstr) == 10 {
		return time.Unix(t, 0)
	}
	if len(tstr) == 13 {
		return time.UnixMilli(t)
	}
	return time.Now()
}

// 获取设备控制指令
func BuildControlXML(cmd string, sn int, id string) string {

	xml := `<?xml version="1.0"?>
	<Query>
		<CmdType>%s</CmdType>
		<SN>%d</SN>
		<DeviceID>%s</DeviceID>
	</Query>`

	return fmt.Sprintf(xml, cmd, sn, id)
}

// 获取设备详情指令
func BuildDeviceInfoXML(sn int, id string) string {
	return BuildControlXML("DeviceInfo", sn, id)
}

// 获取NVR下设备列表指令
func BuildCatalogXML(sn int, id string) string {
	return BuildControlXML("Catalog", sn, id)
}

// 报警订阅
func BuildAlarmXML(sn int, id string) string {

	xml := `<?xml version="1.0"?>
	<Query>
		<CmdType>Alarm</CmdType>
		<SN>%d</SN>
		<DeviceID>%s</DeviceID>
		<StartAlarmPriority>1</StartAlarmPriority>
		<EndAlarmPriority>4</EndAlarmPriority>
		<AlarmMethod>0</AlarmMethod>
	</Query>`

	return fmt.Sprintf(xml, sn, id)
}

// 移动位置订阅
func BuildDevicePositionXML(sn int, id string, interval int) string {

	xml := `<?xml version="1.0"?>
	<Query>
		<CmdType>MobilePosition</CmdType>
		<SN>%d</SN>
		<DeviceID>%s</DeviceID>
		<Interval>%d</Interval>
	</Query>`

	return fmt.Sprintf(xml, sn, id, interval)
}

// 获取录像文件列表指令
func BuildRecordInfoXML(sn int, id string, start, end int64) string {

	// 获取录像文件列表xml样式
	xml := `<?xml version="1.0"?>
	<Query>
		<CmdType>RecordInfo</CmdType>
		<SN>%d</SN>
		<DeviceID>%s</DeviceID>
		<StartTime>%s</StartTime>
		<EndTime>%s</EndTime>
		<Secrecy>0</Secrecy>
		<Type>all</Type>
	</Query>`

	return fmt.Sprintf(xml, sn, id, intTotime(start).Format("2006-01-02T15:04:05"), intTotime(end).Format("2006-01-02T15:04:05"))
}

// 报警订阅结果指令
func BuildAlarmResponseXML(sn int, id string) string {

	// alarm response xml样式
	xml := `<?xml version="1.0"?>
	<Response>
		<CmdType>Alarm</CmdType>
		<SN>%d</SN>
		<DeviceID>%s</DeviceID>
	</Response>`

	return fmt.Sprintf(xml, sn, id)
}

// 获取预置位列表
func BuildPresetXML(sn int, id string) string {

	// 获取预置位列表
	xml := `<?xml version="1.0"?>
	<Query>
		<CmdType>PresetQuery</CmdType>
		<SN>%d</SN>
		<DeviceID>%s</DeviceID>
	</Query>`

	return fmt.Sprintf(xml, sn, id)
}
