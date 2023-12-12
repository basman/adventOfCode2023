package entity

import "fmt"

type Pos struct {
	X, Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("[%v,%v]", p.X, p.Y)
}

type Symbol struct {
	Pos
	Symbol rune
}

func (s Symbol) String() string {
	return fmt.Sprintf("%v %v", string(s.Symbol), s.Pos)
}

type Number struct {
	Pos
	Box
	Value int
}

type Box struct {
	Start Pos
	End   Pos
}

func (b Box) Contains(p Pos) bool {
	return p.X >= b.Start.X && p.X <= b.End.X && p.Y >= b.Start.Y && p.Y <= b.End.Y
}
