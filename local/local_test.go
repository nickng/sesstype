package local

import (
	"strings"
	"testing"

	"go.nickng.io/sesstype"
)

func TestParseEnd(t *testing.T) {
	l, err := Parse(strings.NewReader("end"))
	if err != nil {
		t.Fatal(err)
	}
	if want, got := "end", l.String(); want != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParseBranch(t *testing.T) {
	l, err := Parse(strings.NewReader("B &{ ?().end, ?l(int).end }"))
	if err != nil {
		t.Fatal(err)
	}
	if want, want2, got := "B &{ ?().end, ?l(int).end }", "B &{ ?l(int).end, ?().end }", l.String(); want != got && want2 != got {
		t.Errorf("Parse error: want either %s or %s but got %s", want, want2, got)
	}
}

func TestParseSelect(t *testing.T) {
	l, err := Parse(strings.NewReader("A +{!().end,!l(int).end}"))
	if err != nil {
		t.Fatal(err)
	}
	if want, want2, got := "A ⊕{ !().end, !l(int).end }", "A ⊕{ !l(int).end, !().end }", l.String(); want != got && want2 != got {
		t.Errorf("Parse error: want either %s or %s but got %s", want, want2, got)
	}
}

func TestParseRecur(t *testing.T) {
	l, err := Parse(strings.NewReader("*T . T"))
	if err != nil {
		t.Fatal(err)
	}
	if want, got := "μT.T", l.String(); want != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParseTypeVar(t *testing.T) {
	// This is a tricky, "T" should be recognised as TypeVar
	// but it's possible to mistaken it as Role
	l, err := Parse(strings.NewReader("T"))
	if err != nil {
		t.Fatal(err)
	}
	if l, ok := l.(sesstype.LTypeVar); !ok {
		t.Errorf("Parse error: expected to get LTypeVar but got %T", l)
	}
	if want, got := "T", l.(sesstype.LTypeVar).T; want != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParse(t *testing.T) {
	l, err := Parse(strings.NewReader(`
	*T . A  !l(int).
		A &{ ?l().end, ?().T }
	`))
	if err != nil {
		t.Fatal(err)
	}
	if want, want2, got := "μT.A !l(int).A &{ ?l().end, ?().T }", "μT.A !l(int).A &{ ?().T, ?l().end }", l.String(); want != got && want2 != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}
