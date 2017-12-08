%{
package local

import (
	"io"

	"go.nickng.io/sesstype"
)

var local sesstype.Local // Temporary holder for parsed Local
%}

%union {
	strval    string
	msg       sesstype.Message
	role      sesstype.Role
	local     sesstype.Local
	sendrecvs []struct{m sesstype.Message; l sesstype.Local}
	branch    map[sesstype.Message]sesstype.Local
}

%token LPAREN RPAREN LBRACE RBRACE COLON ARROW DOT COMMA MU UNIT END EXCLAIMMARK QUESTIONMARK AMPERSAND OPLUS
%token <strval> IDENT
%type <msg> message
%type <role> role
%type <local> local lrecur ltypevar lend
%type <branch> selects branches
%type <sendrecvs> send sends recv recvs

%%

top : local { local = $1 }
    ;

role : IDENT { $$ = sesstype.NewRole($1) }
     ;

message : IDENT UNIT                { $$ = sesstype.Message{Label: $1}              }
        |       UNIT                { $$ = sesstype.Message{Label: ""}              }
        | IDENT LPAREN IDENT RPAREN { $$ = sesstype.Message{Label: $1, Payload: $3} }
        ;

local : role branches { $$ = sesstype.NewBranch($1, $2) }
      | role selects  { $$ = sesstype.NewSelect($1, $2) }
      | lrecur               { $$ = $1 }
      | ltypevar             { $$ = $1 }
      | lend                 { $$ = $1 }
      ;

branches : AMPERSAND LBRACE recvs RBRACE  { $$ = make(map[sesstype.Message]sesstype.Local); for _, l := range $3 { $$[l.m] = l.l } }
         |                  recv          { $$ = make(map[sesstype.Message]sesstype.Local); for _, l := range $1 { $$[l.m] = l.l } }
         ;

recv : QUESTIONMARK message DOT local { $$ = []struct{m sesstype.Message; l sesstype.Local}{ struct{m sesstype.Message; l sesstype.Local}{$2, $4} } }
     ;

recvs : recvs COMMA recv { $$ = append($1, $3...) }
      |             recv { $$ = []struct{m sesstype.Message; l sesstype.Local}{ struct{m sesstype.Message; l sesstype.Local}{m: $1[0].m, l: $1[0].l} } }
      ;

selects : OPLUS LBRACE sends RBRACE { $$ = make(map[sesstype.Message]sesstype.Local); for _, s := range $3 { $$[s.m] = s.l } }
        |              send         { $$ = make(map[sesstype.Message]sesstype.Local); for _, s := range $1 { $$[s.m] = s.l } }
        ;

send : EXCLAIMMARK message DOT local { $$ = []struct{m sesstype.Message; l sesstype.Local}{ struct{m sesstype.Message; l sesstype.Local}{$2, $4} } }
     ;

sends : sends COMMA send { $$ = append($1, $3...) }
      |             send { $$ = []struct{m sesstype.Message; l sesstype.Local}{ struct{m sesstype.Message; l sesstype.Local}{m: $1[0].m, l: $1[0].l} } }
      ;

lrecur : MU IDENT DOT local { $$ = sesstype.NewLRecur($2, $4) }
       ;

ltypevar : IDENT { $$ = sesstype.NewLTypeVar($1) }
         ;

lend : END { $$ = sesstype.NewLEnd() }
     ;

%%

// Parse is the entry point to the local type parser.
func Parse(r io.Reader) (sesstype.Local, error) {
	l := NewLexer(r)
	sesstypeParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return local, nil
	}
}
