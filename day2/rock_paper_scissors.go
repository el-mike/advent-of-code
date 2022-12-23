package day2

import (
	"el-mike/advent-of-code/common"
	"strings"
)

const (
	ROCK = iota
	PAPER
	SCISSORS
)

const (
	LOSE = iota
	DRAW
	WIN
)

var OPPONENT_MOVES = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var DESIRED_RESULTS = map[string]int{
	"X": LOSE,
	"Y": DRAW,
	"Z": WIN,
}

var SHAPE_SCORES = map[int]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

var RESULT_SCORES = map[int]int{
	LOSE: 0,
	DRAW: 3,
	WIN:  6,
}

var WINS_TO = map[int]int{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

func getPlayerShape(opponentShape int, desiredResult int) int {
	if desiredResult == DRAW {
		// For the draw, we simply return opponent's shape.
		return opponentShape
	} else if desiredResult == WIN {
		// For the win, we simply return a move that beats opponent's shape.
		return WINS_TO[opponentShape]
	} else {
		// For the lose, we invert the win scenario, by getting the key by value.
		return common.GetKeyByValue(WINS_TO, opponentShape)
	}
}

func RockPaperScissors() int {
	scanner, err := common.GetFileScanner("./day2/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		moves := strings.Split(line, " ")
		opponentMoveSymbol, resultSymbol := moves[0], moves[1]

		opponentShape, desiredResult := OPPONENT_MOVES[opponentMoveSymbol], DESIRED_RESULTS[resultSymbol]
		playerShape := getPlayerShape(opponentShape, desiredResult)

		totalScore += RESULT_SCORES[desiredResult] + SHAPE_SCORES[playerShape]
	}

	return totalScore
}
