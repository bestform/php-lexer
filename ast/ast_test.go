package ast

import (
	"io/ioutil"
	"php-lexer/lexer"
	"strings"
	"testing"
)

func TestSimpleNodeName(t *testing.T) {
	s := NewSimpleNode("foo")

	if s.Name() != "foo" {
		t.Error("Name should be foo")
	}
}

func TestSimpleNodeChildren(t *testing.T) {
	s := NewSimpleNode("foo")

	if len(s.Children()) != 0 {
		t.Error("New SimpleNode should not have children")
	}

	s.AddChild(NewSimpleNode("bar"))

	if len(s.Children()) != 1 {
		t.Error("Children should be added")
	}

	if s.Children()[0].Name() != "bar" {
		t.Error("Added node should be the first in the list")
	}
}

func TestAST(t *testing.T) {
	content, err := ioutil.ReadFile("../lexer/testfixtures/in.php")
	if err != nil {
		t.Error(err)
	}
	parts := strings.Split(string(content), "/*===EXPECTED===\n")
	if len(parts) != 3 {
		t.Error("Faulty fixture input")
	}
	s, err := lexer.Lexer.Scanner([]byte(parts[0]))
	if err != nil {
		t.Error(err)
	}
	MakeASTFromLexer(s)

}
