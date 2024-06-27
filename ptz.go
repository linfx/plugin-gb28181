package gb28181

import "fmt"

var (
	name2code = map[string]uint8{
		"stop":      0, // 停止
		"right":     1, // 右转
		"left":      2, // 左转
		"down":      4, // 下转
		"downright": 5,
		"downleft":  6,
		"up":        8,
		"upright":   9,
		"upleft":    10,
		"zoomin":    16,
		"zoomout":   32,

		// preset
		"set":    0x81,
		"goto":   0x82,
		"remove": 0x83,
	}
)

func toPtzCode(cmd string) (uint8, error) {
	if code, ok := name2code[cmd]; ok {
		return code, nil
	} else {
		return 0, fmt.Errorf("invalid ptz cmd %q", cmd)
	}
}

func toPTZCmdByName(cmdName string, horizontalSpeed, verticalSpeed, zoomSpeed uint8) (string, error) {
	code, err := toPtzCode(cmdName)
	if err != nil {
		return "", err
	}

	checkCode := uint16(0xA5+0x0F+0x01+code+horizontalSpeed+verticalSpeed+(zoomSpeed&0xF0)) % 0x100
	return fmt.Sprintf("A50F01%02X%02X%02X%01X0%02X",
		code,
		horizontalSpeed,
		verticalSpeed,
		zoomSpeed>>4, // 根据 GB28181 协议，zoom 只取 4 bit
		checkCode,
	), err
}

func toPTZCmdByName_preset(cmdName string, preset uint8) (string, error) {
	code, err := toPtzCode(cmdName)
	if err != nil {
		return "", err
	}

	checkCode := uint16(0xA5+0x0F+0x01+code+0x00+preset+0x00) % 0x100
	return fmt.Sprintf("A50F01%02X00%02X00%02X",
		code,
		preset,
		checkCode,
	), err
}
