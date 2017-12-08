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

package sesstype

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Tokens for use with lexer and parser.

// Token is a lexical token.
type Token int

const (
	// ILLEGAL is a special token for errors.
	ILLEGAL Token = iota
)

var eof = rune(0)

// TokenPos is a pair of coordinate to identify start of token.
type TokenPos struct {
	Char  int
	Lines []int
}

// CaretDiag returns the input b with caret to locate error position for diagnosis.
func (pos TokenPos) CaretDiag(b []byte) []byte {
	var lastLine bytes.Buffer

	for _, l := range pos.Lines {
		if l > 0 {
			lastLine.Reset() // New line will replace last line.
		}
		for c := 0; c < l; {
			r, size := utf8.DecodeRune(b)
			if r != '\n' {
				lastLine.WriteRune(r)
			}
			b = b[size:]
			c += size
		}
		b = b[1:] // newline
	}

	var errbuf bytes.Buffer
	var caret bytes.Buffer
	column := 0
LINE:
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r == '\n' {
			if column == 0 {
			} else {
				break LINE
			}
		} else {
			errbuf.WriteRune(r)
		}
		if column == pos.Char-1 {
			caret.WriteRune('↑')
		} else if column < pos.Char-1 {
			caret.WriteRune(' ')
		}
		b = b[size:]
		column += size
	}

	var diag bytes.Buffer
	prefix := strings.Repeat(" ", len(pos.String()))
	if lastLine.String() != "" {
		diag.WriteString(fmt.Sprintf("%s   %s\n", prefix, lastLine.String()))
	}
	diag.WriteString(fmt.Sprintf("%s → %s\n", pos.String(), errbuf.String()))
	diag.WriteString(fmt.Sprintf("%s   %s\n", prefix, caret.String()))
	return diag.Bytes()
}

func (p TokenPos) String() string {
	return fmt.Sprintf("%d:%d", len(p.Lines)+1, p.Char)
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlphaNum(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9')
}
