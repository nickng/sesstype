//line local.y:2
package local

import __yyfmt__ "fmt"

//line local.y:2
import (
	"io"

	"go.nickng.io/sesstype"
)

var local sesstype.Local // Temporary holder for parsed Local

//line local.y:13
type sesstypeSymType struct {
	yys       int
	strval    string
	msg       sesstype.Message
	role      sesstype.Role
	local     sesstype.Local
	sendrecvs []struct {
		m sesstype.Message
		l sesstype.Local
	}
	branch map[sesstype.Message]sesstype.Local
}

const LPAREN = 57346
const RPAREN = 57347
const LBRACE = 57348
const RBRACE = 57349
const COLON = 57350
const ARROW = 57351
const DOT = 57352
const COMMA = 57353
const MU = 57354
const UNIT = 57355
const END = 57356
const EXCLAIMMARK = 57357
const QUESTIONMARK = 57358
const AMPERSAND = 57359
const OPLUS = 57360
const IDENT = 57361

var sesstypeToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"COLON",
	"ARROW",
	"DOT",
	"COMMA",
	"MU",
	"UNIT",
	"END",
	"EXCLAIMMARK",
	"QUESTIONMARK",
	"AMPERSAND",
	"OPLUS",
	"IDENT",
}
var sesstypeStatenames = [...]string{}

const sesstypeEofCode = 1
const sesstypeErrCode = 2
const sesstypeInitialStackSize = 16

//line local.y:81

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

//line yacctab:1
var sesstypeExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	1, 22,
	7, 22,
	11, 22,
	-2, 2,
}

const sesstypePrivate = 57344

const sesstypeLast = 46

var sesstypeAct = [...]int{

	15, 2, 8, 13, 9, 40, 18, 21, 16, 7,
	17, 16, 12, 14, 23, 17, 33, 32, 37, 30,
	22, 29, 38, 27, 35, 24, 31, 34, 36, 25,
	20, 19, 39, 44, 1, 41, 26, 28, 10, 43,
	42, 11, 6, 5, 4, 3,
}
var sesstypePact = [...]int{

	-10, -1000, -1000, -5, -1000, -1000, -1000, -1000, -13, -1000,
	-1000, -1000, 25, -1000, 24, -1000, 1, 1, 19, -8,
	0, 9, 13, -1000, 6, -10, 17, -1000, 11, -1000,
	-10, -1000, -14, -10, -1000, -1000, -8, -1000, 0, -1000,
	28, -1000, -1000, -1000, -1000,
}
var sesstypePgo = [...]int{

	0, 7, 45, 1, 44, 43, 42, 41, 38, 0,
	37, 3, 36, 34,
}
var sesstypeR1 = [...]int{

	0, 13, 2, 1, 1, 1, 3, 3, 3, 3,
	3, 8, 8, 11, 12, 12, 7, 7, 9, 10,
	10, 4, 5, 6,
}
var sesstypeR2 = [...]int{

	0, 1, 1, 2, 1, 4, 2, 2, 1, 1,
	1, 4, 1, 4, 3, 1, 4, 1, 4, 3,
	1, 4, 1, 1,
}
var sesstypeChk = [...]int{

	-1000, -13, -3, -2, -4, -5, -6, 19, 12, 14,
	-8, -7, 17, -11, 18, -9, 16, 15, 19, 6,
	6, -1, 19, 13, -1, 10, -12, -11, -10, -9,
	10, 13, 4, 10, -3, 7, 11, 7, 11, -3,
	19, -3, -11, -9, 5,
}
var sesstypeDef = [...]int{

	0, -2, 1, 0, 8, 9, 10, -2, 0, 23,
	6, 7, 0, 12, 0, 17, 0, 0, 0, 0,
	0, 0, 0, 4, 0, 0, 0, 15, 0, 20,
	0, 3, 0, 0, 21, 11, 0, 16, 0, 13,
	0, 18, 14, 19, 5,
}
var sesstypeTok1 = [...]int{

	1,
}
var sesstypeTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19,
}
var sesstypeTok3 = [...]int{
	0,
}

var sesstypeErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	sesstypeDebug        = 0
	sesstypeErrorVerbose = false
)

type sesstypeLexer interface {
	Lex(lval *sesstypeSymType) int
	Error(s string)
}

type sesstypeParser interface {
	Parse(sesstypeLexer) int
	Lookahead() int
}

type sesstypeParserImpl struct {
	lval  sesstypeSymType
	stack [sesstypeInitialStackSize]sesstypeSymType
	char  int
}

func (p *sesstypeParserImpl) Lookahead() int {
	return p.char
}

func sesstypeNewParser() sesstypeParser {
	return &sesstypeParserImpl{}
}

const sesstypeFlag = -1000

func sesstypeTokname(c int) string {
	if c >= 1 && c-1 < len(sesstypeToknames) {
		if sesstypeToknames[c-1] != "" {
			return sesstypeToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func sesstypeStatname(s int) string {
	if s >= 0 && s < len(sesstypeStatenames) {
		if sesstypeStatenames[s] != "" {
			return sesstypeStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func sesstypeErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !sesstypeErrorVerbose {
		return "syntax error"
	}

	for _, e := range sesstypeErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + sesstypeTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := sesstypePact[state]
	for tok := TOKSTART; tok-1 < len(sesstypeToknames); tok++ {
		if n := base + tok; n >= 0 && n < sesstypeLast && sesstypeChk[sesstypeAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if sesstypeDef[state] == -2 {
		i := 0
		for sesstypeExca[i] != -1 || sesstypeExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; sesstypeExca[i] >= 0; i += 2 {
			tok := sesstypeExca[i]
			if tok < TOKSTART || sesstypeExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if sesstypeExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += sesstypeTokname(tok)
	}
	return res
}

func sesstypelex1(lex sesstypeLexer, lval *sesstypeSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = sesstypeTok1[0]
		goto out
	}
	if char < len(sesstypeTok1) {
		token = sesstypeTok1[char]
		goto out
	}
	if char >= sesstypePrivate {
		if char < sesstypePrivate+len(sesstypeTok2) {
			token = sesstypeTok2[char-sesstypePrivate]
			goto out
		}
	}
	for i := 0; i < len(sesstypeTok3); i += 2 {
		token = sesstypeTok3[i+0]
		if token == char {
			token = sesstypeTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = sesstypeTok2[1] /* unknown char */
	}
	if sesstypeDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", sesstypeTokname(token), uint(char))
	}
	return char, token
}

func sesstypeParse(sesstypelex sesstypeLexer) int {
	return sesstypeNewParser().Parse(sesstypelex)
}

func (sesstypercvr *sesstypeParserImpl) Parse(sesstypelex sesstypeLexer) int {
	var sesstypen int
	var sesstypeVAL sesstypeSymType
	var sesstypeDollar []sesstypeSymType
	_ = sesstypeDollar // silence set and not used
	sesstypeS := sesstypercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	sesstypestate := 0
	sesstypercvr.char = -1
	sesstypetoken := -1 // sesstypercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		sesstypestate = -1
		sesstypercvr.char = -1
		sesstypetoken = -1
	}()
	sesstypep := -1
	goto sesstypestack

ret0:
	return 0

ret1:
	return 1

sesstypestack:
	/* put a state and value onto the stack */
	if sesstypeDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", sesstypeTokname(sesstypetoken), sesstypeStatname(sesstypestate))
	}

	sesstypep++
	if sesstypep >= len(sesstypeS) {
		nyys := make([]sesstypeSymType, len(sesstypeS)*2)
		copy(nyys, sesstypeS)
		sesstypeS = nyys
	}
	sesstypeS[sesstypep] = sesstypeVAL
	sesstypeS[sesstypep].yys = sesstypestate

sesstypenewstate:
	sesstypen = sesstypePact[sesstypestate]
	if sesstypen <= sesstypeFlag {
		goto sesstypedefault /* simple state */
	}
	if sesstypercvr.char < 0 {
		sesstypercvr.char, sesstypetoken = sesstypelex1(sesstypelex, &sesstypercvr.lval)
	}
	sesstypen += sesstypetoken
	if sesstypen < 0 || sesstypen >= sesstypeLast {
		goto sesstypedefault
	}
	sesstypen = sesstypeAct[sesstypen]
	if sesstypeChk[sesstypen] == sesstypetoken { /* valid shift */
		sesstypercvr.char = -1
		sesstypetoken = -1
		sesstypeVAL = sesstypercvr.lval
		sesstypestate = sesstypen
		if Errflag > 0 {
			Errflag--
		}
		goto sesstypestack
	}

sesstypedefault:
	/* default state action */
	sesstypen = sesstypeDef[sesstypestate]
	if sesstypen == -2 {
		if sesstypercvr.char < 0 {
			sesstypercvr.char, sesstypetoken = sesstypelex1(sesstypelex, &sesstypercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if sesstypeExca[xi+0] == -1 && sesstypeExca[xi+1] == sesstypestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			sesstypen = sesstypeExca[xi+0]
			if sesstypen < 0 || sesstypen == sesstypetoken {
				break
			}
		}
		sesstypen = sesstypeExca[xi+1]
		if sesstypen < 0 {
			goto ret0
		}
	}
	if sesstypen == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			sesstypelex.Error(sesstypeErrorMessage(sesstypestate, sesstypetoken))
			Nerrs++
			if sesstypeDebug >= 1 {
				__yyfmt__.Printf("%s", sesstypeStatname(sesstypestate))
				__yyfmt__.Printf(" saw %s\n", sesstypeTokname(sesstypetoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for sesstypep >= 0 {
				sesstypen = sesstypePact[sesstypeS[sesstypep].yys] + sesstypeErrCode
				if sesstypen >= 0 && sesstypen < sesstypeLast {
					sesstypestate = sesstypeAct[sesstypen] /* simulate a shift of "error" */
					if sesstypeChk[sesstypestate] == sesstypeErrCode {
						goto sesstypestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if sesstypeDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", sesstypeS[sesstypep].yys)
				}
				sesstypep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if sesstypeDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", sesstypeTokname(sesstypetoken))
			}
			if sesstypetoken == sesstypeEofCode {
				goto ret1
			}
			sesstypercvr.char = -1
			sesstypetoken = -1
			goto sesstypenewstate /* try again in the same state */
		}
	}

	/* reduction by production sesstypen */
	if sesstypeDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", sesstypen, sesstypeStatname(sesstypestate))
	}

	sesstypent := sesstypen
	sesstypept := sesstypep
	_ = sesstypept // guard against "declared and not used"

	sesstypep -= sesstypeR2[sesstypen]
	// sesstypep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if sesstypep+1 >= len(sesstypeS) {
		nyys := make([]sesstypeSymType, len(sesstypeS)*2)
		copy(nyys, sesstypeS)
		sesstypeS = nyys
	}
	sesstypeVAL = sesstypeS[sesstypep+1]

	/* consult goto table to find next state */
	sesstypen = sesstypeR1[sesstypen]
	sesstypeg := sesstypePgo[sesstypen]
	sesstypej := sesstypeg + sesstypeS[sesstypep].yys + 1

	if sesstypej >= sesstypeLast {
		sesstypestate = sesstypeAct[sesstypeg]
	} else {
		sesstypestate = sesstypeAct[sesstypej]
		if sesstypeChk[sesstypestate] != -sesstypen {
			sesstypestate = sesstypeAct[sesstypeg]
		}
	}
	// dummy call; replaced with literal code
	switch sesstypent {

	case 1:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:32
		{
			local = sesstypeDollar[1].local
		}
	case 2:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:35
		{
			sesstypeVAL.role = sesstype.NewRole(sesstypeDollar[1].strval)
		}
	case 3:
		sesstypeDollar = sesstypeS[sesstypept-2 : sesstypept+1]
		//line local.y:38
		{
			sesstypeVAL.msg = sesstype.Message{Label: sesstypeDollar[1].strval}
		}
	case 4:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:39
		{
			sesstypeVAL.msg = sesstype.Message{Label: ""}
		}
	case 5:
		sesstypeDollar = sesstypeS[sesstypept-4 : sesstypept+1]
		//line local.y:40
		{
			sesstypeVAL.msg = sesstype.Message{Label: sesstypeDollar[1].strval, Payload: sesstypeDollar[3].strval}
		}
	case 6:
		sesstypeDollar = sesstypeS[sesstypept-2 : sesstypept+1]
		//line local.y:43
		{
			sesstypeVAL.local = sesstype.NewBranch(sesstypeDollar[1].role, sesstypeDollar[2].branch)
		}
	case 7:
		sesstypeDollar = sesstypeS[sesstypept-2 : sesstypept+1]
		//line local.y:44
		{
			sesstypeVAL.local = sesstype.NewSelect(sesstypeDollar[1].role, sesstypeDollar[2].branch)
		}
	case 8:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:45
		{
			sesstypeVAL.local = sesstypeDollar[1].local
		}
	case 9:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:46
		{
			sesstypeVAL.local = sesstypeDollar[1].local
		}
	case 10:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:47
		{
			sesstypeVAL.local = sesstypeDollar[1].local
		}
	case 11:
		sesstypeDollar = sesstypeS[sesstypept-4 : sesstypept+1]
		//line local.y:50
		{
			sesstypeVAL.branch = make(map[sesstype.Message]sesstype.Local)
			for _, l := range sesstypeDollar[3].sendrecvs {
				sesstypeVAL.branch[l.m] = l.l
			}
		}
	case 12:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:51
		{
			sesstypeVAL.branch = make(map[sesstype.Message]sesstype.Local)
			for _, l := range sesstypeDollar[1].sendrecvs {
				sesstypeVAL.branch[l.m] = l.l
			}
		}
	case 13:
		sesstypeDollar = sesstypeS[sesstypept-4 : sesstypept+1]
		//line local.y:54
		{
			sesstypeVAL.sendrecvs = []struct {
				m sesstype.Message
				l sesstype.Local
			}{struct {
				m sesstype.Message
				l sesstype.Local
			}{sesstypeDollar[2].msg, sesstypeDollar[4].local}}
		}
	case 14:
		sesstypeDollar = sesstypeS[sesstypept-3 : sesstypept+1]
		//line local.y:57
		{
			sesstypeVAL.sendrecvs = append(sesstypeDollar[1].sendrecvs, sesstypeDollar[3].sendrecvs...)
		}
	case 15:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:58
		{
			sesstypeVAL.sendrecvs = []struct {
				m sesstype.Message
				l sesstype.Local
			}{struct {
				m sesstype.Message
				l sesstype.Local
			}{m: sesstypeDollar[1].sendrecvs[0].m, l: sesstypeDollar[1].sendrecvs[0].l}}
		}
	case 16:
		sesstypeDollar = sesstypeS[sesstypept-4 : sesstypept+1]
		//line local.y:61
		{
			sesstypeVAL.branch = make(map[sesstype.Message]sesstype.Local)
			for _, s := range sesstypeDollar[3].sendrecvs {
				sesstypeVAL.branch[s.m] = s.l
			}
		}
	case 17:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:62
		{
			sesstypeVAL.branch = make(map[sesstype.Message]sesstype.Local)
			for _, s := range sesstypeDollar[1].sendrecvs {
				sesstypeVAL.branch[s.m] = s.l
			}
		}
	case 18:
		sesstypeDollar = sesstypeS[sesstypept-4 : sesstypept+1]
		//line local.y:65
		{
			sesstypeVAL.sendrecvs = []struct {
				m sesstype.Message
				l sesstype.Local
			}{struct {
				m sesstype.Message
				l sesstype.Local
			}{sesstypeDollar[2].msg, sesstypeDollar[4].local}}
		}
	case 19:
		sesstypeDollar = sesstypeS[sesstypept-3 : sesstypept+1]
		//line local.y:68
		{
			sesstypeVAL.sendrecvs = append(sesstypeDollar[1].sendrecvs, sesstypeDollar[3].sendrecvs...)
		}
	case 20:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:69
		{
			sesstypeVAL.sendrecvs = []struct {
				m sesstype.Message
				l sesstype.Local
			}{struct {
				m sesstype.Message
				l sesstype.Local
			}{m: sesstypeDollar[1].sendrecvs[0].m, l: sesstypeDollar[1].sendrecvs[0].l}}
		}
	case 21:
		sesstypeDollar = sesstypeS[sesstypept-4 : sesstypept+1]
		//line local.y:72
		{
			sesstypeVAL.local = sesstype.NewLRecur(sesstypeDollar[2].strval, sesstypeDollar[4].local)
		}
	case 22:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:75
		{
			sesstypeVAL.local = sesstype.NewLTypeVar(sesstypeDollar[1].strval)
		}
	case 23:
		sesstypeDollar = sesstypeS[sesstypept-1 : sesstypept+1]
		//line local.y:78
		{
			sesstypeVAL.local = sesstype.NewLEnd()
		}
	}
	goto sesstypestack /* stack new state and value */
}
