package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"

	//Identifiers + literals
	IDENT     = "IDENT"  //variable and function names
	INT       = "INT"    //123456789
    STRING    = "STRING" //"Hello World"

	//Operators
	ASSIGN    = "="
	PLUS      = "+"
    MINUS     = "-"
    BANG      = "!"
    ASTRISK   = "*"
    SLASH     = "/"

    LT        = "<"
    GT        = ">"
    EQ        = "=="
    NOT_EQ    = "!="

	//Delimiters
	COMMA      = ","
	SEMICOLON  = ";"
    COLON      = ":"

	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
    LBRACKET  = "["
    RBRACKET  = "]"

	//Keywords
	FUNCTION  = "FUNCTION"
    MACRO     = "MACRO"
	LET       = "LET"
    IF        = "IF"
    ELSE      = "ELSE"
    TRUE      = "TRUE"
    FALSE     = "FALSE"
    RETURN    = "RETURN"

)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
    "if": IF,
    "else": ELSE,
    "true": TRUE,
    "false": FALSE,
    "return": RETURN,
    "macro": MACRO,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

