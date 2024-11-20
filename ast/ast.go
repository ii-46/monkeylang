package ast

import "github.com/ii-46/monkeylang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// This going to be root node of AST
type Program struct {
	Statements []Statement // Contain all statements
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // token.INDENT
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal

}

func (i *Identifier) expressionNode() {}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) statementNode() {
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statementNode() {
}

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) statementNode() {
}
