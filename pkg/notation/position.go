package notation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	PieceList  PieceList
	WhitesTurn bool
	castling   struct {
		whiteShort bool
		whiteLong  bool
		blackShort bool
		blackLong  bool
	}
	EnPassantSquare Square
	HalfmoveCount   int
	FullmoveCount   int
}

func (p Position) String() string {
	return string(p.FEN())
}

// ==================== Piece Functions ====================

func (p Position) PieceAt(s Square) Piece {
	return p.PieceList.PieceAt(s)
}

func (p Position) PieceAtIsSame(s Square) bool {
	// inverter used to determine same/opposite color
	inverter := func() int {
		if p.WhitesTurn {
			return 1
		} else {
			return -1
		}
	}()

	piece := p.PieceList.PieceAt(s)
	return piece*Piece(inverter) > 0
}

// ==================== Castling Functions ====================

func (p Position) CanWhiteCastle() bool {
	return p.castling.whiteLong || p.castling.whiteShort
}

func (p Position) CanBlackCastle() bool {
	return p.castling.blackLong || p.castling.blackShort
}

func (p Position) CanCastle(isWhite, isShort bool) bool {
	if isWhite && isShort {
		return p.castling.whiteShort
	} else if isWhite && !isShort {
		return p.castling.whiteLong
	} else if !isWhite && isShort {
		return p.castling.blackShort
	} else if !isWhite && !isShort {
		return p.castling.blackLong
	}
	return false
}

// ==================== Fen Functions ====================

func (p *Position) parseFEN(fenStr FEN) error {
	// Parse FEN with regexp
	submatches := FenRegExp.FindStringSubmatch(string(fenStr))
	if submatches == nil {
		return fmt.Errorf("failed to parse regexp")
	}

	// Parse PieceList from fenPiecePlacementStr
	p.PieceList = make(PieceList, 64)
	index := 0
	pieceRows := strings.Split(submatches[1], "/")
	for i := len(pieceRows) - 1; i >= 0; i-- {
		for _, c := range []byte(pieceRows[i]) {
			if regexp.MustCompile("[1-8]").Match([]byte{c}) {
				val, _ := strconv.Atoi(string(c))
				index += val
			} else if regexp.MustCompile("[pnbrqkPNBRQK]").Match([]byte{c}) {
				v, err := PieceChar(c).Val()
				if err != nil {
					return err
				}
				p.PieceList[index] = v
				index++
			} else {
				return fmt.Errorf("invalid piece syntax")
			}
		}
	}

	// Parse Side to Move
	p.WhitesTurn = (submatches[2] == "w")

	// Parse Castling
	p.castling.whiteShort = strings.Contains(submatches[3], "K")
	p.castling.whiteLong = strings.Contains(submatches[3], "Q")
	p.castling.blackShort = strings.Contains(submatches[3], "k")
	p.castling.blackLong = strings.Contains(submatches[3], "q")

	// Parse En Passant Square
	var err error
	if submatches[4] == "-" {
		p.EnPassantSquare = Square(-1)
	} else {
		p.EnPassantSquare, err = NewSquareFromString(submatches[4])
		if err != nil {
			return fmt.Errorf("failed to parse square: %w", err)
		}
	}

	// Parse Halfmove Count
	p.HalfmoveCount, err = strconv.Atoi(submatches[5])
	if err != nil {
		return fmt.Errorf("failed to parse halfmove count: %w", err)
	}

	// Parse Fullmove Count
	p.FullmoveCount, err = strconv.Atoi(submatches[6])
	if err != nil {
		return fmt.Errorf("failed to parse fullmove count: %w", err)
	}

	return nil
}

func (p Position) FEN() FEN {
	// Print Piece Placement
	pieceRows := []string{}
	emptyCount := 0
	for r := 0; r < 8; r++ {
		pieceRow := ""
		for f := 0; f < 8; f++ {
			piece := p.PieceList[r*8+f]
			if piece == Piece_Empty {
				emptyCount++
				continue
			}
			if emptyCount != 0 {
				pieceRow += fmt.Sprint(emptyCount)
				emptyCount = 0
			}
			pieceRow += fmt.Sprint(piece.String())
		}
		if emptyCount != 0 {
			pieceRow += fmt.Sprint(emptyCount)
			emptyCount = 0
		}
		pieceRows = append(pieceRows, pieceRow)
	}
	piecePlacementStr := pieceRows[len(pieceRows)-1]
	for i := len(pieceRows) - 2; i >= 0; i-- {
		piecePlacementStr += "/" + pieceRows[i]
	}

	// Print Side to Move
	sideToMoveStr := "b"
	if p.WhitesTurn {
		sideToMoveStr = "w"
	}

	// Print Castling
	castlingStr := ""
	if p.castling.whiteShort {
		castlingStr += "K"
	}
	if p.castling.whiteLong {
		castlingStr += "Q"
	}
	if p.castling.blackShort {
		castlingStr += "k"
	}
	if p.castling.blackLong {
		castlingStr += "q"
	}
	if castlingStr == "" {
		castlingStr = "-"
	}

	// Print En Passant Square
	enPassantTargetSquareStr := "-"
	if p.EnPassantSquare != Square(-1) {
		enPassantTargetSquareStr = p.EnPassantSquare.String()
	}

	// Parse Halfmove Count
	halfmoveCountStr := strconv.Itoa(p.HalfmoveCount)

	// Parse Fullmove Count
	fullmoveCountStr := strconv.Itoa(p.FullmoveCount)

	return FEN(fmt.Sprintf("%s %s %s %s %s %s",
		piecePlacementStr, sideToMoveStr, castlingStr,
		enPassantTargetSquareStr, halfmoveCountStr, fullmoveCountStr))
}
