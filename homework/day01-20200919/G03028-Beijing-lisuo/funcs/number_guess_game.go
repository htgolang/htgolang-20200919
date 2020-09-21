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
	criteria := rand.Intn(100)
	// fmt.Printf("The target is: %v\n", criteria)
	for i := 1; i <= threshold; i++ {
		fmt.Print("input the guess[1-100]: ")
		_, err := fmt.Scan(&guess)
		// make sure the input is legal
		if err != nil {
			fmt.Println("You are not input a number.")
			break
		} else if guess < 1 || guess > 100 {
			fmt.Println("Input the number between 1 and 100.")
			// give one more chance
			threshold += 1
			continue
		}
		// guess
		if guess > criteria {
			fmt.Println("Exceeded, guess again.\n")
		} else if guess < criteria {
			fmt.Println("Below the criteria, guess again.\n")
		} else {
			fmt.Println("Bingo!!\n")
			break
		}
		if i == threshold {
			fmt.Println("Oh no, all your luck have been taken!")
			break
		}
	}

}
