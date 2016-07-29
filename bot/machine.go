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
var Thermistor *gpio.AnalogSensorDriver

var LightSensorValue interface{}
var ThermistorValue interface{}
var mwork func()

func init() {
	MACHINE = models.NewMachine("Test Machine")

	MACHINE.AddSensor(models.NewAnalogSensor("LightSensor"))
	MACHINE.AddSensor(models.NewAnalogSensor("Thermistor"))

	mwork = func() {
		gobot.Every(1 * time.Second, func() {
			Beat()

			MACHINE.Sensors[0].Value = int32(LightSensorValue.(int))
			fmt.Println("LightSensor Value: ", LightSensorValue)

			MACHINE.Sensors[1].Value = int32(ThermistorValue.(int))
			fmt.Println("Thermistor Value: ", ThermistorValue)
		})

		gobot.On(LightSensor.Event("data"), func(data interface{}) {
			LightSensorValue = data
			//LsValue = uint8(
			//	gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			//)
			//fmt.Println("sensor", data)
		})
		gobot.On(Thermistor.Event("data"), func(data interface{}) {
			ThermistorValue = data
		})
	}
}

func NewMachineBot(b *beaglebone.BeagleboneAdaptor) *gobot.Robot {
 	HeartBeatt = gpio.NewLedDriver(b, "led", "P9_12")
	ReportIndicator = gpio.NewLedDriver(b, "led", "P9_14")
	LightSensor = gpio.NewAnalogSensorDriver(b, "LightSensor", "P9_33")
	Thermistor = gpio.NewAnalogSensorDriver(b, "LightSensor", "P9_35")

	robot := gobot.NewRobot("Peripheral Bot", []gobot.Connection{b},
		[]gobot.Device {
			HeartBeatt,
			ReportIndicator,
			LightSensor,
			Thermistor,
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