%{
package sesstype

import (
	"io"
)

var global Global // Temporary holder for parsed Global
%}

%union {
	strval     string
	msg        Message
	role       Role
	global     Global
	branch     map[Message]Global
	sendrecvs  []struct{m Message; g Global}
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

role : IDENT { $$ = NewRole($1) }
     ;

message : IDENT UNIT                { $$ = Message{Label: $1}              }
        |       UNIT                { $$ = Message{Label: ""}              }
        | IDENT LPAREN IDENT RPAREN { $$ = Message{Label: $1, Payload: $3} }
        ;

global : role ARROW role COLON interact { $$ = NewInteract($1, $3, $5) }
       | recur                          { $$ = $1 }
       | typevar                        { $$ = $1 }
       | end                            { $$ = $1 }
       ;

interact : LBRACE sendrecvs RBRACE { $$ = make(map[Message]Global); for _, sr := range $2 { $$[sr.m] = sr.g } }
         |        sendrecv         { $$ = make(map[Message]Global); for _, sr := range $1 { $$[sr.m] = sr.g } }
         ;

sendrecv : message DOT global { $$ = []struct{m Message; g Global}{ struct{m Message; g Global}{$1, $3} } }
         ;

sendrecvs : sendrecvs COMMA sendrecv { $$ = append($1, $3...) }
          |                 sendrecv { $$ = []struct{m Message; g Global}{ struct{m Message; g Global}{m: $1[0].m, g: $1[0].g} } }
          ;

recur : MU IDENT DOT global { $$ = NewRecur($2, $4) }
      ;

typevar : IDENT { $$ = NewTypeVar($1) }
        ;

end : END { $$ = NewEnd() }
    ;

%%

// Parse is the entry point to the global type parser.
func Parse(r io.Reader) (Global, error) {
	l := NewLexer(r)
	sesstypeParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return global, nil
	}
}
