package main

import "fmt"

type Medical struct {
	DepartmentBase
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient:", p.name)
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient:", p.name)
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
