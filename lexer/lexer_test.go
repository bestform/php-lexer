package lexer

import (
	"testing"

	lex "github.com/timtadh/lexmachine"
)

// TestParse tests parsing of literals, keywords and tokens
func TestParse(t *testing.T) {
	cases := []struct {
		Name, Input   string
		ExpectedTypes []string
	}{
		{"Brackets", "(){}", []string{`(`, `)`, `{`, `}`}},
		{"Misc", "=,;", []string{`=`, `,`, `;`}},
		{"php tag and arrow", "-><?php", []string{`->`, `<?php`}},
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
