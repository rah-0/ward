package bools_test

import (
	"testing"

	"github.com/rah-0/ward/types/bools"
)

func run(rule bools.Rule, value bool) bool {
	return rule.Fn(&value) == nil
}

func TestIsTrue(t *testing.T) {
	if !run(bools.RuleIsTrue(), true) {
		t.Error("true should pass IsTrue")
	}
	if run(bools.RuleIsTrue(), false) {
		t.Error("false should fail IsTrue")
	}
}

func TestIsFalse(t *testing.T) {
	if !run(bools.RuleIsFalse(), false) {
		t.Error("false should pass IsFalse")
	}
	if run(bools.RuleIsFalse(), true) {
		t.Error("true should fail IsFalse")
	}
}
