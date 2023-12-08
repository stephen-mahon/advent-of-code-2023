package day

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func Seven() (int, int) {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	cardsBids := make(map[string]int)
	for i := range data {
		line := strings.Split(data[i], " ")
		bid, _ := strconv.Atoi(line[1])
		cardsBids[line[0]] = bid
	}

	return DaySevenPartOne(cardsBids), DaySevenPartTwo(cardsBids)

}
