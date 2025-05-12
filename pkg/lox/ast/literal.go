package ast

func NewLiteral[T any](v any) *Literal[T] {
	return &Literal[T]{
		value: v,
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Literal[T any] struct {
	value any
}

//----------------------------------------------------------------------------------------------------------------------

func (l *Literal[T]) Accept(v Visitor[T]) T {
	return v.visitLiteralExpr(l)
}
