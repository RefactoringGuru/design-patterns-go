package main

type collection interface {
	createIterator() iterator
}
