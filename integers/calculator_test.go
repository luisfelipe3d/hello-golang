package integers

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	assertCorrectValue := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Addition function", func(t *testing.T) {
		got := Add(3, 4)
		expected := 7

		assertCorrectValue(t, got, expected)
	})

	t.Run("Subtraction function", func(t *testing.T) {
		got := Sub(10, 3)
		expected := 7

		assertCorrectValue(t, got, expected)
	})

	t.Run("Multiplication function", func(t *testing.T) {
		got := Multi(5, 10)
		expected := 50

		assertCorrectValue(t, got, expected)
	})

	t.Run("Division function", func(t *testing.T) {
		got := Div(10, 5)
		expected := 2

		assertCorrectValue(t, got, expected)
	})

}
