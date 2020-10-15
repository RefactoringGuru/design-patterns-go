package main

type server interface {
	handleRequest(string, string) (int, string)
}
