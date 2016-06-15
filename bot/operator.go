package bot

import(
	"fmt"
	"time"
	"log"

	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot"

	"github.com/gobott-web/models"
)

var mqttAdaptor *mqtt.MqttAdaptor
var owork func()
var Timer int
var On bool

func init() {
	On = true

	owork = func() {
		gobot.Every(5 * time.Second, func() {
			if On == true {
				testReport := models.NewReport(MACHINE)

				json, err := testReport.MarshalJson()

				if err != nil {
					log.Fatal(err)
				}

				sendMessage("bot_to_web", json)
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

func handleMessage(data []byte) {
	fmt.Println("Handling Message")

	if string(data) == "start_bot" {
		ResumeOutboundReport()

		fmt.Println("Starting Bot")
	} else if string(data) == "stop_bot" {
		HaltOutboundReport()

		fmt.Println("Stopping Bot")
	} else {
		//d := json.Unmarshal(data, models.Button{})

		fmt.Println(data)
	}
}

func sendMessage(topic string, b []byte) {
	mqttAdaptor.Publish(topic, b)
	fmt.Println("Sending Json")
}