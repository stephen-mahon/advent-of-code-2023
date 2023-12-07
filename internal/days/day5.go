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

	seeds, soil, fert, water, light, temp, humid, loc := dataFive(data)

	// soil, fert, water, light, temp, humid, loc

	soilMap := createMap(soil)
	fertMap := createMap(fert)
	waterMap := createMap(water)
	lightMap := createMap(light)
	tempMap := createMap(temp)
	humidMap := createMap(humid)
	locMap := createMap(loc)

	var locs []int
	for _, seed := range seeds {
		_, ok := soilMap[seed]
		if ok {
			seed = soilMap[seed]
		}

		_, ok = fertMap[seed]
		if ok {
			seed = fertMap[seed]
		}

		_, ok = waterMap[seed]
		if ok {
			seed = waterMap[seed]
		}

		_, ok = lightMap[seed]
		if ok {
			seed = lightMap[seed]
		}

		_, ok = tempMap[seed]
		if ok {
			seed = tempMap[seed]
		}

		_, ok = humidMap[seed]
		if ok {
			seed = humidMap[seed]
		}

		_, ok = locMap[seed]
		if ok {
			seed = locMap[seed]
		}

		locs = append(locs, seed)
	}

	fmt.Println(findMin(locs))

}

func createMap(vs []seed) map[int]int {
	mapArr := make(map[int]int)
	for _, v := range vs {
		j := v.des
		for i := v.src; i < v.src+v.span; i++ {
			mapArr[i] = j
			j++
		}
	}
	return mapArr
}

type seed struct {
	des  int // destination range start
	src  int // source range start
	span int // range is reserved
}

func dataFive(data []string) (seeds []int, soilMap, fertMap, waterMap, lightMap, tempMap, humidMap, locMap []seed) {
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

	for i := range data {
		vals := strings.Split(data[i], " ")
		d, _ := strconv.Atoi(vals[0])
		s, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])

		locMap = append(locMap, seed{d, s, r})
	}

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
