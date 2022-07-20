package main

type Department interface {
	execute(*Patient)
	setNext(Department)
}
