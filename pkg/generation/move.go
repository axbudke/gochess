package generation

import (
	"fmt"
	"gochess/pkg/board"
)

type Move struct {
	BoardI     board.BoardInterface
	From       board.Square
	To         board.Square
	Piece      board.PieceVal
	IsCapture  bool
	PromotedTo board.PieceVal
}

func (m Move) String() string {
	return m.LAN()
}

// PCN - Pure Coordinate Notation
// <move descriptor> ::= <from square><to square>[<promoted to>]
// <square>        ::= <file letter><rank number>
// <file letter>   ::= 'a'|'b'|'c'|'d'|'e'|'f'|'g'|'h'
// <rank number>   ::= '1'|'2'|'3'|'4'|'5'|'6'|'7'|'8'
// <promoted to>   ::= 'q'|'r'|'b'|'n'
func (m Move) PCN() string {
	return fmt.Sprintf("%s%s%s", m.From, m.To, m.PromotedTo)
}

// LAN - Long Algebraic Notation
// <LAN move descriptor piece moves> ::= <Piece symbol><from square>['-'|'x']<to square>
// <LAN move descriptor pawn moves>  ::= <from square>['-'|'x']<to square>[<promoted to>]
// <Piece symbol> ::= 'N' | 'B' | 'R' | 'Q' | 'K'
func (m Move) LAN() string {
	capStr := ""
	if m.IsCapture {
		capStr = "x"
	}

	if m.Piece != board.PieceVal_WhitePawn {
		return fmt.Sprintf("%s%s%s%s", m.Piece.PieceSymbol(), m.From, capStr, m.To)
	} else {
		return fmt.Sprintf("%s%s%s%s", m.From, capStr, m.To, m.PromotedTo)
	}
}

// SAN - Standard Algebraic Notation
// <SAN move descriptor piece moves>   ::= <Piece symbol>[<from file>|<from rank>|<from square>]['x']<to square>
// <SAN move descriptor pawn captures> ::= <from file>[<from rank>] 'x' <to square>[<promoted to>]
// <SAN move descriptor pawn push>     ::= <to square>[<promoted to>]
func (m Move) SAN() string {
	if m.Piece != board.PieceVal_WhitePawn {
		capStr := ""
		if m.IsCapture {
			capStr = "x"
		}
		fromStr := ""
		// TODO solve ambiguities
		return fmt.Sprintf("%s%s%s%s", m.Piece.PieceSymbol(), fromStr, capStr, m.To)
	} else {
		if m.IsCapture {
			fromF, fromR := m.From.FileRank()
			return fmt.Sprintf("%s%sx%s%s", fromF, fromR, m.To, m.PromotedTo)
		}
		return fmt.Sprintf("%s%s", m.To, m.PromotedTo)
	}
}

type MoveList []*Move

func (m MoveList) String() string {
	str := ""
	for _, move := range m {
		str += fmt.Sprintf("Move: %s ", *move)
	}
	return str
}
