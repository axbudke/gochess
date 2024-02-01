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
	for _, pieceRow := range pieceList {
		for _, pieceVal := range pieceRow {
			switch pieceVal {
			case notation.PieceVal_BlackPawn:
				total += -1
			case notation.PieceVal_BlackBishop, notation.PieceVal_BlackKnight:
				total += -3
			case notation.PieceVal_BlackRook:
				total += -5
			case notation.PieceVal_BlackQueen:
				total += -9
			case notation.PieceVal_BlackKing:
				total += -200
			case notation.PieceVal_WhitePawn:
				total += 1
			case notation.PieceVal_WhiteBishop, notation.PieceVal_WhiteKnight:
				total += 3
			case notation.PieceVal_WhiteRook:
				total += 5
			case notation.PieceVal_WhiteQueen:
				total += 9
			case notation.PieceVal_WhiteKing:
				total += 200
			case notation.PieceVal_Empty:
			default:
			}
		}
	}
	return total
}
