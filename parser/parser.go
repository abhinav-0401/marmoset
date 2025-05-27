package parser

import (
	"github.com/abhinav-0401/marmoset/ast"
	"github.com/abhinav-0401/marmoset/lexer"
	"github.com/abhinav-0401/marmoset/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	var p = &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	var program = &ast.Program{}
	program.Stmts = []ast.Stmt{}

	for p.currToken.Type != token.EOF {
		var stmt ast.Stmt = p.parseStmt()

		program.Stmts = append(program.Stmts, stmt)
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStmt()
	default:
		return nil
	}
}

func (p *Parser) parseLetStmt() *ast.LetStmt {
	var stmt = &ast.LetStmt{}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	var ident = &ast.Ident{Name: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(kind token.TokenType) bool {
	if p.peekTokenIs(kind) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) currTokenIs(kind token.TokenType) bool {
	return p.currToken.Type == kind
}

func (p *Parser) peekTokenIs(kind token.TokenType) bool {
	return p.peekToken.Type == kind
}
