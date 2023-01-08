package day_13

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"sort"
)

type ComparisonResult uint8

const (
	Same ComparisonResult = iota
	LeftSmaller
	RightSmaller
)

func DistressSignal() {
	scanner, err := common.GetFileScanner("./year_2022/day_13/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	var packets []*Packet

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		packet := NewPacket(line, false)
		packets = append(packets, packet)
	}

	packets = append(
		packets,
		NewPacket("[[2]]", true),
		NewPacket("[[6]]", true),
	)

	sort.Slice(packets, func(i, j int) bool {
		left, right := packets[i], packets[j]

		comparisonResult := compare(left.ToList(), right.ToList())

		if comparisonResult == LeftSmaller {
			return true
		} else {
			return false
		}
	})

	decoderKey := 1

	for i, packet := range packets {
		if packet.IsDivider {
			// Packets are one-indexed, therefore we need to add one.
			decoderKey *= i + 1
		}
	}

	fmt.Println(decoderKey)
}

func compare(left, right Element) ComparisonResult {
	if isInteger(left) && isInteger(right) {
		return compareInt(left.(*IntegerElement), right.(*IntegerElement))
	}

	var leftList, rightList *ListElement

	if !isList(left) {
		leftList = &ListElement{Elements: []Element{left}}
	} else {
		leftList = left.(*ListElement)
	}

	if !isList(right) {
		rightList = &ListElement{Elements: []Element{right}}
	} else {
		rightList = right.(*ListElement)
	}

	leftLen := len(leftList.Elements)
	rightLen := len(rightList.Elements)
	maxLen := 0

	if leftLen > rightLen {
		maxLen = len(leftList.Elements)
	} else {
		maxLen = len(rightList.Elements)
	}

	for i := 0; i < maxLen; i += 1 {
		if i >= leftLen {
			return LeftSmaller
		}

		if i >= rightLen {
			return RightSmaller
		}

		leftElement := leftList.Elements[i]
		rightElement := rightList.Elements[i]

		result := compare(leftElement, rightElement)

		if result == LeftSmaller || result == RightSmaller {
			return result
		}
	}

	return Same
}

func compareInt(left, right *IntegerElement) ComparisonResult {
	leftValue := left.Value
	rightValue := right.Value

	if leftValue == rightValue {
		return Same
	}

	if leftValue < rightValue {
		return LeftSmaller
	} else {
		return RightSmaller
	}
}

func isInteger(element Element) bool {
	return element.Type() == IntegerType
}

func isList(element Element) bool {
	return element.Type() == ListType
}
