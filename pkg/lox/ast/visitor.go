package ast

type Visitor[T any] interface {
	visitBinaryExpr(expr *Binary[T]) T
	visitGroupingExpr(expr *Grouping[T]) T
	visitLiteralExpr(expr *Literal[T]) T
	visitUnaryExpr(expr *Unary[T]) T
}
