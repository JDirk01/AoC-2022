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
		//fmt.Println(split_string[i])
		//check if draw
		//draw if Rock, draw = 3pts, rock = 1
		if split_string[i] == "A" && split_string[i+1] == "X" {
			final_score += 4
			continue
		}
		//draw if Paper, draw = 3pts, paper = 2
		if split_string[i] == "B" && split_string[i+1] == "Y" {
			final_score += 5
			continue
		}
		//draw if Scissors, draw = 3pts, scissors = 3
		if split_string[i] == "C" && split_string[i+1] == "Z" {
			final_score += 6
			continue
		}

		//Check if win
		//Win if opponent plays rock and user plays paper, win = 6 pts, paper = 2 pts
		if split_string[i] == "A" && split_string[i+1] == "Y" {
			final_score += 8
			continue
		}

		//Win if opponent plays paper and user plays scissors, win = 6 pts, scissors = 3 pts
		if split_string[i] == "B" && split_string[i+1] == "Z" {
			final_score += 9
			continue
		}

		//Win if opponent plays scissors and user plays rock, win = 6 pts, rock = 1 pts
		if split_string[i] == "C" && split_string[i+1] == "X" {
			final_score += 7
			continue
		}

		// Losses
		if split_string[i+1] == "X" {
			final_score += 1
			continue
		}

		if split_string[i+1] == "Y" {
			final_score += 2
			continue
		}

		if split_string[i+1] == "Z" {
			final_score += 3
			continue
		}

	}

	fmt.Println(final_score)

}
