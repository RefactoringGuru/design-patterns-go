package main

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}
