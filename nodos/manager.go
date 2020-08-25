package nodos

type NeuralRed = map[string]chan Event

type Manager struct {
	Nodos       []NodoBuilder
	connections NeuralRed
}

func (m *Manager) Start() {
	if m.connections == nil {
		m.connections = NeuralRed{}
	}

	// First build each one of the nodos, so we can have all setup for
	// bind the communications channels (connections)
	for _, nodo := range m.Nodos {
		// Get the channel for each node.
		if _, ok := m.connections[nodo.Name]; !ok {
			ch := nodo.Build()
			m.connections[nodo.Name] = ch
		}
	}

	//go func(n Nodo) {
	//
	//	n.SetNotificationHandler(func(ev Event) {
	//		log.Printf("SENDING :: %s", ev)
	//		net <- ev
	//	})
	//	n.ListenEvents(net)
	//}(nodo)
}
