package bot

import (
	"github.com/gobott/store"

	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"fmt"
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

	fmt.Println("Gateway Retrieved")
	fmt.Println(gateway)

	return gateway, err
}

func (g *Gateway) Save() error {
	gatewayJson, err := json.Marshal(g)
	if err != nil {
		return err
	}

	store.AddToDb([]byte("machine"), []byte("machine"), gatewayJson)
	println(gatewayJson)

	return err
}

func (g *Gateway) Update() error {
	gatewayJson, err := json.Marshal(g)
	if err !=  nil {
		return err
	}

	fmt.Println("Updating Gateway")
	store.UpdateDb([]byte("machine"), []byte("machine"), gatewayJson)

	return err
}