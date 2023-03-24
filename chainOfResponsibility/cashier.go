package main

import "fmt"

type Cashier struct {
	DepartmentBase
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient:", p.name)
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
