package ast

func NewGrouping[T any](e Expr[T]) *Grouping[T] {
	return &Grouping[T]{
		expression: e,
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Grouping[T any] struct {
	expression Expr[T]
}

//----------------------------------------------------------------------------------------------------------------------

func (g *Grouping[T]) Accept(v Visitor[T]) T {
	return v.visitGroupingExpr(g)
}
