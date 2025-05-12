package scanner

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/dvaxert/golox/pkg/lox/loxerr"
	"github.com/dvaxert/golox/pkg/lox/token"
	"github.com/dvaxert/golox/pkg/lox/token/tokentype"
)

func New(source []rune) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  nil,
		start:   0,
		current: 0,
		line:    0,
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Scanner struct {
	source  []rune
	tokens  []token.Token
	start   int
	current int
	line    int
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) ScanTokens() ([]token.Token, []error) {
	var errs []error
	sc.tokens = nil

	for !sc.isAtEnd() {
		sc.start = sc.current
		err := sc.ScanToken()

		var le *loxerr.LoxError
		if err != nil {
			if errors.As(err, &le) {
				errs = append(errs, err)
			} else {
				panic(err)
			}
		}
	}

	sc.tokens = append(sc.tokens, token.New(tokentype.EOF, token.WithLine(sc.line)))
	return sc.tokens, errs
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) ScanToken() error {
	c := sc.advance()
	switch c {
	case '(':
		sc.addToken(token.New(tokentype.LeftParen))
	case ')':
		sc.addToken(token.New(tokentype.RightParen))
	case '{':
		sc.addToken(token.New(tokentype.LeftBrace))
	case '}':
		sc.addToken(token.New(tokentype.RightBrace))
	case ',':
		sc.addToken(token.New(tokentype.Comma))
	case '.':
		sc.addToken(token.New(tokentype.Dot))
	case '-':
		sc.addToken(token.New(tokentype.Minus))
	case '+':
		sc.addToken(token.New(tokentype.Plus))
	case ';':
		sc.addToken(token.New(tokentype.Semicolon))
	case '*':
		sc.addToken(token.New(tokentype.Star))
	case '!':
		if sc.match('=') {
			sc.addToken(token.New(tokentype.BangEqual))
		} else {
			sc.addToken(token.New(tokentype.Bang))
		}
	case '=':
		if sc.match('=') {
			sc.addToken(token.New(tokentype.EqualEqual))
		} else {
			sc.addToken(token.New(tokentype.Equal))
		}
	case '<':
		if sc.match('=') {
			sc.addToken(token.New(tokentype.LessEqual))
		} else {
			sc.addToken(token.New(tokentype.Less))
		}
	case '>':
		if sc.match('=') {
			sc.addToken(token.New(tokentype.GreaterEqual))
		} else {
			sc.addToken(token.New(tokentype.Greater))
		}
	case '/':
		if sc.match('/') {
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else if sc.match('*') {
			sc.advance()
			err := sc.skipComment()
			if err != nil {
				return err
			}
		} else {
			sc.addToken(token.New(tokentype.Slash))
		}
	case ' ', '\r', '\t':
		// do nothing
	case '\n':
		sc.line++
	case '"':
		sc.readString()
	default:
		if unicode.IsDigit(c) {
			sc.readNumber()
		} else if unicode.IsLetter(c) || c == '_' {
			sc.readIdentifier()
		} else {
			return loxerr.New(loxerr.WithLine(sc.line), loxerr.WithMessage("unexpected character"))
		}
	}

	return nil
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) addToken(t token.Token) {
	sc.tokens = append(sc.tokens, t)
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) advance() rune {
	r := sc.source[sc.current]
	sc.current++
	return r
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) match(expected rune) bool {
	if sc.isAtEnd() {
		return false
	}

	if sc.source[sc.current] != expected {
		return false
	}

	sc.current++
	return true
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) peek() rune {
	if sc.isAtEnd() {
		return 0
	}
	return sc.source[sc.current]
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) readString() error {
	for sc.peek() != '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
		sc.advance()
	}

	if sc.isAtEnd() {
		return loxerr.New(loxerr.WithLine(sc.line), loxerr.WithMessage("Unterminated string."))
	}

	sc.advance()
	sc.addToken(token.New(tokentype.String, token.WithLiteral(sc.source[sc.start+1:sc.current-1])))

	return nil
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) readNumber() error {
	for unicode.IsDigit(sc.peek()) {
		sc.advance()
	}

	if sc.peek() == '.' && unicode.IsDigit(sc.peekNext()) {
		sc.advance()

		for unicode.IsDigit(sc.peek()) {
			sc.advance()
		}
	}

	value, err := strconv.ParseFloat(string(sc.source[sc.start:sc.current]), 64)
	if err != nil {
		panic(err)
	}

	sc.addToken(token.New(tokentype.Number, token.WithLiteral(value)))
	return nil
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) peekNext() rune {
	if sc.current+1 >= len(sc.source) {
		return 0
	}
	return sc.source[sc.current+1]
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) readIdentifier() {
	r := sc.peek()
	for unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		sc.advance()
		r = sc.peek()
	}

	token_type, ok := keywords[string(sc.source[sc.start:sc.current])]
	if !ok {
		token_type = tokentype.Identifier
	}

	sc.addToken(token.New(token_type))
}

//----------------------------------------------------------------------------------------------------------------------

func (sc *Scanner) skipComment() error {
	for {
		if sc.isAtEnd() {
			return loxerr.New(loxerr.WithLine(sc.line), loxerr.WithMessage("comment not terminated"))
		}

		if sc.peek() == '*' && sc.peekNext() == '/' {
			break
		}

		if sc.peek() == '/' && sc.peekNext() == '*' {
			sc.advance()
			sc.advance()

			err := sc.skipComment()

			if err != nil {
				return err
			}
		}

		sc.advance()
	}

	sc.advance()
	sc.advance()
	return nil
}
