// Package bools provides bool validation rules for ward.
package bools

import "github.com/rah-0/ward"

const (
	IDIsTrue  uint32 = 2
	IDIsFalse uint32 = 3
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDIsTrue:  "IsTrue",
	IDIsFalse: "IsFalse",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
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
