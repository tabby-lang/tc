// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

// Package bounds defines type boundaries (such as overflow values).
package bounds

type IntBoundary struct {
	MaxVal int64
	MinVal int64
}

type FloatBoundary struct {
	MaxVal float64
	MinVal float64
}

var IntBoundaries = map[string]IntBoundary{
	"int": {
		MaxVal: 0x7fffffff,
		MinVal: -0x80000000,
	},
}
