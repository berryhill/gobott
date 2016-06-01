package bot

import(
	"fmt"
	"time"
	"encoding/json"

	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot"
	"github.com/gobott/models"
)

var mqttAdaptor *mqtt.MqttAdaptor
var owork func()

func init() {
	owork = func() {

		mqttAdaptor.On("web_to_bot", func(data []byte) {
			d := json.Unmarshal(data, models.Button{})
			fmt.Println(d)
		})


		gobot.Every(1*time.Second, func() {
			json := Buttons[0].MarshalJson()

			//bm := models.NewBaseModel()
			//json, _ := json.Marshal(bm)

			SendMessage(json)
			fmt.Print("Send Message")
			fmt.Println(json)
		})
	}
}

func NewOperator() *gobot.Robot {
	mqttAdaptor = mqtt.NewMqttAdaptor("server", "tcp://test.mosquitto.org:1883", "operator")
	robot := gobot.NewRobot("mqttBot", []gobot.Connection{mqttAdaptor}, owork, )

	return robot
}

func SendMessage(b []byte) {
	mqttAdaptor.Publish("bot_to_web", b)
}