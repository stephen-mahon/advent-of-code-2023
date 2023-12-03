package day

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

type die struct {
	red   int
	green int
	blue  int
}

const maxRed int = 12
const maxGreen int = 13
const maxBlue int = 14

func Two() (int, int) {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var maxColors []die

	for i := range data {
		line := strings.Split(data[i], ": ")
		sets := strings.Split(line[1], "; ") // Split after "Game %: " on each semicolon
		var dice []die
		for j := range sets {
			values := strings.Split(sets[j], ", ") // Split on each comma
			dice = append(dice, colorDie(values))
		}

		maxColors = append(maxColors, maxDieColors(dice))
	}

	var part1 int
	var part2 int

	for i := range maxColors {
		if possible(maxColors[i]) {
			part1 += (i + 1)
		}
		part2 += maxColors[i].red * maxColors[i].green * maxColors[i].blue
	}

	return part1, part2

}

func possible(die die) bool {
	return die.red <= maxRed && die.green <= maxGreen && die.blue <= maxBlue
}

func colorDie(val []string) die {
	var die die
	for i := range val {
		line := strings.Split(val[i], " ")
		num, _ := strconv.Atoi(line[0])
		color := line[1]

		switch color {
		case "red":
			die.red = num
		case "green":
			die.green = num
		case "blue":
			die.blue = num
		}
	}
	return die
}

func maxDieColors(dice []die) die {
	var maxDie die
	for j := range dice {
		if dice[j].red > maxDie.red {
			maxDie.red = dice[j].red
		}
		if dice[j].green > maxDie.green {
			maxDie.green = dice[j].green
		}
		if dice[j].blue > maxDie.blue {
			maxDie.blue = dice[j].blue
		}
	}
	return maxDie
}
