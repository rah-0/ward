// Package floats provides float64 validation and sanitization rules for ward.
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
	IDPositiveOrZero     uint32 = 8
	IDIsFinite           uint32 = 9
	IDMaxDecimalPlaces   uint32 = 10
	IDOneOf              uint32 = 11
	IDNotOneOf           uint32 = 12
	IDNegative           uint32 = 13
	IDNegativeOrZero     uint32 = 14
	IDIsInteger          uint32 = 15
	IDIsNaN              uint32 = 16
	IDIsInf              uint32 = 17
	IDRound              uint32 = 18
	IDFloor              uint32 = 19
	IDCeil               uint32 = 20
	IDClamp              uint32 = 21
	IDAbs                uint32 = 22
)

// IDs lists all rule IDs in this package.
var IDs = []uint32{
	IDGreaterThan, IDGreaterThanOrEqual,
	IDLessThan, IDLessThanOrEqual,
	IDInRange,
	IDPositive, IDPositiveOrZero,
	IDIsFinite,
	IDMaxDecimalPlaces,
	IDOneOf, IDNotOneOf,
	IDNegative, IDNegativeOrZero,
	IDIsInteger,
	IDIsNaN, IDIsInf,
	IDRound, IDFloor, IDCeil,
	IDClamp, IDAbs,
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

// RulePositiveOrZero passes when v >= 0.
func RulePositiveOrZero() Rule {
	return Rule{ID: IDPositiveOrZero, Fn: func(v *float64) *Result {
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

// RuleNegative passes when v < 0.
func RuleNegative() Rule {
	return Rule{ID: IDNegative, Fn: func(v *float64) *Result {
		if *v < 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleNegativeOrZero passes when v <= 0.
func RuleNegativeOrZero() Rule {
	return Rule{ID: IDNegativeOrZero, Fn: func(v *float64) *Result {
		if *v <= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsInteger passes when v has no fractional part (e.g. 1.0, -3.0).
// NaN and ±Inf fail.
func RuleIsInteger() Rule {
	return Rule{ID: IDIsInteger, Fn: func(v *float64) *Result {
		if !math.IsNaN(*v) && !math.IsInf(*v, 0) && *v == math.Trunc(*v) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsNaN passes when v is NaN.
func RuleIsNaN() Rule {
	return Rule{ID: IDIsNaN, Fn: func(v *float64) *Result {
		if math.IsNaN(*v) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsInf passes when v is +Inf or -Inf.
func RuleIsInf() Rule {
	return Rule{ID: IDIsInf, Fn: func(v *float64) *Result {
		if math.IsInf(*v, 0) {
			return nil
		}
		return &Result{}
	}}
}

// -----------------------------------------------------------------------------
// Sanitizers — the following rules mutate *v
// -----------------------------------------------------------------------------

// RuleRound is a sanitizer that rounds *v to n decimal places using
// half-away-from-zero rounding (math.Round). Negative n is treated as 0.
// NaN and ±Inf are left unchanged.
func RuleRound(n int) Rule {
	return Rule{ID: IDRound, Fn: func(v *float64) *Result {
		if math.IsNaN(*v) || math.IsInf(*v, 0) {
			return nil
		}
		if n < 0 {
			n = 0
		}
		shift := math.Pow(10, float64(n))
		*v = math.Round(*v*shift) / shift
		return nil
	}}
}

// RuleFloor is a sanitizer that replaces *v with math.Floor(*v).
func RuleFloor() Rule {
	return Rule{ID: IDFloor, Fn: func(v *float64) *Result {
		*v = math.Floor(*v)
		return nil
	}}
}

// RuleCeil is a sanitizer that replaces *v with math.Ceil(*v).
func RuleCeil() Rule {
	return Rule{ID: IDCeil, Fn: func(v *float64) *Result {
		*v = math.Ceil(*v)
		return nil
	}}
}

// RuleClamp is a sanitizer that clamps *v into the inclusive range [min, max].
// If min > max, the rule is a no-op. NaN is left unchanged.
func RuleClamp(min, max float64) Rule {
	return Rule{ID: IDClamp, Fn: func(v *float64) *Result {
		if min > max || math.IsNaN(*v) {
			return nil
		}
		if *v < min {
			*v = min
		} else if *v > max {
			*v = max
		}
		return nil
	}}
}

// RuleAbs is a sanitizer that replaces *v with its absolute value.
func RuleAbs() Rule {
	return Rule{ID: IDAbs, Fn: func(v *float64) *Result {
		*v = math.Abs(*v)
		return nil
	}}
}
