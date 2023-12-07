package day

import (
	"flag"
	"log"
	"math"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func Four() (int, int) {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	winningNos, myNums := cleanData(data)

	var part1 int
	n := make(map[int]int)

	for i := range winningNos {
		n[i] += 1
		numMatches := intersection(winningNos[i], myNums[i])

		if numMatches > 0 {
			part1 += int(math.Pow(2, float64(numMatches-1)))
		}

		for j := 0; j < numMatches; j++ {
			n[i+j+1] += n[i]
		}
	}

	var part2 int
	for _, v := range n {
		part2 += v
	}
	return part1, part2
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

func intersection(line []string, check []string) int {
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
