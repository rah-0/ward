package floats_test

import (
	"math"
	"testing"

	"github.com/rah-0/ward/types/floats"
)

func run(rule floats.Rule, value float64) bool {
	return rule.Fn(&value) == nil
}

func TestGreaterThan(t *testing.T) {
	if !run(floats.RuleGreaterThan(1.0), 1.1) {
		t.Error("1.1 > 1.0 should pass")
	}
	if run(floats.RuleGreaterThan(1.0), 1.0) {
		t.Error("1.0 > 1.0 should fail")
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	if !run(floats.RuleGreaterThanOrEqual(1.0), 1.0) {
		t.Error("1.0 >= 1.0 should pass")
	}
	if run(floats.RuleGreaterThanOrEqual(1.0), 0.9) {
		t.Error("0.9 >= 1.0 should fail")
	}
}

func TestLessThan(t *testing.T) {
	if !run(floats.RuleLessThan(1.0), 0.9) {
		t.Error("0.9 < 1.0 should pass")
	}
	if run(floats.RuleLessThan(1.0), 1.0) {
		t.Error("1.0 < 1.0 should fail")
	}
}

func TestLessThanOrEqual(t *testing.T) {
	if !run(floats.RuleLessThanOrEqual(1.0), 1.0) {
		t.Error("1.0 <= 1.0 should pass")
	}
	if run(floats.RuleLessThanOrEqual(1.0), 1.1) {
		t.Error("1.1 <= 1.0 should fail")
	}
}

func TestInRange(t *testing.T) {
	if !run(floats.RuleInRange(0.0, 1.0), 0.5) {
		t.Error("0.5 in [0.0,1.0] should pass")
	}
	if !run(floats.RuleInRange(0.0, 1.0), 0.0) {
		t.Error("0.0 in [0.0,1.0] (inclusive) should pass")
	}
	if !run(floats.RuleInRange(0.0, 1.0), 1.0) {
		t.Error("1.0 in [0.0,1.0] (inclusive) should pass")
	}
	if run(floats.RuleInRange(0.0, 1.0), -0.1) {
		t.Error("-0.1 in [0.0,1.0] should fail")
	}
	if run(floats.RuleInRange(0.0, 1.0), 1.1) {
		t.Error("1.1 in [0.0,1.0] should fail")
	}
}

func TestPositive(t *testing.T) {
	if !run(floats.RulePositive(), 0.001) {
		t.Error("0.001 should be positive")
	}
	if run(floats.RulePositive(), 0.0) {
		t.Error("0.0 should not be positive")
	}
	if run(floats.RulePositive(), -0.1) {
		t.Error("-0.1 should not be positive")
	}
}

func TestNonNegative(t *testing.T) {
	if !run(floats.RuleNonNegative(), 0.0) {
		t.Error("0.0 should be non-negative")
	}
	if !run(floats.RuleNonNegative(), 1.5) {
		t.Error("1.5 should be non-negative")
	}
	if run(floats.RuleNonNegative(), -0.1) {
		t.Error("-0.1 should fail non-negative")
	}
}

func TestIsFinite(t *testing.T) {
	if !run(floats.RuleIsFinite(), 3.14) {
		t.Error("3.14 should be finite")
	}
	if !run(floats.RuleIsFinite(), 0.0) {
		t.Error("0.0 should be finite")
	}
	if run(floats.RuleIsFinite(), math.NaN()) {
		t.Error("NaN should fail IsFinite")
	}
	if run(floats.RuleIsFinite(), math.Inf(1)) {
		t.Error("+Inf should fail IsFinite")
	}
	if run(floats.RuleIsFinite(), math.Inf(-1)) {
		t.Error("-Inf should fail IsFinite")
	}
}

func TestMaxDecimalPlaces(t *testing.T) {
	if !run(floats.RuleMaxDecimalPlaces(2), 1.5) {
		t.Error("1.5 has 1 decimal place, should pass max=2")
	}
	if !run(floats.RuleMaxDecimalPlaces(2), 1.25) {
		t.Error("1.25 has 2 decimal places, should pass max=2")
	}
	if run(floats.RuleMaxDecimalPlaces(2), 1.123) {
		t.Error("1.123 has 3 decimal places, should fail max=2")
	}
	if !run(floats.RuleMaxDecimalPlaces(0), 5.0) {
		t.Error("5.0 has 0 decimal places, should pass max=0")
	}
}

func TestOneOf(t *testing.T) {
	if !run(floats.RuleOneOf(1.0, 2.5, 3.14), 2.5) {
		t.Error("2.5 is in list, should pass")
	}
	if run(floats.RuleOneOf(1.0, 2.5, 3.14), 4.0) {
		t.Error("4.0 is not in list, should fail")
	}
	if run(floats.RuleOneOf(1.0, 2.5, 3.14), 0.0) {
		t.Error("0.0 is not in list, should fail")
	}
}

func TestNotOneOf(t *testing.T) {
	if !run(floats.RuleNotOneOf(1.0, 2.5), 3.14) {
		t.Error("3.14 is not excluded, should pass")
	}
	if run(floats.RuleNotOneOf(1.0, 2.5), 1.0) {
		t.Error("1.0 is excluded, should fail")
	}
}
