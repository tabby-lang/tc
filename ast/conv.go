// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package ast

import "strconv"

// TypedValue returns the int64 value representation of an IntegerLiteral.
func (lit *IntegerLiteral) TypedValue() (int64, error) {
	return strconv.ParseInt(string(lit.Token.Lit), 0, 64)
}
