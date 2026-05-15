package ints_test

import (
	"testing"

	"github.com/rah-0/ward/types/ints"
)

func run(rule ints.Rule, value int64) bool {
	return rule.Fn(&value) == nil
}

func TestGreaterThan(t *testing.T) {
	if !run(ints.RuleGreaterThan(5), 6) {
		t.Error("6 > 5 should pass")
	}
	if run(ints.RuleGreaterThan(5), 5) {
		t.Error("5 > 5 should fail")
	}
	if run(ints.RuleGreaterThan(5), 4) {
		t.Error("4 > 5 should fail")
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	if !run(ints.RuleGreaterThanOrEqual(5), 5) {
		t.Error("5 >= 5 should pass")
	}
	if !run(ints.RuleGreaterThanOrEqual(5), 6) {
		t.Error("6 >= 5 should pass")
	}
	if run(ints.RuleGreaterThanOrEqual(5), 4) {
		t.Error("4 >= 5 should fail")
	}
}

func TestLessThan(t *testing.T) {
	if !run(ints.RuleLessThan(5), 4) {
		t.Error("4 < 5 should pass")
	}
	if run(ints.RuleLessThan(5), 5) {
		t.Error("5 < 5 should fail")
	}
	if run(ints.RuleLessThan(5), 6) {
		t.Error("6 < 5 should fail")
	}
}

func TestLessThanOrEqual(t *testing.T) {
	if !run(ints.RuleLessThanOrEqual(5), 5) {
		t.Error("5 <= 5 should pass")
	}
	if !run(ints.RuleLessThanOrEqual(5), 4) {
		t.Error("4 <= 5 should pass")
	}
	if run(ints.RuleLessThanOrEqual(5), 6) {
		t.Error("6 <= 5 should fail")
	}
}

func TestInRange(t *testing.T) {
	if !run(ints.RuleInRange(1, 10), 5) {
		t.Error("5 in [1,10] should pass")
	}
	if !run(ints.RuleInRange(1, 10), 1) {
		t.Error("1 in [1,10] (inclusive) should pass")
	}
	if !run(ints.RuleInRange(1, 10), 10) {
		t.Error("10 in [1,10] (inclusive) should pass")
	}
	if run(ints.RuleInRange(1, 10), 0) {
		t.Error("0 in [1,10] should fail")
	}
	if run(ints.RuleInRange(1, 10), 11) {
		t.Error("11 in [1,10] should fail")
	}
}

func TestPositive(t *testing.T) {
	if !run(ints.RulePositive(), 1) {
		t.Error("1 should be positive")
	}
	if run(ints.RulePositive(), 0) {
		t.Error("0 should not be positive")
	}
	if run(ints.RulePositive(), -1) {
		t.Error("-1 should not be positive")
	}
}

func TestPositiveOrZero(t *testing.T) {
	if !run(ints.RulePositiveOrZero(), 0) {
		t.Error("0 should pass PositiveOrZero")
	}
	if !run(ints.RulePositiveOrZero(), 1) {
		t.Error("1 should pass PositiveOrZero")
	}
	if run(ints.RulePositiveOrZero(), -1) {
		t.Error("-1 should fail PositiveOrZero")
	}
}

func TestMultipleOf(t *testing.T) {
	if !run(ints.RuleMultipleOf(3), 9) {
		t.Error("9 is multiple of 3")
	}
	if !run(ints.RuleMultipleOf(3), 0) {
		t.Error("0 is multiple of 3")
	}
	if run(ints.RuleMultipleOf(3), 7) {
		t.Error("7 is not multiple of 3")
	}
	if run(ints.RuleMultipleOf(0), 5) {
		t.Error("n=0 should always fail")
	}
}

func TestOneOf(t *testing.T) {
	if !run(ints.RuleOneOf(1, 2, 3), 2) {
		t.Error("2 is in [1,2,3], should pass")
	}
	if run(ints.RuleOneOf(1, 2, 3), 4) {
		t.Error("4 is not in [1,2,3], should fail")
	}
	if run(ints.RuleOneOf(1, 2, 3), 0) {
		t.Error("0 is not in [1,2,3], should fail")
	}
}

func TestNotOneOf(t *testing.T) {
	if !run(ints.RuleNotOneOf(1, 2, 3), 5) {
		t.Error("5 is not excluded, should pass")
	}
	if run(ints.RuleNotOneOf(1, 2, 3), 2) {
		t.Error("2 is excluded, should fail")
	}
}

func TestNegative(t *testing.T) {
	if !run(ints.RuleNegative(), -1) {
		t.Error("-1 should be negative")
	}
	if run(ints.RuleNegative(), 0) {
		t.Error("0 should not be negative")
	}
	if run(ints.RuleNegative(), 1) {
		t.Error("1 should not be negative")
	}
}

func TestNegativeOrZero(t *testing.T) {
	if !run(ints.RuleNegativeOrZero(), 0) {
		t.Error("0 should pass NegativeOrZero")
	}
	if !run(ints.RuleNegativeOrZero(), -1) {
		t.Error("-1 should pass NegativeOrZero")
	}
	if run(ints.RuleNegativeOrZero(), 1) {
		t.Error("1 should fail NegativeOrZero")
	}
}

func TestIsEven(t *testing.T) {
	for _, v := range []int64{-4, -2, 0, 2, 4} {
		if !run(ints.RuleIsEven(), v) {
			t.Errorf("%d should be even", v)
		}
	}
	for _, v := range []int64{-3, -1, 1, 3} {
		if run(ints.RuleIsEven(), v) {
			t.Errorf("%d should not be even", v)
		}
	}
}

func TestIsOdd(t *testing.T) {
	for _, v := range []int64{-3, -1, 1, 3} {
		if !run(ints.RuleIsOdd(), v) {
			t.Errorf("%d should be odd", v)
		}
	}
	for _, v := range []int64{-4, -2, 0, 2} {
		if run(ints.RuleIsOdd(), v) {
			t.Errorf("%d should not be odd", v)
		}
	}
}

func TestClamp(t *testing.T) {
	v := int64(5)
	ints.RuleClamp(1, 10).Fn(&v)
	if v != 5 {
		t.Errorf("in-range value should be unchanged, got %d", v)
	}

	v = -3
	ints.RuleClamp(1, 10).Fn(&v)
	if v != 1 {
		t.Errorf("below min should clamp to 1, got %d", v)
	}

	v = 99
	ints.RuleClamp(1, 10).Fn(&v)
	if v != 10 {
		t.Errorf("above max should clamp to 10, got %d", v)
	}

	// invalid range: min > max → no-op
	v = 5
	ints.RuleClamp(10, 1).Fn(&v)
	if v != 5 {
		t.Errorf("invalid range should be a no-op, got %d", v)
	}
}

func TestClampMin(t *testing.T) {
	v := int64(5)
	ints.RuleClampMin(0).Fn(&v)
	if v != 5 {
		t.Errorf("above min should be unchanged, got %d", v)
	}
	v = -10
	ints.RuleClampMin(0).Fn(&v)
	if v != 0 {
		t.Errorf("below min should be raised, got %d", v)
	}
}

func TestClampMax(t *testing.T) {
	v := int64(5)
	ints.RuleClampMax(10).Fn(&v)
	if v != 5 {
		t.Errorf("below max should be unchanged, got %d", v)
	}
	v = 100
	ints.RuleClampMax(10).Fn(&v)
	if v != 10 {
		t.Errorf("above max should be lowered, got %d", v)
	}
}

func TestAbs(t *testing.T) {
	v := int64(-5)
	ints.RuleAbs().Fn(&v)
	if v != 5 {
		t.Errorf("expected 5, got %d", v)
	}

	v = 5
	ints.RuleAbs().Fn(&v)
	if v != 5 {
		t.Errorf("positive unchanged, got %d", v)
	}

	v = 0
	ints.RuleAbs().Fn(&v)
	if v != 0 {
		t.Errorf("zero unchanged, got %d", v)
	}

	// math.MinInt64 — no positive counterpart, left as-is
	v = -1 << 63
	ints.RuleAbs().Fn(&v)
	if v != -1<<63 {
		t.Errorf("MinInt64 should be left as-is, got %d", v)
	}
}
