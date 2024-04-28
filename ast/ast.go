package ast

import "soooota1201/go_interpreter/token"

type Program struct {
	Statements []Statement
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface { // 動的に式を代入できる？？というか代入させるために interface にしている？
	Node
	expressionNode()
}

// Identifier は Expression を実装する
type Identifier struct {
	Token token.Token //token.IDENT トークン
	Value string
}

type LetStatement struct {
	Token token.Token //token.LET トークン
	Name *Identifier
	Value Expression // インターフェースへのポインタを保持する。これによりインターフェースの実装を参照できる。
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents
func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents
func (i *Identifier) expressionNode()
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

