package generation

import (
	"math"

	"gochess/pkg/notation/move"
	"gochess/pkg/notation/piece"
	"gochess/pkg/notation/position"
	"gochess/pkg/notation/square"
)

func GenerateMoves(p *position.Position) (moveList move.MoveList) {
	defer func() {
		moveList.Sort()
	}()

	// Find King
	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}
	var kingSquare square.Square
	for squareInt, pc := range p.PieceList {
		if pc == piece.Piece_WhiteKing*piece.Piece(inverter) {
			kingSquare = square.Square(squareInt)
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
		moves := make(move.MoveList, 0, len(psuedoLegalMoves))

		// King Moves
		moves = append(moves, psuedoLegalMoves.FindMovesFrom(kingSquare)...)

		// Moves to capture checking piece
		moves = append(moves, psuedoLegalMoves.FindMovesTo(checkMoves[0].From)...)

		// Blocking Moves
		for _, square := range square.SquaresInBetween(kingSquare, checkMoves[0].From) {
			moves = append(moves, psuedoLegalMoves.FindMovesTo(square)...)
		}

		return moves
	}

	return psuedoLegalMoves
}

func GenerateChecksAndPins(p *position.Position, kingSquare square.Square) (move.MoveList, []square.Square) {
	checkMoves := move.MoveList{}
	pinnedSquares := []square.Square{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	f, r := kingSquare.FileRank()

	// Find Knight/Pawn checks
	for _, pair := range append(KnightMovementPairs, []MovementPair{{1 * inverter, 1}, {1 * inverter, -1}}...) {
		// Check square is valid
		fromSquare, err := square.NewSquareCheck(f+square.File(pair.FP), r+square.Rank(pair.RP))
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
		var possiblePinnedPiece *square.Square
		for i := 1; i <= 7; i++ {
			// Check square is valid
			fromSquare, err := square.NewSquareCheck(f+square.File(pair.FP*i), r+square.Rank(pair.RP*i))
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
			pieceAbsVal := piece.Piece(math.Abs(float64(move.Piece)))
			// TODO fix
			if move.IsCapture && (pieceAbsVal == piece.Piece_WhiteBishop || pieceAbsVal == piece.Piece_WhiteQueen) {
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

func GeneratePseudoLegalMoves(p *position.Position) move.MoveList {
	moves := move.MoveList{}

	for i, pieceVal := range p.PieceList {
		fromSquare := square.Square(i)
		if p.WhitesTurn {
			switch pieceVal {
			case piece.Piece_WhitePawn:
				moves = append(moves, GeneratePawnMoves(p, fromSquare)...)
			case piece.Piece_WhiteKnight:
				moves = append(moves, GenerateKnightMoves(p, fromSquare)...)
			case piece.Piece_WhiteBishop:
				moves = append(moves, GenerateBishopMoves(p, fromSquare)...)
			case piece.Piece_WhiteRook:
				moves = append(moves, GenerateRookMoves(p, fromSquare)...)
			case piece.Piece_WhiteQueen:
				moves = append(moves, GenerateQueenMoves(p, fromSquare)...)
			case piece.Piece_WhiteKing:
				moves = append(moves, GenerateKingMoves(p, fromSquare)...)
			}
		} else {
			switch pieceVal {
			case piece.Piece_BlackPawn:
				moves = append(moves, GeneratePawnMoves(p, fromSquare)...)
			case piece.Piece_BlackKnight:
				moves = append(moves, GenerateKnightMoves(p, fromSquare)...)
			case piece.Piece_BlackBishop:
				moves = append(moves, GenerateBishopMoves(p, fromSquare)...)
			case piece.Piece_BlackRook:
				moves = append(moves, GenerateRookMoves(p, fromSquare)...)
			case piece.Piece_BlackQueen:
				moves = append(moves, GenerateQueenMoves(p, fromSquare)...)
			case piece.Piece_BlackKing:
				moves = append(moves, GenerateKingMoves(p, fromSquare)...)
			}
		}
	}

	return moves
}

func GeneratePawnMoves(p *position.Position, fromSquare square.Square) move.MoveList {
	moves := move.MoveList{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	f, r := fromSquare.FileRank()

	// Check pawn movements
	forwardMovements := []int{1}
	if (p.WhitesTurn && r == square.Rank2) || (!p.WhitesTurn && r == square.Rank7) {
		forwardMovements = append(forwardMovements, 2)
	}
	for _, rp := range forwardMovements {
		// Check square is valid
		toSquare, err := square.NewSquareCheck(f, r+square.Rank(rp*inverter))
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
			if (r == 8 && move.Piece == piece.Piece_WhitePawn) ||
				(r == 1 && move.Piece == piece.Piece_BlackPawn) {
				moves = append(moves, GeneratePromotionMoves(p, move)...)
			} else {
				moves = append(moves, move)
			}
		}
	}

	// Check pawn captures
	for _, fp := range []int{1, -1} {
		// Check square is valid
		toSquare, err := square.NewSquareCheck(f+square.File(fp*inverter), r+square.Rank(1*inverter))
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
			if (r == 8 && move.Piece == piece.Piece_WhitePawn) ||
				(r == 1 && move.Piece == piece.Piece_BlackPawn) {
				GeneratePromotionMoves(p, move)
			} else {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func GeneratePromotionMoves(p *position.Position, m *move.Move) move.MoveList {
	moves := move.MoveList{}

	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	for _, promotionPiece := range piece.PromotionPieces {
		promotionMove := &move.Move{
			PieceList:  m.PieceList,
			From:       m.From,
			To:         m.To,
			Piece:      m.Piece,
			PromotedTo: promotionPiece * piece.Piece(inverter),
		}
		moves = append(moves, promotionMove)
	}

	return moves
}

func GenerateKingMoves(p *position.Position, kingSquare square.Square) move.MoveList {
	moves := GenerateNoSlideMoves(p, kingSquare, KingMovementPairs)

	if !(p.WhitesTurn && kingSquare == square.Square_e1) && !(!p.WhitesTurn && kingSquare == square.Square_e8) {
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
		toSquare, err := square.NewSquareCheck(f+square.File(pair.FP), r+square.Rank(pair.RP))
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

func GenerateKnightMoves(p *position.Position, fromSquare square.Square) move.MoveList {
	return GenerateNoSlideMoves(p, fromSquare, KnightMovementPairs)
}

func GenerateBishopMoves(p *position.Position, fromSquare square.Square) move.MoveList {
	return GenerateSlideMoves(p, fromSquare, BishopMovementPairs)
}

func GenerateRookMoves(p *position.Position, fromSquare square.Square) move.MoveList {
	return GenerateSlideMoves(p, fromSquare, RookMovementPairs)
}

func GenerateQueenMoves(p *position.Position, fromSquare square.Square) move.MoveList {
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

func GenerateNoSlideMoves(p *position.Position, fromSquare square.Square, pairs []MovementPair) move.MoveList {
	return GenerateMovementMoves(p, fromSquare, pairs, 1)
}

func GenerateSlideMoves(p *position.Position, fromSquare square.Square, pairs []MovementPair) move.MoveList {
	return GenerateMovementMoves(p, fromSquare, pairs, 7)
}

func GenerateMovementMoves(p *position.Position, fromSquare square.Square, pairs []MovementPair, slideCount int) move.MoveList {
	moves := move.MoveList{}

	// Movements
	f, r := fromSquare.FileRank()
	for _, pair := range pairs {
		for i := 1; i <= slideCount; i++ {
			// Is valid square
			toSquare, err := square.NewSquareCheck(f+square.File(pair.FP*i), r+square.Rank(pair.RP*i))
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

func GenerateMove(p *position.Position, fromSquare, toSquare square.Square) *move.Move {
	// inverter used to determine same/opposite color
	inverter := func() int {
		if p.WhitesTurn {
			return 1
		} else {
			return -1
		}
	}()

	pc := p.PieceAt(toSquare)
	if pc*piece.Piece(inverter) > 0 {
		// Same color piece
		return nil
	}

	fromPiece := p.PieceAt(fromSquare)

	// Create move
	move := &move.Move{
		PieceList: p.PieceList,
		From:      fromSquare,
		To:        toSquare,
		Piece:     fromPiece,
	}

	if pc*piece.Piece(inverter) < 0 {
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
