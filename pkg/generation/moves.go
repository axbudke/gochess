package generation

import (
	"gochess/pkg/board"
)

type Move struct {
	From board.Square
	To   board.Square
}

// func (m Move) String() string {
// 	return ""
// }

type MoveList []Move

func GenerateMoves(b board.BoardInterface) MoveList {
	moves := MoveList{}

	for i, pieceVal := range b.PieceList() {
		square := board.Square(i)
		if b.IsWhitesTurn() {
			switch pieceVal {
			case board.PieceVal_WhitePawn:
				moves = append(moves, GeneratePawnMoves(b, square)...)
			case board.PieceVal_WhiteKnight:
				moves = append(moves, GenerateKnightMoves(b, square)...)
			}
		} else {
			switch pieceVal {
			case board.PieceVal_BlackPawn:
				moves = append(moves, GeneratePawnMoves(b, square)...)
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

	// Get Pawn movements
	f, r := fromSquare.FileRank()
	r_forward1 := r + 1*board.Rank(inverter)
	r_forward2 := r + 2*board.Rank(inverter)

	// Check pawn movements
	// Forward 1
	foward1, err := board.NewSquare(r_forward1, f)
	if err == nil && b.PieceList().PieceAt(foward1) == 0 {
		moves = append(moves, Move{From: fromSquare, To: foward1})

		// Forward 2
		if _, fromR := fromSquare.FileRank(); fromR == board.Rank2 {
			forward2, err := board.NewSquare(r_forward2, f)
			if err == nil && b.PieceList().PieceAt(forward2) == 0 {
				moves = append(moves, Move{From: fromSquare, To: forward2})
			}
		}
	}

	// Check pawn takes
	for _, fp := range []board.File{f + 1*board.File(inverter), f - 1*board.File(inverter)} {
		toSquare, err := board.NewSquare(r_forward1, fp)
		if err != nil {
			continue
		} else if b.PieceList().PieceAt(toSquare)*board.PieceVal(inverter) < 0 {
			moves = append(moves, Move{From: fromSquare, To: toSquare})
		}
		// TODO check for enpassant take
	}

	// Promote
	// TODO

	return moves
}

func GenerateKnightMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	// Check knight movements
	f, r := fromSquare.FileRank()
	rs := []board.Rank{r + 1, r + 2, r + 1, r + 2, r - 1, r - 2, r - 1, r - 2}
	fs := []board.File{f + 2, f + 1, f - 2, f - 1, f + 2, f + 1, f - 2, f - 1}
	for i := range rs {
		toSquare, err := board.NewSquare(rs[i], fs[i])
		if err == nil && (b.PieceList().PieceAt(toSquare)*board.PieceVal(inverter)) < 0 {
			moves = append(moves, Move{From: fromSquare, To: toSquare})
		} else if err == nil && b.PieceList().PieceAt(toSquare) == 0 {
			moves = append(moves, Move{From: fromSquare, To: toSquare})
		}
	}

	return moves
}

func GenerateBishopMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	// Check bishop movements
	f, r := fromSquare.FileRank()
	for _, rp := range []int{1, -1} {
		for _, fp := range []int{1, -1} {
			for i := 1; i < 8; i++ {
				toSquare, err := board.NewSquare(r+board.Rank(rp*i), f+board.File(fp*i))
				if err != nil {
					// Invalid square, Can't move further
					break
				}
				toSquarePieceVal := b.PieceList().PieceAt(toSquare)
				if toSquarePieceVal*board.PieceVal(inverter) > 0 {
					// Same color piece in the way, Can't move further
					break
				} else if toSquarePieceVal*board.PieceVal(inverter) < 0 {
					moves = append(moves, Move{From: fromSquare, To: toSquare})
					// Opposite color piece in the way, Can't move further
					break
				} else if toSquarePieceVal == 0 {
					moves = append(moves, Move{From: fromSquare, To: toSquare})
				}
			}
		}
	}

	return moves
}

func GenerateRookMoves(b board.BoardInterface, fromSquare board.Square) MoveList {
	moves := MoveList{}

	inverter := 1
	if !b.IsWhitesTurn() {
		inverter = -1
	}

	// Check rook movements
	f, r := fromSquare.FileRank()
	for _, r_or_f := range []bool{true, false} {
		for _, p := range []int{1, -1} {
			var rp, fp int
			if r_or_f {
				rp = p
				fp = 0
			} else {
				rp = 0
				fp = p
			}
			for i := 1; i < 8; i++ {
				toSquare, err := board.NewSquare(r+board.Rank(rp*i), f+board.File(fp*i))
				if err != nil {
					// Invalid square, Can't move further
					break
				}
				toSquarePieceVal := b.PieceList().PieceAt(toSquare)
				if toSquarePieceVal*board.PieceVal(inverter) > 0 {
					// Same color piece in the way, Can't move further
					break
				} else if toSquarePieceVal*board.PieceVal(inverter) < 0 {
					moves = append(moves, Move{From: fromSquare, To: toSquare})
					// Opposite color piece in the way, Can't move further
					break
				} else if toSquarePieceVal == 0 {
					moves = append(moves, Move{From: fromSquare, To: toSquare})
				}
			}
		}
	}

	return moves
}
