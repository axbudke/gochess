package generation

import (
	"gochess/pkg/board"
)

func GenerateMoves(b board.BoardInterface) MoveList {

	// Generate sudo-legal moves
	moves := GenerateSudoLegalMoves(b)

	// Check that the king is not in check for any of those moves
	// TODO

	return moves
}

func GenerateSudoLegalMoves(b board.BoardInterface) MoveList {
	moves := MoveList{}

	for i, pieceVal := range b.PieceList() {
		fromSquare := board.Square(i)
		if b.IsWhitesTurn() {
			switch pieceVal {
			case board.PieceVal_WhitePawn:
				moves = append(moves, GeneratePawnMoves(b, fromSquare)...)
			case board.PieceVal_WhiteKnight:
				moves = append(moves, GenerateKnightMoves(b, fromSquare)...)
			case board.PieceVal_WhiteBishop:
				moves = append(moves, GenerateBishopMoves(b, fromSquare)...)
			case board.PieceVal_WhiteRook:
				moves = append(moves, GenerateRookMoves(b, fromSquare)...)
			case board.PieceVal_WhiteQueen:
				moves = append(moves, GenerateQueenMoves(b, fromSquare)...)
			case board.PieceVal_WhiteKing:
				moves = append(moves, GenerateKingMoves(b, fromSquare)...)
			}
		} else {
			switch pieceVal {
			case board.PieceVal_BlackPawn:
				moves = append(moves, GeneratePawnMoves(b, fromSquare)...)
			case board.PieceVal_BlackKnight:
				moves = append(moves, GenerateKnightMoves(b, fromSquare)...)
			case board.PieceVal_BlackBishop:
				moves = append(moves, GenerateBishopMoves(b, fromSquare)...)
			case board.PieceVal_BlackRook:
				moves = append(moves, GenerateRookMoves(b, fromSquare)...)
			case board.PieceVal_BlackQueen:
				moves = append(moves, GenerateQueenMoves(b, fromSquare)...)
			case board.PieceVal_BlackKing:
				moves = append(moves, GenerateKingMoves(b, fromSquare)...)
			}
		}
	}

	return moves
}

func GeneratePawnMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	f, r := fromSquare.FileRank()

	// Check pawn movements
	forwardMovements := []int{1}
	if r == board.Rank2 {
		forwardMovements = append(forwardMovements, 2)
	}
	for _, rp := range forwardMovements {
		move, err := GenerateMove(b, fromSquare, r+board.Rank(rp*inverter), f)
		if err != nil || move == nil {
			break
		}
		if !move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == board.PieceVal_WhitePawn) ||
				(r == 1 && move.Piece == board.PieceVal_BlackPawn) {
				GeneratePromotionMoves(b, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	// Check pawn captures
	for _, fp := range []int{1, -1} {
		move, err := GenerateMove(b, fromSquare, r+board.Rank(1*inverter), f+board.File(fp*inverter))
		if err != nil || move == nil {
			continue
		}
		if move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == board.PieceVal_WhitePawn) ||
				(r == 1 && move.Piece == board.PieceVal_BlackPawn) {
				GeneratePromotionMoves(b, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func GeneratePromotionMoves(b board.BoardInterface, move *Move) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	for _, promotionPiece := range board.PromotionPieceVals {
		promotionMove := &Move{
			BoardI:     b,
			From:       move.From,
			To:         move.To,
			Piece:      move.Piece,
			PromotedTo: promotionPiece * board.PieceVal(inverter),
		}
		moves = append(moves, promotionMove)
	}

	return moves
}

func GenerateKnightMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	return GenerateNoSlideMoves(b, fromSquare, knightPairs)
}

func GenerateKingMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	return GenerateNoSlideMoves(b, fromSquare, kingPairs)
}

func GenerateBishopMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	return GenerateSlideMoves(b, fromSquare, bishopPairs)
}

func GenerateRookMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	return GenerateSlideMoves(b, fromSquare, rookPairs)
}

func GenerateQueenMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	return GenerateSlideMoves(b, fromSquare, queenPairs)
}

var (
	knightPairs = []movementPair{{1, 2}, {2, 1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {-1, -2}, {-2, -1}}
	bishopPairs = []movementPair{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	rookPairs   = []movementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queenPairs  = []movementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	kingPairs   = []movementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)

type movementPair struct{ rp, fp int }

func GenerateNoSlideMoves(b board.BoardInterface, fromSquare board.Square, pairs []movementPair) MoveList {
	return GenerateMovementMoves(b, fromSquare, pairs, 1)
}

func GenerateSlideMoves(b board.BoardInterface, fromSquare board.Square, pairs []movementPair) MoveList {
	return GenerateMovementMoves(b, fromSquare, pairs, 7)
}

func GenerateMovementMoves(b board.BoardInterface, fromSquare board.Square, pairs []movementPair, slideCount int) MoveList {
	moves := MoveList{}

	// Check movements
	f, r := fromSquare.FileRank()
	for _, pair := range pairs {
		for i := 1; i <= slideCount; i++ {
			move, err := GenerateMove(b, fromSquare, r+board.Rank(pair.rp*i), f+board.File(pair.fp*i))
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

func GenerateMove(b board.BoardInterface, fromSquare board.Square, toR board.Rank, toF board.File) (move *Move, err error) {
	pVal := b.PieceList().PieceAt(fromSquare)

	// inverter used to determine same/opposite color
	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	// Check square is valid
	toSquare, err := board.NewSquare(toR, toF)
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
	} else if toSquarePieceVal*board.PieceVal(inverter) < 0 { // Opposite color piece
		move.IsCapture = true
		return move, nil
	} else if toSquarePieceVal*board.PieceVal(inverter) > 0 { // Same color piece
		return nil, nil
	}

	return nil, nil
}
