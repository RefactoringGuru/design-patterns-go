package main

type inode interface {
	print(string)
	clone() inode
}
