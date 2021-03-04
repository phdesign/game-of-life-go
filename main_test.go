package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should return product of numbers that sum 2020", func(t *testing.T) {
		want := "Hi, Gladys. Welcome!"
		got := Hello("Gladys")

		if got != want {
			t.Errorf("Want %q, got %q", want, got)
		}
	})
}
