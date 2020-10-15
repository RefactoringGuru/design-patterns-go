package main

import "fmt"

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
