package day10

//goland:noinspection SpellCheckingInspection
const (
	NoopOperation = "noop"
	AddOperation  = "addx"
)

var executionTimesMap = map[string]int{
	NoopOperation: 1,
	AddOperation:  2,
}

type Operation struct {
	Name          string
	Argument      string
	ExecutionTime int
}

func NewOperation(name, argument string) *Operation {
	return &Operation{
		Name:          name,
		Argument:      argument,
		ExecutionTime: executionTimesMap[name],
	}
}
