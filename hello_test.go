package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrecMessage(t, got, want)
	})
	t.Run("warning when an empty string is suplied", func(t * testing.T) {
		got := Hello("", "")
		want := "you left the name field empty!"
		assertCorrecMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrecMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Brenda", "French")
		want := "Bonjour, Brenda"
		assertCorrecMessage(t, got, want)
	})
}

func assertCorrecMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}