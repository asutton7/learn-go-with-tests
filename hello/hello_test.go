package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	}

	t.Run("Saying hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, got, expected)
	})
}
