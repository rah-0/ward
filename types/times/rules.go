// Package times provides time.Time validation rules for ward.
package times

import "time"

const (
	IDAfter         uint32 = 2
	IDAfterOrEqual  uint32 = 3
	IDBefore        uint32 = 4
	IDBeforeOrEqual uint32 = 5
	IDInRange       uint32 = 6
	IDIsZero        uint32 = 7
	IDIsNotZero     uint32 = 8
	IDOneOf         uint32 = 9
	IDNotOneOf      uint32 = 10
	IDIsPast        uint32 = 11
	IDIsFuture      uint32 = 12
	IDIsWeekday     uint32 = 13
	IDIsWeekend     uint32 = 14
)

// IDs lists all rule IDs in this package.
var IDs = []uint32{
	IDAfter, IDAfterOrEqual,
	IDBefore, IDBeforeOrEqual,
	IDInRange,
	IDIsZero, IDIsNotZero,
	IDOneOf, IDNotOneOf,
	IDIsPast, IDIsFuture,
	IDIsWeekday, IDIsWeekend,
}

// RuleAfter passes when v is strictly after threshold.
func RuleAfter(threshold time.Time) Rule {
	return Rule{ID: IDAfter, Fn: func(v *time.Time) *Result {
		if v.After(threshold) {
			return nil
		}
		return &Result{Arg1: threshold}
	}}
}

// RuleAfterOrEqual passes when v is after or equal to threshold.
func RuleAfterOrEqual(threshold time.Time) Rule {
	return Rule{ID: IDAfterOrEqual, Fn: func(v *time.Time) *Result {
		if !v.Before(threshold) {
			return nil
		}
		return &Result{Arg1: threshold}
	}}
}

// RuleBefore passes when v is strictly before threshold.
func RuleBefore(threshold time.Time) Rule {
	return Rule{ID: IDBefore, Fn: func(v *time.Time) *Result {
		if v.Before(threshold) {
			return nil
		}
		return &Result{Arg1: threshold}
	}}
}

// RuleBeforeOrEqual passes when v is before or equal to threshold.
func RuleBeforeOrEqual(threshold time.Time) Rule {
	return Rule{ID: IDBeforeOrEqual, Fn: func(v *time.Time) *Result {
		if !v.After(threshold) {
			return nil
		}
		return &Result{Arg1: threshold}
	}}
}

// RuleInRange passes when start <= v <= end (inclusive on both ends).
func RuleInRange(start, end time.Time) Rule {
	return Rule{ID: IDInRange, Fn: func(v *time.Time) *Result {
		if !v.Before(start) && !v.After(end) {
			return nil
		}
		return &Result{Arg1: start, Arg2: end}
	}}
}

// RuleIsZero passes when v is the zero time.
func RuleIsZero() Rule {
	return Rule{ID: IDIsZero, Fn: func(v *time.Time) *Result {
		if v.IsZero() {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsNotZero passes when v is not the zero time.
func RuleIsNotZero() Rule {
	return Rule{ID: IDIsNotZero, Fn: func(v *time.Time) *Result {
		if !v.IsZero() {
			return nil
		}
		return &Result{}
	}}
}

// RuleOneOf passes when *v equals one of the allowed time values.
func RuleOneOf(allowed ...time.Time) Rule {
	return Rule{ID: IDOneOf, Fn: func(v *time.Time) *Result {
		for _, a := range allowed {
			if v.Equal(a) {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOf passes when *v does not equal any of the excluded time values.
func RuleNotOneOf(excluded ...time.Time) Rule {
	return Rule{ID: IDNotOneOf, Fn: func(v *time.Time) *Result {
		for _, e := range excluded {
			if v.Equal(e) {
				return &Result{Arg1: excluded}
			}
		}
		return nil
	}}
}

// RuleIsPast passes when v is strictly before time.Now() at validation time.
// Comparison is done with time.Now() so a fresh "now" is observed on each call;
// the same field validated repeatedly may cross the boundary as wall time advances.
func RuleIsPast() Rule {
	return Rule{ID: IDIsPast, Fn: func(v *time.Time) *Result {
		if v.Before(time.Now()) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsFuture passes when v is strictly after time.Now() at validation time.
func RuleIsFuture() Rule {
	return Rule{ID: IDIsFuture, Fn: func(v *time.Time) *Result {
		if v.After(time.Now()) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsWeekday passes when v falls on Monday through Friday in its own location.
func RuleIsWeekday() Rule {
	return Rule{ID: IDIsWeekday, Fn: func(v *time.Time) *Result {
		d := v.Weekday()
		if d != time.Saturday && d != time.Sunday {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsWeekend passes when v falls on Saturday or Sunday in its own location.
func RuleIsWeekend() Rule {
	return Rule{ID: IDIsWeekend, Fn: func(v *time.Time) *Result {
		d := v.Weekday()
		if d == time.Saturday || d == time.Sunday {
			return nil
		}
		return &Result{}
	}}
}
