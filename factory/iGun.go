package main

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
