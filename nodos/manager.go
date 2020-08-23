package nodos

import "log"

type Manager struct {
	Nodos []Nodo
}

func (m *Manager) Start() {

	net := make(chan Event)

	for _, nodo := range m.Nodos {
		go func(n Nodo) {

			n.SetNotificationHandler(func(ev Event) {
				log.Printf("SENDING :: %s", ev)
				net <- ev
			})
			n.ListenEvents(net)
		}(nodo)
	}
}
