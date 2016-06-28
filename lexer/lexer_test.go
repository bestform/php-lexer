package lexer

import (
	"io/ioutil"
	"strings"
	"testing"

	lex "github.com/timtadh/lexmachine"
)

// TestParse tests parsing of literals, keywords and tokens
func TestParse(t *testing.T) {
	cases := []struct {
		Name, Input   string
		ExpectedTypes []string
	}{
		{"Brackets",
			"(){}",
			[]string{`(`, `)`, `{`, `}`}},
		{"Punctuation",
			"=,;",
			[]string{`=`, `,`, `;`}},
		{"php tag and arrow",
			"-><?php",
			[]string{`->`, `<?php`}},
		{"numbers",
			"123 3.45",
			[]string{`NUMBER`, `NUMBER`}},
		{"variables",
			"$foo $bar",
			[]string{`VAR`, `VAR`}},
		{"comments",
			`
		 // line comment
		 /* block comment */
		 /** doc block comment */
		 `, []string{`COMMENT`, `BLOCKCOMMENT`, `DOCCOMMENT`}},
		{"key words",
			"use return function private protected public class extends implements true false",
			[]string{`USE`, `RETURN`, `FUNCTION`, `PRIVATE`, `PROTECTED`, `PUBLIC`, `CLASS`, `EXTENDS`, `IMPLEMENTS`, `TRUE`, `FALSE`}},
	}

	for _, c := range cases {
		s, err := Lexer.Scanner([]byte(c.Input))
		if err != nil {
			t.Error(err)
		}
		var actualTypes []string
		for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
			if err != nil {
				t.Error(err)
			}
			token := tok.(*lex.Token)
			actualTypes = append(actualTypes, Tokens[token.Type])
		}

		if len(actualTypes) != len(c.ExpectedTypes) {
			t.Errorf("error in Test '%v': expected: %v, actual: %v", c.Name, c.ExpectedTypes, actualTypes)
		}

		for i, ac := range actualTypes {
			if c.ExpectedTypes[i] != ac {
				t.Errorf("error in Test '%v': expected: %v, actual: %v", c.Name, c.ExpectedTypes, actualTypes)
			}
		}

	}
}

func TestFixture(t *testing.T) {
	content, err := ioutil.ReadFile("testfixtures/in.php")
	if err != nil {
		t.Error(err)
	}
	parts := strings.Split(string(content), "/*===EXPECTED===\n")
	if len(parts) != 3 {
		t.Error("Faulty fixture input")
	}
	// @todo: this is a lot of code duplication across tests. Should consolidate
	s, err := Lexer.Scanner([]byte(parts[0]))
	if err != nil {
		t.Error(err)
	}
	var actualTypes []string
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			t.Error(err)
		}
		token := tok.(*lex.Token)
		actualTypes = append(actualTypes, Tokens[token.Type])
	}
	expectedTypes := strings.Split(parts[1], "\n")
	if len(actualTypes) != len(expectedTypes) {
		t.Errorf("error in Fixture Test: expected: \n%v, \nactual: \n%v", expectedTypes, actualTypes)
	}
	for i, ac := range actualTypes {
		if expectedTypes[i] != ac {
			t.Errorf("error in Fixture Test: expected: \n%v, \nactual: \n%v", expectedTypes, actualTypes)
		}
	}
}
