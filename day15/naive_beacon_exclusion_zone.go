package day15

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

func NaiveBeaconExclusionZone() {
	scanner, err := common.GetFileScanner("./day15/" + InputFilename)
	if err != nil {
		panic(err)
	}

	var sensors [][]int
	var beacons [][]int

	//target := 10
	target := 2000000

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

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	scannedMap := map[int][]int{}
	added := map[int]bool{}

	for i, sensor := range sensors {
		beacon := beacons[i]

		distance := getDistance(sensor, beacon)

		sx, sy := sensor[0], sensor[1]

		for j := 0; j <= distance; j += 1 {
			if sy-j == target || sy+j == target {
				for k := 0; k <= (distance - j); k += 1 {
					x1, x2 := sx-k, sx+k
					if !added[x1] {
						scannedMap[sy-j] = append(scannedMap[sy-j], x1)
						scannedMap[sy+j] = append(scannedMap[sy+j], x1)

						added[x1] = true
					}

					if !added[x2] {
						scannedMap[sy-j] = append(scannedMap[sy-j], x2)
						scannedMap[sy+j] = append(scannedMap[sy+j], x2)

						added[x2] = true
					}
				}
			}
		}
	}

	targetRow := scannedMap[target]

	positions := len(targetRow)

	var targetBeacons [][]int
	var targetSensors [][]int
	var checkedX []int
	var checkedY []int

	for _, beacon := range beacons {
		if beacon[1] == target {
			targetBeacons = append(targetBeacons, beacon)
		}
	}

	for _, sensor := range sensors {
		if sensor[1] == target {
			targetSensors = append(targetSensors, sensor)
		}
	}

	for _, tx := range targetRow {
		for _, beacon := range targetBeacons {
			if tx == beacon[0] && !common.Contains(checkedX, beacon[0]) && !common.Contains(checkedY, beacon[1]) {
				positions -= 1
				checkedX = append(checkedX, beacon[0])
				checkedY = append(checkedY, beacon[1])
			}
		}

		for _, sensor := range targetSensors {
			if tx == sensor[0] && !common.Contains(checkedX, sensor[0]) && !common.Contains(checkedY, sensor[1]) {
				positions -= 1
				checkedX = append(checkedX, sensor[0])
				checkedY = append(checkedY, sensor[1])
			}
		}
	}

	fmt.Print(positions)
}
