package main

import "fmt"

//import "strings"

//import "io"
import "os"
import "log"
import "bufio"
import "strconv"
import "sort"

func main() {
	//fmt.Println("Hello World!")
	file, err := os.Open("b_in.txt")
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
	// PART B: most obvious solution would be to just sort and take top3 from sorted list.
	sort.Ints(elf_sums)
	elf_sums_length := len(elf_sums)

	total_cals := elf_sums[elf_sums_length-1] + elf_sums[elf_sums_length-2] + elf_sums[elf_sums_length-3]
	fmt.Println(total_cals)
}
