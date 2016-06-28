package lexer

/*
tokens for now:

CLASS: class
TYPE: classname
IMPLEMENTS: implements
EXTENDS: extends
LEFTCURLY: {
RIGHTCURLY: }
LEFTBRACKET: (
RIGHTBRACKET: )
VARIABLE: $foo
EQ: =
COMMENTSTART: //
BLOCKCOMMENTSTART: /*
DOCBOCKSTART: /**
INNERBLOCKCOMMENTSTART: /
STRING: "foobar"
SQSTRING: 'foobar'
PLUS: +
MINUS: -
DIVIDE: /
MULTIPLY: *
NUMBER: 1234
SEMI: ;

*/
import (
	"strings"

	lex "github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

// Literals contains the tokens representing literal strings
var Literals []string

// Keywords contains the keyword tokens
var Keywords []string

// Tokens contains all of the tokens (incl. literals and keywords)
var Tokens []string

// TokenIds is a map from the token names to their int ids
var TokenIds map[string]int

// Lexer is used to construct a scanner
var Lexer *lex.Lexer

func initTokens() {
	Literals = []string{
		"(",
		")",
		"{",
		"}",
		"=",
		",",
		";",
		"->",
		"<?php",
	}
	Keywords = []string{
		"CLASS",
		"FUNCTION",
		"PRIVATE",
		"PUBLIC",
		"PROTECTED",
		"RETURN",
	}
	Tokens = []string{
		"COMMENT",
		"DOCCOMMENT",
		"BLOCKCOMMENT",
		"NAME",
		"VAR",
	}
	Tokens = append(Tokens, Literals...)
	Tokens = append(Tokens, Keywords...)
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

func initLexer() (*lex.Lexer, error) {
	lexer := lex.NewLexer()

	for _, lit := range Literals {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(lit))
	}
	for _, name := range Keywords {
		lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	lexer.Add([]byte(`/\*([^*]|\r|\n|)*\*+/`), token("BLOCKCOMMENT"))
	lexer.Add([]byte(`/\*([^*]|\r|\n|(\*+([^*/]|\r|\n)))*\*+/`), token("DOCCOMMENT"))
	lexer.Add([]byte(`//([^\n\r])*`), token("COMMENT"))
	lexer.Add([]byte(`\$([a-z]|[A-Z]|[0-9])*`), token("VAR"))
	lexer.Add([]byte(`([a-z]|[A-Z]|[0-9])*`), token("NAME"))
	lexer.Add([]byte("( |\t|\n|\r)+"), skip)

	err := lexer.Compile()
	if err != nil {
		return nil, err
	}
	return lexer, nil
}

func skip(*lex.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

func token(name string) lex.Action {
	return func(s *lex.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}

func init() {
	initTokens()
	var err error
	Lexer, err = initLexer()
	if err != nil {
		panic(err)
	}
}
