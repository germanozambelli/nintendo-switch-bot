package nxbt_joycon

import (
	"bufio"
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"net"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller/stick"
)

type VirtualJoyCon struct {
	connection *net.Conn
	logger     logger.Logger
}

func NewVirtualJoyCon(host string, logger logger.Logger) (*VirtualJoyCon, error) {
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:65431", host))

	if err != nil {
		return nil, err
	}

	return &VirtualJoyCon{
		connection: &connection,
		logger:     logger,
	}, nil
}

func (virtualJoyCon *VirtualJoyCon) PressButton(button button.Button) {
	duration := 81 * time.Millisecond

	jButton := convertToButton(button)

	pressAction := fmt.Sprintf("%s %fs", jButton, duration.Seconds())

	virtualJoyCon.sendMessage(pressAction)

	time.Sleep(100 * time.Millisecond)
}

func (virtualJoyCon *VirtualJoyCon) HoldButton(button button.Button, duration time.Duration) {
	if duration < 75*time.Millisecond {
		virtualJoyCon.PressButton(button)
		return
	}

	jButton := convertToButton(button)

	pressAction := fmt.Sprintf("%s %fs", jButton, duration.Seconds())

	virtualJoyCon.sendMessage(pressAction)

	time.Sleep(75 * time.Millisecond)
}

func (virtualJoyCon *VirtualJoyCon) MoveStick(
	stick stick.Stick,
	xPosition int,
	yPosition int,
	duration time.Duration,
) {
	if duration < 75*time.Millisecond {
		duration = 75 * time.Millisecond
	}

	jStick := convertToStick(stick)

	moveAction := fmt.Sprintf("%s@%+04d%+04d %fs", jStick, xPosition, yPosition, duration.Seconds())

	virtualJoyCon.sendMessage(moveAction)

	time.Sleep(75 * time.Millisecond)
}

func (virtualJoyCon *VirtualJoyCon) sendMessage(message string) {
	message = fmt.Sprintf("%s\n", message)

	if virtualJoyCon.connection == nil {
		panic("controller not connected")
	}

	(*virtualJoyCon.connection).Write([]byte(message))

	bufio.NewReader(*virtualJoyCon.connection).ReadString('\n')

	virtualJoyCon.logger.Debug(fmt.Sprintf("sent message: %s", message))

	time.Sleep(30 * time.Millisecond)
}

func convertToStick(s stick.Stick) string {
	switch s {
	case stick.LEFT:
		return "L_STICK"
	case stick.RIGHT:
		return "R_STICK"
	}

	panic("unknown stick")
}

func convertToButton(b button.Button) string {
	switch b {
	case button.A:
		return "A"
	case button.B:
		return "B"
	case button.X:
		return "X"
	case button.Y:
		return "Y"
	case button.HOME:
		return "HOME"
	case button.UP:
		return "DPAD_UP"
	case button.DOWN:
		return "DPAD_DOWN"
	case button.LEFT:
		return "DPAD_LEFT"
	case button.RIGHT:
		return "DPAD_RIGHT"
	case button.CAPTURE:
		return "CAPTURE"
	case button.MINUS:
		return "MINUS"
	case button.PLUS:
		return "PLUS"
	case button.L:
		return "L"
	}

	panic("unknown button")
}
