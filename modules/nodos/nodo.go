package nodos

type Nodo interface {
	ListenEvents(net chan Event)
}
