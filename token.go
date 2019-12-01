package c2

// A parsed token of input
type Token struct {
	symbol Terminal
	lexeme string
	bytes  []byte
	file   string
	row    int
	column int
}

// The lexeme for the token.
func (t Token) Lexeme() string {
	return t.lexeme
}

// The actual bytes making up the token.
func (t Token) Bytes() []byte {
	return t.bytes[:]
}

// The position of the token in its file.
func (t Token) Loc() (file string, row, column int) {
	return t.file, t.row, t.column
}
