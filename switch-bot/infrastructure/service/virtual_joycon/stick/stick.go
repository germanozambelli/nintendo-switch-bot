package stick

import "github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/stick"

type Stick string

const (
	LEFT  Stick = "L_STICK"
	RIGHT       = "R_STICK"
)

func ConvertToStick(s stick.Stick) Stick {
	switch s {
	case stick.LEFT:
		return LEFT
	case stick.RIGHT:
		return RIGHT
	}

	return "unknown"
}
