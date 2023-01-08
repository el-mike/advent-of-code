package day_16

import (
	"regexp"
	"strconv"
	"strings"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseLine(line string) *Valve {
	re := regexp.MustCompile(`Valve ([A-Z]{2}).*=(\d+);.*valves? (.*)`)

	results := re.FindAllStringSubmatch(line, -1)[0]

	// index "0" contains entire match.
	name := results[1]
	flowRate, err := strconv.Atoi(results[2])
	if err != nil {
		panic(err)
	}

	leadsTo := strings.Split(results[3], ", ")

	return NewValve(name, flowRate, leadsTo)
}
