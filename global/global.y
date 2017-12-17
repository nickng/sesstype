%{
package global

import (
	"io"

	"go.nickng.io/sesstype"
	"go.nickng.io/sesstype/local"
)

var global Type // Temporary holder for parsed Global
%}

%union {
	strval     string
	msg        sesstype.Message
	role       sesstype.Role
	gtype      Type
	branch     map[sesstype.Message]Type
	sendrecvs  []parseSendRecvs
	ltype      local.Type
	lbranch    map[sesstype.Message]local.Type
	lsendrecvs []struct{m sesstype.Message; l local.Type}
}

%token LPAREN RPAREN LBRACE RBRACE COLON ARROW DOT COMMA MU UNIT END EXCLAIMMARK QUESTIONMARK AMPERSAND OPLUS
%token <strval> IDENT TYPEVAR
%type <msg> message
%type <role> role
%type <gtype> global recur typevar end
%type <sendrecvs> sendrecv sendrecvs
%type <branch> interact
%type <ltype> local lrecur ltypevar lend
%type <lbranch> selects branches
%type <lsendrecvs> send sends recv recvs

%%

top : global { global = $1 }
    ;

role : IDENT { $$ = sesstype.NewRole($1) }
     ;

message : IDENT UNIT                { $$ = sesstype.Message{Label: $1}              }
        |       UNIT                { $$ = sesstype.Message{Label: ""}              }
        | IDENT LPAREN IDENT RPAREN { $$ = sesstype.Message{Label: $1, Payload: sesstype.BaseType{$3}} }
        | IDENT LPAREN local RPAREN { $$ = sesstype.Message{Label: $1, Payload: $3} }
        ;

global : role ARROW role COLON interact { $$ = NewInteract($1, $3, $5) }
       | recur                          { $$ = $1 }
       | typevar                        { $$ = $1 }
       | end                            { $$ = $1 }
       ;

interact : LBRACE sendrecvs RBRACE { $$ = make(map[sesstype.Message]Type); for _, sr := range $2 { $$[sr.m] = sr.g } }
         |        sendrecv         { $$ = make(map[sesstype.Message]Type); for _, sr := range $1 { $$[sr.m] = sr.g } }
         ;

sendrecv : message DOT global { $$ = []parseSendRecvs{parseSendRecvs{$1, $3}} }
         ;

sendrecvs : sendrecvs COMMA sendrecv { $$ = append($1, $3...) }
          |                 sendrecv { $$ = []parseSendRecvs{parseSendRecvs{m: $1[0].m, g: $1[0].g}} }
          ;

recur : MU IDENT DOT global { $$ = NewRecur($2, $4) }
      ;

typevar : TYPEVAR {  $$ = NewTypeVar($1) }
        ;

end : END { $$ = NewEnd() }
    ;

local : role branches { $$ = local.NewBranch($1, $2) }
      | role selects  { $$ = local.NewSelect($1, $2) }
      | lrecur        { $$ = $1 }
      | ltypevar      { $$ = $1 }
      | lend          { $$ = $1 }
      ;

branches : AMPERSAND LBRACE recvs RBRACE  { $$ = make(map[sesstype.Message]local.Type); for _, l := range $3 { $$[l.m] = l.l } }
         |                  recv          { $$ = make(map[sesstype.Message]local.Type); for _, l := range $1 { $$[l.m] = l.l } }
         ;

recv : QUESTIONMARK message DOT local { $$ = []struct{m sesstype.Message; l local.Type}{ struct{m sesstype.Message; l local.Type}{$2, $4} } }
     ;

recvs : recvs COMMA recv { $$ = append($1, $3...) }
      |             recv { $$ = []struct{m sesstype.Message; l local.Type}{ struct{m sesstype.Message; l local.Type}{m: $1[0].m, l: $1[0].l} } }
      ;

selects : OPLUS LBRACE sends RBRACE { $$ = make(map[sesstype.Message]local.Type); for _, s := range $3 { $$[s.m] = s.l } }
        |              send         { $$ = make(map[sesstype.Message]local.Type); for _, s := range $1 { $$[s.m] = s.l } }
        ;

send : EXCLAIMMARK message DOT local { $$ = []struct{m sesstype.Message; l local.Type}{ struct{m sesstype.Message; l local.Type}{$2, $4} } }
     ;

sends : sends COMMA send { $$ = append($1, $3...) }
      |             send { $$ = []struct{m sesstype.Message; l local.Type}{ struct{m sesstype.Message; l local.Type}{m: $1[0].m, l: $1[0].l} } }
      ;

lrecur : MU IDENT DOT local { $$ = local.NewRecur($2, $4) }
       ;

ltypevar : TYPEVAR { $$ = local.NewTypeVar($1) }
         ;

lend : END { $$ = local.NewEnd() }
     ;

%%

// Parse is the entry point to the Global Type parser.
// To parse Local Type, use the Parse function of the
// local subpackage (i.e. local.Parse).
func Parse(r io.Reader) (Type, error) {
	l := NewLexer(r)
	sesstypeParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return global, nil
	}
}
