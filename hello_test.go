package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty strigs default to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Espanish", func(t *testing.T) {
		got := Hello("Maluma", "Spanish")
		want := "Hola, Maluma"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Frech", func(t *testing.T) {
		got := Hello("Silvie", "French")
		want := "Bonjour, Silvie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Portuguese", func(t *testing.T) {
		got := Hello("Mario", "Portuguese")
		want := "Olá, Mario"
		assertCorrectMessage(t, got, want)
	})
}
