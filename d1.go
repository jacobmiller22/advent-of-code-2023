package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func FindFirstDigit(s string) string {
	runes := []rune(s)
	for _, r := range runes {
		if unicode.IsDigit(r) {
			return string(r)
		}
	}

	log.Fatal("Didn't find a digit")
	return ""
}

func FindLastDigit(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return string(runes[i])
		}
	}

	log.Fatal("Didn't find a digit")
	return ""
}

func main() {
	fp, err := os.Open("input-d1.txt")
	if err != nil {
		log.Fatal("Failed to read input file")
	}

	scanner := bufio.NewScanner(fp)

	if err != nil {
		log.Fatal("Failed to read input file")
	}

	sum := 0

	numbers := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("Before replace %s\n", line)
		for k, v := range numbers {
			line = strings.ReplaceAll(line, k, v)
		}
		log.Printf("After replace %s\n", line)
		char1 := FindFirstDigit(line)
		charLast := FindLastDigit(line)
		combination := fmt.Sprintf("%s%s", char1, charLast)
		calibration_num, err := strconv.Atoi(combination)
		if err != nil {
			log.Fatal("Failed to convert digits to numbers")
		}
		fmt.Printf("%d\n", calibration_num)

		sum += calibration_num
	}

	log.Printf("The sum of the calibration values is: %d", sum)
}
