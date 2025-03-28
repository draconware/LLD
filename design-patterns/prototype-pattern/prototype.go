package main

type INode interface {
	print(string)
	clone() INode
}
