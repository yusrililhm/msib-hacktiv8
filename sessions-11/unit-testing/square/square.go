package square

import "math"

type Square struct {
	s float64
}

type SquareInterface interface {
	Area() float64
	Volume() float64
	Keliling() float64
}

func (s *Square) Area() float64 {
	return math.Pow(s.s, 2)
}

func (s *Square) Volume() float64 {
	return math.Pow(s.s, 3)
}

func (s *Square) Keliling() float64 {
	return 4 * s.s
}

func SetSisi(sisi float64) SquareInterface {
	return &Square{
		s: sisi,
	}
}
