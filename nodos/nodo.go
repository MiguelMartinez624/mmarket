package nodos

type Nodo interface {

	SetNotificationHandler(EventHandler)

	ListenEvents(net chan Event)
}
