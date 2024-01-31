package position

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BoardInterface interface {
	PieceList() PieceList
	IsWhitesTurn() bool
	CanCastle(isWhite, isShort bool) bool
	EnPassantSquare() string
	HalfmoveCount() int
	FullmoveCount() int

	String() string
}

func New(fenStr FEN) (BoardInterface, error) {
	b := &Board{}
	err := b.parseFEN(fenStr)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type Board struct {
	pieceList  PieceList
	whitesTurn bool
	castling   struct {
		whiteShort bool
		whiteLong  bool
		blackShort bool
		blackLong  bool
	}
	enPassantSquare string
	halfmoveCount   int
	fullmoveCount   int
}

func (b Board) PieceList() PieceList    { return b.pieceList }
func (b Board) IsWhitesTurn() bool      { return b.whitesTurn }
func (b Board) EnPassantSquare() string { return b.enPassantSquare }
func (b Board) HalfmoveCount() int      { return b.halfmoveCount }
func (b Board) FullmoveCount() int      { return b.fullmoveCount }
func (b Board) CanCastle(isWhite, isShort bool) bool {
	if isWhite && isShort {
		return b.castling.whiteShort
	} else if isWhite && !isShort {
		return b.castling.whiteLong
	} else if !isWhite && isShort {
		return b.castling.blackShort
	} else if !isWhite && !isShort {
		return b.castling.blackLong
	}
	return false
}

func (b Board) String() string {
	return string(b.FEN())
}

func (b *Board) parseFEN(fenStr FEN) error {
	// Parse FEN with regexp
	submatches := fenRegExp.FindStringSubmatch(string(fenStr))
	if submatches == nil {
		return fmt.Errorf("failed to parse regexp")
	}

	// Parse PieceList from fenPiecePlacementStr
	b.pieceList = make(PieceList, 64)
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
				b.pieceList[index] = v
				index++
			} else {
				return fmt.Errorf("invalid piece syntax")
			}
		}
	}

	// Parse Active Color
	b.whitesTurn = submatches[2] == "w"

	// Parse Castling
	b.castling.whiteShort = strings.Contains(submatches[3], "K")
	b.castling.whiteLong = strings.Contains(submatches[3], "Q")
	b.castling.blackShort = strings.Contains(submatches[3], "k")
	b.castling.blackLong = strings.Contains(submatches[3], "q")

	// Parse En Passant Square
	b.enPassantSquare = submatches[4]

	// Parse Halfmove Count
	halfmoveCount, err := strconv.Atoi(submatches[5])
	if err != nil {
		return err
	}
	b.halfmoveCount = halfmoveCount

	// Parse Fullmove Count
	fullmoveCount, err := strconv.Atoi(submatches[6])
	if err != nil {
		return err
	}
	b.fullmoveCount = fullmoveCount

	return nil
}

func (b Board) FEN() FEN {
	piecePlacementStr := ""
	pieceRows := []string{}
	emptyCount := 0
	for r := 0; r < 8; r++ {
		pieceRow := ""
		for f := 0; f < 8; f++ {
			pieceVal := b.pieceList[r*8+f]
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
