package nodos

type Manager struct {
	Nodos []Nodo
}

func (m *Manager) Start() {

	net := make(chan Event)

	for _, nodo := range m.Nodos {
		nodo.SetNotificationHandler(func(ev Event) {
			net <- ev
		})
		nodo.ListenEvents(net)
	}
}
