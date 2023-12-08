package day

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

type Node struct {
	L string
	R string
}

func Eight() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	input := data[0]
	path := make(map[string]Node)

	data = data[2:]
	var startingA []string
	for i := range data {
		dat := strings.Split(data[i], " = ")
		vals := strings.Split(dat[1], ", ")
		if string(dat[0][2]) == "A" {
			startingA = append(startingA, dat[0])
		}
		path[dat[0]] = Node{vals[0][1:], vals[1][:len(vals[1])-1]}
	}

	/*
		next := "AAA"
		part1 := 0
		for next != "ZZZ" {
			direction := string(input[part1%len(input)])
			next = choosePath(direction, path[next])
			part1 += 1
		}

		fmt.Println(part1)
	*/

	part2 := 0
	allZs := false
	for !allZs {
		allZs = true
		direction := string(input[part2%len(input)])
		for i, current := range startingA {
			next := choosePath(direction, path[current])
			startingA[i] = next
			if string(next[2]) != "Z" {
				allZs = false
			}
		}
		part2 += 1
	}
	fmt.Println(part2)

}

func choosePath(direction string, node Node) string {
	if direction == "R" {
		return node.R
	}
	return node.L
}
