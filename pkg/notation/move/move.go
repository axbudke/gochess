package move

import (
	"cmp"
	"fmt"
	"slices"

	"gochess/pkg/notation/piece"
	"gochess/pkg/notation/square"
)

// ==================== Move List ====================

type MoveList []*Move

func (m MoveList) Sort() {
	slices.SortFunc(m, func(a, b *Move) int {
		var pieceCmp int
		if a.Piece.IsWhite() {
			pieceCmp = cmp.Compare(a.Piece, b.Piece)
		} else {
			pieceCmp = cmp.Compare(b.Piece, a.Piece)
		}
		if pieceCmp != 0 {
			return pieceCmp
		}
		if fromCmp := cmp.Compare(a.From.String(), b.From.String()); fromCmp != 0 {
			return fromCmp
		}
		return cmp.Compare(a.To.String(), b.To.String())
	})
}

func (m MoveList) FindMovesFrom(fromSquare square.Square) MoveList {
	moves := make(MoveList, 0, len(m))
	for _, move := range m {
		if move.From == fromSquare {
			moves = append(moves, move)
		}
	}
	return moves
}

func (m MoveList) FindMovesTo(toSquare square.Square) MoveList {
	moves := make(MoveList, 0, len(m))
	for _, move := range m {
		if move.To == toSquare {
			moves = append(moves, move)
		}
	}
	return moves
}

func (m MoveList) FindMovesNotFrom(fromSquare square.Square) MoveList {
	moves := make(MoveList, 0, len(m))
	for _, move := range m {
		if move.From != fromSquare {
			moves = append(moves, move)
		}
	}
	return moves
}

// ==================== Move ====================

type Move struct {
	From         square.Square
	To           square.Square
	Piece        piece.Piece
	PromotedTo   piece.Piece
	IsCapture    bool
	IsDoublePush bool
	IsCastling   bool
	PieceList    []piece.Piece
}

func (m Move) String() string {
	return string(m.LAN())
}

// PCN - Pure Coordinate Notation
// <move descriptor> ::= <from square><to square>[<promoted to>]
// <square>        ::= <file letter><rank number>
// <file letter>   ::= 'a'|'b'|'c'|'d'|'e'|'f'|'g'|'h'
// <rank number>   ::= '1'|'2'|'3'|'4'|'5'|'6'|'7'|'8'
// <promoted to>   ::= 'q'|'r'|'b'|'n'
type PCN string

func NewMoveFromPCN(pcn PCN) (Move, error) {
	return Move{}, nil
}

func (m Move) PCN() PCN {
	promotedToStr := ""
	if m.PromotedTo != piece.Piece_None {
		promotedToStr = m.PromotedTo.Symbol()
	}
	return PCN(fmt.Sprintf("%s%s%s", m.From, m.To, promotedToStr))
}

// LAN - Long Algebraic Notation
// <LAN move descriptor piece moves> ::= <Piece symbol><from square>['-'|'x']<to square>
// <LAN move descriptor pawn moves>  ::= <from square>['-'|'x']<to square>[<promoted to>]
// <Piece symbol> ::= 'N' | 'B' | 'R' | 'Q' | 'K'
type LAN string

func NewMoveFromLAN(lan LAN) (Move, error) {
	return Move{}, nil
}

func (m Move) LAN() LAN {
	capStr := ""
	if m.IsCapture {
		capStr = "x"
	}
	if !m.Piece.IsPawn() {
		return LAN(fmt.Sprintf("%s%s%s%s", m.Piece.Symbol(), m.From, capStr, m.To))
	} else {
		promotedToStr := ""
		if m.PromotedTo != piece.Piece_None {
			promotedToStr = m.PromotedTo.Symbol()
		}
		return LAN(fmt.Sprintf("%s%s%s%s", m.From, capStr, m.To, promotedToStr))
	}
}

// SAN - Standard Algebraic Notation
// <SAN move descriptor piece moves>   ::= <Piece symbol>[<from file>|<from rank>|<from square>]['x']<to square>
// <SAN move descriptor pawn captures> ::= <from file>[<from rank>] 'x' <to square>[<promoted to>]
// <SAN move descriptor pawn push>     ::= <to square>[<promoted to>]
type SAN string

func NewMoveFromSAN(san SAN) (Move, error) {
	return Move{}, nil
}

func (m Move) SAN() SAN {
	if !m.Piece.IsPawn() {
		capStr := ""
		if m.IsCapture {
			capStr = "x"
		}
		fromStr := ""
		// TODO solve ambiguities
		return SAN(fmt.Sprintf("%s%s%s%s", m.Piece.Symbol(), fromStr, capStr, m.To))
	} else {
		promotedToStr := ""
		if m.PromotedTo != piece.Piece_None {
			promotedToStr = m.PromotedTo.Symbol()
		}
		if m.IsCapture {
			fromF, fromR := m.From.FileRank()
			// TODO solve ambiguities
			return SAN(fmt.Sprintf("%s%sx%s%s", fromF, fromR, m.To, promotedToStr))
		}
		return SAN(fmt.Sprintf("%s%s", m.To, promotedToStr))
	}
}
