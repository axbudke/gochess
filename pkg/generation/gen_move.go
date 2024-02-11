package generation

import (
	"gochess/pkg/notation"
	"math"
)

func GenerateMoves(p *notation.Position) notation.MoveList {
	// Find King
	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}
	var kingSquare notation.Square
	for squareInt, piece := range p.PieceList {
		if piece == notation.Piece_WhiteKing*notation.Piece(inverter) {
			kingSquare = notation.Square(squareInt)
		}
	}

	// Generate all Checks on King
	checkMoves, pinnedPieces := GenerateChecksAndPins(p, kingSquare)

	if len(checkMoves) >= 2 {
		// Double Check, Only King moves are valid
		return GenerateKingMoves(p, kingSquare)
	}

	// Generate pseudo legal moves
	psuedoLegalMoves := GeneratePseudoLegalMoves(p)

	// Filter moves of pinned pieces
	for _, pinnedPiece := range pinnedPieces {
		psuedoLegalMoves = psuedoLegalMoves.FindMovesNotFrom(pinnedPiece)
	}

	if len(checkMoves) == 1 {
		// Single Check, Only King moves or blocking moves

		// Pre-allocate move list
		moves := make(notation.MoveList, 0, len(psuedoLegalMoves))

		// King Moves
		moves = append(moves, psuedoLegalMoves.FindMovesFrom(kingSquare)...)

		// Moves to capture checking piece
		moves = append(moves, psuedoLegalMoves.FindMovesTo(checkMoves[0].From)...)

		// Blocking Moves
		for _, square := range notation.SquaresInBetween(kingSquare, checkMoves[0].From) {
			moves = append(moves, psuedoLegalMoves.FindMovesTo(square)...)
		}

		return moves
	}

	return psuedoLegalMoves
}

func GenerateChecksAndPins(p *notation.Position, kingSquare notation.Square) (notation.MoveList, []notation.Square) {
	checkMoves := notation.MoveList{}
	pinnedSquares := []notation.Square{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	f, r := kingSquare.FileRank()

	// Find Knight/Pawn checks
	for _, pair := range append(KnightMovementPairs, []MovementPair{{1 * inverter, 1}, {1 * inverter, -1}}...) {
		// Check square is valid
		fromSquare, err := notation.NewSquareCheck(f+notation.File(pair.FP), r+notation.Rank(pair.RP))
		if err != nil {
			continue
		}

		// Create Move
		move := GenerateMove(p, fromSquare, kingSquare)
		if move == nil {
			continue
		}

		// Add move
		if move.IsCapture {
			checkMoves = append(checkMoves, move)
		}
	}

	// Find Bishop/Queen checks and pinned pieces
	for _, pair := range BishopMovementPairs {
		var possiblePinnedPiece *notation.Square
		for i := 1; i <= 7; i++ {
			// Check square is valid
			fromSquare, err := notation.NewSquareCheck(f+notation.File(pair.FP*i), r+notation.Rank(pair.RP*i))
			if err != nil {
				break
			}

			// Create move
			move := GenerateMove(p, fromSquare, kingSquare)
			if move == nil {
				possiblePinnedPiece = &fromSquare
				continue
			}

			// Add move
			pieceAbsVal := notation.Piece(math.Abs(float64(move.Piece)))
			// TODO fix
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

// ==================== Pseudo-Legal Moves ====================

func GeneratePseudoLegalMoves(p *notation.Position) notation.MoveList {
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
		toSquare, err := notation.NewSquareCheck(f, r+notation.Rank(rp*inverter))
		if err != nil {
			break
		}

		// Check type of move
		move := GenerateMove(p, fromSquare, toSquare)
		if move == nil {
			break
		}

		// Add move
		if !move.IsCapture {
			// Check for promotion
			_, r := move.To.FileRank()
			if (r == 8 && move.Piece == notation.Piece_WhitePawn) ||
				(r == 1 && move.Piece == notation.Piece_BlackPawn) {
				moves = append(moves, GeneratePromotionMoves(p, move)...)
			} else {
				moves = append(moves, move)
			}
		}
	}

	// Check pawn captures
	for _, fp := range []int{1, -1} {
		// Check square is valid
		toSquare, err := notation.NewSquareCheck(f+notation.File(fp*inverter), r+notation.Rank(1*inverter))
		if err != nil {
			continue
		}

		// Check type of move
		move := GenerateMove(p, fromSquare, toSquare)
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

func GenerateKingMoves(p *notation.Position, kingSquare notation.Square) notation.MoveList {
	moves := GenerateNoSlideMoves(p, kingSquare, KingMovementPairs)

	if !(p.WhitesTurn && kingSquare == notation.Square_e1) && !(!p.WhitesTurn && kingSquare == notation.Square_e8) {
		// King has moved off starting square
		return moves
	}
	if !p.CanCastle(p.WhitesTurn, true) && !p.CanCastle(p.WhitesTurn, false) {
		// Can't castle either way
		return moves
	}

	// Castling moves
	f, r := kingSquare.FileRank()
	for _, pair := range []MovementPair{{2, 0}, {-2, 0}} {
		// Is valid square
		toSquare, err := notation.NewSquareCheck(f+notation.File(pair.FP), r+notation.Rank(pair.RP))
		if err != nil {
			break
		}

		// Create move
		move := GenerateMove(p, kingSquare, toSquare)
		if move == nil {
			break
		}

		// Add move
		moves = append(moves, move)
		if move.IsCapture {
			break
		}
	}

	return moves
}

func GenerateKnightMoves(p *notation.Position, fromSquare notation.Square) notation.MoveList {
	return GenerateNoSlideMoves(p, fromSquare, KnightMovementPairs)
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
	KnightMovementPairs   = []MovementPair{{1, 2}, {2, 1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {-1, -2}, {-2, -1}}
	BishopMovementPairs   = []MovementPair{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	RookMovementPairs     = []MovementPair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	QueenMovementPairs    = append(BishopMovementPairs, RookMovementPairs...)
	KingMovementPairs     = QueenMovementPairs
	KingCastlingMovements = []MovementPair{{2, 0}, {-2, 0}}
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

	// Movements
	f, r := fromSquare.FileRank()
	for _, pair := range pairs {
		for i := 1; i <= slideCount; i++ {
			// Is valid square
			toSquare, err := notation.NewSquareCheck(f+notation.File(pair.FP*i), r+notation.Rank(pair.RP*i))
			if err != nil {
				break
			}

			// Create move
			move := GenerateMove(p, fromSquare, toSquare)
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

func GenerateMove(p *notation.Position, fromSquare, toSquare notation.Square) *notation.Move {
	pieceIsSameColor := p.PieceAtIsSame(toSquare)
	if pieceIsSameColor {
		// Same color piece
		return nil
	}

	fromPiece := p.PieceAt(fromSquare)

	// Create move
	move := &notation.Move{
		PieceList: p.PieceList,
		From:      fromSquare,
		To:        toSquare,
		Piece:     fromPiece,
	}

	if !pieceIsSameColor {
		// Opposite color piece
		move.IsCapture = true
		return move
	}
	// Empty Square

	// Check for EnPassant
	if fromPiece.IsPawn() && toSquare == p.EnPassantSquare {
		move.IsCapture = true
		return move
	}

	return move
}
