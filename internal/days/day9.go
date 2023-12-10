package day

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func Nine() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var vals [][]int
	for i := range data {
		line := strings.Split(data[i], " ")
		var lineArr []int
		for j := range line {
			val, _ := strconv.Atoi(string(line[j]))
			lineArr = append(lineArr, val)
		}
		vals = append(vals, lineArr)
	}

	var part1 int
	var part2 int

	for _, current := range vals {
		var nextElem int
		var next []int
		allZeros := false
		for !allZeros {
			next, allZeros = seqOfDiff(current, false)
			nextElem += lastElem(current)
			current = next
		}
		part1 += nextElem

	}

	for _, current := range vals {
		var firstPairs [][]int
		var next []int
		allZeros := false

		fmt.Println(current[0], current)
		for !allZeros {
			next, allZeros = seqOfDiff(current, true)
			elems := []int{current[0], next[0]}
			current = next
			fmt.Println(current[0], current, allZeros)
			firstPairs = append(firstPairs, elems)

		}
		fmt.Println(firstPairs)
		fmt.Println()

	}

	fmt.Println(part1, part2)
}

func seqOfDiff(arr []int, reverse bool) ([]int, bool) {
	if len(arr) == 1 {
		return []int{}, false
	}

	nextLine := make([]int, len(arr)-1)
	allZeros := true

	if reverse {
		for i := len(arr) - 1; i > 0; i-- {
			nextLine[i-1] = arr[i] - arr[i-1]
			if i != len(arr)-1 {
				allZeros = nextLine[i-1] == 0 && nextLine[i] == 0
			}
		}

		return nextLine, allZeros
	}

	for i := 1; i < len(arr); i++ {
		nextLine[i-1] = arr[i] - arr[i-1]
		if arr[i]-arr[i-1] == 0 {
			allZeros = true
		} else {
			allZeros = false
		}
	}

	return nextLine, allZeros
}

func lastElem(arr []int) int {
	return arr[len(arr)-1]
}
