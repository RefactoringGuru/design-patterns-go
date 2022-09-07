package main

type Server interface {
	handleRequest(string, string) (int, string)
}
