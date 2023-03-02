package main

import "fmt"
import "./iteration"

const helloPrefix = "Hello, "

func Hello(name string) string {

	if name == "" {
		name = "World"
	}

	return helloPrefix + name + "!"
}

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Printls(iteration.Repeat())
	fmt.Println(Hello("Ehsan and Zara"))
}
