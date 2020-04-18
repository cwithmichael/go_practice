package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	secretNumber := rand.Intn(100 - 1) + 1
	fmt.Println("Guess the number!")
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input your guess.")
		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Println("Something went terribly wrong")
			os.Exit(1)
		}

		guess, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Println("That's not a number :( | Try again")
			continue
		}
		fmt.Println("You guessed: ", guess)

		switch {
		case guess < secretNumber:
			fmt.Println("Too small!")
		case guess > secretNumber:
			fmt.Println("Too big!")
		default:
			fmt.Println("You win!")
			os.Exit(0)
		}
	}
}
