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

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
