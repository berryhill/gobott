package bot

import	(
	"time"

	"github.com/gobott-web/models"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
)

var MACHINE *models.Machine
var HeartBeatt *gpio.LedDriver
var mwork func()

func init() {
	MACHINE = models.NewMachine("Test Machine")

	mwork = func() {

	}

	//var sensor models.Sensor
	//
	//sensor = models.NewAnalogSensor("PH Sensor")
	//MACHINE.AddSensor(&sensor)
	//
	//sensor = models.NewAnalogSensor("PPM Sensor")
	//MACHINE.AddSensor(&sensor)
	//
	//sensor = models.NewBooleanSensor("Button")
	//MACHINE.AddSensor(&sensor)
}

func NewMachineBot(b *beaglebone.BeagleboneAdaptor) *gobot.Robot {
 	HeartBeatt = gpio.NewLedDriver(b, "led", "P9_12")

	robot := gobot.NewRobot("Peripheral Bot", []gobot.Connection{b},
		[]gobot.Device {
			HeartBeatt,
		}, mwork,
	)

	return robot
}

func Beat() {
	for k := 0; k < 4; k++ {
		HeartBeatt.Toggle()
		time.Sleep(100 * time.Millisecond)
	}
}