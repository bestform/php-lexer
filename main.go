package main

import (
	"fmt"
	"log"
	"php-lexer/lexer"

	lex "github.com/timtadh/lexmachine"
)

func main() {
	s, err := lexer.Lexer.Scanner([]byte(`
    <?php


    class foo() {

			/** I am a DOC Comment */
      public function foobar() {
				$a = 4; // I am a comment

				/* I am a block comment */

				return $a;
      }

    }

    `))

	if err != nil {
		log.Fatal(err)
	}

	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lex.Token)
		fmt.Printf("%-7v -> %-10v | %v:%v-%v:%v\n",
			lexer.Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
}
