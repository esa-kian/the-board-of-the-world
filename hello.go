package main

import "fmt"

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

func Sum(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	return nil
}

func main() {
	fmt.Println(nil)
	fmt.Println(Sum([5]int{4,3,6,5,9}))
	fmt.Println(Add(5, 7))
	fmt.Println(Hello("Ehsan and Zara"))
}
