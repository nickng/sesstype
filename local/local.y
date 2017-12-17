%{
package local

import (
	"io"

	"go.nickng.io/sesstype"
)

var local Type // Temporary holder for parsed Local
%}

%union {
	strval     string
	msg        sesstype.Message
	role       sesstype.Role
	ltype      Type
	lbranch    map[sesstype.Message]Type
	lsendrecvs []parseSendRecvs
}

%token LPAREN RPAREN LBRACE RBRACE COLON ARROW DOT COMMA MU UNIT END EXCLAIMMARK QUESTIONMARK AMPERSAND OPLUS
%token <strval> IDENT TYPEVAR
%type <msg> message
%type <role> role
%type <ltype> local lrecur ltypevar lend
%type <lbranch> selects branches
%type <lsendrecvs> send sends recv recvs

%%

top : local { local = $1 }
    ;

role : IDENT { $$ = sesstype.NewRole($1) }
     ;

message : IDENT UNIT                { $$ = sesstype.Message{Label: $1}              }
        |       UNIT                { $$ = sesstype.Message{Label: ""}              }
        | IDENT LPAREN IDENT RPAREN { $$ = sesstype.Message{Label: $1, Payload: sesstype.BaseType{Type: $3}} }
        | IDENT LPAREN local RPAREN { $$ = sesstype.Message{Label: $1, Payload: $3} }
        ;

local : role branches { $$ = NewBranch($1, $2) }
      | role selects  { $$ = NewSelect($1, $2) }
      | lrecur        { $$ = $1 }
      | ltypevar      { $$ = $1 }
      | lend          { $$ = $1 }
      ;

branches : AMPERSAND LBRACE recvs RBRACE  { $$ = make(map[sesstype.Message]Type); for _, l := range $3 { $$[l.m] = l.l } }
         |                  recv          { $$ = make(map[sesstype.Message]Type); for _, l := range $1 { $$[l.m] = l.l } }
         ;

recv : QUESTIONMARK message DOT local { $$ = []parseSendRecvs{parseSendRecvs{$2, $4} } }
     ;

recvs : recvs COMMA recv { $$ = append($1, $3...) }
      |             recv { $$ = []parseSendRecvs{parseSendRecvs{m: $1[0].m, l: $1[0].l}} }
      ;

selects : OPLUS LBRACE sends RBRACE { $$ = make(map[sesstype.Message]Type); for _, s := range $3 { $$[s.m] = s.l } }
        |              send         { $$ = make(map[sesstype.Message]Type); for _, s := range $1 { $$[s.m] = s.l } }
        ;

send : EXCLAIMMARK message DOT local { $$ = []parseSendRecvs{parseSendRecvs{$2, $4}} }
     ;

sends : sends COMMA send { $$ = append($1, $3...) }
      |             send { $$ = []parseSendRecvs{parseSendRecvs{m: $1[0].m, l: $1[0].l}} }
      ;

lrecur : MU IDENT DOT local { $$ = NewRecur($2, $4) }
       ;

ltypevar : TYPEVAR { $$ = NewTypeVar($1) }
         ;

lend : END { $$ = NewEnd() }
     ;

%%

// Parse is the entry point to the local type parser.
func Parse(r io.Reader) (Type, error) {
	l := NewLexer(r)
	sesstypeParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return local, nil
	}
}
