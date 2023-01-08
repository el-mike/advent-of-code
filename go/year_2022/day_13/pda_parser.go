package day_13

import "strconv"

type PdaParser struct{}

func NewPdaParser() *PdaParser {
	return &PdaParser{}
}

func (pp *PdaParser) parseInt(intStr string) int {
	value, err := strconv.Atoi(intStr)
	if err != nil {
		panic(err)
	}

	return value
}

func (pp *PdaParser) Parse(line string, startIndex int) ([]Element, int) {
	var elements []Element
	currentIntStr := ""

	i := startIndex

	for ; i < len(line); i += 1 {
		r := line[i]

		if r == '[' {
			childElements, resumeAt := pp.Parse(line, i+1)
			listElement := &ListElement{
				Elements: childElements,
			}

			elements = append(elements, listElement)
			i = resumeAt

			continue
		} else if r == ']' || r == ',' {
			if currentIntStr != "" {
				elements = append(elements, &IntegerElement{Value: pp.parseInt(currentIntStr)})
				currentIntStr = ""
			}

			if r == ',' {
				continue
			} else {
				break
			}
		} else {
			currentIntStr += string(r)
		}
	}

	if currentIntStr != "" {
		elements = append(elements, &IntegerElement{Value: pp.parseInt(currentIntStr)})
		currentIntStr = ""
	}

	return elements, i + 1
}
