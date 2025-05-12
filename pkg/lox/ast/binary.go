package ast

import (
	"github.com/dvaxert/golox/pkg/lox/token"
)

func NewBinary[T any](l Expr[T], op token.Token, r Expr[T]) *Binary[T] {
	return &Binary[T]{
		Left:     l,
		Operator: op,
		Right:    r,
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Binary[T any] struct {
	Left     Expr[T]
	Operator token.Token
	Right    Expr[T]
}

//----------------------------------------------------------------------------------------------------------------------

func (b *Binary[T]) Accept(v Visitor[T]) T {
	return v.visitBinaryExpr(b)
}
