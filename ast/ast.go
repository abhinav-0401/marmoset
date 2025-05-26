package ast

import (
	"github.com/abhinav-0401/marmoset/token"
)

type Node interface {
	TokenLiteral() string
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type Program struct {
	Stmts []Stmt
}

func (p *Program) TokenLiteral() string {
	if len(p.Stmts) > 0 {
		return p.Stmts[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStmt struct {
	Token token.Token
}
