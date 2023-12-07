package day

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func Six() (int, int) {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	line1 := strings.Split(data[0], " ")
	times := getArr(line1[1:])

	line2 := strings.Split(data[1], " ")
	distances := getArr(line2[1:])

	part1 := 1

	for i := range times {
		var val int
		for j := 0; j <= times[i]; j++ {
			if j*(times[i]-j) > distances[i] {
				val += 1
			}
		}
		part1 *= val
	}

	time, distance := joinInts(times), joinInts(distances)

	var part2 int
	for j := 0; j <= time; j++ {
		if j*(time-j) > distance {
			part2 += 1
		}
	}

	return part1, part2

}

func getArr(line []string) []int {
	var arr []int

	for i := range line {
		if line[i] != "" {
			t, _ := strconv.Atoi(line[i])
			arr = append(arr, t)
		}
	}
	return arr
}

func joinInts(arr []int) int {
	var tmp string
	for i := range arr {
		tmp += string(fmt.Sprint(arr[i]))
	}

	val, _ := strconv.Atoi(tmp)
	return val
}
