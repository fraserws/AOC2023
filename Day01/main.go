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
	lines := strings.Split(inputStrings, "\n")
	extractedNumbers := make([][]string, len(lines))
	re := regexp.MustCompile(`\d+`)

	total := 0
	for i, line := range lines {
		extractedNumbers[i] = re.FindAllString(line, -1)
	}

	for _, numbers := range extractedNumbers {
		if len(numbers) > 0 {
			concatenated := string(numbers[0][0]) + string(numbers[len(numbers)-1][len(numbers[len(numbers)-1])-1])
			num, err := strconv.Atoi(concatenated)
			if err != nil {
				panic(err)
			}
			total += num
		}
	}

	fmt.Println("Total:", total)
}
