package bot

import	(
	"github.com/gobott-web/models"
)

var MACHINE *models.Machine

func init() {
	MACHINE = models.NewMachine("Test Machine")

	var sensor models.Sensor

	sensor = models.NewAnalogSensor("PH Sensor")
	MACHINE.AddSensor(&sensor)

	sensor = models.NewAnalogSensor("PPM Sensor")
	MACHINE.AddSensor(&sensor)

	sensor = models.NewBooleanSensor("Button")
	MACHINE.AddSensor(&sensor)
}