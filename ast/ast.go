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
	Token  token.Token
	Symbol *Ident
	Value  Expr
}

func (ls *LetStmt) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStmt) stmtNode()            {}

type Ident struct {
	Token token.Token
	Name  string
}

func (i *Ident) TokenLiteral() string { return i.Token.Literal }
func (i *Ident) exprNode()            {}

type ReturnStmt struct {
	Token       token.Token
	ReturnValue Expr
}

func (rs *ReturnStmt) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStmt) stmtNode()            {}

type ExprStmt struct {
	Token      token.Token
	Expression Expr
}

func (es *ExprStmt) TokenLiteral() string { return es.Token.Literal }
func (es *ExprStmt) stmtNode()            {}
