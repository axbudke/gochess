package notation

type Castling uint8

const (
	Castling_None     Castling = 0
	Castling_WhiteOO  Castling = 1
	Castling_WhiteOOO Castling = 1 << 1
	Castling_BlackOO  Castling = 1 << 2
	Castling_BlackOOO Castling = 1 << 3

	Castling_White Castling = Castling_WhiteOO | Castling_WhiteOOO
	Castling_Black Castling = Castling_BlackOO | Castling_BlackOOO

	Castling_KingSide   Castling = Castling_WhiteOO | Castling_BlackOO
	Castling_QueeenSide Castling = Castling_WhiteOOO | Castling_BlackOOO

	Castling_Any Castling = Castling_White | Castling_Black
)
