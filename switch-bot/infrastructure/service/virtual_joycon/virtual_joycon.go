package virtual_joycon

import (
	"bufio"
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/stick"
	joyconButton "github.com/germanozambelli/nintendo-switch-bot/switch-bot/infrastructure/service/virtual_joycon/button"
	joyconStick "github.com/germanozambelli/nintendo-switch-bot/switch-bot/infrastructure/service/virtual_joycon/stick"
	"net"
	"time"
)

type VirtualJoyCon struct {
	connection *net.Conn
}

func NewVirtualJoyCon() *VirtualJoyCon {
	return &VirtualJoyCon{}
}

func (virtualJoyCon *VirtualJoyCon) PressButton(button button.Button, duration time.Duration) {

	if duration < 75*time.Millisecond {
		duration = 75 * time.Millisecond
	}

	jButton := joyconButton.ConvertToButton(button)

	pressAction := fmt.Sprintf("%s %fs", jButton, duration.Seconds())

	virtualJoyCon.sendMessage(pressAction)

	//println(fmt.Sprintf("pressed %s", jButton))
}

func (virtualJoyCon *VirtualJoyCon) MoveStick(stick stick.Stick, xPosition int, yPosition int, duration time.Duration) {
	if duration < 75*time.Millisecond {
		duration = 75 * time.Millisecond
	}

	jStick := joyconStick.ConvertToStick(stick)

	moveAction := fmt.Sprintf("%s@%+04d%+04d %fs", jStick, xPosition, yPosition, duration.Seconds())

	virtualJoyCon.sendMessage(moveAction)

	println(fmt.Sprintf("moved %s stick", stick))
}

func (virtualJoyCon *VirtualJoyCon) Connect() error {
	connection, err := net.Dial("tcp", "127.0.0.1:65431")

	if err != nil {
		return err
	}

	virtualJoyCon.connection = &connection

	return nil
}

func (virtualJoyCon *VirtualJoyCon) sendMessage(message string) {
	message = fmt.Sprintf("%s\n", message)

	if virtualJoyCon.connection == nil {
		panic("controller not connected")
	}

	(*virtualJoyCon.connection).Write([]byte(message))

	//wait response
	bufio.NewReader(*virtualJoyCon.connection).ReadString('\n')

	//println(msg)

	time.Sleep(30 * time.Millisecond)
}
