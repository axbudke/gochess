package generation

import "gochess/pkg/notation"

func MakeMove(p *notation.Position, move notation.Move) *notation.Position {
	newP := &notation.Position{}

	// Update PieceList
	newP.PieceList = make(notation.PieceList, len(p.PieceList))
	copy(newP.PieceList, p.PieceList)
	newP.PieceList[int(move.From)] = notation.Piece_None
	if move.PromotedTo == notation.Piece_None {
		newP.PieceList[int(move.To)] = move.Piece
	} else {
		newP.PieceList[int(move.To)] = move.PromotedTo
	}

	// Update SideToMove
	newP.WhitesTurn = !p.WhitesTurn

	// Update Castling
	newP.Castling = p.Castling

	// Update EnPassantSquare
	newP.EnPassantSquare = notation.Square_Invalid

	// Update FullmoveCount
	newP.HalfmoveCount = p.HalfmoveCount
	if move.IsCapture || move.Piece == notation.Piece_WhitePawn {
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
