package main

import "fmt"
import "math"
import "errors"
import "bytes"

const helloPrefix = "Hello, "

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("Address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}


func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

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
