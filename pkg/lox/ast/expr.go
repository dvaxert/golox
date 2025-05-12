package ast

type Expr[T any] interface {
	Accept(v Visitor[T]) T
}
