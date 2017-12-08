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
)

// Role represents an endpoint or participant in a session interaction.
type Role struct {
	name string
}

// Name returns the identifying name of the Role.
func (r Role) Name() string {
	return r.name
}

// NewRole returns a Role value with the given name.
func NewRole(name string) Role {
	return Role{name: name}
}

// Message represents a message passed between Roles in a session interaction.
type Message struct {
	Label   string // Message label
	Payload string // Freeform string
}

func (m Message) String() string {
	return fmt.Sprintf("%s(%s)", m.Label, m.Payload)
}

// Global represents a global type.
type Global interface {
	global()
	String() string
}

// Interact is a global type interaction.
type Interact struct {
	From, To Role               // Interaction source and destination roles
	Globals  map[Message]Global // Mapping from message to continuations
}

func (*Interact) global() {}

// NewInteract returns a new global type interaction from p to q,
// where g is a mapping from message to send/receive to continuations.
func NewInteract(p, q Role, g map[Message]Global) *Interact {
	return &Interact{From: p, To: q, Globals: g}
}

func (s *Interact) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s → %s: ", s.From.Name(), s.To.Name()))
	if len(s.Globals) > 1 {
		buf.WriteString("{ ")
		first := true
		for m, g := range s.Globals {
			if !first {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%s.%s", m.String(), g.String()))
			first = false
		}
		buf.WriteString(" }")
	} else {
		for m, g := range s.Globals {
			buf.WriteString(fmt.Sprintf("%s.%s", m.String(), g.String()))
		}
	}
	return buf.String()
}

// Recur is a global type recursion.
type Recur struct {
	T string
	G Global
}

func (*Recur) global() {}

// NewRecur returns a new recursion with given recur label t and body g.
func NewRecur(t string, g Global) *Recur {
	return &Recur{T: t, G: g}
}

func (s *Recur) String() string {
	return fmt.Sprintf("μ%s.%s", s.T, s.G.String())
}

// TypeVar is a global type type variable.
type TypeVar struct {
	T string
}

func (TypeVar) global() {}

// NewTypeVar returns a new type variable with for recursion labelled t.
func NewTypeVar(t string) TypeVar {
	return TypeVar{T: t}
}

func (s TypeVar) String() string {
	return s.T
}

// End is a global type End.
type End struct{}

func (End) global() {}

// NewEnd returns a new global type End.
func NewEnd() End {
	return End{}
}

func (End) String() string {
	return "end"
}

// Local represents a local type.
type Local interface {
	local()
	String() string
}

// Branch is a local type branching.
type Branch struct {
	From   Role              // Interaction source Role
	Locals map[Message]Local // Mapping from message to continuations
}

func (*Branch) local() {}

// NewBranch returns a new local type branching from p,
// where l is a mapping from message to continuations.
func NewBranch(p Role, l map[Message]Local) *Branch {
	return &Branch{From: p, Locals: l}
}

func (s *Branch) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s ", s.From.Name()))
	if len(s.Locals) > 1 {
		buf.WriteString("&{ ")
		first := true
		for m, g := range s.Locals {
			if !first {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("?%s.%s", m.String(), g.String()))
			first = false
		}
		buf.WriteString(" }")
	} else {
		for m, g := range s.Locals {
			buf.WriteString(fmt.Sprintf("?%s.%s", m.String(), g.String()))
		}
	}
	return buf.String()
}

// Select is a local type selection.
type Select struct {
	To     Role              // Interaction destination Role
	Locals map[Message]Local // Mapping from message to continuations
}

func (*Select) local() {}

// NewSelect returns a new local type selection by p,
// where l is a mapping from message to continuations.
func NewSelect(p Role, l map[Message]Local) *Select {
	return &Select{To: p, Locals: l}
}

func (s *Select) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s ", s.To.Name()))
	if len(s.Locals) > 1 {
		buf.WriteString("⊕{ ")
		first := true
		for m, g := range s.Locals {
			if !first {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("!%s.%s", m.String(), g.String()))
			first = false
		}
		buf.WriteString(" }")
	} else {
		for m, g := range s.Locals {
			buf.WriteString(fmt.Sprintf("!%s.%s", m.String(), g.String()))
		}
	}
	return buf.String()
}

// LRecur is a local type recursion.
type LRecur struct {
	T string
	L Local
}

func (*LRecur) local() {}

// NewLRecur returns a new recursion with given recur label t and body l.
func NewLRecur(t string, l Local) *LRecur {
	return &LRecur{T: t, L: l}
}

func (s *LRecur) String() string {
	return fmt.Sprintf("μ%s.%s", s.T, s.L.String())
}

// LTypeVar is a local type type variable.
type LTypeVar struct {
	T string
}

func (LTypeVar) local() {}

// NewLTypeVar returns a new type variable with for recursion labelled t.
func NewLTypeVar(t string) LTypeVar {
	return LTypeVar{T: t}
}

func (s LTypeVar) String() string {
	return s.T
}

// LEnd is a local type End.
type LEnd struct{}

func (LEnd) local() {}

// NewEnd returns a new local type End.
func NewLEnd() LEnd {
	return LEnd{}
}

func (LEnd) String() string {
	return "end"
}
