// Package bools provides bool validation rules for ward.
package bools

const (
	IDIsTrue  uint32 = 2
	IDIsFalse uint32 = 3
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDIsTrue:  "IsTrue",
	IDIsFalse: "IsFalse",
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
