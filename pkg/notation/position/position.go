package position

import (
	"gochess/pkg/notation/piece"
	"gochess/pkg/notation/square"
)

func NewPosition(fenStr FEN) (*Position, error) {
	p := &Position{}
	err := p.parseFEN(fenStr)
	if err != nil {
		return nil, err
	}
	return p, nil
}

type Position struct {
	PieceList  []piece.Piece
	WhitesTurn bool
	Castling   struct {
		whiteShort bool
		whiteLong  bool
		blackShort bool
		blackLong  bool
	}
	EnPassantSquare square.Square
	HalfmoveCount   int
	FullmoveCount   int
}

func (p Position) String() string {
	return string(p.FEN())
}

// ==================== Piece Functions ====================

func (p Position) PieceAt(s square.Square) piece.Piece {
	// TODO check if square is valid
	return p.PieceList[int(s)]
}

// ==================== Castling Functions ====================

func (p Position) CanWhiteCastle() bool {
	return p.Castling.whiteLong || p.Castling.whiteShort
}

func (p Position) CanBlackCastle() bool {
	return p.Castling.blackLong || p.Castling.blackShort
}

func (p Position) CanCastle(isWhite, isShort bool) bool {
	if isWhite && isShort {
		return p.Castling.whiteShort
	} else if isWhite && !isShort {
		return p.Castling.whiteLong
	} else if !isWhite && isShort {
		return p.Castling.blackShort
	} else if !isWhite && !isShort {
		return p.Castling.blackLong
	}
	return false
}

// ==================== Ascii ====================

func (p Position) AsciiString() string {
	var asciiString string

	// Top of board
	asciiString += "\n +---+---+---+---+---+---+---+---+\n"

	// Each Rank
	for r := square.Rank8; r >= square.Rank1; r-- {
		// Each File
		for f := square.FileA; f <= square.FileH; f++ {
			// Each Piece
			asciiString += " | " + p.PieceAt(square.NewSquare(f, r)).String()
		}
		asciiString += " | " + r.String() + "\n"
		asciiString += " +---+---+---+---+---+---+---+---+\n"
	}

	// Bottom of board
	asciiString += "   a   b   c   d   e   f   g   h \n\n"

	// Extras
	asciiString += "FEN: " + string(p.FEN()) + ""

	return asciiString
}
