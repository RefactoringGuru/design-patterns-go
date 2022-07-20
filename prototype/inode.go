package main

type Inode interface {
	print(string)
	clone() Inode
}
