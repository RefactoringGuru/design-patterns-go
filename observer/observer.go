package main

type observer interface {
	update(string)
	getID() string
}
