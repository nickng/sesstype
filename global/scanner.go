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

package global

import (
	"bufio"
	"bytes"
	"io"

	"go.nickng.io/sesstype"
)

// Scanner is a lexical scanner.
type Scanner struct {
	r   *bufio.Reader
	pos sesstype.TokenPos
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), pos: sesstype.TokenPos{Char: 0, Lines: []int{}}}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if reached the end or error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if ch == '\n' {
		s.pos.Lines = append(s.pos.Lines, s.pos.Char)
		s.pos.Char = 0
	} else {
		s.pos.Char++
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	if s.pos.Char == 0 {
		s.pos.Char = s.pos.Lines[len(s.pos.Lines)-1]
		s.pos.Lines = s.pos.Lines[:len(s.pos.Lines)-1]
	} else {
		s.pos.Char--
	}
}

// Scan returns the next token and parsed value.
func (s *Scanner) Scan() (token Token, value string, startPos, endPos sesstype.TokenPos) {
	ch := s.read()

	if isWhitespace(ch) {
		s.skipWhitespace()
		ch = s.read()
	}
	if isAlphaNum(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Track token positions.
	startPos = s.pos
	defer func() { endPos = s.pos }()

	switch ch {
	case eof:
		return 0, "", startPos, endPos
	case 'μ', '*':
		return MU, "μ", startPos, endPos
	case ':':
		return COLON, string(ch), startPos, endPos
	case ',':
		return COMMA, string(ch), startPos, endPos
	case '.':
		return DOT, string(ch), startPos, endPos
	case '&':
		return AMPERSAND, string(ch), startPos, endPos
	case '!':
		return EXCLAIMMARK, string(ch), startPos, endPos
	case '?':
		return QUESTIONMARK, string(ch), startPos, endPos
	case '⊕', '+':
		return OPLUS, string(ch), startPos, endPos
	case '{':
		return LBRACE, string(ch), startPos, endPos
	case '}':
		return RBRACE, string(ch), startPos, endPos
	case '(':
		if ch2 := s.read(); ch2 == ')' {
			return UNIT, "()", startPos, endPos
		}
		s.unread()
		return LPAREN, string(ch), startPos, endPos
	case ')':
		return RPAREN, string(ch), startPos, endPos
	case '-':
		if ch2 := s.read(); ch2 == '>' {
			return ARROW, "→", startPos, endPos
		}
		s.unread()
		// Illegal
	case '→':
		return ARROW, "→", startPos, endPos
	}

	return ILLEGAL, string(ch), startPos, endPos
}

func (s *Scanner) scanIdent() (token Token, value string, startPos, endPos sesstype.TokenPos) {
	var buf bytes.Buffer
	startPos = s.pos
	defer func() { endPos = s.pos }()
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isAlphaNum(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch buf.String() {
	case "end":
		return END, buf.String(), startPos, endPos
	}
	return IDENT, buf.String(), startPos, endPos
}

func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
