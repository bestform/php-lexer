package ast

import (
	"fmt"

	"github.com/timtadh/lexmachine"
)

// Node is a node in the AST
type Node interface {
	Name() string
	Children() []Node
	AddChild(Node)
	String() string
}

// SimpleNode is the simplest form of a node.
// It just has a name and a list of children
type SimpleNode struct {
	name     string
	children []Node
}

// Name returns the name of the Node
func (s *SimpleNode) Name() string {
	return s.name
}

// Children returns a plain list of pointers to its children nodes
func (s *SimpleNode) Children() []Node {
	return s.children
}

// AddChild adds a Node to the list of children
func (s *SimpleNode) AddChild(n Node) {
	s.children = append(s.children, n)
}

func (s *SimpleNode) String() string {
	st := s.Name()
	for _, c := range s.Children() {
		st += "->" + c.String()
	}
	return st
}

// NewSimpleNode returns a pointer to a new SimpleNode with the name set
// and an empty children list
func NewSimpleNode(name string) *SimpleNode {
	s := SimpleNode{name: name}
	s.children = []Node{}

	return &s
}

// MakeASTFromLexer will produce an AST and return the root node
func MakeASTFromLexer(s *lexmachine.Scanner) {
	root := NewSimpleNode("root")

	current := root
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			panic("Something went terribly wrong while lexing: " + err.Error())
		}
		token := tok.(*lexmachine.Token)

		node := NewSimpleNode(string(token.Lexeme))

		(*current).AddChild(node)

		current = node
	}

	fmt.Println(root.String())
}
