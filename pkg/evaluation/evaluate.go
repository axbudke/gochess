package evaluation

import (
	"gochess/pkg/position"
)

func Evaluate(b position.BoardInterface) int {
	var total int
	total += GetMaterialCount(b)
	return total
}

func GetMaterialCount(b position.BoardInterface) int {
	var total int
	pieceList := position.PieceList{}
	for _, pieceRow := range pieceList {
		for _, pieceVal := range pieceRow {
			switch pieceVal {
			case position.PieceVal_BlackPawn:
				total += -1
			case position.PieceVal_BlackBishop, position.PieceVal_BlackKnight:
				total += -3
			case position.PieceVal_BlackRook:
				total += -5
			case position.PieceVal_BlackQueen:
				total += -9
			case position.PieceVal_BlackKing:
				total += -200
			case position.PieceVal_WhitePawn:
				total += 1
			case position.PieceVal_WhiteBishop, position.PieceVal_WhiteKnight:
				total += 3
			case position.PieceVal_WhiteRook:
				total += 5
			case position.PieceVal_WhiteQueen:
				total += 9
			case position.PieceVal_WhiteKing:
				total += 200
			case position.PieceVal_Empty:
			default:
			}
		}
	}
	return total
}
