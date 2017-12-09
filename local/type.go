package local

import (
	"bytes"
	"fmt"

	"go.nickng.io/sesstype"
)

// Local represents a local type.
type Type interface {
	local()
	String() string
}

// Branch is a local type branching.
type Branch struct {
	From   sesstype.Role             // Interaction source Role
	Locals map[sesstype.Message]Type // Mapping from message to continuations
}

func (*Branch) local() {}

// NewBranch returns a new local type branching from p,
// where l is a mapping from message to continuations.
func NewBranch(p sesstype.Role, l map[sesstype.Message]Type) *Branch {
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
	To     sesstype.Role             // Interaction destination Role
	Locals map[sesstype.Message]Type // Mapping from message to continuations
}

func (*Select) local() {}

// NewSelect returns a new local type selection by p,
// where l is a mapping from message to continuations.
func NewSelect(p sesstype.Role, l map[sesstype.Message]Type) *Select {
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

// Recur is a local type recursion.
type Recur struct {
	T string
	L Type
}

func (*Recur) local() {}

// NewLRecur returns a new recursion with given recur label t and body l.
func NewRecur(t string, l Type) *Recur {
	return &Recur{T: t, L: l}
}

func (s *Recur) String() string {
	return fmt.Sprintf("μ%s.%s", s.T, s.L.String())
}

// TypeVar is a local type type variable.
type TypeVar struct {
	T string
}

func (TypeVar) local() {}

// NewTypeVar returns a new type variable with for recursion labelled t.
func NewTypeVar(t string) TypeVar {
	return TypeVar{T: t}
}

func (s TypeVar) String() string {
	return s.T
}

// End is a local type End.
type End struct{}

func (End) local() {}

// NewEnd returns a new local type End.
func NewEnd() End {
	return End{}
}

func (End) String() string {
	return "end"
}
