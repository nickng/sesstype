// Copyright 2017 Nicholas Ng <nickng@nickng.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package local

//go:generate goyacc -p sesstype -o parser.y.go local.y

import (
	"io"

	"go.nickng.io/sesstype"
)

// Lexer for sesstype.
type Lexer struct {
	scanner *Scanner
	Errors  chan error
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{scanner: NewScanner(r), Errors: make(chan error, 1)}
}

// Lex is provided for yacc-compatible parser.
func (l *Lexer) Lex(yylval *sesstypeSymType) int {
	var token sesstype.Token
	token, yylval.strval, _, _ = l.scanner.Scan()
	return int(token)
}

// Error handles error.
func (l *Lexer) Error(err string) {
	l.Errors <- &sesstype.ErrParse{Err: err, Pos: l.scanner.pos}
}
