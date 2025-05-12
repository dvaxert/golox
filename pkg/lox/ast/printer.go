package ast

import (
	"bytes"
	"fmt"
)

type AstPrinter struct{}

//----------------------------------------------------------------------------------------------------------------------

func (p *AstPrinter) Print(e Expr[string]) string {
	return e.Accept(p)
}

//----------------------------------------------------------------------------------------------------------------------

func (p *AstPrinter) visitBinaryExpr(b *Binary[string]) string {
	return p.parenthesize(b.Operator.Lexeme(), b.Left, b.Right)
}

//----------------------------------------------------------------------------------------------------------------------

func (p *AstPrinter) visitGroupingExpr(g *Grouping[string]) string {
	return p.parenthesize("group", g.expression)
}

//----------------------------------------------------------------------------------------------------------------------

func (p *AstPrinter) visitLiteralExpr(l *Literal[string]) string {
	return fmt.Sprint(l.value)
}

//----------------------------------------------------------------------------------------------------------------------

func (p *AstPrinter) visitUnaryExpr(u *Unary[string]) string {
	return p.parenthesize(u.operator.Lexeme(), u.right)
}

//----------------------------------------------------------------------------------------------------------------------

func (p *AstPrinter) parenthesize(name string, exprs ...Expr[string]) string {
	buf := bytes.Buffer{}

	buf.WriteRune('(')
	buf.WriteString(name)
	for _, e := range exprs {
		buf.WriteRune(' ')
		buf.WriteString(e.Accept(p))
	}
	buf.WriteRune(')')
	return buf.String()
}
