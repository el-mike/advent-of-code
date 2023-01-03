package day20

type FileNumber struct {
	OriginalIndex int
	Value         int
}

func NewFileNumber(originalIndex, value int) *FileNumber {
	return &FileNumber{
		OriginalIndex: originalIndex,
		Value:         value,
	}
}
