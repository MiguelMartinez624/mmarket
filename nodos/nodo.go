package nodos

type Nodo interface {
	SetNotificationHandler(EventHandler)

	ListenEvents(net NeuralRed)
}

//NodoBuilder build the nodo
type NodoBuilder struct {
	Name string
	Nodo Nodo
}

func (n NodoBuilder) Build() chan Event {
	ch := make(chan Event)
	return ch

}
