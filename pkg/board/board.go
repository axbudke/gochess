package board

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
}

func New(fenStr FEN) (*Board, error) {
	b := &Board{}
	err := b.parse(fenStr)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type Board struct {
	piecesStr  string
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

func (b *Board) parse(fenStr FEN) error {
	// Parse FEN with regexp
	submatches := fullRegExp.FindStringSubmatch(string(fenStr))
	if submatches == nil {
		return fmt.Errorf("failed to parse regexp")
	}
	fmt.Printf("FEN submatches: %#v\n", submatches)

	// Parse PieceList from fenPiecePlacementStr
	b.piecesStr = submatches[1]
	b.pieceList = make(PieceList, 64)
	index := 0
	for _, c := range []byte(submatches[1]) {
		// Parse Piece from string
		if c == '/' {
			// This char is just a row divider, means nothing in this parsing
			continue
		} else if regexp.MustCompile("[1-8]").Match([]byte{c}) {
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
	fmt.Printf("PiecePlacement.pieceList: %#v\n", b.pieceList)

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
	fullmoveCount, err := strconv.Atoi(submatches[5])
	if err != nil {
		return err
	}
	b.fullmoveCount = fullmoveCount

	return nil
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
