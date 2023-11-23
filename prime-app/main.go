package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("Goodbye.")
}

func readUserInput(r io.Reader, c chan bool) {
	scanner := bufio.NewScanner(r)

	for {
		res, done := checkNumbers(scanner)

		if done {
			c <- true
			return
		}

		fmt.Println(res)
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	prompt()
	// read user input
	scanner.Scan()

	// check if user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return fmt.Sprintf("%s is not a number!", scanner.Text()), false
	}

	_, msg := isPrime(input)

	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are exempt
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition!", n)
	}
	// negative numbers are not prime
	if n < 0 {
		return false, fmt.Sprint("Negative numbers are not prime by definition!")
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", n)
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is prime or not. Enter q to quit.")
}

func prompt() {
	fmt.Print("-> ")
}
