package bot

import (
	"fmt"
	"time"
	"log"

	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot"

	"github.com/gobott-web/models"
	"strings"
	"strconv"
)

var mqttAdaptor *mqtt.MqttAdaptor
var owork func()
var On bool
var Timer *models.Timer
var Counter int

func init() {
	On = true
	Timer = new(models.Timer)
	Timer.Seconds = 5
	Counter = 1

	owork = func() {
		gobot.Every(1 * time.Second, func() {
			if On == true {
				if Counter >= Timer.Seconds {
					report := models.NewReport(MACHINE)
					report.Name = "Green1"

					json, err := report.MarshalJson()
					if err != nil {
						log.Fatal(err)
					}

					SendMessage("bot_to_web", json)
					ReportSent()
					Counter = 0
				}

				Counter += 1
			}
		})

		mqttAdaptor.On("web_to_bot", func(data []byte) {
			handleMessage(data)
		})
	}
}

func ResumeOutboundReport() {
	On = true
}

func HaltOutboundReport() {
	On = false
}

func NewOperator() *gobot.Robot {
	mqttAdaptor = mqtt.NewMqttAdaptor("server", "tcp://test.mosquitto.org:1883", "operator")
	robot := gobot.NewRobot("mqttBot", []gobot.Connection{mqttAdaptor}, owork, )

	return robot
}

func SendMessage(topic string, b []byte) {
	mqttAdaptor.Publish(topic, b)
	//fmt.Println("Sending Json")
}

func handleMessage(data []byte) error {
	fmt.Println("Handling Message")
	fmt.Println(data)

	dataStrs := strings.Split(string(data), " ")

	if dataStrs[0] == "start_bot" {
		ResumeOutboundReport()
		fmt.Println("Starting Bot")

	} else if dataStrs[0] == "stop_bot" {
		HaltOutboundReport()
		fmt.Println("Stopping Bot")

	} else if dataStrs[0] == "timer" {
		fmt.Println("Timer" + string(Counter))
		/*
		t := new(models.Timer)

		if err := json.Unmarshal(data, t); err != nil {
			return err
		}
		*/
		seconds, err := strconv.Atoi(dataStrs[1])
		if err != nil {
			return err
		}

		Timer.Seconds = seconds
	}

	return nil
}

