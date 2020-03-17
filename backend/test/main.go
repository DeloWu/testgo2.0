package main

import "fmt"

type File struct {
	name string
	size int
	attr struct {
		perm  int
		owner int
	}
}

type User struct {
	name string
	age  int
}

func main() {
	f := File{
		name: "test.txt",
		//size: 1024,
	}
	fmt.Println(f)
	u1 := User{age: 1}
	fmt.Println(u1)
}
