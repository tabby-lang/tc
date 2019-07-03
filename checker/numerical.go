// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package checker

import (
	"fmt"
	"reflect"

	"github.com/tabby-lang/tc/bounds"
	"github.com/tabby-lang/tc/token"
)

// CheckOverflow checks a numeric literal against its type constraints,
// returning an overflow error when the value is out of range.
func CheckOverflow(num interface{}, lit *token.Token) error {
	t := reflect.ValueOf(token.TokMap)
	v := t.FieldByName("typeMap").Index(int(lit.Type))
	typeString := fmt.Sprintf("%v", v)

	switch n := num.(type) {
	case int64:
		if n > bounds.IntBoundaries[typeString].MaxVal || n < bounds.IntBoundaries[typeString].MinVal {
			return fmt.Errorf("Arithmetic overflow detected! %v overflows %v (Line: %v, Column: %v)", n, typeString, lit.Pos.Line, lit.Pos.Column)
		}
	}
	return nil
}
