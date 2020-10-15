package main

import (
	"fmt"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}
