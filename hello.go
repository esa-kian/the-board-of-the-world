package main

import "fmt"
import "math"

const helloPrefix = "Hello, "


type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Hello(name string) string {

	if name == "" {
		name = "World"
	}

	return helloPrefix + name + "!"
}

func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}


func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius 
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func main() {
	//rectangle := Rectangle{12.0, 6.0}

//	fmt.Println(Area(rectangle))
//	fmt.Println(Perimeter(rectangle))
	fmt.Println(Sum([]int{4,3,6,5,9}))
	fmt.Println(Add(5, 7))
	fmt.Println(Hello("Ehsan and Zara"))
}
