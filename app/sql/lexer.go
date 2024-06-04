package sql

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"strings"
)

type TokenType int

const (
	EOF TokenType = iota

	// Keywords
	SELECT
	FROM
	WHERE

	// Operators
	EQ
	NEQ
	GT
	GTE
	LT
	LTE
	AND
	OR

	// Reserved symbols
	ASTERISK
	COMMA
	OPEN_PAREN
	CLOSE_PAREN

	// Constants
	STRING
	INT
	FLOAT
)

var reservedKeywords = map[string]TokenType{
	"SELECT": SELECT,
	"FROM":   FROM,
	"WHERE":  WHERE,
	"AND":    AND,
	"OR":     OR,
}

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input  string
	curPos int
	length int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		length: len(input),
		curPos: 0,
	}
}

// Returns the next byte in the input string without advancing the lexer's position.
func (l *Lexer) Peek() byte {
	if l.curPos >= l.length {
		return 0
	}

	return l.input[l.curPos]
}

// Returns the next byte in the input string and advances the lexer's position.
func (l *Lexer) Next() byte {
	if l.curPos >= l.length {
		return 0
	}

	b := l.input[l.curPos]
	l.curPos++

	return b
}

// Returns the next token in the input string and advances the lexer's position.
// Returns an error if an unexpected token is found.
func (l *Lexer) NextToken() (Token, error) {
	tok := Token{}

	// Skip whitespace in the input
	for l.Peek() == ' ' {
		l.Next()
	}

	chr := l.Peek()
	switch chr {
	case 0:
		tok.Type = EOF
		tok.Value = ""

	case '=':
		tok.Type = EQ
		tok.Value = "="
		l.Next()

	case '>':
		l.Next()
		if l.Peek() == '=' {
			tok.Type = GTE
			tok.Value = ">="
			l.Next()
		} else {
			tok.Type = GT
			tok.Value = ">"
		}

	case '<':
		l.Next()
		if l.Peek() == '=' {
			tok.Type = LTE
			tok.Value = "<="
			l.Next()
		} else {
			tok.Type = LT
			tok.Value = "<"
		}

	case '!':
		l.Next()
		if l.Peek() == '=' {
			tok.Type = NEQ
			tok.Value = "!="
			l.Next()
		} else {
			return tok, fmt.Errorf("unexpected token: %c after ! operator", l.Peek())
		}

	case '*':
		l.Next()
		tok.Type = ASTERISK
		tok.Value = "*"

	case ',':
		l.Next()
		tok.Type = COMMA
		tok.Value = ","

	case '(':
		l.Next()
		tok.Type = OPEN_PAREN
		tok.Value = "("

	case ')':
		l.Next()
		tok.Type = CLOSE_PAREN
		tok.Value = ")"

	default:
		if utils.IsDigit(chr) {
			return l.parseIntOrFloat()
		} else if utils.IsAlpha(chr) || chr == '"' {
			return l.parseString()
		} else {
			return tok, fmt.Errorf("unexpected token: %c", chr)
		}
	}

	return tok, nil
}

// Internal function to parse a string token.
// Evaluates the same to be either a "" enclosed string or a reserved keyword.
func (l *Lexer) parseString() (Token, error) {
	tok := Token{Type: STRING, Value: ""}

	if l.Peek() == '"' {
		l.Next() // Skip the opening quote

		for l.Peek() != '"' {
			if l.Peek() == 0 {
				return tok, fmt.Errorf("unexpected EOF while parsing string")
			}
			tok.Value += string(l.Next())
		}

		l.Next() // Skip the closing quote
	} else {
		// Read till we get alphanumeric characters
		for utils.IsAlphaNumeric(l.Peek()) {
			tok.Value += string(l.Next())
		}

		// The string must be a reserved keyword
		if keyword, ok := reservedKeywords[strings.ToUpper(tok.Value)]; ok {
			tok.Type = keyword
			tok.Value = strings.ToUpper(tok.Value)
		}
	}

	return tok, nil
}

func (l *Lexer) parseIntOrFloat() (Token, error) {
	tok := Token{Type: INT, Value: ""}

	// Parse the integer part of the number
	val := int(l.Next() - '0')
	for utils.IsDigit(l.Peek()) {
		val = val*10 + int(l.Next()-'0')
	}
	tok.Value = fmt.Sprintf("%d", val)

	// Check if the integer is a DOUBLE
	if l.Peek() == '.' {
		l.Next()

		tok.Type = FLOAT
		valFloat := float64(val)
		dec := 0.1
		for utils.IsDigit(l.Peek()) {
			valFloat += float64(l.Next()-'0') * dec
			dec /= 10
		}
		tok.Value = fmt.Sprintf("%f", valFloat)
	}

	return tok, nil
}
