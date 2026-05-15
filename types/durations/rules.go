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
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

// RuleGreaterThan passes when v > min.
func RuleGreaterThan(min time.Duration) Rule {
	return Rule{TypeID: TypeID, ID: IDGreaterThan, Fn: func(v *time.Duration) *Result {
		if *v > min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleGreaterThanOrEqual passes when v >= min.
func RuleGreaterThanOrEqual(min time.Duration) Rule {
	return Rule{TypeID: TypeID, ID: IDGreaterThanOrEqual, Fn: func(v *time.Duration) *Result {
		if *v >= min {
			return nil
		}
		return &Result{Arg1: min}
	}}
}

// RuleLesserThan passes when v < max.
func RuleLesserThan(max time.Duration) Rule {
	return Rule{TypeID: TypeID, ID: IDLesserThan, Fn: func(v *time.Duration) *Result {
		if *v < max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleLesserThanOrEqual passes when v <= max.
func RuleLesserThanOrEqual(max time.Duration) Rule {
	return Rule{TypeID: TypeID, ID: IDLesserThanOrEqual, Fn: func(v *time.Duration) *Result {
		if *v <= max {
			return nil
		}
		return &Result{Arg1: max}
	}}
}

// RuleInRange passes when min <= v <= max (inclusive on both ends).
func RuleInRange(min, max time.Duration) Rule {
	return Rule{TypeID: TypeID, ID: IDInRange, Fn: func(v *time.Duration) *Result {
		if *v >= min && *v <= max {
			return nil
		}
		return &Result{Arg1: min, Arg2: max}
	}}
}

// RulePositive passes when v > 0.
func RulePositive() Rule {
	return Rule{TypeID: TypeID, ID: IDPositive, Fn: func(v *time.Duration) *Result {
		if *v > 0 {
			return nil
		}
		return &Result{}
	}}
}

// RulePositiveOrZero passes when v >= 0.
func RulePositiveOrZero() Rule {
	return Rule{TypeID: TypeID, ID: IDPositiveOrZero, Fn: func(v *time.Duration) *Result {
		if *v >= 0 {
			return nil
		}
		return &Result{}
	}}
}

// RuleOneOf passes when *v equals one of the allowed duration values.
func RuleOneOf(allowed ...time.Duration) Rule {
	return Rule{TypeID: TypeID, ID: IDOneOf, Fn: func(v *time.Duration) *Result {
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
	return Rule{TypeID: TypeID, ID: IDNotOneOf, Fn: func(v *time.Duration) *Result {
		for _, e := range excluded {
			if *v == e {
				return &Result{Arg1: excluded}
			}
		}
		return nil
	}}
}
