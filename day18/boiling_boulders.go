package day18

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"strconv"
	"strings"
)

func BoilingBoulders() {
	scanner, err := common.GetFileScanner("./day18/" + common.TestInputFilename)
	if err != nil {
		panic(err)
	}

	var droplets []*Droplet

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var coords Coords

		for i, numberStr := range strings.Split(line, ",") {
			value, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}

			coords[i] = value
		}

		droplets = append(droplets, NewDroplet(coords))
	}

	fmt.Println()
}
