package nodos

type Connectable interface {
	Join(net NeuralRed)

	ListenEvents(net NeuralRed)
}

//NodoBuilder build the nodo
type Neuron struct {
	Name string
	Cell  Connectable
	// oun channel to pull out
	outCh chan Event
}

func (n *Neuron) BroadcastToRedev(ev Event) {
	n.outCh <- ev
}

func (n *Neuron) Build() chan Event {
	if n.outCh == nil {
		n.outCh = make(chan Event)
	}

	return n.outCh

}
