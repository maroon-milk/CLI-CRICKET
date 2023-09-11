package main

import (
	"fmt"
	"math/rand"
)

// try to utilise all that you have learnt so far

func main() {
	var name string

	fmt.Printf("SWITCHCASES - A Game Of CRICKET\n")
	fmt.Println("Enter your name, batsman!")
	fmt.Scanln(&name)

	possibleRuns := []int{2, 4, 6}                //Slice of int used to store the runs
	runsByBatsman := rand.Intn(len(possibleRuns)) //Should be 6 but since last digit ignored in range, 7 is used

	fmt.Println("\nLasith Malinga bowls a viscous ball and...\n----------------------------------------------")
	fmt.Printf("%v hits the ball hardly with his bat!\n", name)

	// for {
	switch runsByBatsman {
	case 2:
		fmt.Println("The 2 batmen on the ground are running b/w the wickets and have just made\n--2 RUNS--\n")
	case 4:
		fmt.Printf("A short-shot off the field and %v effortlessly scores a sweeping 4\n--RUNS--!\n", name)
	case 6:
		fmt.Printf("A long shot sends the ball flying out of the stadium! What a shot by %v as he scores\n--6 RUNS--!", name)
	default:
		fmt.Println("OH! What a tragedy! %v has been caught out... It's a game over for him!")
		break
	}
	// }
}

// add trigger such as play lower cut/uppercut/sideswee/etc
