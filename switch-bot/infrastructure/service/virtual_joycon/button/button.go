package button

import "github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"

type Button string

const (
	A       Button = "A"
	B              = "B"
	X              = "X"
	Y              = "Y"
	HOME           = "HOME"
	UP             = "DPAD_UP"
	DOWN           = "DPAD_DOWN"
	LEFT           = "DPAD_LEFT"
	RIGHT          = "DPAD_RIGHT"
	CAPTURE        = "CAPTURE"
	MINUS          = "MINUS"
	PLUS           = "PLUS"
	L              = "L"
)

func ConvertToButton(b button.Button) Button {
	switch b {
	case button.A:
		return A
	case button.B:
		return B
	case button.X:
		return X
	case button.Y:
		return Y
	case button.HOME:
		return HOME
	case button.UP:
		return UP
	case button.DOWN:
		return DOWN
	case button.LEFT:
		return LEFT
	case button.RIGHT:
		return RIGHT
	case button.CAPTURE:
		return CAPTURE
	case button.MINUS:
		return MINUS
	case button.PLUS:
		return PLUS
	case button.L:
		return L
	}

	return "unknown"
}
