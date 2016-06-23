package bot

import (
	"github.com/hybridgroup/gobot/platforms/raspi"
	"github.com/siddontang/go/bson"
)

type Gateway struct {
	Id 					bson.ObjectId
	raspiAdaptor		*raspi.RaspiAdaptor
}

func NewGateway() *Gateway {
	g := new(Gateway)
	g.Id = bson.NewObjectId()
	g.raspiAdaptor = raspi.NewRaspiAdaptor("raspi")

	return g
}