package nodos

type Connectable interface {
	SetNotificationHandler(EventHandler)

	ListenEvents(net NeuralRed)
}

//NodoBuilder build the nodo
type Cell struct {
	Name string
}

func (n *Cell) Build() chan Event {
	ch := make(chan Event)
	return ch

}
