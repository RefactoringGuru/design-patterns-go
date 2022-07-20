package main

type Iterator interface {
	hasNext() bool
	getNext() *User
}
