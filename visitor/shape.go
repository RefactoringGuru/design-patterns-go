package main

type shape interface {
	getType() string
	accept(visitor)
}
