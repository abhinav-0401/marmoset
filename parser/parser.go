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

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expr
	infixParseFn  func(left ast.Expr) ast.Expr
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
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

		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExprStmt()
	}
}

func (p *Parser) parseLetStmt() *ast.LetStmt {
	var stmt = &ast.LetStmt{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	var ident = &ast.Ident{
		Token: p.currToken,
		Name:  p.currToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	stmt.Symbol = ident

	return stmt
}

func (p *Parser) parseReturnStmt() *ast.ReturnStmt {
	var stmt = &ast.ReturnStmt{Token: p.currToken}

	p.nextToken()
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	var stmt = &ast.ExprStmt{Token: p.currToken}
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
