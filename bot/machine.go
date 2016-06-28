package bot

import	(
	"time"

	"github.com/gobott-web/models"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"fmt"
)

var MACHINE *models.Machine
var HeartBeatt *gpio.LedDriver
var ReportIndicator *gpio.LedDriver
var LightSensor *gpio.AnalogSensorDriver
var LsValue uint8
var mwork func()

func init() {
	MACHINE = models.NewMachine("Test Machine")

	mwork = func() {
		gobot.Every(1 * time.Second, func() {
			Beat()
			fmt.Println(LsValue)
		})

		gobot.On(LightSensor.Event("data"), func(data interface{}) {
			LsValue = data
			//LsValue = uint8(
			//	gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			//)
			//fmt.Println("sensor", data)
		})
	}
}

func NewMachineBot(b *beaglebone.BeagleboneAdaptor) *gobot.Robot {
 	HeartBeatt = gpio.NewLedDriver(b, "led", "P9_12")
	ReportIndicator = gpio.NewLedDriver(b, "led", "P9_14")
	LightSensor = gpio.NewAnalogSensorDriver(b, "LightSensor", "P9_33")

	robot := gobot.NewRobot("Peripheral Bot", []gobot.Connection{b},
		[]gobot.Device {
			HeartBeatt,
			ReportIndicator,
			LightSensor,
		}, mwork,
	)

	return robot
}

func ReportSent() {
	ReportIndicator.Toggle()
	time.Sleep(500 * time.Millisecond)
	ReportIndicator.Toggle()
}

func Beat() {
	for k := 0; k < 4; k++ {
		HeartBeatt.Toggle()
		time.Sleep(100 * time.Millisecond)
	}
}