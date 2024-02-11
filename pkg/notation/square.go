package notation

import (
	"fmt"
	"strconv"
)

func SquaresInBetween(fromSquare, toSquare Square) []Square {
	squaresInBetween := make([]Square, 0, 6)

	// Same Square
	if fromSquare == toSquare {
		return squaresInBetween
	}

	fromF, fromR := fromSquare.FileRank()
	toF, toR := toSquare.FileRank()

	// On the same file
	if fromF == toF {
		lowerR, higherR := fromR, toR
		if fromR > toR {
			lowerR, higherR = toR, fromR
		}
		for rp := lowerR + 1; rp < higherR; rp++ {
			squaresInBetween = append(squaresInBetween, NewSquare(fromF, rp))
		}
		return squaresInBetween
	}

	// On the same rank
	if fromR == toR {
		lowerF, higherF := fromF, toF
		if fromF > toF {
			lowerF, higherF = toF, fromF
		}
		for fp := lowerF + 1; fp < higherF; fp++ {
			squaresInBetween = append(squaresInBetween, NewSquare(fp, fromR))
		}
		return squaresInBetween
	}

	// On the positive diagonal
	if int(fromR-toR) == int(fromF-toF) {
		lowerR, _ := fromR, toR
		lowerF, higherF := fromF, toF
		if fromF > toF {
			lowerR, _ = toR, fromR
			lowerF, higherF = toF, fromF
		}
		for p := 1; p < int(higherF-lowerF); p++ {
			squaresInBetween = append(squaresInBetween, NewSquare(lowerF+File(p), lowerR-Rank(p)))
		}
		return squaresInBetween
	}

	// On the negative diagonal
	if int(fromR-toR) == -int(fromF-toF) {
		lowerR, _ := fromR, toR
		lowerF, higherF := fromF, toF
		if fromF > toF {
			lowerR, _ = toR, fromR
			lowerF, higherF = toF, fromF
		}
		for p := 1; p < int(higherF-lowerF); p++ {
			squaresInBetween = append(squaresInBetween, NewSquare(lowerF+File(p), lowerR-Rank(p)))
		}
		return squaresInBetween
	}

	// No valid in between squares

	return squaresInBetween
}

// ==================== Square ====================

type Square int

const (
	Square_a1 Square = iota
	Square_b1
	Square_c1
	Square_d1
	Square_e1
	Square_f1
	Square_g1
	Square_h1
	Square_a2
	Square_b2
	Square_c2
	Square_d2
	Square_e2
	Square_f2
	Square_g2
	Square_h2
	Square_a3
	Square_b3
	Square_c3
	Square_d3
	Square_e3
	Square_f3
	Square_g3
	Square_h3
	Square_a4
	Square_b4
	Square_c4
	Square_d4
	Square_e4
	Square_f4
	Square_g4
	Square_h4
	Square_a5
	Square_b5
	Square_c5
	Square_d5
	Square_e5
	Square_f5
	Square_g5
	Square_h5
	Square_a6
	Square_b6
	Square_c6
	Square_d6
	Square_e6
	Square_f6
	Square_g6
	Square_h6
	Square_a7
	Square_b7
	Square_c7
	Square_d7
	Square_e7
	Square_f7
	Square_g7
	Square_h7
	Square_a8
	Square_b8
	Square_c8
	Square_d8
	Square_e8
	Square_f8
	Square_g8
	Square_h8

	Square_Invalid
)

func NewSquare(f File, r Rank) Square {
	return Square(int(r)<<3 + int(f))
}

func NewSquareCheck(f File, r Rank) (Square, error) {
	if f < 0 || f >= 8 {
		return -1, fmt.Errorf("invalid file")
	} else if r < 0 || r >= 8 {
		return -1, fmt.Errorf("invalid rank")
	}
	return NewSquare(f, r), nil
}

func (s Square) String() string {
	f, r := s.FileRank()
	return f.String() + r.String()
}

func (s Square) FileRank() (File, Rank) {
	return File(int(s) & 7), Rank(int(s) >> 3)
}

// Square is a simple notation for a square on a chess board
// <square>        ::= <file letter><rank number>
// <file letter>   ::= 'a'|'b'|'c'|'d'|'e'|'f'|'g'|'h'
// <rank number>   ::= '1'|'2'|'3'|'4'|'5'|'6'|'7'|'8'

func NewSquareFromString(str string) (Square, error) {
	bs := []byte(str)
	if len(bs) != 2 {
		return Square_Invalid, fmt.Errorf("invalid square")
	}
	f, err := NewFileFromString(string(bs[0]))
	if err != nil {
		return Square_Invalid, err
	}
	r, err := NewRankFromString(string(bs[1]))
	if err != nil {
		return Square_Invalid, err
	}
	return NewSquare(f, r), nil
}

// ==================== Rank ====================

type Rank int8

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8

	RankInvalid
)

func NewRankFromString(str string) (Rank, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return RankInvalid, err
	}
	r := i - 1
	if r < 0 || r >= 8 {
		return RankInvalid, fmt.Errorf("invalid rank: out of range")
	}
	return Rank(r), nil
}

func (r Rank) String() string {
	if r < 0 || r >= 8 {
		return "_"
	}
	return strconv.Itoa(int(r) + 1)
}

// ==================== File ====================

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

	FileInvalid
)

type FileStr string

const (
	FileStrA FileStr = "a"
	FileStrB FileStr = "b"
	FileStrC FileStr = "c"
	FileStrD FileStr = "d"
	FileStrE FileStr = "e"
	FileStrF FileStr = "f"
	FileStrG FileStr = "g"
	FileStrH FileStr = "h"

	FileStrInvalid FileStr = "_"
)

func NewFileFromString(str string) (File, error) {
	switch FileStr(str) {
	case FileStrA:
		return FileA, nil
	case FileStrB:
		return FileB, nil
	case FileStrC:
		return FileC, nil
	case FileStrD:
		return FileD, nil
	case FileStrE:
		return FileE, nil
	case FileStrF:
		return FileF, nil
	case FileStrG:
		return FileG, nil
	case FileStrH:
		return FileH, nil
	default:
		return FileInvalid, fmt.Errorf("invalid file")
	}
}

func (f File) String() string {
	switch f {
	case FileA:
		return string(FileStrA)
	case FileB:
		return string(FileStrB)
	case FileC:
		return string(FileStrC)
	case FileD:
		return string(FileStrD)
	case FileE:
		return string(FileStrE)
	case FileF:
		return string(FileStrF)
	case FileG:
		return string(FileStrG)
	case FileH:
		return string(FileStrH)
	default:
		return string(FileStrInvalid)
	}
}
