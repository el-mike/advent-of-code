package day13

const (
	IntegerType uint8 = iota
	ListType
)

type Element interface {
	Type() uint8
}

type IntegerElement struct {
	Value int
}

func (e *IntegerElement) Type() uint8 {
	return IntegerType
}

type ListElement struct {
	Elements []Element
}

func (e *ListElement) Type() uint8 {
	return ListType
}
