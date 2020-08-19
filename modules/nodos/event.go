package nodos

type EventHandler func(ev Event)

type Event struct {
	Name string
	Data interface{}
}
