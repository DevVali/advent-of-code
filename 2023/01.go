package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// FIRST PART

func part1(input *bufio.Scanner) int {
	calibrations := 0

	for input.Scan() {
		filteredDigits := regexp.MustCompile("[0-9]+").FindAllString(input.Text(), -1)

		firstDigit := filteredDigits[0][0:1]
		lastDigit := filteredDigits[len(filteredDigits)-1:][0]
		lastDigit = lastDigit[len(lastDigit)-1:]

		combination, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println(err)
		}

		calibrations += combination
	}

	return calibrations
}

func part2(input *bufio.Scanner) int {
	calibrations := 0

	for input.Scan() {
		values := input.Text()

		overlapedNumbers := map[string]string{
			"oneight": "18",
			"twone":   "21",
			"eightwo": "82",
		}
		numbers := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}

		for s, r := range overlapedNumbers {
			values = strings.Replace(values, s, r, -1)
		}
		for s, r := range numbers {
			values = strings.Replace(values, s, r, -1)
		}

		filteredDigits := regexp.MustCompile("[0-9]+").FindAllString(values, -1)

		firstDigit := filteredDigits[0][0:1]
		lastDigit := filteredDigits[len(filteredDigits)-1:][0]
		lastDigit = lastDigit[len(lastDigit)-1:]

		combination, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println(err)
		}

		calibrations += combination
	}

	return calibrations
}

// OUTPUT

func main() {
	// PREPARE FIRST PART
	readFile, err := os.Open("./01_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := bufio.NewScanner(readFile)
	input.Split(bufio.ScanLines)

	// OUTPUT FIRST PART
	fmt.Println(part1(input))

	// PREPARE SECOND PART
	readFile, err = os.Open("./01_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input = bufio.NewScanner(readFile)
	input.Split(bufio.ScanLines)

	// OUTPUT SECOND PART
	fmt.Println(part2(input))

	// CLOSE FILE
	readFile.Close()
}
