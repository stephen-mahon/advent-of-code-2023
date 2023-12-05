package day

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/stephen-mahon/advent-of-code-2023/internal/read"
)

type pos struct {
	i int
	j int
}

func Three() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	data, err := read.Lines(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	posArr := findPartNumbers(data)
	posMap := findGears(posArr, data)

	var part1 int

	for k, v := range posMap {
		val, _ := strconv.Atoi(k)
		part1 += val
		fmt.Println(k, v, string(data[v.i][v.j]))
	}
	fmt.Println(part1)
}

func findPartNumbers(arr []string) map[string]pos {
	posArr := make(map[string]pos)
	for i := 0; i < len(arr); i++ {
		var partNum string
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] > 47 && arr[i][j] < 58 {
				for arr[i][j] > 47 && arr[i][j] < 58 && j < len(arr[0]) {
					partNum += string(arr[i][j])
					j++
					if j == len(arr[0]) {
						break
					}
				}
				posArr[partNum] = pos{i, j}
				partNum = ""
			}
		}
	}
	return posArr
}

func adjacentPos(partNo string, coor pos) []pos {
	var arr []pos
	arr = append(arr, pos{coor.i, coor.j - (len(partNo) + 1)})
	for i := coor.i - 1; i <= coor.i+1; i += 2 {
		for j := coor.j; j >= coor.j-(len(partNo)+1); j-- {
			arr = append(arr, pos{i, j})
		}
		if i == coor.i-1 {
			arr = append(arr, pos{coor.i, coor.j})
		}
	}

	return arr
}

func findGears(mapArr map[string]pos, arr []string) map[string]pos {
	gearMap := make(map[string]pos)
	for k, v := range mapArr {
		posArr := adjacentPos(k, v)
		for x := range posArr {
			if posArr[x].i >= 0 && posArr[x].i < len(arr) && posArr[x].j >= 0 && posArr[x].j < len(arr[0]) {
				if arr[posArr[x].i][posArr[x].j] != 46 {
					gearMap[k] = posArr[x]
				}
			}
		}
	}
	return gearMap
}