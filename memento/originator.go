package main

type Originator struct {
	state string
}

func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}
