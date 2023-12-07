package day

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

func Five() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	seedRanges, conditions := dataFive(data)

	var seeds []int
	for i := 0; i < len(seedRanges); i += 2 {
		for j := seedRanges[i]; j < seedRanges[i]+seedRanges[i+1]; j++ {
			seeds = append(seeds, j)
		}
	}

	var locs []int
	for i := range seeds {
		for _, condition := range conditions {
			for j := 0; j < len(condition); j++ {
				v := condition[j]
				if seeds[i] >= v.src && seeds[i] < v.src+v.span {
					seeds[i] = v.des + seeds[i] - v.src
					break
				}
			}
		}
		locs = append(locs, seeds[i])
	}

	fmt.Println(findMin(locs))
}

type seed struct {
	des  int // destination range start
	src  int // source range start
	span int // range is reserved
}

func dataFive(data []string) (seeds []int, conditions [][]seed) {
	var soilMap, fertMap, waterMap, lightMap, tempMap, humidMap, locMap []seed
	for _, v := range strings.Split(data[0][7:], " ") {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("could not read seed value %s: %v", v, err)
		}
		seeds = append(seeds, val)
	}

	data = data[3:]
	for i := range data {
		if data[i] == "" {
			data = data[i+2:]
			break
		}
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		soilMap = append(soilMap, seed{d, s, r})
	}
	conditions = append(conditions, soilMap)

	for i := range data {
		if data[i] == "" {
			data = data[i+2:]
			break
		}
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		fertMap = append(fertMap, seed{d, s, r})
	}
	conditions = append(conditions, fertMap)

	for i := range data {
		if data[i] == "" {
			data = data[i+2:]
			break
		}
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		waterMap = append(waterMap, seed{d, s, r})
	}
	conditions = append(conditions, waterMap)

	for i := range data {
		if data[i] == "" {
			data = data[i+2:]
			break
		}
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		lightMap = append(lightMap, seed{d, s, r})
	}
	conditions = append(conditions, lightMap)

	for i := range data {
		if data[i] == "" {
			data = data[i+2:]
			break
		}
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		tempMap = append(tempMap, seed{d, s, r})
	}
	conditions = append(conditions, tempMap)

	for i := range data {
		if data[i] == "" {
			data = data[i+2:]
			break
		}
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		humidMap = append(humidMap, seed{d, s, r})
	}
	conditions = append(conditions, humidMap)

	for i := range data {
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		locMap = append(locMap, seed{d, s, r})
	}
	conditions = append(conditions, locMap)

	return
}

func findMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
