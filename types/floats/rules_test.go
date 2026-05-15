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

func TestLesserThan(t *testing.T) {
	if !run(floats.RuleLesserThan(1.0), 0.9) {
		t.Error("0.9 < 1.0 should pass")
	}
	if run(floats.RuleLesserThan(1.0), 1.0) {
		t.Error("1.0 < 1.0 should fail")
	}
}

func TestLesserThanOrEqual(t *testing.T) {
	if !run(floats.RuleLesserThanOrEqual(1.0), 1.0) {
		t.Error("1.0 <= 1.0 should pass")
	}
	if run(floats.RuleLesserThanOrEqual(1.0), 1.1) {
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

func TestPositiveOrZero(t *testing.T) {
	if !run(floats.RulePositiveOrZero(), 0.0) {
		t.Error("0.0 should pass PositiveOrZero")
	}
	if !run(floats.RulePositiveOrZero(), 1.5) {
		t.Error("1.5 should pass PositiveOrZero")
	}
	if run(floats.RulePositiveOrZero(), -0.1) {
		t.Error("-0.1 should fail PositiveOrZero")
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

func TestNegative(t *testing.T) {
	if !run(floats.RuleNegative(), -0.1) {
		t.Error("-0.1 should be negative")
	}
	if run(floats.RuleNegative(), 0.0) {
		t.Error("0.0 should not be negative")
	}
	if run(floats.RuleNegative(), 1.0) {
		t.Error("1.0 should not be negative")
	}
}

func TestNegativeOrZero(t *testing.T) {
	if !run(floats.RuleNegativeOrZero(), 0.0) {
		t.Error("0.0 should pass NegativeOrZero")
	}
	if !run(floats.RuleNegativeOrZero(), -0.5) {
		t.Error("-0.5 should pass NegativeOrZero")
	}
	if run(floats.RuleNegativeOrZero(), 0.001) {
		t.Error("0.001 should fail NegativeOrZero")
	}
}

func TestIsInteger(t *testing.T) {
	for _, v := range []float64{0.0, 1.0, -1.0, 100.0, -100.0} {
		if !run(floats.RuleIsInteger(), v) {
			t.Errorf("%v should be integer-valued", v)
		}
	}
	for _, v := range []float64{0.5, -0.5, 1.1, 3.14} {
		if run(floats.RuleIsInteger(), v) {
			t.Errorf("%v should not be integer-valued", v)
		}
	}
	if run(floats.RuleIsInteger(), math.NaN()) {
		t.Error("NaN should fail IsInteger")
	}
	if run(floats.RuleIsInteger(), math.Inf(1)) {
		t.Error("+Inf should fail IsInteger")
	}
}

func TestIsNaN(t *testing.T) {
	if !run(floats.RuleIsNaN(), math.NaN()) {
		t.Error("NaN should pass IsNaN")
	}
	if run(floats.RuleIsNaN(), 0.0) {
		t.Error("0.0 should fail IsNaN")
	}
	if run(floats.RuleIsNaN(), math.Inf(1)) {
		t.Error("+Inf should fail IsNaN")
	}
}

func TestIsInf(t *testing.T) {
	if !run(floats.RuleIsInf(), math.Inf(1)) {
		t.Error("+Inf should pass IsInf")
	}
	if !run(floats.RuleIsInf(), math.Inf(-1)) {
		t.Error("-Inf should pass IsInf")
	}
	if run(floats.RuleIsInf(), 0.0) {
		t.Error("0.0 should fail IsInf")
	}
	if run(floats.RuleIsInf(), math.NaN()) {
		t.Error("NaN should fail IsInf")
	}
}

func TestRound(t *testing.T) {
	v := 1.23456
	floats.RuleRound(2).Fn(&v)
	if v != 1.23 {
		t.Errorf("expected 1.23, got %v", v)
	}

	v = 1.5
	floats.RuleRound(0).Fn(&v)
	if v != 2.0 {
		t.Errorf("expected 2.0 (half-away-from-zero), got %v", v)
	}

	v = -1.5
	floats.RuleRound(0).Fn(&v)
	if v != -2.0 {
		t.Errorf("expected -2.0 (half-away-from-zero), got %v", v)
	}

	// negative n treated as 0
	v = 3.7
	floats.RuleRound(-1).Fn(&v)
	if v != 4.0 {
		t.Errorf("expected 4.0, got %v", v)
	}

	// NaN/Inf unchanged
	v = math.NaN()
	floats.RuleRound(2).Fn(&v)
	if !math.IsNaN(v) {
		t.Error("NaN should be left unchanged")
	}
	v = math.Inf(1)
	floats.RuleRound(2).Fn(&v)
	if !math.IsInf(v, 1) {
		t.Error("+Inf should be left unchanged")
	}
}

func TestFloor(t *testing.T) {
	v := 1.7
	floats.RuleFloor().Fn(&v)
	if v != 1.0 {
		t.Errorf("expected 1.0, got %v", v)
	}
	v = -1.2
	floats.RuleFloor().Fn(&v)
	if v != -2.0 {
		t.Errorf("expected -2.0, got %v", v)
	}
}

func TestCeil(t *testing.T) {
	v := 1.2
	floats.RuleCeil().Fn(&v)
	if v != 2.0 {
		t.Errorf("expected 2.0, got %v", v)
	}
	v = -1.7
	floats.RuleCeil().Fn(&v)
	if v != -1.0 {
		t.Errorf("expected -1.0, got %v", v)
	}
}

func TestClamp(t *testing.T) {
	v := 0.5
	floats.RuleClamp(0.0, 1.0).Fn(&v)
	if v != 0.5 {
		t.Errorf("in-range should be unchanged, got %v", v)
	}

	v = -0.5
	floats.RuleClamp(0.0, 1.0).Fn(&v)
	if v != 0.0 {
		t.Errorf("below min should clamp to 0.0, got %v", v)
	}

	v = 2.0
	floats.RuleClamp(0.0, 1.0).Fn(&v)
	if v != 1.0 {
		t.Errorf("above max should clamp to 1.0, got %v", v)
	}

	// invalid range no-op
	v = 5.0
	floats.RuleClamp(10.0, 1.0).Fn(&v)
	if v != 5.0 {
		t.Errorf("invalid range should be a no-op, got %v", v)
	}

	// NaN unchanged
	v = math.NaN()
	floats.RuleClamp(0.0, 1.0).Fn(&v)
	if !math.IsNaN(v) {
		t.Error("NaN should be left unchanged")
	}
}

func TestAbs(t *testing.T) {
	v := -3.5
	floats.RuleAbs().Fn(&v)
	if v != 3.5 {
		t.Errorf("expected 3.5, got %v", v)
	}
	v = 3.5
	floats.RuleAbs().Fn(&v)
	if v != 3.5 {
		t.Errorf("positive unchanged, got %v", v)
	}
	v = 0.0
	floats.RuleAbs().Fn(&v)
	if v != 0.0 {
		t.Errorf("zero unchanged, got %v", v)
	}
}
