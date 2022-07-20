package main

type Train interface {
	arrive()
	depart()
	permitArrival()
}
