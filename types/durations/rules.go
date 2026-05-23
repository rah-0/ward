// Package durations provides time.Duration validation rules for ward.
package durations

import (
	"time"

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
	IDOneOf              uint32 = 9
	IDNotOneOf           uint32 = 10
	IDNegative           uint32 = 11
	IDNegativeOrZero     uint32 = 12
	IDMultipleOf         uint32 = 13
	IDClamp              uint32 = 14
	IDClampMin           uint32 = 15
	IDClampMax           uint32 = 16
	IDAbs                uint32 = 17
	IDRound              uint32 = 18
	IDTruncate           uint32 = 19
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
	IDOneOf:              "OneOf",
	IDNotOneOf:           "NotOneOf",
	IDNegative:           "Negative",
	IDNegativeOrZero:     "NegativeOrZero",
	IDMultipleOf:         "MultipleOf",
	IDClamp:              "Clamp",
	IDClampMin:           "ClampMin",
	IDClampMax:           "ClampMax",
	IDAbs:                "Abs",
	IDRound:              "Round",
	IDTruncate:           "Truncate",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

// RuleGreaterThan passes when v > min.
func RuleGreaterThan(min time.Duration) Rule {
	return Rule{ID: IDGreaterThan, Fn: func(v *time.Duration) *Result {
		if *v > min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleGreaterThanOrEqual passes when v >= min.
func RuleGreaterThanOrEqual(min time.Duration) Rule {
	return Rule{ID: IDGreaterThanOrEqual, Fn: func(v *time.Duration) *Result {
		if *v >= min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleLesserThan passes when v < max.
func RuleLesserThan(max time.Duration) Rule {
	return Rule{ID: IDLesserThan, Fn: func(v *time.Duration) *Result {
		if *v < max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleLesserThanOrEqual passes when v <= max.
func RuleLesserThanOrEqual(max time.Duration) Rule {
	return Rule{ID: IDLesserThanOrEqual, Fn: func(v *time.Duration) *Result {
		if *v <= max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleInRange passes when min <= v <= max (inclusive on both ends).
func RuleInRange(min, max time.Duration) Rule {
	return Rule{ID: IDInRange, Fn: func(v *time.Duration) *Result {
		if *v >= min && *v <= max {
			return nil
		}
		return &Result{Arg1: min, Arg2: max}
	}}
}

// RulePositive passes when v > 0.
func RulePositive() Rule {
	return Rule{ID: IDPositive, Fn: func(v *time.Duration) *Result {
		if *v > 0 {
			return nil
		}
		return &Result{}
	}}
}

// RulePositiveOrZero passes when v >= 0.
func RulePositiveOrZero() Rule {
	return Rule{ID: IDPositiveOrZero, Fn: func(v *time.Duration) *Result {
		if *v >= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleOneOf passes when *v equals one of the allowed duration values.
func RuleOneOf(allowed ...time.Duration) Rule {
	return Rule{ID: IDOneOf, Fn: func(v *time.Duration) *Result {
		for _, a := range allowed {
			if *v == a {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOf passes when *v does not equal any of the excluded duration values.
func RuleNotOneOf(excluded ...time.Duration) Rule {
	return Rule{ID: IDNotOneOf, Fn: func(v *time.Duration) *Result {
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
	return Rule{ID: IDNegative, Fn: func(v *time.Duration) *Result {
		if *v < 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleNegativeOrZero passes when v <= 0.
func RuleNegativeOrZero() Rule {
	return Rule{ID: IDNegativeOrZero, Fn: func(v *time.Duration) *Result {
		if *v <= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleMultipleOf passes when v is an integer multiple of n.
// n == 0 is a no-op (always passes), matching the no-op convention used
// by RuleClamp for nonsensical arguments.
func RuleMultipleOf(n time.Duration) Rule {
	return Rule{ID: IDMultipleOf, Fn: func(v *time.Duration) *Result {
		if n == 0 {
			return nil
		}
		if *v%n == 0 {
			return nil
		}
		return &Result{Arg1: n}
	}}
}

// -----------------------------------------------------------------------------
// Sanitizers — the following rules mutate *v
// -----------------------------------------------------------------------------

// RuleClamp is a sanitizer that clamps *v into the inclusive range [min, max].
// If min > max, the rule is a no-op to avoid producing nonsensical results.
func RuleClamp(min, max time.Duration) Rule {
	return Rule{ID: IDClamp, Fn: func(v *time.Duration) *Result {
		if min > max {
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

// RuleClampMin is a sanitizer that raises *v to min if it is below.
func RuleClampMin(min time.Duration) Rule {
	return Rule{ID: IDClampMin, Fn: func(v *time.Duration) *Result {
		if *v < min {
			*v = min
		}
		return nil
	}}
}

// RuleClampMax is a sanitizer that lowers *v to max if it is above.
func RuleClampMax(max time.Duration) Rule {
	return Rule{ID: IDClampMax, Fn: func(v *time.Duration) *Result {
		if *v > max {
			*v = max
		}
		return nil
	}}
}

// RuleAbs is a sanitizer that replaces *v with its absolute value.
// math.MinInt64 nanoseconds has no positive counterpart and is left unchanged.
func RuleAbs() Rule {
	return Rule{ID: IDAbs, Fn: func(v *time.Duration) *Result {
		if *v < 0 && *v != time.Duration(-1<<63) {
			*v = -*v
		}
		return nil
	}}
}

// RuleRound is a sanitizer that rounds *v to the nearest multiple of m using
// time.Duration.Round (half-away-from-zero). m <= 0 is a no-op.
func RuleRound(m time.Duration) Rule {
	return Rule{ID: IDRound, Fn: func(v *time.Duration) *Result {
		if m <= 0 {
			return nil
		}
		*v = v.Round(m)
		return nil
	}}
}

// RuleTruncate is a sanitizer that rounds *v toward zero to a multiple of m
// using time.Duration.Truncate. m <= 0 is a no-op.
func RuleTruncate(m time.Duration) Rule {
	return Rule{ID: IDTruncate, Fn: func(v *time.Duration) *Result {
		if m <= 0 {
			return nil
		}
		*v = v.Truncate(m)
		return nil
	}}
}
