package ast

import (
	"github.com/Rexbrainz/Rex/token"
)

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

type Identifier struct {
	Token	token.Token
	Value	string
}

type LetStatement struct {
	Token	token.Token
	Name	*Identifier
	Value	Expression
}

type ReturnStatement struct {
	Token		token.Token // the 'return' token
	ReturnValue	Expression
}

type Program struct {
	Statements []Statement
}

// Methods
// Program method
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement methods
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier method
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement methods
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
