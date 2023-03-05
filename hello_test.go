package main

import "testing"
import "fmt"

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
	got := Sum([5]int{1, 2, 3, 4, 5})

	want := 15

	assertCorrectMessageInt(t, got, want)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if got != want {
		t.Errorf("got %v want %v", got, want)
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
