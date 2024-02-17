package generation

import (
	"gochess/pkg/notation/move"
	"gochess/pkg/notation/piece"
	"gochess/pkg/notation/position"
	"gochess/pkg/notation/square"
)

func MakeMove(p *position.Position, m move.Move) *position.Position {
	newP := &position.Position{}

	// Update PieceList
	newP.PieceList = make([]piece.Piece, len(p.PieceList))
	copy(newP.PieceList, p.PieceList)
	newP.PieceList[int(m.From)] = piece.Piece_None
	if m.PromotedTo == piece.Piece_None {
		newP.PieceList[int(m.To)] = m.Piece
	} else {
		newP.PieceList[int(m.To)] = m.PromotedTo
	}

	// Update SideToMove
	newP.WhitesTurn = !p.WhitesTurn

	// Update Castling
	newP.Castling = p.Castling

	// Update EnPassantSquare
	newP.EnPassantSquare = square.Square_Invalid

	// Update FullmoveCount
	newP.HalfmoveCount = p.HalfmoveCount
	if m.IsCapture || m.Piece == piece.Piece_WhitePawn {
		newP.HalfmoveCount = 0
	} else {
		newP.HalfmoveCount++
	}

	// Update FullmoveCount
	newP.FullmoveCount = p.FullmoveCount
	if !p.WhitesTurn {
		newP.FullmoveCount++
	}

	return newP
}
