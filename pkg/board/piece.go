package board

import "fmt"

type PieceList []PieceVal

func (l PieceList) PieceAt(s Square) PieceVal {
	return l[int(s)]
}

type PieceVal int8

// pieceVal enums
const (
	PieceVal_Empty       PieceVal = 0
	PieceVal_WhitePawn   PieceVal = 1
	PieceVal_WhiteKnight PieceVal = 2
	PieceVal_WhiteBishop PieceVal = 3
	PieceVal_WhiteRook   PieceVal = 4
	PieceVal_WhiteQueen  PieceVal = 5
	PieceVal_WhiteKing   PieceVal = 6
	PieceVal_BlackPawn   PieceVal = -1
	PieceVal_BlackKnight PieceVal = -2
	PieceVal_BlackBishop PieceVal = -3
	PieceVal_BlackRook   PieceVal = -4
	PieceVal_BlackQueen  PieceVal = -5
	PieceVal_BlackKing   PieceVal = -6
)

func (v PieceVal) Char() (PieceChar, error) {
	switch v {
	case PieceVal_WhitePawn:
		return PieceChar_WhitePawn, nil
	case PieceVal_WhiteKnight:
		return PieceChar_WhiteKnight, nil
	case PieceVal_WhiteBishop:
		return PieceChar_WhiteBishop, nil
	case PieceVal_WhiteRook:
		return PieceChar_WhiteRook, nil
	case PieceVal_WhiteQueen:
		return PieceChar_WhiteQueen, nil
	case PieceVal_WhiteKing:
		return PieceChar_WhiteKing, nil
	case PieceVal_BlackPawn:
		return PieceChar_BlackPawn, nil
	case PieceVal_BlackKnight:
		return PieceChar_BlackKnight, nil
	case PieceVal_BlackBishop:
		return PieceChar_BlackBishop, nil
	case PieceVal_BlackRook:
		return PieceChar_BlackRook, nil
	case PieceVal_BlackQueen:
		return PieceChar_BlackQueen, nil
	case PieceVal_BlackKing:
		return PieceChar_BlackKing, nil
	default:
		return PieceChar('a'), fmt.Errorf("invalid value")
	}
}

type PieceChar byte

// pieceChar enums
const (
	PieceChar_WhitePawn   PieceChar = 'p'
	PieceChar_WhiteKnight PieceChar = 'n'
	PieceChar_WhiteBishop PieceChar = 'b'
	PieceChar_WhiteRook   PieceChar = 'r'
	PieceChar_WhiteQueen  PieceChar = 'q'
	PieceChar_WhiteKing   PieceChar = 'k'
	PieceChar_BlackPawn   PieceChar = 'P'
	PieceChar_BlackKnight PieceChar = 'N'
	PieceChar_BlackBishop PieceChar = 'B'
	PieceChar_BlackRook   PieceChar = 'R'
	PieceChar_BlackQueen  PieceChar = 'Q'
	PieceChar_BlackKing   PieceChar = 'K'
)

func (c PieceChar) Val() (PieceVal, error) {
	switch c {
	case PieceChar_WhitePawn:
		return PieceVal_WhitePawn, nil
	case PieceChar_WhiteKnight:
		return PieceVal_WhiteKnight, nil
	case PieceChar_WhiteBishop:
		return PieceVal_WhiteBishop, nil
	case PieceChar_WhiteRook:
		return PieceVal_WhiteRook, nil
	case PieceChar_WhiteQueen:
		return PieceVal_WhiteQueen, nil
	case PieceChar_WhiteKing:
		return PieceVal_WhiteKing, nil
	case PieceChar_BlackPawn:
		return PieceVal_BlackPawn, nil
	case PieceChar_BlackKnight:
		return PieceVal_BlackKnight, nil
	case PieceChar_BlackBishop:
		return PieceVal_BlackBishop, nil
	case PieceChar_BlackRook:
		return PieceVal_BlackRook, nil
	case PieceChar_BlackQueen:
		return PieceVal_BlackQueen, nil
	case PieceChar_BlackKing:
		return PieceVal_BlackKing, nil
	default:
		return PieceVal(0), fmt.Errorf("invalid char")
	}
}
