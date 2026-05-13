// Package bools provides bool validation rules for ward.
package bools

const (
	IDIsTrue  uint32 = 2
	IDIsFalse uint32 = 3
)

// IDs lists all rule IDs in this package.
var IDs = []uint32{
	IDIsTrue, IDIsFalse,
}

// RuleIsTrue passes when v is true.
func RuleIsTrue() Rule {
	return Rule{ID: IDIsTrue, Fn: func(v *bool) *Result {
		if *v {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsFalse passes when v is false.
func RuleIsFalse() Rule {
	return Rule{ID: IDIsFalse, Fn: func(v *bool) *Result {
		if !*v {
			return nil
		}
		return &Result{}
	}}
}
