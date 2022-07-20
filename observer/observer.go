package main

type Observer interface {
	update(string)
	getID() string
}
