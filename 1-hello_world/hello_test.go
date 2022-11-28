package main

import "testing"

// t of type *testing.T
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectmessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectmessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectmessage(t, got, want)
	})
}

// Accepting testin.TB so both helper functions form a test and benchmark can be called
func assertCorrectmessage(t testing.TB, got, want string) {
	// Needed to tell the test suite that this method is a helper
	t.Helper()
	if got != want {
		// f = format, allows for building of string with values inserted into the placeholder values %q
		t.Errorf("got %q want %q", got, want)
	}
}
