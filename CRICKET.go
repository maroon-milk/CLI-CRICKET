package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var name string
	var empty string
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("| WELCOME TO CRICKET CLI! |")
	fmt.Println("---------------------------\n")

	fmt.Println("\nEnter your name, batsman!")
	fmt.Scanln(&name)

	possibleRuns := []int{1, 2, 3, 4, 5, 6}
	totalRun := 0 //Slice of int used to store the runs
	fmt.Printf("\n\nNice to have you on the team, %v! You're our operning player.\n", name)
	fmt.Println("RULE: You will be playing until you are declared 'Out' by the bowler.\nAre you ready?")
	fmt.Scanln(&empty)
	fmt.Println("\n----------------------------------------------")

	for {
		fmt.Println("\nChoose batting style\n- HS for Helicopter shot\n- SW for Side-Sweep\n- UC for Upper-Cut\n")
		var battingStyle string
		fmt.Scanln(&battingStyle)

		battingStyles := [3]string{"HS", "SW", "UC"}

		if battingStyle == battingStyles[0] {
			fmt.Printf("You Chose Helicopter shot!")
		} else if battingStyle == battingStyles[1] {
			fmt.Printf("You Chose Side-Sweep!")
		} else if battingStyle == battingStyles[2] {
			fmt.Printf("You Chose Upper-Cut!")
		} else {
			fmt.Printf("Invalid choice - try again")
			continue
		}

		fmt.Printf("\nLasith Malinga bowls a viscous ball and %v hits the ball hardly with his bat!\n", name)
		fmt.Printf("\n----------------------------------------------\n")

		runsByBatsman := rand.Intn(len(possibleRuns))
		totalRun += runsByBatsman

		switch runsByBatsman {
		case 1:
			fmt.Println("Run by Runs\n1 RUN\nTotal runs: ", totalRun)
		case 2:
			fmt.Println("Run by Runs\n2 RUN\nTotal runs: ", totalRun)
		case 3:
			fmt.Println("Run by Runs\n3 RUN\nTotal runs: ", totalRun)
		case 4:
			fmt.Printf("A sweeping shot through the boundary!\n4 RUN\nTotal runs: ", totalRun)
		case 5:
			fmt.Println("Run by Runs\n5 RUN\nTotal runs: ", totalRun)
		case 6:
			fmt.Printf("%v sends the ball flying out of the stadium! What a shot!\n6 RUN\nTotal runs: ", name, totalRun)
		default:
			fmt.Printf("OH! What a tragedy! %v has been caught out... It's a game over for him!\n", name)
			goto RESULT
		}
	}

RESULT:
	fmt.Printf("Total runs made by %v were: %d\nBetter luck next time!\n", name, totalRun)
	// fmt.Printf("Total balls played were: %d\n", name, totalRun)

}
