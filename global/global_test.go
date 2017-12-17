package global

import (
	"strings"
	"testing"

	"go.nickng.io/sesstype"
)

func TestParseEnd(t *testing.T) {
	g, err := Parse(strings.NewReader("end"))
	if err != nil {
		t.Fatal(err)
	}
	if want, got := "end", g.String(); want != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParseInteract(t *testing.T) {
	g, err := Parse(strings.NewReader("A -> B : { ().end, l(int).end }"))
	if err != nil {
		t.Fatal(err)
	}
	if want, want2, got := "A → B: { ().end, l(int).end }", "A → B: { l(int).end, ().end }", g.String(); want != got && want2 != got {
		t.Errorf("Parse error: want either %s or %s but got %s", want, want2, got)
	}
}

func TestParseRecur(t *testing.T) {
	g, err := Parse(strings.NewReader("*T . T"))
	if err != nil {
		t.Fatal(err)
	}
	if want, got := "μT.T", g.String(); want != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParseTypeVar(t *testing.T) {
	// This is a tricky, "T" should be recognised as an invalid TypeVar
	// but it's possible to mistaken it as Role (which makes type incomplete).
	_, err := Parse(strings.NewReader("T"))
	if err != nil {
		if _, ok := err.(*sesstype.ErrParse); !ok {
			t.Fatalf("Expecting parse error but got %v", err)
		}
	}
}

func TestParse(t *testing.T) {
	g, err := Parse(strings.NewReader(`
	*T . A -> B: l(int).
		B->A: { l().end, ().T }
	`))
	if err != nil {
		t.Fatal(err)
	}
	if want, want2, got := "μT.A → B: l(int).B → A: { l().end, ().T }", "μT.A → B: l(int).B → A: { ().T, l().end }", g.String(); want != got && want2 != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParseSessionPayload(t *testing.T) {
	g, err := Parse(strings.NewReader(`
	*T . *x . A->B: l(C&{ ?l(B!L().end).end, ?l2(y).end }).end
	`))
	if err != nil {
		t.Fatal(err)
	}
	if want, want2, got := "μT.μx.A → B: l(C &{ ?l(B !L().end).end, ?l2(y).end }).end", "μT.μx.A → B: l(C &{ ?l2(y).end, ?l(B !L().end).end }).end", g.String(); want != got && want2 != got {
		t.Errorf("Parse error: want %s but got %s", want, got)
	}
}

func TestParseUnclosedSessionPayload(t *testing.T) {
	_, err := Parse(strings.NewReader(`
	*T . *x . A->B: l(C&{ ?l(B!L().T).end, ?l2(end).x }).end
	`))
	if err != nil {
		if _, ok := err.(*sesstype.ErrParse); !ok {
			t.Fatalf("Expecting parse error but got %v", err)
		}
	}
}
