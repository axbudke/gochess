package notation

import (
	"fmt"
)

var PromotionPieces = []Piece{
	Piece_WhiteQueen,
	Piece_WhiteRook,
	Piece_WhiteBishop,
	Piece_WhiteKnight,
}

// ==================== Piece List ====================

type PieceList []Piece

func (l PieceList) PieceAt(s Square) Piece {
	return l[int(s)]
}

// ==================== Piece ====================

type Piece2 uint8

const (
	Piece2_None Piece2 = 0

	// General Pieces
	Piece2_Pawn   Piece2 = 1
	Piece2_Knight Piece2 = 2
	Piece2_Bishop Piece2 = 3
	Piece2_Rook   Piece2 = 4
	Piece2_Queen  Piece2 = 5
	Piece2_King   Piece2 = 6

	// Piece Color
	Piece2_NoColor Piece2 = 0b01111111
	Piece2_Color   Piece2 = 0b10000000
	Piece2_White   Piece2 = 0b00000000
	Piece2_Black   Piece2 = 0b10000000

	// White Pieces
	Piece2_WhitePawn   Piece2 = Piece2_White | Piece2_Pawn
	Piece2_WhiteKnight Piece2 = Piece2_White | Piece2_Knight
	Piece2_WhiteBishop Piece2 = Piece2_White | Piece2_Bishop
	Piece2_WhiteRook   Piece2 = Piece2_White | Piece2_Rook
	Piece2_WhiteQueen  Piece2 = Piece2_White | Piece2_Queen
	Piece2_WhiteKing   Piece2 = Piece2_White | Piece2_King

	// Black Pieces
	Piece2_BlackPawn   Piece2 = Piece2_Black | Piece2_Pawn
	Piece2_BlackKnight Piece2 = Piece2_Black | Piece2_Knight
	Piece2_BlackBishop Piece2 = Piece2_Black | Piece2_Bishop
	Piece2_BlackRook   Piece2 = Piece2_Black | Piece2_Rook
	Piece2_BlackQueen  Piece2 = Piece2_Black | Piece2_Queen
	Piece2_BlackKing   Piece2 = Piece2_Black | Piece2_King
)

func (p Piece2) IsWhite() bool { return (p & Piece2_Color) == Piece2_White }
func (p Piece2) IsBlack() bool { return (p & Piece2_Color) == Piece2_Black }

func (p Piece2) IsPawn() bool { return (p & Piece2_NoColor) == Piece2_Pawn }

func (p Piece2) Symbol() string {
	switch p & Piece2_NoColor {
	case Piece2_Pawn:
		return string(PieceStr_Pawn)
	case Piece2_Knight:
		return string(PieceStr_Knight)
	case Piece2_Bishop:
		return string(PieceStr_Bishop)
	case Piece2_Rook:
		return string(PieceStr_Rook)
	case Piece2_Queen:
		return string(PieceStr_Queen)
	case Piece2_King:
		return string(PieceStr_King)
	default:
		return ""
	}
}

func (c Piece2) String() string {
	switch c {
	// White Pieces
	case Piece2_WhitePawn:
		return string(PieceStr_WhitePawn)
	case Piece2_WhiteKnight:
		return string(PieceStr_WhiteKnight)
	case Piece2_WhiteBishop:
		return string(PieceStr_WhiteBishop)
	case Piece2_WhiteRook:
		return string(PieceStr_WhiteRook)
	case Piece2_WhiteQueen:
		return string(PieceStr_WhiteQueen)
	case Piece2_WhiteKing:
		return string(PieceStr_WhiteKing)
	// Black Pieces
	case Piece2_BlackPawn:
		return string(PieceStr_BlackPawn)
	case Piece2_BlackKnight:
		return string(PieceStr_BlackKnight)
	case Piece2_BlackBishop:
		return string(PieceStr_BlackBishop)
	case Piece2_BlackRook:
		return string(PieceStr_BlackRook)
	case Piece2_BlackQueen:
		return string(PieceStr_BlackQueen)
	case Piece2_BlackKing:
		return string(PieceStr_BlackKing)
	// Default
	default:
		return ""
	}
}

type PieceStr string

const (
	PieceStr_Pawn   PieceStr = "P"
	PieceStr_Knight PieceStr = "N"
	PieceStr_Bishop PieceStr = "B"
	PieceStr_Rook   PieceStr = "R"
	PieceStr_Queen  PieceStr = "Q"
	PieceStr_King   PieceStr = "K"

	PieceStr_WhitePawn   PieceStr = "P"
	PieceStr_WhiteKnight PieceStr = "N"
	PieceStr_WhiteBishop PieceStr = "B"
	PieceStr_WhiteRook   PieceStr = "R"
	PieceStr_WhiteQueen  PieceStr = "Q"
	PieceStr_WhiteKing   PieceStr = "K"

	PieceStr_BlackPawn   PieceStr = "p"
	PieceStr_BlackKnight PieceStr = "n"
	PieceStr_BlackBishop PieceStr = "b"
	PieceStr_BlackRook   PieceStr = "r"
	PieceStr_BlackQueen  PieceStr = "q"
	PieceStr_BlackKing   PieceStr = "k"
)

func NewPieceFromStr(str string) (Piece, error) {
	switch PieceStr(str) {
	// White Pieces
	case PieceStr_WhitePawn:
		return Piece_WhitePawn, nil
	case PieceStr_WhiteKnight:
		return Piece_WhiteKnight, nil
	case PieceStr_WhiteBishop:
		return Piece_WhiteBishop, nil
	case PieceStr_WhiteRook:
		return Piece_WhiteRook, nil
	case PieceStr_WhiteQueen:
		return Piece_WhiteQueen, nil
	case PieceStr_WhiteKing:
		return Piece_WhiteKing, nil
	// Default
	default:
		return Piece_None, nil
	}
}

type Piece int8

const (
	Piece_None Piece = 0

	Piece_Pawn   Piece = 1
	Piece_Knight Piece = 2
	Piece_Bishop Piece = 3
	Piece_Rook   Piece = 4
	Piece_Queen  Piece = 5
	Piece_King   Piece = 6

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

func (p Piece) Abs() Piece {
	if p < 0 {
		return -p
	}
	return p
}

func (p Piece) IsEmpty() bool { return p == 0 }
func (p Piece) IsWhite() bool { return p > 0 }
func (p Piece) IsBlack() bool { return p < 0 }

func (p Piece) IsPawn() bool   { return p.Abs() == Piece_Pawn }
func (p Piece) IsKnight() bool { return p.Abs() == Piece_Knight }
func (p Piece) IsBishop() bool { return p.Abs() == Piece_Bishop }
func (p Piece) IsRook() bool   { return p.Abs() == Piece_Rook }
func (p Piece) IsQueen() bool  { return p.Abs() == Piece_Queen }
func (p Piece) IsKing() bool   { return p.Abs() == Piece_King }

func (v Piece) String() string {
	c, err := v.Char()
	if err != nil {
		return " "
	}
	return fmt.Sprint(c)
}

func (v Piece) Symbol() string {
	c, err := v.Abs().Char()
	if err != nil {
		return ""
	}
	return fmt.Sprint(c)
}

func (v Piece) Char() (PieceChar, error) {
	switch v {
	// White Pieces
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
	// Black Pieces
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
	// Default
	default:
		return PieceChar('a'), fmt.Errorf("invalid value")
	}
}

type PieceChar byte

const (
	PieceChar_Empty PieceChar = ' '

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
		return Piece_None, fmt.Errorf("invalid char")
	}
}
