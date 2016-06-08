package bot

import(
	"fmt"
	"time"
	"encoding/json"
	"log"

	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot"
	m "github.com/gobott/models"

	"github.com/gobott-web/models"
)

var mqttAdaptor *mqtt.MqttAdaptor
var owork func()

func init() {
	owork = func() {
		mqttAdaptor.On("web_to_bot_report", func(data []byte) {
			d := json.Unmarshal(data, m.Button{})
			fmt.Println(d)
		})

		gobot.Every(1*time.Second, func() {
			testReport := models.NewReport()

			json, err := testReport.MarshalJson()

			if err != nil {
				log.Fatal(err)
			} else {
				SendMessage("bot_to_web_report", json)
			}
		})
	}
}

func NewOperator() *gobot.Robot {
	mqttAdaptor = mqtt.NewMqttAdaptor("server", "tcp://test.mosquitto.org:1883", "operator")
	robot := gobot.NewRobot("mqttBot", []gobot.Connection{mqttAdaptor}, owork, )

	return robot
}

func SendMessage(topic string, b []byte) {
	mqttAdaptor.Publish(topic, b)
	fmt.Println("Sending Json")
}