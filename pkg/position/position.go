package position

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func New(fenStr FEN) (*Position, error) {
	p := &Position{}
	err := p.parseFEN(fenStr)
	if err != nil {
		return nil, err
	}
	return p, nil
}

type Position struct {
	pieceList  PieceList
	whitesTurn bool
	castling   struct {
		whiteShort bool
		whiteLong  bool
		blackShort bool
		blackLong  bool
	}
	enPassantSquare Square
	halfmoveCount   int
	fullmoveCount   int
}

func (p Position) PieceList() PieceList    { return p.pieceList }
func (p Position) IsWhitesTurn() bool      { return p.whitesTurn }
func (p Position) EnPassantSquare() Square { return p.enPassantSquare }
func (p Position) HalfmoveCount() int      { return p.halfmoveCount }
func (p Position) FullmoveCount() int      { return p.fullmoveCount }
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

func (p Position) String() string {
	return string(p.FEN())
}

func (p *Position) parseFEN(fenStr FEN) error {
	// Parse FEN with regexp
	submatches := FenRegExp.FindStringSubmatch(string(fenStr))
	if submatches == nil {
		return fmt.Errorf("failed to parse regexp")
	}

	// Parse PieceList from fenPiecePlacementStr
	p.pieceList = make(PieceList, 64)
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
				p.pieceList[index] = v
				index++
			} else {
				return fmt.Errorf("invalid piece syntax")
			}
		}
	}

	// Parse Active Color
	p.whitesTurn = (submatches[2] == "w")

	// Parse Castling
	p.castling.whiteShort = strings.Contains(submatches[3], "K")
	p.castling.whiteLong = strings.Contains(submatches[3], "Q")
	p.castling.blackShort = strings.Contains(submatches[3], "k")
	p.castling.blackLong = strings.Contains(submatches[3], "q")

	// Parse En Passant Square
	var err error
	p.enPassantSquare, err = NewSquareFromString(submatches[4])
	if err != nil {
		return err
	}

	// Parse Halfmove Count
	p.halfmoveCount, err = strconv.Atoi(submatches[5])
	if err != nil {
		return err
	}

	// Parse Fullmove Count
	p.fullmoveCount, err = strconv.Atoi(submatches[6])
	if err != nil {
		return err
	}

	return nil
}

func (p Position) FEN() FEN {
	piecePlacementStr := ""
	pieceRows := []string{}
	emptyCount := 0
	for r := 0; r < 8; r++ {
		pieceRow := ""
		for f := 0; f < 8; f++ {
			pieceVal := p.pieceList[r*8+f]
			if pieceVal == PieceVal_Empty {
				emptyCount++
				continue
			}
			if emptyCount != 0 {
				pieceRow += fmt.Sprint(emptyCount)
				emptyCount = 0
			}
			pieceRow += fmt.Sprint(pieceVal.String())
		}
		pieceRows = append(pieceRows, pieceRow)
	}
	for i := len(pieceRows) - 1; i >= 0; i-- {
		piecePlacementStr += pieceRows[i]
	}
	fmt.Printf("%s", piecePlacementStr)

	fenStr := ""

	return FEN(fenStr)
}
