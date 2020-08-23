package nodos

import (
	"encoding/json"
)

type EventHandler func(ev Event)

type Event struct {
	Name string
	Data interface{}
}

func (ev *Event) GetData(output interface{}) error {

	jsonString, err := json.Marshal(ev.Data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonString, output); err != nil {
		return err
	}
	return nil
}
