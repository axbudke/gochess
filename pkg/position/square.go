package position

import (
	"fmt"
	"strconv"
)

type Rank int

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) String() string {
	if r < 0 || r >= 8 {
		return "_"
	}
	return strconv.Itoa(int(r) + 1)
}

type File int

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) String() string {
	if f < 0 || f >= 8 {
		return "_"
	}
	return []string{"a", "b", "c", "d", "e", "f", "g", "h"}[f]
}

type Square int

const (
	Square_a1 Square = iota
	Square_a2
	Square_a3
	Square_a4
	Square_a5
	Square_a6
	Square_a7
	Square_a8
	Square_b1
	Square_b2
	Square_b3
	Square_b4
	Square_b5
	Square_b6
	Square_b7
	Square_b8
	Square_c1
	Square_c2
	Square_c3
	Square_c4
	Square_c5
	Square_c6
	Square_c7
	Square_c8
	Square_d1
	Square_d2
	Square_d3
	Square_d4
	Square_d5
	Square_d6
	Square_d7
	Square_d8
	Square_e1
	Square_e2
	Square_e3
	Square_e4
	Square_e5
	Square_e6
	Square_e7
	Square_e8
	Square_f1
	Square_f2
	Square_f3
	Square_f4
	Square_f5
	Square_f6
	Square_f7
	Square_f8
	Square_g1
	Square_g2
	Square_g3
	Square_g4
	Square_g5
	Square_g6
	Square_g7
	Square_g8
	Square_h1
	Square_h2
	Square_h3
	Square_h4
	Square_h5
	Square_h6
	Square_h7
	Square_h8
)

func NewSquare(r Rank, f File) (Square, error) {
	if r < 0 || r >= 8 {
		return -1, fmt.Errorf("invalid rank")
	} else if f < 0 || f >= 8 {
		return -1, fmt.Errorf("invalid file")
	}
	return Square(int(r)*8 + int(f)), nil
}

func (s Square) String() string {
	f, r := s.FileRank()
	return f.String() + r.String()
}

func (s Square) FileRank() (File, Rank) {
	return File(int(s) % 8), Rank(int(s) / 8)
}
