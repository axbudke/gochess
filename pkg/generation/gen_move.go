package generation

import (
	"gochess/pkg/notation"
	"math"
)

func GenerateMoves(p *notation.Position) notation.MoveList {

	// Generate sudo-legal moves
	moves := GenerateSudoLegalMoves(p)

	// Check that the king is not in check for any of those moves
	// TODO
	// checkMoves, pinnedSquares := FindReverseKingMoves(p)

	return moves
}

func FindReverseKingMoves(p *notation.Position) (notation.MoveList, []notation.Square) {
	checkMoves := notation.MoveList{}
	pinnedSquares := []notation.Square{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	// Find King
	var kingSquare notation.Square
	for squareInt, piece := range p.PieceList {
		if piece == notation.Piece_WhiteKing*notation.Piece(inverter) {
			kingSquare = notation.Square(squareInt)
		}
	}
	f, r := kingSquare.FileRank()

	// Find knight checks
	for _, pair := range KnightMovementPairs {
		// Check square is valid
		fromSquare, err := notation.NewSquare(r+notation.Rank(pair.RP), f+notation.File(pair.FP))
		if err != nil {
			continue
		}

		// Check type of move
		move := GenerateReverseMove(p, kingSquare, fromSquare)
		if move == nil {
			continue
		}

		// Add move
		if move.IsCapture {
			checkMoves = append(checkMoves, move)
		}
	}

	// Find Pawn checks

	// Find Bishop or Queen checks and pinned pieces
	for _, pair := range BishopMovementPairs {
		var possiblePinnedPiece *notation.Square
		for i := 1; i <= 7; i++ {
			// Check square is valid
			fromSquare, err := notation.NewSquare(r+notation.Rank(pair.RP*i), f+notation.File(pair.FP*i))
			if err != nil {
				break
			}

			// Check type of move
			move := GenerateReverseMove(p, kingSquare, fromSquare)
			if move == nil {
				possiblePinnedPiece = &fromSquare
				break
			}

			// Add move
			pieceAbsVal := notation.Piece(math.Abs(float64(move.Piece)))
			if move.IsCapture && (pieceAbsVal == notation.Piece_WhiteBishop || pieceAbsVal == notation.Piece_WhiteQueen) {
				if possiblePinnedPiece != nil {
					pinnedSquares = append(pinnedSquares, *possiblePinnedPiece)
				} else {
					checkMoves = append(checkMoves, move)
					break
				}
			}
		}
	}

	// Find Rook or Queen checks

	return checkMoves, pinnedSquares
}

// ==================== Sudo-Legal Moves ====================

func GenerateSudoLegalMoves(p *notation.Position) notation.MoveList {
	moves := notation.MoveList{}

	for i, pieceVal := range p.PieceList {
		fromSquare := notation.Square(i)
		if p.WhitesTurn {
			switch pieceVal {
			case notation.Piece_WhitePawn:
				moves = append(moves, GeneratePawnMoves(p, fromSquare)...)
			case notation.Piece_WhiteKnight:
				moves = append(moves, GenerateKnightMoves(p, fromSquare)...)
			case notation.Piece_WhiteBishop:
				moves = append(moves, GenerateBishopMoves(p, fromSquare)...)
			case notation.Piece_WhiteRook:
				moves = append(moves, GenerateRookMoves(p, fromSquare)...)
			case notation.Piece_WhiteQueen:
				moves = append(moves, GenerateQueenMoves(p, fromSquare)...)
			case notation.Piece_WhiteKing:
				moves = append(moves, GenerateKingMoves(p, fromSquare)...)
			}
		} else {
			switch pieceVal {
			case notation.Piece_BlackPawn:
				moves = append(moves, GeneratePawnMoves(p, fromSquare)...)
			case notation.Piece_BlackKnight:
				moves = append(moves, GenerateKnightMoves(p, fromSquare)...)
			case notation.Piece_BlackBishop:
				moves = append(moves, GenerateBishopMoves(p, fromSquare)...)
			case notation.Piece_BlackRook:
				moves = append(moves, GenerateRookMoves(p, fromSquare)...)
			case notation.Piece_BlackQueen:
				moves = append(moves, GenerateQueenMoves(p, fromSquare)...)
			case notation.Piece_BlackKing:
				moves = append(moves, GenerateKingMoves(p, fromSquare)...)
			}
		}
	}

	return moves
}

func GeneratePawnMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	moves := notation.MoveList{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	f, r := fromSquare.FileRank()

	// Check pawn movements
	forwardMovements := []int{1}
	if r == notation.Rank2 {
		forwardMovements = append(forwardMovements, 2)
	}
	for _, rp := range forwardMovements {
		// Check square is valid
		toSquare, err := notation.NewSquare(r+notation.Rank(rp*inverter), f)
		if err != nil {
			break
		}

		// Check type of move
		move := GenerateNormalMove(p, fromSquare, toSquare)
		if move == nil {
			break
		}

		// Add move
		if !move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == notation.Piece_WhitePawn) ||
				(r == 1 && move.Piece == notation.Piece_BlackPawn) {
				GeneratePromotionMoves(p, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	// Check pawn captures
	for _, fp := range []int{1, -1} {
		// Check square is valid
		toSquare, err := notation.NewSquare(r+notation.Rank(1*inverter), f+notation.File(fp*inverter))
		if err != nil {
			continue
		}

		// Check type of move
		move := GenerateNormalMove(p, fromSquare, toSquare)
		if move == nil {
			continue
		}

		// Add move
		if move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == notation.Piece_WhitePawn) ||
				(r == 1 && move.Piece == notation.Piece_BlackPawn) {
				GeneratePromotionMoves(p, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func GeneratePromotionMoves(p *notation.Position, move *notation.Move) notation.MoveList {
	moves := notation.MoveList{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	for _, promotionPiece := range notation.PromotionPieces {
		promotionMove := &notation.Move{
			PieceList:  move.PieceList,
			From:       move.From,
			To:         move.To,
			Piece:      move.Piece,
			PromotedTo: promotionPiece * notation.Piece(inverter),
		}
		moves = append(moves, promotionMove)
	}

	return moves
}

func GenerateKnightMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	return GenerateNoSlideMoves(p, fromSquare, KnightMovementPairs)
}

func GenerateKingMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	return GenerateNoSlideMoves(p, fromSquare, KingMovementPairs)
}

func GenerateBishopMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	return GenerateSlideMoves(p, fromSquare, BishopMovementPairs)
}

func GenerateRookMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	return GenerateSlideMoves(p, fromSquare, RookMovementPairs)
}

func GenerateQueenMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	return GenerateSlideMoves(p, fromSquare, QueenMovementPairs)
}

// ========================= Movement Moves ====================

var (
	KnightMovementPairs = []MovementPair{{1, 2}, {2, 1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {-1, -2}, {-2, -1}}
	BishopMovementPairs = []MovementPair{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	RookMovementPairs   = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	QueenMovementPairs  = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	KingMovementPairs   = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)

type MovementPair struct{ RP, FP int }

func GenerateNoSlideMoves(p *notation.Position, fromSquare notation.Square, pairs []MovementPair) notation.MoveList {
	return GenerateMovementMoves(p, fromSquare, pairs, 1)
}

func GenerateSlideMoves(p *notation.Position, fromSquare notation.Square, pairs []MovementPair) notation.MoveList {
	return GenerateMovementMoves(p, fromSquare, pairs, 7)
}

func GenerateMovementMoves(p *notation.Position, fromSquare notation.Square, pairs []MovementPair, slideCount int) notation.MoveList {
	moves := notation.MoveList{}

	// Check movements
	f, r := fromSquare.FileRank()
	for _, pair := range pairs {
		for i := 1; i <= slideCount; i++ {
			// Check square is valid
			toSquare, err := notation.NewSquare(r+notation.Rank(pair.RP*i), f+notation.File(pair.FP*i))
			if err != nil {
				break
			}

			// Check type of move
			move := GenerateNormalMove(p, fromSquare, toSquare)
			if move == nil {
				break
			}

			// Add move
			moves = append(moves, move)
			if move.IsCapture {
				break
			}
		}
	}

	return moves
}

// ==================== Basic Generate Move ====================

func GenerateNormalMove(p *notation.Position, fromSquare, toSquare notation.Square) (move *notation.Move) {
	return GenerateMove(p, fromSquare, toSquare, false)
}

func GenerateReverseMove(p *notation.Position, toSquare, fromSquare notation.Square) (move *notation.Move) {
	return GenerateMove(p, toSquare, fromSquare, true)
}

func GenerateMove(p *notation.Position, square, newSquare notation.Square, reverse bool) (move *notation.Move) {
	// Check if square is same color piece
	pieceIsSameColor := p.PieceAtIsSame(newSquare)
	if pieceIsSameColor {
		return nil
	} else {
		var fromSquare, toSquare notation.Square
		if reverse {
			fromSquare = newSquare
			toSquare = square
		} else {
			fromSquare = square
			toSquare = newSquare
		}
		fromPiece := p.PieceAt(fromSquare)

		// Create basic move
		move = &notation.Move{
			PieceList: p.PieceList,
			From:      fromSquare,
			To:        toSquare,
			Piece:     fromPiece,
		}

		// Check if square is empty or opposite color piece
		if !pieceIsSameColor { // Opposite color piece
			move.IsCapture = true
			return move
		} else { // Empty Square
			if fromPiece.IsPawn() && toSquare == p.EnPassantSquare {
				move.IsCapture = true
				return move
			}
			return move
		}
	}
}
