package main

import "testing"
import "fmt"
import "reflect"


func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "it is just a test"}
	
	got := dictionary.Search("test")
	want := "it is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestWallet(t *testing.T){
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("Address of balance in test is %v \n", &wallet.balance)
	want := Bitcoin(10)
	
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ehsan")
		want := "Hello, Ehsan!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello when an empty string is supplied", func(t *testing.T){
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})


}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	
	fmt.Println("Answer is: ", sum)
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4, 5})

	want := 15

	assertCorrectMessageInt(t, got, want)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	//want := "Boob"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
	

		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)

		}
	}

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessageInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
