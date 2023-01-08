package day_15

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"log"
	"time"
)

func BeaconExclusionZone2() {
	start := time.Now()

	scanner, err := common.GetFileScanner("./year_2022/day_15/" + common.InputFilename)
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

	var borderPoints [][]int

	for i, sensor := range sensors {
		sx, sy := sensor[0], sensor[1]
		distance := distances[i]

		for dY := 0; dY <= distance; dY += 1 {
			xDiff := distance - dY

			xpLeft := (sx - xDiff) - 1
			xpRight := (sx + xDiff) + 1
			ypUp := sy - dY
			ypDown := sy + dY

			pLeftUp, pRightUp := []int{xpLeft, ypUp, sx, sy}, []int{xpRight, ypUp, sx, sy}
			borderPoints = append(borderPoints, pLeftUp, pRightUp)

			if dY == 0 {
				continue
			}

			pLeftDown, pRightDown := []int{xpLeft, ypDown, sx, sy}, []int{xpRight, ypDown, sx, sy}
			borderPoints = append(borderPoints, pLeftDown, pRightDown)

			if dY == distance {
				maxUp := (sy - distance) - 1
				maxDown := (sy + distance) + 1

				pMaxUp, pMaxDown := []int{sx, maxUp, sx, sy}, []int{sx, maxDown, sx, sy}

				borderPoints = append(borderPoints, pMaxUp, pMaxDown)
			}
		}
	}

	targetX, targetY := 0, 0

	for _, point := range borderPoints {
		if point[0] < 0 || point[0] > SearchSize || point[1] < 0 || point[1] > SearchSize {
			continue
		}

		covered := false

		for i, sensor := range sensors {
			if point[2] == sensor[0] && point[3] == sensor[1] {
				continue
			}

			distance := distances[i]

			if getDistance(sensor, point) <= distance {
				covered = true
			}
		}

		if !covered {
			targetX, targetY = point[0], point[1]
			break
		}
	}

	result := (targetX * Multiplier) + targetY

	fmt.Println(result)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
