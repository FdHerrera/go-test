package main

import (
	"bufio"
	"bytes"
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

func Test_checkNumbers(t *testing.T) {
	checkNumberTests := []struct {
		input         string
		expectedMsg   string
		expectedQuits bool
	}{
		{"q", "", true},
		{"foo", "foo is not a number!", false},
		{"0", "0 is not prime by definition!", false},
		{"1", "1 is not prime by definition!", false},
		{"7", "7 is a prime number!", false},
		{"-1", "Negative numbers are not prime by definition!", false},
	}

	for _, e := range checkNumberTests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		actualResult, wantsToQuit := checkNumbers(reader)

		if actualResult != e.expectedMsg || e.expectedQuits != wantsToQuit {
			t.Error(actualResult)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var in bytes.Buffer

	in.Write([]byte("1\nq\n"))

	go readUserInput(&in, doneChan)

	<-doneChan
	close(doneChan)
}
