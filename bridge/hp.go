package main

import "fmt"

type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
