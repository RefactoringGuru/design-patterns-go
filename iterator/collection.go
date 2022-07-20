package main

type Collection interface {
	createIterator() Iterator
}
