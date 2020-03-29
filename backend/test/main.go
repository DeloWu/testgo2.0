package main

import "fmt"

func main() {
	m := make(map[string]int)
	m2 := map[string]int{"two": 2, "three": 3}
	m["one"] = 1
	fmt.Println(m)
	fmt.Println(m2)
}
