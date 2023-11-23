package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name        string
		testNum     int
		expected    bool
		expectedMsg string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime by definition!"},
		{"one", 1, false, "1 is not prime by definition!"},
		{"negative number", -1, false, "Negative numbers are not prime by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected != result {
			t.Errorf("%s: expected: [%t], but got: [%t]", e.name, e.expected, result)
		}

		if e.expectedMsg != msg {
			t.Errorf("%s: expected [%s], but got [%s]", e.name, e.expectedMsg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()
	w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("Incorrect promt, got [%s]", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()
	w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number, and we'll tell you if it is prime or not. Enter q to quit.") {
		t.Error("Wrong message displayed, got: []", string(out))
	}
}
