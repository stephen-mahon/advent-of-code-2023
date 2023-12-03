package day

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func One() (ans1, ans2 int) {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	dat, err := read.Data(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	return part1(dat), part2(dat)
}

func parseNum(s string) (i int, err error) {
	var vals []string
	for i := range s {
		_, err := strconv.Atoi(string(s[i]))
		if err == nil {
			vals = append(vals, string(s[i]))
		}
	}

	return firstAndLast(vals)
}

func part1(input []string) int {
	var total int

	for i := range input {
		val, err := parseNum(input[i])
		if err != nil {
			log.Fatalf("%v\ncould not data line %d: %s", err, i, input[i])
		}
		total += val
	}

	return total

}

func firstAndLast(arr []string) (int, error) {
	return strconv.Atoi(arr[0] + arr[len(arr)-1])
}

func part2(input []string) int {
	total := 0

	for i := range input {
		arr := findDigits(input[i])
		val, _ := firstAndLast(arr)
		total += val
	}

	return total
}

func findDigits(line string) []string {
	digits := map[string]string{
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

	var arr []string

	for i := 0; i < len(line); i++ {
		digit := string(line[i])

	out:
		for j := i; j < len(line); j++ {
			_, err := strconv.Atoi(string(line[j]))
			if err == nil {
				arr = append(arr, string(line[j]))
				break out
			}
			digit += string(line[j])
			for k, v := range digits {
				if strings.Contains(digit, k) {
					arr = append(arr, v)
					i = j - 1
					break out
				}
			}
		}
	}

	return arr
}
