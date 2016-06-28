package bot

import (
	"github.com/gobott/store"

	"github.com/siddontang/go/bson"
	"encoding/json"
)

type Gateway struct {
	Id 					bson.ObjectId
}

func NewGateway() *Gateway {
	g := new(Gateway)
	g.Id = bson.NewObjectId()
	g.Save()

	return g
}

func (g *Gateway) Retrieve() (*Gateway, error) {
	gatewayJson, err := store.RetrieveFromDb([]byte("machine"), []byte("machine"))
	gateway := new(Gateway)
	err = json.Unmarshal(gatewayJson, gateway)

	return gateway, err
}

func (g *Gateway) Save() error {
	gatewayJson, err := json.Marshal(g)
	if err != nil {
		return err
	}

	store.AddToDb([]byte("machine"), []byte("machine"), gatewayJson)

	return err
}