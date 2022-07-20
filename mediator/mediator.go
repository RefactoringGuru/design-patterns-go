package main

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
