// Package floats provides float64 validation rules for ward.
package floats

import (
	"math"
	"strconv"
	"strings"
)

const (
	IDGreaterThan        uint32 = 2
	IDGreaterThanOrEqual uint32 = 3
	IDLessThan           uint32 = 4
	IDLessThanOrEqual    uint32 = 5
	IDInRange            uint32 = 6
	IDPositive           uint32 = 7
	IDNonNegative        uint32 = 8
	IDIsFinite           uint32 = 9
	IDMaxDecimalPlaces   uint32 = 10
	IDOneOf              uint32 = 11
	IDNotOneOf           uint32 = 12
)

// IDs lists all rule IDs in this package.
var IDs = []uint32{
	IDGreaterThan, IDGreaterThanOrEqual,
	IDLessThan, IDLessThanOrEqual,
	IDInRange,
	IDPositive, IDNonNegative,
	IDIsFinite,
	IDMaxDecimalPlaces,
	IDOneOf, IDNotOneOf,
}

// RuleGreaterThan passes when v > min.
func RuleGreaterThan(min float64) Rule {
	return Rule{ID: IDGreaterThan, Fn: func(v *float64) *Result {
		if *v > min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleGreaterThanOrEqual passes when v >= min.
func RuleGreaterThanOrEqual(min float64) Rule {
	return Rule{ID: IDGreaterThanOrEqual, Fn: func(v *float64) *Result {
		if *v >= min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleLessThan passes when v < max.
func RuleLessThan(max float64) Rule {
	return Rule{ID: IDLessThan, Fn: func(v *float64) *Result {
		if *v < max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleLessThanOrEqual passes when v <= max.
func RuleLessThanOrEqual(max float64) Rule {
	return Rule{ID: IDLessThanOrEqual, Fn: func(v *float64) *Result {
		if *v <= max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleInRange passes when min <= v <= max (inclusive on both ends).
func RuleInRange(min, max float64) Rule {
	return Rule{ID: IDInRange, Fn: func(v *float64) *Result {
		if *v >= min && *v <= max {
			return nil
		}
		return &Result{Arg1: min, Arg2: max}
	}}
}

// RulePositive passes when v > 0.
func RulePositive() Rule {
	return Rule{ID: IDPositive, Fn: func(v *float64) *Result {
		if *v > 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleNonNegative passes when v >= 0.
func RuleNonNegative() Rule {
	return Rule{ID: IDNonNegative, Fn: func(v *float64) *Result {
		if *v >= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsFinite passes when v is neither NaN nor ±Inf.
func RuleIsFinite() Rule {
	return Rule{ID: IDIsFinite, Fn: func(v *float64) *Result {
		if !math.IsNaN(*v) && !math.IsInf(*v, 0) {
			return nil
		}
		return &Result{}
	}}
}

// RuleMaxDecimalPlaces passes when v has at most n digits after the decimal point.
// Uses the shortest decimal representation, so 1.5 has 1 decimal place, not 15.
func RuleMaxDecimalPlaces(n int) Rule {
	return Rule{ID: IDMaxDecimalPlaces, Fn: func(v *float64) *Result {
		s := strconv.FormatFloat(*v, 'f', -1, 64)
		if idx := strings.Index(s, "."); idx != -1 {
			if len(s)-idx-1 > n {
				return &Result{Arg1: n}
			}
		}
		return nil
	}}
}


// RuleOneOf passes when *v equals one of the allowed values.
func RuleOneOf(allowed ...float64) Rule {
	return Rule{ID: IDOneOf, Fn: func(v *float64) *Result {
		for _, a := range allowed {
			if *v == a {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOf passes when *v does not equal any of the excluded values.
func RuleNotOneOf(excluded ...float64) Rule {
	return Rule{ID: IDNotOneOf, Fn: func(v *float64) *Result {
		for _, e := range excluded {
			if *v == e {
				return &Result{Arg1: excluded}
			}
		}
		return nil
	}}
}
