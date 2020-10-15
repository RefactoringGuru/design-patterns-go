package main

import "fmt"

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
