package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for x := 0; x < 10; x++ {
		fmt.Println("Number of guesses left:", 10-x)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputTrimmed := strings.TrimSpace(input)
		guess, err := strconv.Atoi(inputTrimmed)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Your guess was right!")
			success = true
			break
		}
		fmt.Println("Guess:", guess)
	}

	if !success {
		fmt.Println("You haven't guessed it within the number of allowed guesses.")
	}
	fmt.Println("The game is now finished. Press Enter to close the application.")
	reader.ReadString('\n')
}
