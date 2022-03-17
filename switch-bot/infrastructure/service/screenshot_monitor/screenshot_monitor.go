package screenshot_monitor

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/action"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/stick"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type ScreenshotMonitor struct {
	controller controller.Controller
}

func NewScreenshotMonitor(controller controller.Controller) ScreenshotMonitor {
	return ScreenshotMonitor{controller: controller}
}

func (monitor ScreenshotMonitor) CurrentDialogContains(texts []string) bool {
	monitor.takeScreenshot()
	monitor.startUsbSharing()
	screenshot := monitor.getLastScreenshot()
	imageText := monitor.getTextFromImage(screenshot)
	monitor.stopUsbSharing()

	fmt.Println(imageText)

	for _, text := range texts {
		if strings.Contains(imageText, text) {
			return true
		}
	}

	return false
}

func (monitor ScreenshotMonitor) getTextFromImage(image string) string {
	request := map[string]string{"base64": image, "trim": "\n"}
	jsonData, err := json.Marshal(request)

	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:8080/base64", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	result := map[string]string{}

	json.NewDecoder(response.Body).Decode(&result)

	return result["result"]
}

func (monitor ScreenshotMonitor) getLastScreenshot() string {
	dir := "/run/user/1000/gvfs/mtp:host=Nintendo_Nintendo_Switch_XTJ10384125725/Album/The Legend of Zelda: Breath of the Wild/"

	files, err1 := ioutil.ReadDir(dir)

	if err1 != nil {
		monitor.startUsbSharing()
		return monitor.getLastScreenshot()
	}

	var newestFile string
	var newestTime int64 = 0
	for _, f := range files {
		fi, err := os.Stat(dir + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		currTime := fi.ModTime().Unix()
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}

	fmt.Println(newestFile)
	cmd := exec.Command("gio", "copy", dir+newestFile, "/tmp/"+newestFile)

	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		panic("error in gio command")
	}

	image_file, err := os.Open("/tmp/" + newestFile)
	if err != nil {
		log.Fatal(err)
	}
	my_image, err := jpeg.Decode(image_file)
	if err != nil {
		log.Fatal(err)
	}

	croppedImg, err := cutter.Crop(my_image, cutter.Config{
		Width:  550,
		Height: 98,
		Anchor: image.Point{365, 545},
		Mode:   cutter.TopLeft,
	})

	output_file, outputErr := os.Create("/tmp/" + newestFile)
	if outputErr != nil {
		log.Fatal(outputErr)
	}
	jpeg.Encode(output_file, croppedImg, nil)

	f, _ := os.Open("/tmp/" + newestFile)

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded
}

func (monitor ScreenshotMonitor) startUsbSharing() {
	action.NewCombine([]action.Action{
		action.NewPressButton(button.HOME, 75*time.Millisecond, monitor.controller),
		action.NewNothing(1 * time.Second),
		action.NewRepeat(
			action.NewPressButton(button.DOWN, 75*time.Millisecond, monitor.controller),
			2,
		),
		action.NewRepeat(
			action.NewPressButton(button.RIGHT, 75*time.Millisecond, monitor.controller),
			5,
		),
		action.NewPressButton(button.A, 75*time.Millisecond, monitor.controller),
		action.NewNothing(1 * time.Second),
		action.NewRepeat(
			action.NewPressButton(button.DOWN, 75*time.Millisecond, monitor.controller),
			5,
		),
		action.NewPressButton(button.RIGHT, 75*time.Millisecond, monitor.controller),
		action.NewMoveStick(stick.LEFT, 0, -100, 1*time.Second, monitor.controller),
		action.NewPressButton(button.UP, 75*time.Millisecond, monitor.controller),
		action.NewPressButton(button.A, 75*time.Millisecond, monitor.controller),
		action.NewNothing(1 * time.Second),
		action.NewRepeat(
			action.NewPressButton(button.DOWN, 75*time.Millisecond, monitor.controller),
			4,
		),
		action.NewPressButton(button.A, 75*time.Millisecond, monitor.controller),
		action.NewNothing(3 * time.Second),
	}).Play()
}

func (monitor ScreenshotMonitor) stopUsbSharing() {
	action.NewCombine([]action.Action{
		action.NewRepeat(
			action.NewPressButton(button.DOWN, 75*time.Millisecond, monitor.controller),
			2,
		),
		action.NewPressButton(button.A, 75*time.Millisecond, monitor.controller),
		action.NewPressButton(button.HOME, 75*time.Millisecond, monitor.controller),
		action.NewNothing(1 * time.Second),
		action.NewPressButton(button.HOME, 75*time.Millisecond, monitor.controller),
		action.NewNothing(1 * time.Second),
	}).Play()
}

func (monitor ScreenshotMonitor) takeScreenshot() {
	action.NewPressButton(button.CAPTURE, 75*time.Millisecond, monitor.controller).Play()
}
