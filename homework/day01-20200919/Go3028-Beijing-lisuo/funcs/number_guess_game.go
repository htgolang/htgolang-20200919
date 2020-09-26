package funcs

/*
a short code for a guess game
author: lisuo
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func NumGuess() {

	var guess int
	var threshold int = 5

	// using UnixNano cause the random int more random
	rand.Seed(time.Now().UnixNano())
	criteria := rand.Intn(100) + 1

	// fmt.Printf("The target is: %v\n", criteria)
	for i := 1; i <= threshold; i++ {
		fmt.Print("input the guess[1-100]: ")
		_, err := fmt.Scan(&guess)

		// make sure the input is legal
		if err != nil {
			fmt.Println("You are not input a number.")
			// give one more chance
			threshold += 1
			continue
		} else if guess < 1 || guess > 100 {
			fmt.Println("Input the number between [1,100].")
			// give one more chance
			threshold += 1
			continue
		}

		// guess
		if guess > criteria {
			fmt.Printf("Exceeded, guess again. %v chance(s) left.\n", threshold-i)
		} else if guess < criteria {
			fmt.Printf("Below the criteria, guess again. %v chance(s) left.\n", threshold-i)
		} else {
			fmt.Println("Bingo!!\n")
			fmt.Printf("The target is: %v\n", criteria)
			break
		}

		if i == threshold {
			fmt.Println("Oh no, all your luck have been taken!")
			fmt.Printf("The target is: %v\n", criteria)
			break
		}

	}

}
