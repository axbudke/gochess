package evaluation

import (
	"gochess/pkg/notation"
)

func Evaluate(b *notation.Position) int {
	var total int
	total += GetMaterialCount(b)
	return total
}

func GetMaterialCount(b *notation.Position) int {
	var total int
	pieceList := notation.PieceList{}
	for _, piece := range pieceList {
		switch piece {
		case notation.Piece_BlackPawn:
			total += -1
		case notation.Piece_BlackBishop, notation.Piece_BlackKnight:
			total += -3
		case notation.Piece_BlackRook:
			total += -5
		case notation.Piece_BlackQueen:
			total += -9
		case notation.Piece_BlackKing:
			total += -200
		case notation.Piece_WhitePawn:
			total += 1
		case notation.Piece_WhiteBishop, notation.Piece_WhiteKnight:
			total += 3
		case notation.Piece_WhiteRook:
			total += 5
		case notation.Piece_WhiteQueen:
			total += 9
		case notation.Piece_WhiteKing:
			total += 200
		case notation.Piece_Empty:
		default:
		}
	}
	return total
}
