package generation

import (
	"gochess/pkg/notation"
)

func GenerateMoves(p *notation.Position) notation.MoveList {

	// Generate sudo-legal moves
	moves := GenerateSudoLegalMoves(p)

	// Check that the king is not in check for any of those moves
	// TODO

	return moves
}

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
		move, err := GenerateMove(p, fromSquare, r+notation.Rank(rp*inverter), f)
		if err != nil || move == nil {
			break
		}
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
		move, err := GenerateMove(p, fromSquare, r+notation.Rank(1*inverter), f+notation.File(fp*inverter))
		if err != nil || move == nil {
			continue
		}
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
			move, err := GenerateMove(p, fromSquare, r+notation.Rank(pair.RP*i), f+notation.File(pair.FP*i))
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

func GenerateMove(p *notation.Position, fromSquare notation.Square, toR notation.Rank, toF notation.File) (move *notation.Move, err error) {
	piece := p.PieceList.PieceAt(fromSquare)

	// inverter used to determine same/opposite color
	inverter := 1
	if !p.WhitesTurn {
		inverter = -1
	}

	// Check square is valid
	toSquare, err := notation.NewSquare(toR, toF)
	if err != nil {
		return nil, err
	}

	// Check it toSquare is same color piece
	toSquarePieceRelative := p.PieceList.PieceAt(toSquare) * notation.Piece(inverter)
	if toSquarePieceRelative > 0 { // Same Color piece
		return nil, nil
	}

	// Create basic move
	move = &notation.Move{
		PieceList: p.PieceList,
		From:      fromSquare,
		To:        toSquare,
		Piece:     piece,
	}

	// Check if toSquare is empty or opposite color piece
	if toSquarePieceRelative < 0 { // Opposite color piece
		move.IsCapture = true
		return move, nil
	} else { // Empty Square
		if piece == notation.Piece_WhitePawn && toSquare == p.EnPassantSquare {
			move.IsCapture = true
			return move, nil
		}
		return move, nil
	}
}
