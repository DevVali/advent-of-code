package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

// FIRST PART

func part1() int {
	readFile, err := os.Open("./01_input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

	calibrations := 0

    for fileScanner.Scan() {
        re := regexp.MustCompile("[0-9]+")
		filteredDigits := re.FindAllString(fileScanner.Text(), -1)

		firstDigit := filteredDigits[0][0:1]
		lastDigit := filteredDigits[len(filteredDigits)-1:][0]
		lastDigit = lastDigit[len(lastDigit)-1:]

		combination, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println(err)
		}

		calibrations += combination
	}

    readFile.Close()

	return calibrations
}

// SECOND PART

func part2() int {
	readFile, err := os.Open("./01_input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

	calibrations := 0

    for fileScanner.Scan() {
		values := fileScanner.Text();
		
		replace := map[string]string{
			"oneight": "18",
			"twone": "21",
			"eightwo": "82",
		}
		
		for s, r := range replace {
			values = strings.Replace(values, s, r, -1)
		}

		replace = map[string]string{
			"one": "1",
			"two": "2",
			"three": "3",
			"four": "4",
			"five": "5",
			"six": "6",
			"seven": "7",
			"eight": "8",
			"nine": "9",
		}
		
		for s, r := range replace {
			values = strings.Replace(values, s, r, -1)
		}

        re := regexp.MustCompile("[0-9]+")
		filteredDigits := re.FindAllString(values, -1)

		firstDigit := filteredDigits[0][0:1]
		lastDigit := filteredDigits[len(filteredDigits)-1:][0]
		lastDigit = lastDigit[len(lastDigit)-1:]

		combination, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println(err)
		}

		calibrations += combination
	}

    readFile.Close()

	return calibrations
}

// OUTPUT

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
