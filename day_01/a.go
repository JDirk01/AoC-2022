package main

import "fmt"

//import "strings"

//import "io"
import "os"
import "log"
import "bufio"
import "strconv"

func main() {
	//fmt.Println("Hello World!")
	file, err := os.Open("a_in.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var split_string []string

	for fileScanner.Scan() {
		split_string = append(split_string, fileScanner.Text())
	}

	file.Close()

	var curr_sum int = 0
	var elf_sums []int
	for _, line := range split_string {
		if line == "" {
			elf_sums = append(elf_sums, curr_sum)
			curr_sum = 0
			continue
		}
		curr_int, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("unable to convert string to int.")
		}
		curr_sum += curr_int
	}
	// calculate max
	var max_elf int = 0
	for _, elf := range elf_sums {
		if elf > max_elf {
			max_elf = elf
		}
	}

	fmt.Println(max_elf)

	//fmt.Println(split_string)
}
