// Package ints provides int64 validation and sanitization rules for ward.
package ints

import "github.com/rah-0/ward"

const (
	IDGreaterThan        uint32 = 2
	IDGreaterThanOrEqual uint32 = 3
	IDLesserThan         uint32 = 4
	IDLesserThanOrEqual  uint32 = 5
	IDInRange            uint32 = 6
	IDPositive           uint32 = 7
	IDPositiveOrZero     uint32 = 8
	IDMultipleOf         uint32 = 9
	IDOneOf              uint32 = 10
	IDNotOneOf           uint32 = 11
	IDNegative           uint32 = 12
	IDNegativeOrZero     uint32 = 13
	IDIsEven             uint32 = 14
	IDIsOdd              uint32 = 15
	IDClamp              uint32 = 16
	IDClampMin           uint32 = 17
	IDClampMax           uint32 = 18
	IDAbs                uint32 = 19
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
	IDMultipleOf:         "MultipleOf",
	IDOneOf:              "OneOf",
	IDNotOneOf:           "NotOneOf",
	IDNegative:           "Negative",
	IDNegativeOrZero:     "NegativeOrZero",
	IDIsEven:             "IsEven",
	IDIsOdd:              "IsOdd",
	IDClamp:              "Clamp",
	IDClampMin:           "ClampMin",
	IDClampMax:           "ClampMax",
	IDAbs:                "Abs",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

// RuleGreaterThan passes when v > min.
func RuleGreaterThan(min int64) Rule {
	return Rule{ID: IDGreaterThan, Fn: func(v *int64) *Result {
		if *v > min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleGreaterThanOrEqual passes when v >= min.
func RuleGreaterThanOrEqual(min int64) Rule {
	return Rule{ID: IDGreaterThanOrEqual, Fn: func(v *int64) *Result {
		if *v >= min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleLesserThan passes when v < max.
func RuleLesserThan(max int64) Rule {
	return Rule{ID: IDLesserThan, Fn: func(v *int64) *Result {
		if *v < max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleLesserThanOrEqual passes when v <= max.
func RuleLesserThanOrEqual(max int64) Rule {
	return Rule{ID: IDLesserThanOrEqual, Fn: func(v *int64) *Result {
		if *v <= max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleInRange passes when min <= v <= max (inclusive on both ends).
func RuleInRange(min, max int64) Rule {
	return Rule{ID: IDInRange, Fn: func(v *int64) *Result {
		if *v >= min && *v <= max {
			return nil
		}
		return &Result{Arg1: min, Arg2: max}
	}}
}

// RulePositive passes when v > 0.
func RulePositive() Rule {
	return Rule{ID: IDPositive, Fn: func(v *int64) *Result {
		if *v > 0 {
			return nil
		}
		return &Result{}
	}}
}

// RulePositiveOrZero passes when v >= 0.
func RulePositiveOrZero() Rule {
	return Rule{ID: IDPositiveOrZero, Fn: func(v *int64) *Result {
		if *v >= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleMultipleOf passes when v is evenly divisible by n.
func RuleMultipleOf(n int64) Rule {
	return Rule{ID: IDMultipleOf, Fn: func(v *int64) *Result {
		if n != 0 && *v%n == 0 {
			return nil
		}
		return &Result{Arg1: n}
	}}
}

// RuleOneOf passes when *v equals one of the allowed values.
func RuleOneOf(allowed ...int64) Rule {
	return Rule{ID: IDOneOf, Fn: func(v *int64) *Result {
		for _, a := range allowed {
			if *v == a {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOf passes when *v does not equal any of the excluded values.
func RuleNotOneOf(excluded ...int64) Rule {
	return Rule{ID: IDNotOneOf, Fn: func(v *int64) *Result {
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
	return Rule{ID: IDNegative, Fn: func(v *int64) *Result {
		if *v < 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleNegativeOrZero passes when v <= 0.
func RuleNegativeOrZero() Rule {
	return Rule{ID: IDNegativeOrZero, Fn: func(v *int64) *Result {
		if *v <= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsEven passes when v is even (divisible by 2, including 0 and negatives).
func RuleIsEven() Rule {
	return Rule{ID: IDIsEven, Fn: func(v *int64) *Result {
		if *v%2 == 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsOdd passes when v is odd.
func RuleIsOdd() Rule {
	return Rule{ID: IDIsOdd, Fn: func(v *int64) *Result {
		if *v%2 != 0 {
			return nil
		}
		return &Result{}
	}}
}

// -----------------------------------------------------------------------------
// Sanitizers — the following rules mutate *v
// -----------------------------------------------------------------------------

// RuleClamp is a sanitizer that clamps *v into the inclusive range [min, max].
// If min > max, the rule is a no-op to avoid producing nonsensical results.
func RuleClamp(min, max int64) Rule {
	return Rule{ID: IDClamp, Fn: func(v *int64) *Result {
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
func RuleClampMin(min int64) Rule {
	return Rule{ID: IDClampMin, Fn: func(v *int64) *Result {
		if *v < min {
			*v = min
		}
		return nil
	}}
}

// RuleClampMax is a sanitizer that lowers *v to max if it is above.
func RuleClampMax(max int64) Rule {
	return Rule{ID: IDClampMax, Fn: func(v *int64) *Result {
		if *v > max {
			*v = max
		}
		return nil
	}}
}

// RuleAbs is a sanitizer that replaces *v with its absolute value.
// math.MinInt64 has no positive counterpart and is left unchanged.
func RuleAbs() Rule {
	return Rule{ID: IDAbs, Fn: func(v *int64) *Result {
		if *v < 0 && *v != -1<<63 {
			*v = -*v
		}
		return nil
	}}
}
