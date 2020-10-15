package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name + "_clone")
}

func (f *file) clone() inode {
	return &file{name: f.name}
}
