package global

import (
	"bytes"
	"fmt"

	"go.nickng.io/sesstype"
)

// Type represents a global type.
type Type interface {
	global()
	String() string
}

// Interact is a global type interaction.
type Interact struct {
	From, To sesstype.Role             // Interaction source and destination roles
	Globals  map[sesstype.Message]Type // Mapping from message to continuations
}

func (*Interact) global() {}

// NewInteract returns a new global type interaction from p to q,
// where g is a mapping from message to send/receive to continuations.
func NewInteract(p, q sesstype.Role, g map[sesstype.Message]Type) *Interact {
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
	G Type
}

func (*Recur) global() {}

// NewRecur returns a new recursion with given recur label t and body g.
func NewRecur(t string, g Type) *Recur {
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
