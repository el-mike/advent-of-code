package day15

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"log"
	"sort"
	"time"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

func overlap(source, target []int) bool {
	return source[0] <= target[1] && source[1] >= target[0]
}

func testRanges(ranges [][]int) (bool, int) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	openRange := ranges[0]

	for i := 1; i < len(ranges); i += 1 {
		current := ranges[i]

		if !overlap(current, openRange) {
			return false, openRange[1]
		}

		if current[1] > openRange[1] {
			openRange[1] = current[1]
		}
	}

	return true, -1
}

func BeaconExclusionZone() {
	start := time.Now()

	scanner, err := common.GetFileScanner("./day15/" + InputFilename)
	if err != nil {
		panic(err)
	}

	var sensors [][]int
	var distances []int

	target := 4000000
	//target := 20

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		numbers, err := common.GetNumbersFromLine(line)
		if err != nil {
			panic(err)
		}

		sensor := []int{numbers[0], numbers[1]}
		beacon := []int{numbers[2], numbers[3]}
		distance := getDistance(sensor, beacon)

		sensors = append(sensors, sensor)
		distances = append(distances, distance)
	}

	x, y := 0, 0

	var rangeMap [][][]int

	for ; y <= target; y += 1 {
		rangeMap = append(rangeMap, [][]int{})

		for i, sensor := range sensors {
			sy, d := sensor[1], distances[i]

			upY := sy - d
			downY := sy + d

			diffX := 0

			if y <= sy && upY <= y {
				diffX = y - upY
			} else if y > sy && downY >= y {
				diffX = downY - y
			} else {
				continue
			}

			min := sensor[0] - diffX
			max := sensor[0] + diffX

			rangeMap[y] = append(rangeMap[y], []int{min, max})
		}

		rangesOverlap, openRangeEndX := testRanges(rangeMap[y])

		if !rangesOverlap {
			x = openRangeEndX + 1
			break
		}
	}

	multiplier := 4000000
	result := (x * multiplier) + y

	fmt.Println(result)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getDistance(sensor, beacon []int) int {
	sx, sy, bx, by := sensor[0], sensor[1], beacon[0], beacon[1]

	dx := 0
	dy := 0

	if sx > bx {
		dx = sx - bx
	} else if sx < bx {
		dx = bx - sx
	}

	if sy > by {
		dy = sy - by
	} else if sy < by {
		dy = by - sy
	}

	return dx + dy
}
