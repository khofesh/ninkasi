package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers + literals
	IDENT ="IDENT" // add, x, y
	INT = "INT" // 1 2 3 

	// operators
	ASSIGN = "="
	PLUS = "+"

	// delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":  LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}