package ast

import (
	"github.com/dvaxert/golox/pkg/lox/token"
)

func NewUnary[T any](op token.Token, r Expr[T]) *Unary[T] {
	return &Unary[T]{
		operator: op,
		right:    r,
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Unary[T any] struct {
	operator token.Token
	right    Expr[T]
}

//----------------------------------------------------------------------------------------------------------------------

func (u *Unary[T]) Accept(v Visitor[T]) T {
	return v.visitUnaryExpr(u)
}
