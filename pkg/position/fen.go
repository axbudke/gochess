package position

import (
	"fmt"
	"regexp"
)

// FEN describes a chess position in a one line ascii string.
type FEN string

// FEN: One FEN string or record consists of six fields separated by a space character.
// <FEN> ::=  <Piece Placement>
//        ' ' <Side to move>
//        ' ' <Castling ability>
//        ' ' <En passant target square>
//        ' ' <Halfmove clock>
//        ' ' <Fullmove counter>

var (
	fenRegExpStr = fmt.Sprintf("(%s) (%s) (%s) (%s) (%s) (%s)",
		piecePlacementRegExpStr, sideToMoveRegExpStr, castlingAbilityRegExpStr,
		enPassantTargetSquareRegExpStr, countRegExpStr, countRegExpStr)
	fenRegExp = regexp.MustCompile(fenRegExpStr)
)

// Piece Placement: The Piece Placement is determined rank by rank in big-endian order, that is starting
// at the 8th rank down to the first rank. Each rank is separated by the terminal symbol '/' (slash).
// One rank, scans piece placement in little-endian file-order from the A to H.
// A decimal digit counts consecutive empty squares, the pieces are identified by a single letter from
// standard English names for chess pieces as used in the Algebraic Chess Notation. Uppercase letters
// are for white pieces, lowercase letters for black pieces.
// <Piece Placement> ::= <rank8>'/'<rank7>'/'<rank6>'/'<rank5>'/'<rank4>'/'<rank3>'/'<rank2>'/'<rank1>
// <ranki>       ::= [<digit17>]<piece> {[<digit17>]<piece>} [<digit17>] | '8'
// <piece>       ::= <white Piece> | <black Piece>
// <digit17>     ::= '1' | '2' | '3' | '4' | '5' | '6' | '7'
// <white Piece> ::= 'P' | 'N' | 'B' | 'R' | 'Q' | 'K'
// <black Piece> ::= 'p' | 'n' | 'b' | 'r' | 'q' | 'k'

var (
	piecePlacementLineRegExpStr = "[pnbrqkPNBRQK1-8]{1,8}"
	piecePlacementRegExpStr     = fmt.Sprintf("%s(?:/%s){7}", piecePlacementLineRegExpStr, piecePlacementLineRegExpStr)
	piecePlacementRegExp        = regexp.MustCompile(piecePlacementRegExpStr)
)

// Side to move: Side to move is one lowercase letter for either White ('w') or Black ('b').
// <Side to move> ::= {'w' | 'b'}

var (
	sideToMoveRegExpStr = "[wb]"
	sideToMoveRegExp    = regexp.MustCompile(sideToMoveRegExpStr)
)

// Castling ability: If neither side can castle, the symbol '-' is used, otherwise each of four individual
// castling rights for king and queen castling for both sides are indicated by a sequence of one to four letters.
// <Castling ability> ::= '-' | ['K'] ['Q'] ['k'] ['q'] (1..4)

var (
	castlingAbilityRegExpStr = "[-KQkq]{0,4}"
	castlingAbilityRegExp    = regexp.MustCompile(castlingAbilityRegExpStr)
)

// En passant target square: The en passant target square is specified after a double push of a pawn,
// no matter whether an en passant capture is really possible or not.
// Other moves than double pawn pushes imply the symbol '-' for this FEN field.
// <En passant target square> ::= '-' | <epsquare>
// <epsquare>   ::= <fileLetter> <eprank>
// <fileLetter> ::= 'a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h'
// <eprank>     ::= '3' | '6'

var (
	enPassantTargetSquareRegExpStr = ".*"
	enPassantTargetSquareRegExp    = regexp.MustCompile(enPassantTargetSquareRegExpStr)
)

// Halfmove Clock: The halfmove clock specifies a decimal number of half moves with respect to the 50
// move draw rule. It is reset to zero after a capture or a pawn move and incremented otherwise.
// <Halfmove Clock> ::= <digit> {<digit>}
// <digit> ::= '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'

var (
	countRegExpStr = "[[:digit:]]{1,3}"
	countRegExp    = regexp.MustCompile(countRegExpStr)
)

// Fullmove counter: The number of the full moves in a game. It starts at 1, and is incremented after
// each Black's move.
// <Fullmove counter> ::= <digit19> {<digit>}
// <digit19> ::= '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'
// <digit>   ::= '0' | <digit19>

// FEN Examples
const (
	StartingFEN FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)
