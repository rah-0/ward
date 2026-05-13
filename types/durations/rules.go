// Package durations provides time.Duration validation rules for ward.
package durations

import "time"

const (
	IDGreaterThan        uint32 = 2
	IDGreaterThanOrEqual uint32 = 3
	IDLessThan           uint32 = 4
	IDLessThanOrEqual    uint32 = 5
	IDInRange            uint32 = 6
	IDPositive           uint32 = 7
	IDNonNegative        uint32 = 8
	IDOneOf              uint32 = 9
	IDNotOneOf           uint32 = 10
)

// IDs lists all rule IDs in this package.
var IDs = []uint32{
	IDGreaterThan, IDGreaterThanOrEqual,
	IDLessThan, IDLessThanOrEqual,
	IDInRange,
	IDPositive, IDNonNegative,
	IDOneOf, IDNotOneOf,
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

// RuleLessThan passes when v < max.
func RuleLessThan(max time.Duration) Rule {
	return Rule{ID: IDLessThan, Fn: func(v *time.Duration) *Result {
		if *v < max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleLessThanOrEqual passes when v <= max.
func RuleLessThanOrEqual(max time.Duration) Rule {
	return Rule{ID: IDLessThanOrEqual, Fn: func(v *time.Duration) *Result {
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

// RuleNonNegative passes when v >= 0.
func RuleNonNegative() Rule {
	return Rule{ID: IDNonNegative, Fn: func(v *time.Duration) *Result {
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
