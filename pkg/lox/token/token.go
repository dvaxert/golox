package token

import (
	"fmt"

	"github.com/dvaxert/golox/pkg/lox/token/tokentype"
)

type Token struct {
	token_type tokentype.TokenType
	lexeme     string
	literal    any
	line       int
}
type TokenOption func(*Token)

//----------------------------------------------------------------------------------------------------------------------

func New(t tokentype.TokenType, opts ...TokenOption) Token {
	token := Token{
		token_type: t,
	}

	for _, opt := range opts {
		opt(&token)
	}

	return token
}

//----------------------------------------------------------------------------------------------------------------------

func WithLexeme(lexeme string) TokenOption {
	return func(t *Token) {
		t.lexeme = lexeme
	}
}

//----------------------------------------------------------------------------------------------------------------------

func WithLiteral(literal any) TokenOption {
	return func(t *Token) {
		t.literal = literal
	}
}

//----------------------------------------------------------------------------------------------------------------------

func WithLine(line int) TokenOption {
	return func(t *Token) {
		t.line = line
	}
}

//----------------------------------------------------------------------------------------------------------------------

func (t Token) String() string {
	makeErrorTokenMessage := func(t Token) string {
		return fmt.Sprintf(
			"Incorrect token: { token_type=%v lexeme=%v literal=%v line=%v",
			t.token_type, t.lexeme, t.literal, t.line,
		)
	}

	switch t.token_type {
	case tokentype.String:
		runes, ok := t.literal.([]rune)
		if !ok {
			panic(makeErrorTokenMessage(t))
		}
		return fmt.Sprintf("%s = \"%s\"", t.token_type.String(), string(runes))
	case tokentype.Number:
		value, ok := t.literal.(float64)
		if !ok {
			panic(makeErrorTokenMessage(t))
		}
		return fmt.Sprintf("%s = %f", t.token_type.String(), value)
	default:
		return fmt.Sprint(t.token_type)
	}
}

//----------------------------------------------------------------------------------------------------------------------

func (t Token) Lexeme() string {
	return t.lexeme
}
