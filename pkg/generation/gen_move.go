package generation

import (
	"gochess/pkg/position"
)

func GenerateMoves(b position.BoardInterface) MoveList {

	// Generate sudo-legal moves
	moves := GenerateSudoLegalMoves(b)

	// Check that the king is not in check for any of those moves
	// TODO

	return moves
}

func GenerateSudoLegalMoves(b position.BoardInterface) MoveList {
	moves := MoveList{}

	for i, pieceVal := range b.PieceList() {
		fromSquare := position.Square(i)
		if b.IsWhitesTurn() {
			switch pieceVal {
			case position.PieceVal_WhitePawn:
				moves = append(moves, GeneratePawnMoves(b, fromSquare)...)
			case position.PieceVal_WhiteKnight:
				moves = append(moves, GenerateKnightMoves(b, fromSquare)...)
			case position.PieceVal_WhiteBishop:
				moves = append(moves, GenerateBishopMoves(b, fromSquare)...)
			case position.PieceVal_WhiteRook:
				moves = append(moves, GenerateRookMoves(b, fromSquare)...)
			case position.PieceVal_WhiteQueen:
				moves = append(moves, GenerateQueenMoves(b, fromSquare)...)
			case position.PieceVal_WhiteKing:
				moves = append(moves, GenerateKingMoves(b, fromSquare)...)
			}
		} else {
			switch pieceVal {
			case position.PieceVal_BlackPawn:
				moves = append(moves, GeneratePawnMoves(b, fromSquare)...)
			case position.PieceVal_BlackKnight:
				moves = append(moves, GenerateKnightMoves(b, fromSquare)...)
			case position.PieceVal_BlackBishop:
				moves = append(moves, GenerateBishopMoves(b, fromSquare)...)
			case position.PieceVal_BlackRook:
				moves = append(moves, GenerateRookMoves(b, fromSquare)...)
			case position.PieceVal_BlackQueen:
				moves = append(moves, GenerateQueenMoves(b, fromSquare)...)
			case position.PieceVal_BlackKing:
				moves = append(moves, GenerateKingMoves(b, fromSquare)...)
			}
		}
	}

	return moves
}

func GeneratePawnMoves(b position.BoardInterface, fromSquare position.Square) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	f, r := fromSquare.FileRank()

	// Check pawn movements
	forwardMovements := []int{1}
	if r == position.Rank2 {
		forwardMovements = append(forwardMovements, 2)
	}
	for _, rp := range forwardMovements {
		move, err := GenerateMove(b, fromSquare, r+position.Rank(rp*inverter), f)
		if err != nil || move == nil {
			break
		}
		if !move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == position.PieceVal_WhitePawn) ||
				(r == 1 && move.Piece == position.PieceVal_BlackPawn) {
				GeneratePromotionMoves(b, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	// Check pawn captures
	for _, fp := range []int{1, -1} {
		move, err := GenerateMove(b, fromSquare, r+position.Rank(1*inverter), f+position.File(fp*inverter))
		if err != nil || move == nil {
			continue
		}
		if move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == position.PieceVal_WhitePawn) ||
				(r == 1 && move.Piece == position.PieceVal_BlackPawn) {
				GeneratePromotionMoves(b, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func GeneratePromotionMoves(b position.BoardInterface, move *Move) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	for _, promotionPiece := range position.PromotionPieceVals {
		promotionMove := &Move{
			BoardI:     b,
			From:       move.From,
			To:         move.To,
			Piece:      move.Piece,
			PromotedTo: promotionPiece * position.PieceVal(inverter),
		}
		moves = append(moves, promotionMove)
	}

	return moves
}

func GenerateKnightMoves(b position.BoardInterface, fromSquare position.Square) MoveList {
	return GenerateNoSlideMoves(b, fromSquare, KnightMovementPairs)
}

func GenerateKingMoves(b position.BoardInterface, fromSquare position.Square) MoveList {
	return GenerateNoSlideMoves(b, fromSquare, KingMovementPairs)
}

func GenerateBishopMoves(b position.BoardInterface, fromSquare position.Square) MoveList {
	return GenerateSlideMoves(b, fromSquare, BishopMovementPairs)
}

func GenerateRookMoves(b position.BoardInterface, fromSquare position.Square) MoveList {
	return GenerateSlideMoves(b, fromSquare, RookMovementPairs)
}

func GenerateQueenMoves(b position.BoardInterface, fromSquare position.Square) MoveList {
	return GenerateSlideMoves(b, fromSquare, QueenMovementPairs)
}

var (
	KnightMovementPairs = []MovementPair{{1, 2}, {2, 1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {-1, -2}, {-2, -1}}
	BishopMovementPairs = []MovementPair{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	RookMovementPairs   = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	QueenMovementPairs  = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	KingMovementPairs   = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)

type MovementPair struct{ RP, FP int }

func GenerateNoSlideMoves(b position.BoardInterface, fromSquare position.Square, pairs []MovementPair) MoveList {
	return GenerateMovementMoves(b, fromSquare, pairs, 1)
}

func GenerateSlideMoves(b position.BoardInterface, fromSquare position.Square, pairs []MovementPair) MoveList {
	return GenerateMovementMoves(b, fromSquare, pairs, 7)
}

func GenerateMovementMoves(b position.BoardInterface, fromSquare position.Square, pairs []MovementPair, slideCount int) MoveList {
	moves := MoveList{}

	// Check movements
	f, r := fromSquare.FileRank()
	for _, pair := range pairs {
		for i := 1; i <= slideCount; i++ {
			move, err := GenerateMove(b, fromSquare, r+position.Rank(pair.RP*i), f+position.File(pair.FP*i))
			if err != nil || move == nil {
				break
			}
			moves = append(moves, move)
			if move.IsCapture {
				break
			}
		}
	}

	return moves
}

func GenerateMove(b position.BoardInterface, fromSquare position.Square, toR position.Rank, toF position.File) (move *Move, err error) {
	pVal := b.PieceList().PieceAt(fromSquare)

	// inverter used to determine same/opposite color
	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	// Check square is valid
	toSquare, err := position.NewSquare(toR, toF)
	if err != nil {
		return nil, err
	}

	// Create basic move
	move = &Move{
		From:  fromSquare,
		To:    toSquare,
		Piece: pVal,
	}

	// Check what piece is on toSquare
	toSquarePieceVal := b.PieceList().PieceAt(toSquare)
	if toSquarePieceVal == 0 { // Empty Square
		// if pVal == board.PieceVal_WhitePawn && toSquare == b.EnPassantSquare() {
		// 	move.IsCapture = true
		// 	return move, nil
		// }
		return move, nil
	} else if toSquarePieceVal*position.PieceVal(inverter) < 0 { // Opposite color piece
		move.IsCapture = true
		return move, nil
	} else if toSquarePieceVal*position.PieceVal(inverter) > 0 { // Same color piece
		return nil, nil
	}

	return nil, nil
}