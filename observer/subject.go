package main

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
