package evaluation

import (
	"gochess/pkg/notation/piece"
	"gochess/pkg/notation/position"
)

func Evaluate(p *position.Position) int {
	var total int
	total += GetMaterialCount(p)
	return total
}

func GetMaterialCount(p *position.Position) int {
	var total int
	for _, pc := range p.PieceList {
		switch pc {
		case piece.Piece_BlackPawn:
			total += -1
		case piece.Piece_BlackBishop, piece.Piece_BlackKnight:
			total += -3
		case piece.Piece_BlackRook:
			total += -5
		case piece.Piece_BlackQueen:
			total += -9
		case piece.Piece_BlackKing:
			total += -200
		case piece.Piece_WhitePawn:
			total += 1
		case piece.Piece_WhiteBishop, piece.Piece_WhiteKnight:
			total += 3
		case piece.Piece_WhiteRook:
			total += 5
		case piece.Piece_WhiteQueen:
			total += 9
		case piece.Piece_WhiteKing:
			total += 200
		case piece.Piece_None:
		default:
		}
	}
	return total
}
