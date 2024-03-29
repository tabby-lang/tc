// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package checker

// operations
const (
	ADD     = "ADD"
	EQUAL   = "EQ"
	LT      = "LT"
	GT      = "GT"
	SUB     = "SUB"
	MUL     = "MUL"
	DIV     = "DIV"
	AND     = "AND"
	OR      = "OR"
	println = "println"
	print   = "print"
)

// variable types
const (
	INT_TYPE     = "Int"
	STRING_TYPE  = "String"
	BOOL_TYPE    = "Bool"
	NOTHING_TYPE = "Nothing"
)

type Signature struct {
	Return string
	Params []string // list of types
}

type Methods map[string]Signature

// type methods
var TypeTable = map[string]Methods{
	INT_TYPE: {
		ADD:     {INT_TYPE, []string{INT_TYPE}},
		SUB:     {INT_TYPE, []string{INT_TYPE}},
		MUL:     {INT_TYPE, []string{INT_TYPE}},
		DIV:     {INT_TYPE, []string{INT_TYPE}},
		LT:      {BOOL_TYPE, []string{INT_TYPE}},
		GT:      {BOOL_TYPE, []string{INT_TYPE}},
		EQUAL:   {BOOL_TYPE, []string{INT_TYPE}},
		println: {NOTHING_TYPE, []string{}},
		print:   {NOTHING_TYPE, []string{}}},
	STRING_TYPE: {
		ADD:     {STRING_TYPE, []string{STRING_TYPE}},
		println: {NOTHING_TYPE, []string{}},
		print:   {NOTHING_TYPE, []string{}}},
	BOOL_TYPE: {
		AND:     {BOOL_TYPE, []string{BOOL_TYPE}},
		OR:      {BOOL_TYPE, []string{BOOL_TYPE}},
		println: {NOTHING_TYPE, []string{}},
		print:   {NOTHING_TYPE, []string{}}},
}

type Environment struct {
	Vals  map[string]string    // map identifier to type
	Funcs map[string]Signature // map function name to return type
	Types map[string]bool      // track valid types
}

var env Environment // set global

func IsBuiltin(name string) bool {
	return name == "println" || name == "print"
}

func NewEnvironment() Environment {
	return Environment{Vals: map[string]string{}, Funcs: map[string]Signature{}, Types: map[string]bool{}}
}

func MethodExist(kind, method string) bool {
	methods, ok := TypeTable[kind]
	if !ok {
		return false
	}

	_, ok = methods[method]
	return ok
}

func GetMethod(kind, method string) (Signature, bool) {
	methods, ok := TypeTable[kind]
	if !ok {
		return Signature{}, false
	}

	sig, ok := methods[method]
	return sig, ok
}

func (e *Environment) Set(name, kind string) {
	e.Vals[name] = kind
}

func SetFunctionSignature(name string, sig Signature) {
	env.Funcs[name] = sig
}

func GetFunctionSignature(name string) (Signature, bool) {
	kind, ok := env.Funcs[name]
	return kind, ok
}

func (e *Environment) Get(name string) (string, bool) {
	kind, ok := e.Vals[name]
	return kind, ok
}

func (e *Environment) IdentExist(kind string) bool {
	_, ok := e.Vals[kind]
	return ok
}

func GetIdentType(name string) (string, bool) {
	kind, ok := env.Vals[name]
	return kind, ok
}

func (e *Environment) TypeExist(kind string) bool {
	_, ok := e.Types[kind]
	return ok
}
