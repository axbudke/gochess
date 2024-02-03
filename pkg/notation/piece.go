package notation

import (
	"fmt"
	"math"
)

type PieceList []Piece

func (l PieceList) PieceAt(s Square) Piece {
	return l[int(s)]
}

var PromotionPieces = []Piece{
	Piece_WhiteQueen,
	Piece_WhiteRook,
	Piece_WhiteBishop,
	Piece_WhiteKnight,
}

type Piece int8

const (
	Piece_Empty       Piece = 0
	Piece_WhitePawn   Piece = 1
	Piece_WhiteKnight Piece = 2
	Piece_WhiteBishop Piece = 3
	Piece_WhiteRook   Piece = 4
	Piece_WhiteQueen  Piece = 5
	Piece_WhiteKing   Piece = 6
	Piece_BlackPawn   Piece = -1
	Piece_BlackKnight Piece = -2
	Piece_BlackBishop Piece = -3
	Piece_BlackRook   Piece = -4
	Piece_BlackQueen  Piece = -5
	Piece_BlackKing   Piece = -6
)

func (v Piece) String() string {
	c, err := v.Char()
	if err != nil {
		return ""
	}
	return fmt.Sprint(c)
}

func (v Piece) Symbol() string {
	c, err := Piece(int(math.Abs(float64(v)))).Char()
	if err != nil {
		return ""
	}
	return fmt.Sprint(c)
}

func (v Piece) Char() (PieceChar, error) {
	switch v {
	case Piece_WhitePawn:
		return PieceChar_WhitePawn, nil
	case Piece_WhiteKnight:
		return PieceChar_WhiteKnight, nil
	case Piece_WhiteBishop:
		return PieceChar_WhiteBishop, nil
	case Piece_WhiteRook:
		return PieceChar_WhiteRook, nil
	case Piece_WhiteQueen:
		return PieceChar_WhiteQueen, nil
	case Piece_WhiteKing:
		return PieceChar_WhiteKing, nil
	case Piece_BlackPawn:
		return PieceChar_BlackPawn, nil
	case Piece_BlackKnight:
		return PieceChar_BlackKnight, nil
	case Piece_BlackBishop:
		return PieceChar_BlackBishop, nil
	case Piece_BlackRook:
		return PieceChar_BlackRook, nil
	case Piece_BlackQueen:
		return PieceChar_BlackQueen, nil
	case Piece_BlackKing:
		return PieceChar_BlackKing, nil
	default:
		return PieceChar('a'), fmt.Errorf("invalid value")
	}
}

type PieceChar byte

const (
	PieceChar_WhitePawn   PieceChar = 'P'
	PieceChar_WhiteKnight PieceChar = 'N'
	PieceChar_WhiteBishop PieceChar = 'B'
	PieceChar_WhiteRook   PieceChar = 'R'
	PieceChar_WhiteQueen  PieceChar = 'Q'
	PieceChar_WhiteKing   PieceChar = 'K'
	PieceChar_BlackPawn   PieceChar = 'p'
	PieceChar_BlackKnight PieceChar = 'n'
	PieceChar_BlackBishop PieceChar = 'b'
	PieceChar_BlackRook   PieceChar = 'r'
	PieceChar_BlackQueen  PieceChar = 'q'
	PieceChar_BlackKing   PieceChar = 'k'
)

func (c PieceChar) String() string {
	return string([]byte{byte(c)})
}

func (c PieceChar) Val() (Piece, error) {
	switch c {
	case PieceChar_WhitePawn:
		return Piece_WhitePawn, nil
	case PieceChar_WhiteKnight:
		return Piece_WhiteKnight, nil
	case PieceChar_WhiteBishop:
		return Piece_WhiteBishop, nil
	case PieceChar_WhiteRook:
		return Piece_WhiteRook, nil
	case PieceChar_WhiteQueen:
		return Piece_WhiteQueen, nil
	case PieceChar_WhiteKing:
		return Piece_WhiteKing, nil
	case PieceChar_BlackPawn:
		return Piece_BlackPawn, nil
	case PieceChar_BlackKnight:
		return Piece_BlackKnight, nil
	case PieceChar_BlackBishop:
		return Piece_BlackBishop, nil
	case PieceChar_BlackRook:
		return Piece_BlackRook, nil
	case PieceChar_BlackQueen:
		return Piece_BlackQueen, nil
	case PieceChar_BlackKing:
		return Piece_BlackKing, nil
	default:
		return Piece_Empty, fmt.Errorf("invalid char")
	}
}
