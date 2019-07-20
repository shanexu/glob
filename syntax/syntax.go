package syntax

import (
	"github.com/shanexu/glob/syntax/ast"
	"github.com/shanexu/glob/syntax/lexer"
)

func Parse(s string) (*ast.Node, error) {
	return ast.Parse(lexer.NewLexer(s))
}

func Special(b byte) bool {
	return lexer.Special(b)
}
