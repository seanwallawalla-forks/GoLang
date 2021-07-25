/*
Author: Brad Myrick
Date: 2021-07-25
Name: numberGuess

Description: A number guessing game.
generate a random number between 1 and 10
ask the user to guess the number, then convert the input to an integer.
if guess is equal to number, winner
if guess is not equal to number tell them if they are too high or too low.
after the game ask the user if they want to play again.
if the user enters "y", generate a new number and start over and increase the difficulty.
if the user answers "n", exit.

To run: go run numberGuess.go from the command line in the directory containing numberGuess.go
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func guess(num int) {
	for {
		fmt.Printf("can you guess it? -> ")
		text, err := readString()
		if err != nil {
			fmt.Println(err)
			continue
		}
		guess, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}
		if guess == num {
			fmt.Println("You got it!")
			break
		} else if guess < num {
			fmt.Println("Too low.")
		} else {
			fmt.Println("Too high.")
		}
	}
}

func readString() (string, error) {
	var text string
	_, err := fmt.Scanf("%s\n", &text)
	return text, err
}

func run(l int) {

	fmt.Println("I am thinking of a number between 1 and", l)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(l) + 1
	guess(num)

}

func main() {
	limit := 10
	for true {
		run(limit)
		fmt.Println("Do you want to play again? (y/n)")
		text, err := readString()
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch text {
		case "y":
			limit *= 2
			continue
		case "n":
			fmt.Println("Bye!")
			break
		default:
			fmt.Println("Please enter y or n.")
			continue
		}
	}
}
