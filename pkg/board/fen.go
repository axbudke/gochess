package board

import (
	"fmt"
	"regexp"
)

// FEN a notation to represent a chess board snapshot
//
//	Syntax:
//		Full: <piece placement> <active color> <castling> <en passant square> <halfmove count> <fullmove count>
//		Piece Placement:
//		Active Color:
//		Castling Availability:
//		En Passant target square:
//		Halfmove count:
//		Fullmove count:
//
//	 Examples:
//	 	Starting position: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var (
	piecePlacementLineRegExpStr = "[pnbrqkPNBRQK1-8]{1,8}"
	piecePlacementRegExpStr     = fmt.Sprintf("%s(?:/%s){7}", piecePlacementLineRegExpStr, piecePlacementLineRegExpStr)
	activeColorRegExpStr        = "[wb]"
	castlingRegExpStr           = "[-KQkq]{0,4}"
	enPassantRegExpStr          = ".*"
	countRegExpStr              = "[[:digit:]]{1,3}"

	fullRegExp = regexp.MustCompile(fmt.Sprintf("(%s) (%s) (%s) (%s) (%s) (%s)",
		piecePlacementRegExpStr, activeColorRegExpStr, castlingRegExpStr,
		enPassantRegExpStr, countRegExpStr, countRegExpStr))
)

type FEN string

const StartingFEN FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
