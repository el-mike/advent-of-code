package day13

type Packet struct {
	Origin    string
	Elements  []Element
	IsDivider bool
}

func NewPacket(line string, isDivider bool) *Packet {
	parser := NewPdaParser()
	elements, _ := parser.Parse(line[1:len(line)-1], 0)

	return &Packet{
		Origin:    line,
		Elements:  elements,
		IsDivider: isDivider,
	}
}

func (p *Packet) ToList() *ListElement {
	return &ListElement{
		Elements: p.Elements,
	}
}
