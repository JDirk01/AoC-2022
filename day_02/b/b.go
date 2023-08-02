package main

//import "strings"

//import "io"
import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hello World!")
	file, err := os.ReadFile("../in.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	raw_input := string(file)
	split_string := strings.Fields(raw_input)

	fmt.Println(len(split_string))

	var final_score int = 0

	for i := 0; i < len(split_string); i += 2 {
		// Part two, check if second char is X (Round must be a loss), Y (Round must be a draw), Z (Round must be a Win)
		// Functions to check loss conditions, draw conditions, win conditions
		if split_string[i+1] == "X" {
			final_score += checkX(split_string[i])
		} else if split_string[i+1] == "Y" {
			final_score += checkY(split_string[i])
		} else if split_string[i+1] == "Z" {
			final_score += checkZ(split_string[i])
		}
	}

	fmt.Println(final_score)

}

func checkX(opponent string) int {
	var score int = 0
	// Opponent plays A (Rock), to lose (+0), player must play (Scissors)(+3)
	// Opponent plays B (Paper), to lose (+0), player must play (Rock) (+1)
	// Opponent plays C (Scissors), to lose (+0), player must play (paper) (+2)
	if opponent == "A" {
		score = 3

	} else if opponent == "B" {
		score = 1
	} else if opponent == "C" {
		score = 2
	}

	return score
}

func checkY(opponent string) int {
	var score int = 3 // Draw is 3 pts

	//opponent plays rock, to draw, user must play rock. draw = 3 pts, rock = 1
	if opponent == "A" {
		score += 1
	} else if opponent == "B" {
		score += 2
	} else if opponent == "C" {
		score += 3
	}

	return score
}

func checkZ(opponent string) int {
	var score int = 6 // Win is 6 pts

	if opponent == "A" {
		score += 2 //Opponent plays rock, user plays paper to win
	} else if opponent == "B" {
		score += 3 // Opponent plays paper, user plays scissors to win
	} else if opponent == "C" {
		score += 1
	}

	return score
}
