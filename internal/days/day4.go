package day

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func Four() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	wins, nums := cleanData(data)

	//fmt.Println(elfsScratchTotal(winning, numbers))

	for i := range wins {
		matches := findNumberMatches(wins[i], nums[i])
		fmt.Printf("%v | %v\n", i, matches)
		for j := i + 1; j < i+matches; j++ {
			fmt.Printf("%v | %v, ", j, findNumberMatches(wins[j], nums[j]))

		}
	}
}

func recursiveFindCards(index int, lines [][]string, checks [][]string, tally int, n int) int {
	return -1
}

func cleanData(dat []string) (winArr [][]string, numArr [][]string) {

	for i := range dat {
		line := strings.Split(dat[i], ": ")
		lineSplit := strings.Split(line[1], " | ")

		left := strings.Split(strings.TrimLeft(lineSplit[0], " "), " ")
		var wins []string
		for j := range left {
			if left[j] != "" {
				wins = append(wins, left[j])
			}
		}
		winArr = append(winArr, wins)

		var nums []string
		right := strings.Split(strings.TrimLeft(lineSplit[1], " "), " ")
		for j := range right {
			if right[j] != "" {
				nums = append(nums, right[j])
			}
		}
		numArr = append(numArr, nums)

	}
	return winArr, numArr
}

func elfsScratchTotal(winning [][]string, numbers [][]string) int {
	var total int
	for k := range winning {
		points := 0
		for j := range winning[k] {
			for i := range numbers[k] {
				if winning[k][j] == numbers[k][i] {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
				}
			}
		}
		total += points
	}
	return total
}

func findNumberMatches(line []string, check []string) int {
	var match int
	for i := range line {
		for j := range check {
			if line[i] == check[j] {
				match++
			}
		}
	}
	return match
}
