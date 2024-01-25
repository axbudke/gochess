package evaluation

import (
	"gochess/pkg/board"
)

func Evaluate(b board.BoardInterface) int {
	var total int
	total += GetMaterialCount(b)
	return total
}

func GetMaterialCount(b board.BoardInterface) int {
	var total int
	pieceList := board.PieceList{}
	for _, pieceRow := range pieceList {
		for _, pieceVal := range pieceRow {
			switch pieceVal {
			case board.PieceVal_BlackPawn:
				total += -1
			case board.PieceVal_BlackBishop, board.PieceVal_BlackKnight:
				total += -3
			case board.PieceVal_BlackRook:
				total += -5
			case board.PieceVal_BlackQueen:
				total += -9
			case board.PieceVal_BlackKing:
				total += -200
			case board.PieceVal_WhitePawn:
				total += 1
			case board.PieceVal_WhiteBishop, board.PieceVal_WhiteKnight:
				total += 3
			case board.PieceVal_WhiteRook:
				total += 5
			case board.PieceVal_WhiteQueen:
				total += 9
			case board.PieceVal_WhiteKing:
				total += 200
			case board.PieceVal_Empty:
			default:
			}
		}
	}
	return total
}
