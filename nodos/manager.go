package nodos

type NeuralRed struct {
	Connections map[string]chan Event
}

func (r *NeuralRed) Emit(chanName string, ev Event) {
	if ch := r.Connections[chanName]; ch != nil {
		ch <- ev
	}
}

type Manager struct {
	Nodos     []Neuron
	neuralRed *NeuralRed
}

func (m *Manager) Start() {
	if m.neuralRed == nil {
		m.neuralRed = &NeuralRed{
			Connections: map[string]chan Event{},
		}
	}

	// First build each one of the nodos, so we can have all setup for
	// bind the communications channels (connections)
	for _, nodo := range m.Nodos {
		// Get the channel for each node.
		if _, ok := m.neuralRed.Connections[nodo.Name]; !ok {
			ch := nodo.Build()
			m.neuralRed.Connections[nodo.Name] = ch
		}
	}

	for _, n := range m.Nodos {

		go func(n Neuron) {
			n.Cell.Join(m.neuralRed)
		}(n)
	}
}
