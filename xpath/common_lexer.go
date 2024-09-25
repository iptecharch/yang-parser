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

// Copyright (c) 2019-2021, AT&T Intellectual Property. All rights reserved.
//
// Copyright (c) 2015-2017 by Brocade Communications Systems, Inc.
// All rights reserved.
//
// SPDX-License-Identifier: MPL-2.0

// Credit for the 'next' function, and initial 'lex' function go to whoever
// wrote the 'expr' YACC example in the Go source code.

// This file implements XPATH lexing / tokenisation for YANG.  Specifically,
// it diverges from a complete XPATH implementation as follows:
//
// (a) As YANG uses only the core function set, we do not accept fully-
//     qualified function names (eg prefix:fn_name())
//
// Different YANG statements use different subsets of XPATH.  This file
// contains the common lexing code, with customisations separated out into
// specific _lexer.go files that live with their associated <prefix>.y YACC
// grammar files.

package xpath

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/sdcio/yang-parser/xpath/xutils"
)

// Allow for different grammars to be compiled and run using Machine,
// sharing common lexing code where possible.
type XpathLexer interface {
	GetError() error
	SetError(err error)
	GetProgBldr() *ProgBuilder
	GetLine() []byte
	Parse()

	Next() rune
	SaveTokenType(tokenType int)
	IsNameChar(c rune) bool
	IsNameStartChar(c rune) bool
	MapTokenValToCommon(tokenType int) int

	LexLiteral(quote rune) (int, TokVal)
	LexDot(c rune) (int, TokVal)
	LexNum(c rune) (int, TokVal)
	LexSlash() (int, TokVal)
	LexColon() (int, TokVal)
	LexAsterisk() (int, TokVal)
	LexRelationalOperator(c rune) (int, TokVal)
	LexName(c rune) (int, TokVal)
	LexPunctuation(c rune) (int, TokVal)
}

type TokVal interface{}

// COMMONLEX
type CommonLex struct {
	// Exported via accessors.
	line     []byte
	err      error
	mapFn    PfxMapFn
	progBldr *ProgBuilder // Used to build the program to be run later.

	// Internal use only
	peek           rune
	precToken      int  // Preceding token type, if any (otherwise EOF)
	allowCustomFns bool // Expr may use custom XPATH functions
	userFnChecker  UserCustomFunctionCheckerFn
}

func NewCommonLex(
	line []byte,
	progBldr *ProgBuilder,
	mapFn PfxMapFn,
) CommonLex {
	return CommonLex{line: line, progBldr: progBldr, mapFn: mapFn}
}

func (lexer *CommonLex) AllowCustomFns() *CommonLex {
	lexer.allowCustomFns = true
	return lexer
}

func (lexer *CommonLex) SetUserFnChecker(
	userFnChecker UserCustomFunctionCheckerFn,
) {
	lexer.userFnChecker = userFnChecker
}

func (lexer *CommonLex) Parse() {
	panic("CommonLex doesn't implement Parse()")
}

func (lexer *CommonLex) CreateProgram(expr string) (prog []Inst, err error) {
	if lexer.progBldr.parseErr == nil && lexer.GetError() == nil {
		return lexer.progBldr.GetMainProg()
	}

	errors := fmt.Sprintf("Failed to compile '%s'\n", expr)
	currentPosInLine :=
		len(string(expr)) - len(string(lexer.progBldr.lineAtErr))
	parsedLine := string(expr)[:currentPosInLine]
	unParsedLine := string(expr)[currentPosInLine:]

	if lexer.GetError() != nil {
		errors += fmt.Sprintf("Lexer Error: %s\n",
			lexer.GetError().Error())
	}

	if lexer.progBldr.parseErr != nil {
		errors += fmt.Sprintf("Parse Error: %s\n",
			lexer.progBldr.parseErr.Error())
	}

	return nil, fmt.Errorf("%s\nGot to approx [X] in '%s [X] %s'\n", errors,
		parsedLine, unParsedLine)
}

func (lexer *CommonLex) GetError() error { return lexer.err }

func (lexer *CommonLex) SetError(err error) { lexer.err = err }

func (lexer *CommonLex) GetProgBldr() *ProgBuilder {
	return lexer.progBldr
}

func (lexer *CommonLex) GetLine() []byte { return lexer.line }

func (lexer *CommonLex) GetMapFn() PfxMapFn { return lexer.mapFn }

// The parser calls this method on a parse error.  It stores the error in the
// machine for later retrieval.
func (x *CommonLex) Error(s string) {
	if x.progBldr.parseErr != nil {
		// Use first error found, if more than one detected.
		return
	}
	x.progBldr.parseErr = fmt.Errorf("%s", s)
	if x.peek != xutils.EOF {
		x.progBldr.lineAtErr = string(x.peek) + string(x.line)
	} else {
		x.progBldr.lineAtErr = string(x.line)
	}
}

// Some parsing will produce different tokens depending on what came before
// so we need to keep track of this.
func (x *CommonLex) SaveTokenType(tokenType int) {
	if tokenType != xutils.EOF && tokenType != xutils.ERR {
		x.precToken = tokenType
	}
}

// The parser calls this method to get each new token.
//
// We store the token value so it is available as the preceding token
// value when parsing the next token.
func LexCommon(x XpathLexer) (tokType int, tokVal TokVal) {
	defer func() {
		x.SaveTokenType(tokType)
	}()
	for {
		c := x.Next()
		switch c {
		case xutils.EOF:
			return xutils.EOF, nil

		case xutils.ERR:
			x.SetError(fmt.Errorf("Invalid UTF-8 input"))
			return xutils.ERR, nil

		case '"', '\'':
			return x.LexLiteral(c)

		case '.':
			return x.LexDot(c)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.LexNum(c)

		case '/':
			return x.LexSlash()

		case ':':
			return x.LexColon()

		case '*':
			return x.LexAsterisk()

		case '+', '-', '(', ')', '@', ',', '[', ']', '|':
			return x.LexPunctuation(c)

		case '=', '>', '<', '!':
			return x.LexRelationalOperator(c)

		case ' ', '\t', '\n', '\r':
			// Deal with whitespace by ignoring it
			continue
		}

		// Names of some form or another ... NameTest, NodeType,
		// OperatorName, FunctionName, or AxisName
		if x.IsNameStartChar(c) {
			return x.LexName(c)
		}

		x.SetError(fmt.Errorf("unrecognised character %q", c))
		return xutils.ERR, nil
	}
}

// Separated out to allow us to override it.
func (x *CommonLex) LexPunctuation(c rune) (int, TokVal) {
	return int(c), nil
}

func (x *CommonLex) LexDot(c rune) (int, TokVal) {
	// Could be '.', '..', or number
	next := x.Next()
	switch next {
	case '.':
		return xutils.DOTDOT, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		x.peek = next
		return x.LexNum(c)
	default:
		x.peek = next
		return '.', nil
	}
}

func (x *CommonLex) LexSlash() (int, TokVal) {
	// Could be '/' or '//'.  NB - this is not 'divide', ever.
	next := x.Next()
	if next == '/' {
		return xutils.DBLSLASH, nil
	}
	x.peek = next
	return '/', nil
}

func (x *CommonLex) LexColon() (int, TokVal) {
	// Should be '::' as single colons are only allowed within QNames and
	// are not detected in main lexer loop.
	next := x.Next()
	if next == ':' {
		return xutils.DBLCOLON, nil
	}
	// Part of a name, should have been detected elsewhere.
	x.peek = next
	x.SetError(fmt.Errorf("':' only supported in QNames"))
	return xutils.ERR, nil
}

func (x *CommonLex) LexAsterisk() (int, TokVal) {
	if x.tokenCanBeOperator() {
		return '*', nil
	}

	// This is the global wildcard representing all child nodes, regardless
	// of module.
	return xutils.NAMETEST, xutils.AllChildren.Name()
}

func (x *CommonLex) LexRelationalOperator(c rune) (int, TokVal) {
	switch c {
	case '=':
		return xutils.EQ, nil
	case '>':
		next := x.Next()
		if next == '=' {
			return xutils.GE, nil
		}
		x.peek = next
		return xutils.GT, nil

	case '<':
		next := x.Next()
		if next == '=' {
			return xutils.LE, nil
		}
		x.peek = next
		return xutils.LT, nil

	case '!':
		next := x.Next()
		if next == '=' {
			return xutils.NE, nil
		}
		x.peek = next
		x.SetError(fmt.Errorf("'!' only valid when followed by '='"))
		return xutils.ERR, nil
	default:
		x.SetError(fmt.Errorf("Invalid relational operator"))
		return xutils.ERR, nil
	}
}

// Lex a non-literal name (ie something textual that isn't quoted).
//
// Rules for disambiguating:
//
// (a) If there is a preceding token, and said token is none of '@', '::',
//
//	'(', '[', ',' or an Operator, then '*' is the MultiplyOperator and
//	NCName must be recognised as an OperatorName
//
// (b) If the character following an NCName (possibly after intervening
//
//	whitespace) is '(', then the token must be recognized as a NodeType
//	or FunctionName
//
// (c) If an NCName is followed by '::' (possibly with intervening whitespace)
//
//	then the NCName must be recognised as an AxisName
//
// (d) In all other cases, the token must NOT be recognised as a Multiply
//
//	Operator, OperatorName, NodeType, FunctionName, or AxisName
func (x *CommonLex) LexName(c rune) (int, TokVal) {
	nameMatcher := func(c rune) bool {
		if x.IsNameChar(c) {
			return true
		}
		return false
	}

	// Next get 'NCName'
	name := x.ConstructToken(c, nameMatcher, "NAME")

	// If there's a preceding token, and it's not '@', '::', '(', '[', ',' or
	// an Operator then NCName is an OperatorName
	if x.tokenCanBeOperator() {
		return x.getOperatorName(name.String()), nil
	}

	// If next non-whitespace character is '(' then this must be a NodeType
	// or a FunctionName
	if x.NextNonWhitespaceStringIs("(") {
		switch name.String() {
		case "text":
			fn, ok := LookupXpathFunction(
				"text",
				false,
				nil)
			if ok {
				return xutils.TEXTFUNC, fn
			}
		case "current":
			return xutils.CURRENTFUNC, nil
		case "deref":
			return xutils.DEREFFUNC, nil
		default:
			if x.nameIsNodeType(name.String()) {
				return xutils.NODETYPE, name.String()
			}
		}

		fn, ok := LookupXpathFunction(
			name.String(),
			x.allowCustomFns,
			x.userFnChecker)
		if ok {
			return xutils.FUNC, fn
		}
		x.SetError(fmt.Errorf("Unknown function or node type: '%s'",
			name.String()))
		return xutils.ERR, nil
	}

	// If next non-whitespace token is '::', NCName is an AxisName.
	if x.NextNonWhitespaceStringIs("::") {
		if x.nameIsAxisName(name.String()) {
			return xutils.AXISNAME, name.String()
		}
		x.SetError(fmt.Errorf("Unknown axis name: '%s'", name.String()))
		return xutils.ERR, nil
	}

	// If none of the above applies, it's a NameTest token.  Question is
	// whether it's a Prefixed or Unprefixed ... so let's see if we have a
	// colon following.  As we already checked for '::', we can safely check
	// for single ':'
	var namespace, localPart, prefix string
	if x.NextNonWhitespaceStringIs(":") {
		// Prefixed, so it's either Prefix:LocalPart or Prefix:*
		// Next token had better be a ':' when formally extracted ...
		if c := x.NextNonWhitespace(); c != ':' {
			x.SetError(fmt.Errorf(
				"Badly formatted QName (exp ':', got '%c'", c))
			return xutils.ERR, nil
		}

		// Now we need the local part - or wildcard (*).  Note that in the
		// latter case this must be 'NCName:*' - the global wildcard '*' is
		// handled by LexAsterisk().
		if x.NextNonWhitespaceStringIs("*") {
			// Next token had better be a '*' when formally extracted ...
			if c := x.NextNonWhitespace(); c != '*' {
				x.SetError(fmt.Errorf("Badly formatted QName (*)."))
				return xutils.ERR, nil
			}
			prefix = name.String()
			localPart = "*"
		} else {
			// We need to extract the second part of the name.
			c := x.NextNonWhitespace()
			if c == xutils.EOF {
				x.err = fmt.Errorf("Name requires local part.")
				return xutils.ERR, nil
			}
			localPartBuf := x.ConstructToken(c, nameMatcher, "NAME")
			localPart = localPartBuf.String()
			prefix = name.String()
		}

	} else {
		localPart = name.String()
	}

	// If we have a mapping function, map the locally-scoped (within namespace)
	// prefix name (if present) to a globally scoped namespace.
	if x.mapFn != nil {
		var err error
		namespace, err = x.mapFn(prefix)
		if err != nil {
			x.SetError(err)
			return xutils.ERR, nil
		}
	}

	return xutils.NAMETEST, xml.Name{Space: namespace, Local: localPart}
}

// Lex 'literal' string contained in single or double quotes
func (x *CommonLex) LexLiteral(quote rune) (int, TokVal) {
	literalMatcher := func(c rune) bool {
		if c != quote {
			return true
		}
		return false
	}

	// Skip initial quote - start from 'next'.  As constructToken always
	// adds first character, we also need to detect empty strings here.
	var b bytes.Buffer
	c := x.Next()
	if c != quote {
		b = x.ConstructToken(c, literalMatcher,
			xutils.GetTokenName(xutils.LITERAL))
		// Skip final quote character.
		x.Next()
	}

	if x.err != nil {
		return xutils.ERR, nil
	}
	return xutils.LITERAL, b.String()
}

// Lex a number.
func (x *CommonLex) LexNum(c rune) (int, TokVal) {
	numMatcher := func(c rune) bool {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			return true
		}
		return false
	}
	b := x.ConstructToken(c, numMatcher, xutils.GetTokenName(xutils.NUM))
	val, err := strconv.ParseFloat(b.String(), 10)

	if err != nil {
		x.SetError(fmt.Errorf("bad number %q", b.String()))
		return xutils.ERR, nil
	}
	return xutils.NUM, val
}

// An operator cannot follow a specific set of other tokens, which include
// other operators (quite reasonably).  See XPATH section 3.7.
func (x *CommonLex) tokenCanBeOperator() bool {
	// Split into 3 cases to avoid wrapping and to match the order in the
	// XPATH spec (ambiguity rules for ExprToken)
	switch x.precToken {
	case xutils.EOF, '@', xutils.DBLCOLON, '(', '[', ',':
		return false

	case xutils.AND, xutils.OR, xutils.MOD, xutils.DIV:
		return false

	case '*', '/', xutils.DBLSLASH, '|', '+', '-',
		xutils.EQ, xutils.NE, xutils.LT, xutils.LE, xutils.GT, xutils.GE:
		return false
	}

	return true
}

// Useful for any multi-character token in conjunction with constructToken()
type tokenMatcherFn func(c rune) bool

// Given first character in token and function to identify further elements,
// return full token and set x.peek to the correct character.
func (x *CommonLex) ConstructToken(
	c rune,
	tokenMatcher tokenMatcherFn,
	tokenName string,
) bytes.Buffer {

	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			x.SetError(fmt.Errorf("WriteRune: %s", err))
		}
	}
	var b bytes.Buffer
	add(&b, c)

	for {
		c = x.Next()
		if tokenMatcher(c) {
			// As a sanity check against rogue tokenMatcher functions that fail
			// to spot EOF and claim a match, trap it here.  It's also rather
			// easier to spot here in the guts of the processing anyway.
			if c == xutils.EOF {
				x.SetError(fmt.Errorf("End of %s token not detected.",
					tokenName))
				break
			}
			add(&b, c)
		} else {
			break
		}
	}

	x.peek = c

	return b
}

func (x *CommonLex) IsNameStartChar(c rune) bool {
	switch {
	case (c >= 'A') && (c <= 'Z'):
		return true
	case c == '_':
		return true
	case (c >= 'a') && (c <= 'z'):
		return true
	case (c >= 0xC0) && (c <= 0xD6):
		return true
	case (c >= 0xD8) && (c <= 0xF6):
		return true
	case (c >= 0xF8) && (c <= 0x2FF):
		return true
	case (c >= 0x370) && (c <= 0x37D):
		return true
	case (c >= 0x37F) && (c <= 0x1FFF):
		return true
	case (c >= 0x200C) && (c <= 0x200D):
		return true
	case (c >= 0x2070) && (c <= 0x218F):
		return true
	case (c >= 0x2C00) && (c <= 0x2FEF):
		return true
	case (c >= 0x3001) && (c <= 0xD7FF):
		return true
	case (c >= 0xF900) && (c <= 0xFDCF):
		return true
	case (c >= 0xFDF0) && (c <= 0xFFFD):
		return true
	case (c >= 0x10000) && (c <= 0xEFFFF):
		return true
	default:
		return false
	}
}

func (x *CommonLex) IsNameChar(c rune) bool {
	switch {
	case x.IsNameStartChar(c):
		return true
	case c == '-' || c == '.':
		return true
	case (c >= '0') && (c <= '9'):
		return true
	case c == 0xB7:
		return true
	case (c >= 0x300) && (c <= 0x36F):
		return true
	case (c >= 0x203F) && (c <= 0x2040):
		return true
	default:
		return false
	}
}

func (x *CommonLex) getOperatorName(name string) int {
	switch name {
	case "and":
		return xutils.AND
	case "or":
		return xutils.OR
	case "mod":
		return xutils.MOD
	case "div":
		return xutils.DIV
	}

	x.SetError(fmt.Errorf("Unrecognised operator name: '%s'", name))
	return xutils.ERR
}

func (x *CommonLex) nameIsNodeType(name string) bool {
	switch name {
	case "comment", "text", "processing-instruction", "node":
		return true
	}

	return false
}

func (x *CommonLex) nameIsAxisName(name string) bool {
	switch name {
	case "ancestor-or-self", "attribute", "child", "descendant",
		"descendant-or-self", "following", "following-sibling",
		"namespace", "parent", "preceding", "preceding-sibling", "self":
		return true
	}

	return false
}

// Return the next rune for the lexer.  'peek' may have been set if we
// needed to look ahead but then didn't consume the character.  In other
// words, what remains to be parsed when we call Next() is:
//
//	x.peek (if not EOF) + x.line
func (x *CommonLex) Next() rune {
	if x.peek != xutils.EOF {
		r := x.peek
		x.peek = xutils.EOF
		return r
	}
	if len(x.line) == 0 {
		return xutils.EOF
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError && size == 1 {
		return xutils.ERR
	}
	return c
}

func (x *CommonLex) isWhitespace(c rune) bool {
	switch c {
	case '\t', '\r', '\n', ' ':
		return true
	}

	return false
}

func (x *CommonLex) NextNonWhitespace() rune {
	c := x.Next()

	for c != xutils.EOF && x.isWhitespace(c) {
		c = x.Next()
	}

	return c
}

func next(line []byte) (rune, []byte) {
	if len(line) == 0 {
		return xutils.EOF, nil
	}
	c, size := utf8.DecodeRune(line)
	line = line[size:]
	if c == utf8.RuneError && size == 1 {
		return xutils.ERR, nil
	}
	return c, line
}

// Won't handle string containing whitespace.
// For now we only need this to match '(', '::', ':' and '*'.
// This assumes the passed in string consists of ASCII bytes
func (x *CommonLex) NextNonWhitespaceStringIs(expr string) bool {

	// First check peek (if in use) and if not whitespace, compare.
	// The ASCII assumption is here, it could be written using
	// utf8.RuneLen(), but that is not necessary.
	if (x.peek != xutils.EOF) && !x.isWhitespace(x.peek) {
		if len(expr) == 0 {
			return true
		}
		if x.peek != rune(expr[0]) {
			return false
		}
		if len(expr) == 1 {
			return true
		}
		expr = expr[1:]
	}

	// Next, skip any whitespace
	lc, line := next(x.line)
	for x.isWhitespace(lc) {
		lc, line = next(line)
	}

	// Now compare the rest of the string against the input
	for _, ec := range expr {
		if lc == xutils.EOF || lc == xutils.ERR {
			return false
		}
		if ec != lc {
			return false
		}
		lc, line = next(line)
	}

	return true
}
