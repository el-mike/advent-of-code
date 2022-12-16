package day15

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

const (
	SearchSize = 4000000
	Multiplier = 4000000
)

// NumberOfBatches - after some testing, seems like around 10 goroutines calculate
// the result fastest.
const NumberOfBatches = 10

func BeaconExclusionZone() {
	start := time.Now()

	scanner, err := common.GetFileScanner("./day15/" + InputFilename)
	if err != nil {
		panic(err)
	}

	var sensors [][]int
	var distances []int

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

	targetX, targetY := 0, 0

	wg := sync.WaitGroup{}
	wg.Add(NumberOfBatches)

	cancelCh := make(chan bool)

	batchSize := SearchSize / NumberOfBatches

	for n := 0; n < NumberOfBatches; n += 1 {
		y := n * batchSize
		yEnd := (y + batchSize) - 1

		go func() {
			for ; y <= yEnd; y += 1 {
				select {
				case <-cancelCh:
					wg.Done()
					return
				default:
					var ranges [][]int

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

						ranges = append(ranges, []int{min, max})
					}

					rangesOverlap, openRangeEndX := testRanges(ranges)

					if !rangesOverlap {
						cancelCh <- true
						targetX, targetY = openRangeEndX+1, y
					}
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()

	result := (targetX * Multiplier) + targetY

	fmt.Println(result)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
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

func overlap(source, target []int) bool {
	return source[0] <= target[1] && source[1] >= target[0]
}
