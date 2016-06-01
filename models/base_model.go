package models

type BaseModel struct {
	Name 		string                `json:"name"`
}

func NewBaseModel() *BaseModel {
	b := new(BaseModel)
	b.Name = "Tester"
	return b
}