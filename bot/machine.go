package bot

import	(
	"github.com/gobott-web/models"
)

var MACHINE *models.Machine

func init() {
	MACHINE = models.NewMachine("Test Machine")

	//MACHINE.AddSensor(models.NewAnalogSensor("PH Sensor"))
	//MACHINE.AddSensor(models.NewAnalogSensor("PPM Sensor"))
}