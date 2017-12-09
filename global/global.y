%{
package global

import (
	"io"

	"go.nickng.io/sesstype"
)

var global Type // Temporary holder for parsed Global
%}

%union {
	strval     string
	msg        sesstype.Message
	role       sesstype.Role
	global     Type
	branch     map[sesstype.Message]Type
	sendrecvs  []struct{m sesstype.Message; g Type}
}

%token LPAREN RPAREN LBRACE RBRACE COLON ARROW DOT COMMA MU UNIT END EXCLAIMMARK QUESTIONMARK AMPERSAND OPLUS
%token <strval> IDENT
%type <msg> message
%type <role> role
%type <global> global recur typevar end
%type <sendrecvs> sendrecv sendrecvs
%type <branch> interact

%%

top : global { global = $1 }
    ;

role : IDENT { $$ = sesstype.NewRole($1) }
     ;

message : IDENT UNIT                { $$ = sesstype.Message{Label: $1}              }
        |       UNIT                { $$ = sesstype.Message{Label: ""}              }
        | IDENT LPAREN IDENT RPAREN { $$ = sesstype.Message{Label: $1, Payload: $3} }
        ;

global : role ARROW role COLON interact { $$ = NewInteract($1, $3, $5) }
       | recur                          { $$ = $1 }
       | typevar                        { $$ = $1 }
       | end                            { $$ = $1 }
       ;

interact : LBRACE sendrecvs RBRACE { $$ = make(map[sesstype.Message]Type); for _, sr := range $2 { $$[sr.m] = sr.g } }
         |        sendrecv         { $$ = make(map[sesstype.Message]Type); for _, sr := range $1 { $$[sr.m] = sr.g } }
         ;

sendrecv : message DOT global { $$ = []struct{m sesstype.Message; g Type}{ struct{m sesstype.Message; g Type}{$1, $3} } }
         ;

sendrecvs : sendrecvs COMMA sendrecv { $$ = append($1, $3...) }
          |                 sendrecv { $$ = []struct{m sesstype.Message; g Type}{ struct{m sesstype.Message; g Type}{m: $1[0].m, g: $1[0].g} } }
          ;

recur : MU IDENT DOT global { $$ = NewRecur($2, $4) }
      ;

typevar : IDENT { $$ = NewTypeVar($1) }
        ;

end : END { $$ = NewEnd() }
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
