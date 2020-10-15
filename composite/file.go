package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}
