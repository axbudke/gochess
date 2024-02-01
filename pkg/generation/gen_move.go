package generation

import (
	"gochess/pkg/position"
)

func GenerateMoves(p *position.Position) MoveList {

	// Generate sudo-legal moves
	moves := GenerateSudoLegalMoves(p)

	// Check that the king is not in check for any of those moves
	// TODO

	return moves
}

func GenerateSudoLegalMoves(p *position.Position) MoveList {
	moves := MoveList{}

	for i, pieceVal := range p.PieceList() {
		fromSquare := position.Square(i)
		if p.IsWhitesTurn() {
			switch pieceVal {
			case position.PieceVal_WhitePawn:
				moves = append(moves, GeneratePawnMoves(p, fromSquare)...)
			case position.PieceVal_WhiteKnight:
				moves = append(moves, GenerateKnightMoves(p, fromSquare)...)
			case position.PieceVal_WhiteBishop:
				moves = append(moves, GenerateBishopMoves(p, fromSquare)...)
			case position.PieceVal_WhiteRook:
				moves = append(moves, GenerateRookMoves(p, fromSquare)...)
			case position.PieceVal_WhiteQueen:
				moves = append(moves, GenerateQueenMoves(p, fromSquare)...)
			case position.PieceVal_WhiteKing:
				moves = append(moves, GenerateKingMoves(p, fromSquare)...)
			}
		} else {
			switch pieceVal {
			case position.PieceVal_BlackPawn:
				moves = append(moves, GeneratePawnMoves(p, fromSquare)...)
			case position.PieceVal_BlackKnight:
				moves = append(moves, GenerateKnightMoves(p, fromSquare)...)
			case position.PieceVal_BlackBishop:
				moves = append(moves, GenerateBishopMoves(p, fromSquare)...)
			case position.PieceVal_BlackRook:
				moves = append(moves, GenerateRookMoves(p, fromSquare)...)
			case position.PieceVal_BlackQueen:
				moves = append(moves, GenerateQueenMoves(p, fromSquare)...)
			case position.PieceVal_BlackKing:
				moves = append(moves, GenerateKingMoves(p, fromSquare)...)
			}
		}
	}

	return moves
}

func GeneratePawnMoves(p *position.Position, fromSquare position.Square) MoveList {
	moves := MoveList{}

	inverter := 1
	if !p.IsWhitesTurn() {
		inverter = -1
	}

	f, r := fromSquare.FileRank()

	// Check pawn movements
	forwardMovements := []int{1}
	if r == position.Rank2 {
		forwardMovements = append(forwardMovements, 2)
	}
	for _, rp := range forwardMovements {
		move, err := GenerateMove(p, fromSquare, r+position.Rank(rp*inverter), f)
		if err != nil || move == nil {
			break
		}
		if !move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == position.PieceVal_WhitePawn) ||
				(r == 1 && move.Piece == position.PieceVal_BlackPawn) {
				GeneratePromotionMoves(p, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	// Check pawn captures
	for _, fp := range []int{1, -1} {
		move, err := GenerateMove(p, fromSquare, r+position.Rank(1*inverter), f+position.File(fp*inverter))
		if err != nil || move == nil {
			continue
		}
		if move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == position.PieceVal_WhitePawn) ||
				(r == 1 && move.Piece == position.PieceVal_BlackPawn) {
				GeneratePromotionMoves(p, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func GeneratePromotionMoves(p *position.Position, move *Move) MoveList {
	moves := MoveList{}

	inverter := 1
	if !p.IsWhitesTurn() {
		inverter = -1
	}

	for _, promotionPiece := range position.PromotionPieceVals {
		promotionMove := &Move{
			PieceList:  move.PieceList,
			From:       move.From,
			To:         move.To,
			Piece:      move.Piece,
			PromotedTo: promotionPiece * position.PieceVal(inverter),
		}
		moves = append(moves, promotionMove)
	}

	return moves
}

func GenerateKnightMoves(p *position.Position, fromSquare position.Square) MoveList {
	return GenerateNoSlideMoves(p, fromSquare, KnightMovementPairs)
}

func GenerateKingMoves(p *position.Position, fromSquare position.Square) MoveList {
	return GenerateNoSlideMoves(p, fromSquare, KingMovementPairs)
}

func GenerateBishopMoves(p *position.Position, fromSquare position.Square) MoveList {
	return GenerateSlideMoves(p, fromSquare, BishopMovementPairs)
}

func GenerateRookMoves(p *position.Position, fromSquare position.Square) MoveList {
	return GenerateSlideMoves(p, fromSquare, RookMovementPairs)
}

func GenerateQueenMoves(p *position.Position, fromSquare position.Square) MoveList {
	return GenerateSlideMoves(p, fromSquare, QueenMovementPairs)
}

var (
	KnightMovementPairs = []MovementPair{{1, 2}, {2, 1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {-1, -2}, {-2, -1}}
	BishopMovementPairs = []MovementPair{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	RookMovementPairs   = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	QueenMovementPairs  = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	KingMovementPairs   = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)

type MovementPair struct{ RP, FP int }

func GenerateNoSlideMoves(p *position.Position, fromSquare position.Square, pairs []MovementPair) MoveList {
	return GenerateMovementMoves(p, fromSquare, pairs, 1)
}

func GenerateSlideMoves(p *position.Position, fromSquare position.Square, pairs []MovementPair) MoveList {
	return GenerateMovementMoves(p, fromSquare, pairs, 7)
}

func GenerateMovementMoves(p *position.Position, fromSquare position.Square, pairs []MovementPair, slideCount int) MoveList {
	moves := MoveList{}

	// Check movements
	f, r := fromSquare.FileRank()
	for _, pair := range pairs {
		for i := 1; i <= slideCount; i++ {
			move, err := GenerateMove(p, fromSquare, r+position.Rank(pair.RP*i), f+position.File(pair.FP*i))
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

func GenerateMove(p *position.Position, fromSquare position.Square, toR position.Rank, toF position.File) (move *Move, err error) {
	pVal := p.PieceList().PieceAt(fromSquare)

	// inverter used to determine same/opposite color
	inverter := 1
	if !p.IsWhitesTurn() {
		inverter = -1
	}

	// Check square is valid
	toSquare, err := position.NewSquare(toR, toF)
	if err != nil {
		return nil, err
	}

	// Create basic move
	move = &Move{
		PieceList: p.PieceList(),
		From:      fromSquare,
		To:        toSquare,
		Piece:     pVal,
	}

	// Check what piece is on toSquare
	toSquarePieceVal := p.PieceList().PieceAt(toSquare)
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
