// Copyright 2024 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by goyacc -o xpath.go -p expr xpath.y. DO NOT EDIT.

//line xpath.y:19

package expr

import __yyfmt__ "fmt"

//line xpath.y:20

import (
	"encoding/xml"

	"github.com/sdcio/yang-parser/xpath"
	"github.com/sdcio/yang-parser/xpath/xutils"
)

//line xpath.y:31
type exprSymType struct {
	yys     int
	sym     *xpath.Symbol /* Symbol table entry */
	val     float64       /* Numeric value */
	name    string        /* NodeType or AxisName */
	xmlname xml.Name      /* For NameTest */
}

const NUM = 57346
const DOTDOT = 57347
const DBLSLASH = 57348
const DBLCOLON = 57349
const ERR = 57350
const FUNC = 57351
const NODETYPE = 57352
const AXISNAME = 57353
const LITERAL = 57354
const NAMETEST = 57355
const OR = 57356
const AND = 57357
const NE = 57358
const EQ = 57359
const GT = 57360
const GE = 57361
const LT = 57362
const LE = 57363
const DIV = 57364
const MOD = 57365
const UNARYMINUS = 57366

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"DOTDOT",
	"DBLSLASH",
	"DBLCOLON",
	"ERR",
	"FUNC",
	"NODETYPE",
	"AXISNAME",
	"LITERAL",
	"NAMETEST",
	"OR",
	"AND",
	"NE",
	"EQ",
	"GT",
	"GE",
	"LT",
	"LE",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"DIV",
	"MOD",
	"UNARYMINUS",
	"'|'",
	"'('",
	"')'",
	"','",
	"'['",
	"']'",
	"'.'",
	"'@'",
}

var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16

//line xpath.y:335

/* Code is in .go files so we get the benefit of gofmt etc.
 * What's above is formatted as best as emacs Bison-mode will allow,
 * with semi-colons added to help Bison-mode think the code is C!
 *
 * If anyone can come up with a better formatting model I'm all ears ... (-:
 */

//line yacctab:1
var exprExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 14,
	6, 30,
	25, 30,
	-2, 27,
}

const exprPrivate = 57344

const exprLast = 168

var exprAct = [...]int8{
	2, 67, 19, 32, 66, 12, 8, 6, 4, 9,
	5, 96, 100, 101, 16, 57, 55, 97, 98, 59,
	61, 54, 103, 37, 63, 90, 64, 53, 41, 33,
	7, 35, 50, 35, 51, 52, 62, 40, 29, 38,
	95, 48, 49, 45, 47, 44, 46, 68, 69, 70,
	72, 73, 71, 36, 39, 78, 79, 85, 60, 83,
	80, 81, 82, 88, 89, 92, 61, 65, 94, 84,
	93, 56, 61, 86, 87, 74, 75, 76, 77, 25,
	37, 38, 38, 34, 26, 27, 33, 24, 35, 30,
	61, 61, 28, 43, 42, 94, 20, 22, 11, 99,
	31, 58, 102, 21, 17, 23, 91, 25, 37, 38,
	36, 39, 26, 27, 33, 24, 35, 18, 15, 14,
	13, 10, 3, 1, 0, 0, 11, 0, 31, 0,
	0, 0, 0, 23, 0, 25, 37, 38, 36, 39,
	26, 27, 33, 24, 35, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 31, 0, 0, 0,
	0, 23, 0, 0, 0, 0, 36, 39,
}

var exprPact = [...]int16{
	103, -1000, -1000, 23, 13, 77, 25, 19, 8, -1000,
	-2, 103, -1000, -1000, -18, 76, 33, -1000, -1000, -1000,
	-1000, 18, -1000, 103, -1000, -1000, -4, -1000, 20, -18,
	-1000, -1000, 18, 41, -1000, -1000, -1000, -1000, -1000, -1000,
	103, 103, 103, 103, 103, 103, 103, 103, 103, 103,
	103, 103, 103, 131, -1000, -1000, 103, -1000, 18, 18,
	18, 18, 33, -6, 75, -18, -18, -1000, 33, -1000,
	13, 77, 25, 25, 19, 19, 19, 19, 8, 8,
	-1000, -1000, -1000, -1000, -23, -1000, 33, 33, -1000, -1000,
	-1000, -1000, -14, -18, -1000, -1000, -1000, -1000, 103, -19,
	-1000, 103, -9, -1000,
}

var exprPgo = [...]int8{
	0, 123, 0, 122, 8, 10, 7, 30, 6, 9,
	121, 5, 120, 119, 118, 14, 3, 117, 1, 104,
	103, 97, 2, 96, 92, 38, 4, 89, 83, 71,
	69, 40,
}

var exprR1 = [...]int8{
	0, 1, 2, 3, 3, 4, 4, 5, 5, 5,
	6, 6, 6, 6, 6, 7, 7, 7, 8, 8,
	8, 8, 9, 9, 10, 10, 11, 11, 11, 11,
	14, 13, 13, 17, 17, 17, 17, 17, 17, 17,
	17, 12, 12, 19, 19, 19, 20, 15, 15, 15,
	22, 22, 22, 22, 22, 24, 24, 25, 26, 26,
	18, 29, 30, 31, 21, 23, 27, 27, 28, 16,
}

var exprR2 = [...]int8{
	0, 1, 1, 1, 3, 1, 3, 1, 3, 3,
	1, 3, 3, 3, 3, 1, 3, 3, 1, 3,
	3, 3, 1, 2, 1, 3, 1, 1, 3, 3,
	1, 1, 2, 3, 1, 1, 3, 4, 6, 8,
	1, 1, 1, 1, 2, 1, 1, 1, 3, 1,
	3, 2, 2, 1, 1, 2, 1, 1, 1, 2,
	3, 1, 1, 1, 2, 3, 1, 1, 1, 1,
}

var exprChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, 23, -11, -12, -13, -14, -15, -19, -17, -22,
	-23, -20, -21, 30, 12, 4, 9, 10, -24, -25,
	-27, 25, -16, 11, -28, 13, 35, 5, 6, 36,
	14, 15, 17, 16, 20, 18, 21, 19, 22, 23,
	24, 26, 27, 29, -9, -18, -29, 33, 25, -16,
	25, -16, -15, -2, 30, -25, -26, -18, -15, 7,
	-4, -5, -6, -6, -7, -7, -7, -7, -8, -8,
	-9, -9, -9, -11, -30, -2, -15, -15, -22, -22,
	31, 31, -2, -26, -18, -31, 34, 31, 32, -2,
	31, 32, -2, 31,
}

var exprDef = [...]int8{
	0, -2, 1, 2, 3, 5, 7, 10, 15, 18,
	22, 0, 24, 26, -2, 0, 41, 42, 31, 47,
	49, 43, 45, 0, 34, 35, 0, 40, 0, 53,
	54, 46, 0, 0, 56, 57, 66, 67, 69, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 32, 0, 61, 0, 0,
	0, 0, 44, 0, 0, 51, 52, 58, 64, 55,
	4, 6, 8, 9, 11, 12, 13, 14, 16, 17,
	19, 20, 21, 25, 0, 62, 28, 29, 48, 65,
	33, 36, 0, 50, 59, 60, 63, 37, 0, 0,
	38, 0, 0, 39,
}

var exprTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	30, 31, 24, 22, 32, 23, 35, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 33, 3, 34, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 29,
}

var exprTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	26, 27, 28,
}

var exprTok3 = [...]int8{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(exprPact[state])
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && int(exprChk[int(exprAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || int(exprExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := int(exprExca[i])
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(exprTok1[0])
		goto out
	}
	if char < len(exprTok1) {
		token = int(exprTok1[char])
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = int(exprTok2[char-exprPrivate])
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = int(exprTok3[i+0])
		if token == char {
			token = int(exprTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(exprTok2[1]) /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = int(exprPact[exprstate])
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = int(exprAct[exprn])
	if int(exprChk[exprn]) == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = int(exprDef[exprstate])
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && int(exprExca[xi+1]) == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = int(exprExca[xi+0])
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = int(exprExca[xi+1])
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = int(exprPact[exprS[exprp].yys]) + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = int(exprAct[exprn]) /* simulate a shift of "error" */
					if int(exprChk[exprstate]) == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= int(exprR2[exprn])
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = int(exprR1[exprn])
	exprg := int(exprPgo[exprn])
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = int(exprAct[exprg])
	} else {
		exprstate = int(exprAct[exprj])
		if int(exprChk[exprstate]) != -exprn {
			exprstate = int(exprAct[exprg])
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:60
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Store, "store")
		}
	case 2:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:67
		{
			getProgBldr(exprlex).CurrentFix()
		}
	case 4:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:74
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Or, "or")
		}
	case 6:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:82
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).And, "and")
		}
	case 8:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:90
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Eq, "eq")
		}
	case 9:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:95
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Ne, "ne")
		}
	case 11:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:103
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Lt, "lt")
		}
	case 12:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:108
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Gt, "gt")
		}
	case 13:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:113
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Le, "le")
		}
	case 14:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:118
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Ge, "ge")
		}
	case 16:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:126
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Add, "add")
		}
	case 17:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:131
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Sub, "sub")
		}
	case 19:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:139
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Mul, "mul")
		}
	case 20:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:144
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Div, "div")
		}
	case 21:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:149
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Mod, "mod")
		}
	case 23:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line xpath.y:157
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Negate, "negate")
		}
	case 25:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:165
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).Union, "union")
		}
	case 26:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:172
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).EvalLocPath, "evalLocPath")
		}
	case 28:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:178
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).EvalLocPath, "evalLocPath")
		}
	case 29:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:183
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).EvalLocPath, "evalLocPath")
		}
	case 30:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:193
		{
			getProgBldr(exprlex).CodeFn(
				getProgBldr(exprlex).FilterExprEnd, "filterExprEnd")
		}
	case 34:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:205
		{
			getProgBldr(exprlex).CodeLiteral(exprDollar[1].name)
		}
	case 35:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:209
		{
			getProgBldr(exprlex).CodeNum(exprDollar[1].val)
		}
	case 36:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line xpath.y:213
		{
			getProgBldr(exprlex).CodeBltin(exprDollar[1].sym, 0)
		}
	case 37:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line xpath.y:217
		{
			getProgBldr(exprlex).CodeBltin(exprDollar[1].sym, 1)
		}
	case 38:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line xpath.y:221
		{
			getProgBldr(exprlex).CodeBltin(exprDollar[1].sym, 2)
		}
	case 39:
		exprDollar = exprS[exprpt-8 : exprpt+1]
//line xpath.y:225
		{
			getProgBldr(exprlex).CodeBltin(exprDollar[1].sym, 3)
		}
	case 40:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:229
		{
			getProgBldr(exprlex).UnsupportedName(xutils.NODETYPE, exprDollar[1].name)
		}
	case 46:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:250
		{
			getProgBldr(exprlex).CodePathOper('/')
		}
	case 55:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line xpath.y:275
		{
			getProgBldr(exprlex).UnsupportedName(xutils.AXISNAME, exprDollar[1].name)
		}
	case 57:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:281
		{
			getProgBldr(exprlex).CodeNameTest(exprDollar[1].xmlname)
		}
	case 61:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:293
		{
			getProgBldr(exprlex).CodePredStart()
		}
	case 63:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:301
		{
			getProgBldr(exprlex).CodePredEnd()
		}
	case 66:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:313
		{
			getProgBldr(exprlex).CodePathOper('.')
		}
	case 67:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:317
		{
			getProgBldr(exprlex).CodePathOper(xutils.DOTDOT)
		}
	case 68:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:323
		{
			getProgBldr(exprlex).UnsupportedName(
				'@', "not yet implemented")
		}
	case 69:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line xpath.y:330
		{
			getProgBldr(exprlex).UnsupportedName(
				xutils.DBLSLASH, "not yet implemented")
		}
	}
	goto exprstack /* stack new state and value */
}
