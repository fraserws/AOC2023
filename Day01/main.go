package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := string(b)
	lines := strings.Split(convertWrittenToNumeric(inputStrings), "\n")

	extractedNumbers := make([][]string, len(lines))
	re := regexp.MustCompile(`\d+`)

	total := 0
	for i, line := range lines {
		extractedNumbers[i] = re.FindAllString(line, -1)
	}

	for _, numbers := range extractedNumbers {
		if len(numbers) > 0 {
			firstDigit := numbers[0][:1]
			lastDigit := numbers[len(numbers)-1][len(numbers[len(numbers)-1])-1:]
			concatenated := firstDigit + lastDigit
			num, err := strconv.Atoi(concatenated)
			if err != nil {
				panic(err)
			}
			total += num
		}
	}

	fmt.Println("Total:", total)
}

func convertWrittenToNumeric(text string) string {
	numberMap := map[string]string{
		"one":   "o1one",
		"two":   "t2two",
		"three": "t3three",
		"four":  "f4four",
		"five":  "f5five",
		"six":   "s6six",
		"seven": "s7seven",
		"eight": "e8eight",
		"nine":  "n9nine",
		"zero":  "z0zero",
	}

	for key, value := range numberMap {
		text = strings.Replace(text, key, value, -1)
	}

	return text
}
