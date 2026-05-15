// Package floats provides float64 validation and sanitization rules for ward.
package floats

import (
	"math"
	"strconv"
	"strings"

	"github.com/rah-0/ward"
)

const (
	IDGreaterThan        uint32 = 2
	IDGreaterThanOrEqual uint32 = 3
	IDLesserThan         uint32 = 4
	IDLesserThanOrEqual  uint32 = 5
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

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDGreaterThan:        "GreaterThan",
	IDGreaterThanOrEqual: "GreaterThanOrEqual",
	IDLesserThan:         "LesserThan",
	IDLesserThanOrEqual:  "LesserThanOrEqual",
	IDInRange:            "InRange",
	IDPositive:           "Positive",
	IDPositiveOrZero:     "PositiveOrZero",
	IDIsFinite:           "IsFinite",
	IDMaxDecimalPlaces:   "MaxDecimalPlaces",
	IDOneOf:              "OneOf",
	IDNotOneOf:           "NotOneOf",
	IDNegative:           "Negative",
	IDNegativeOrZero:     "NegativeOrZero",
	IDIsInteger:          "IsInteger",
	IDIsNaN:              "IsNaN",
	IDIsInf:              "IsInf",
	IDRound:              "Round",
	IDFloor:              "Floor",
	IDCeil:               "Ceil",
	IDClamp:              "Clamp",
	IDAbs:                "Abs",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

// RuleGreaterThan passes when v > min.
func RuleGreaterThan(min float64) Rule {
	return Rule{TypeID: TypeID, ID: IDGreaterThan, Fn: func(v *float64) *Result {
		if *v > min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleGreaterThanOrEqual passes when v >= min.
func RuleGreaterThanOrEqual(min float64) Rule {
	return Rule{TypeID: TypeID, ID: IDGreaterThanOrEqual, Fn: func(v *float64) *Result {
		if *v >= min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleLesserThan passes when v < max.
func RuleLesserThan(max float64) Rule {
	return Rule{TypeID: TypeID, ID: IDLesserThan, Fn: func(v *float64) *Result {
		if *v < max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleLesserThanOrEqual passes when v <= max.
func RuleLesserThanOrEqual(max float64) Rule {
	return Rule{TypeID: TypeID, ID: IDLesserThanOrEqual, Fn: func(v *float64) *Result {
		if *v <= max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleInRange passes when min <= v <= max (inclusive on both ends).
func RuleInRange(min, max float64) Rule {
	return Rule{TypeID: TypeID, ID: IDInRange, Fn: func(v *float64) *Result {
		if *v >= min && *v <= max {
			return nil
		}
		return &Result{Arg1: min, Arg2: max}
	}}
}

// RulePositive passes when v > 0.
func RulePositive() Rule {
	return Rule{TypeID: TypeID, ID: IDPositive, Fn: func(v *float64) *Result {
		if *v > 0 {
			return nil
		}
		return &Result{}
	}}
}

// RulePositiveOrZero passes when v >= 0.
func RulePositiveOrZero() Rule {
	return Rule{TypeID: TypeID, ID: IDPositiveOrZero, Fn: func(v *float64) *Result {
		if *v >= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsFinite passes when v is neither NaN nor ±Inf.
func RuleIsFinite() Rule {
	return Rule{TypeID: TypeID, ID: IDIsFinite, Fn: func(v *float64) *Result {
		if !math.IsNaN(*v) && !math.IsInf(*v, 0) {
			return nil
		}
		return &Result{}
	}}
}

// RuleMaxDecimalPlaces passes when v has at most n digits after the decimal point.
// Uses the shortest decimal representation, so 1.5 has 1 decimal place, not 15.
func RuleMaxDecimalPlaces(n int) Rule {
	return Rule{TypeID: TypeID, ID: IDMaxDecimalPlaces, Fn: func(v *float64) *Result {
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
	return Rule{TypeID: TypeID, ID: IDOneOf, Fn: func(v *float64) *Result {
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
	return Rule{TypeID: TypeID, ID: IDNotOneOf, Fn: func(v *float64) *Result {
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
	return Rule{TypeID: TypeID, ID: IDNegative, Fn: func(v *float64) *Result {
		if *v < 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleNegativeOrZero passes when v <= 0.
func RuleNegativeOrZero() Rule {
	return Rule{TypeID: TypeID, ID: IDNegativeOrZero, Fn: func(v *float64) *Result {
		if *v <= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsInteger passes when v has no fractional part (e.g. 1.0, -3.0).
// NaN and ±Inf fail.
func RuleIsInteger() Rule {
	return Rule{TypeID: TypeID, ID: IDIsInteger, Fn: func(v *float64) *Result {
		if !math.IsNaN(*v) && !math.IsInf(*v, 0) && *v == math.Trunc(*v) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsNaN passes when v is NaN.
func RuleIsNaN() Rule {
	return Rule{TypeID: TypeID, ID: IDIsNaN, Fn: func(v *float64) *Result {
		if math.IsNaN(*v) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsInf passes when v is +Inf or -Inf.
func RuleIsInf() Rule {
	return Rule{TypeID: TypeID, ID: IDIsInf, Fn: func(v *float64) *Result {
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
	return Rule{TypeID: TypeID, ID: IDRound, Fn: func(v *float64) *Result {
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
	return Rule{TypeID: TypeID, ID: IDFloor, Fn: func(v *float64) *Result {
		*v = math.Floor(*v)
		return nil
	}}
}

// RuleCeil is a sanitizer that replaces *v with math.Ceil(*v).
func RuleCeil() Rule {
	return Rule{TypeID: TypeID, ID: IDCeil, Fn: func(v *float64) *Result {
		*v = math.Ceil(*v)
		return nil
	}}
}

// RuleClamp is a sanitizer that clamps *v into the inclusive range [min, max].
// If min > max, the rule is a no-op. NaN is left unchanged.
func RuleClamp(min, max float64) Rule {
	return Rule{TypeID: TypeID, ID: IDClamp, Fn: func(v *float64) *Result {
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
	return Rule{TypeID: TypeID, ID: IDAbs, Fn: func(v *float64) *Result {
		*v = math.Abs(*v)
		return nil
	}}
}
