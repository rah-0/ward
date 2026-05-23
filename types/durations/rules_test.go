package durations_test

import (
	"testing"
	"time"

	"github.com/rah-0/ward/types/durations"
)

func run(rule durations.Rule, value time.Duration) bool {
	return rule.Fn(&value) == nil
}

func TestGreaterThan(t *testing.T) {
	if !run(durations.RuleGreaterThan(time.Second), 2*time.Second) {
		t.Error("2s > 1s should pass")
	}
	if run(durations.RuleGreaterThan(time.Second), time.Second) {
		t.Error("1s > 1s should fail")
	}
	if run(durations.RuleGreaterThan(time.Second), 0) {
		t.Error("0 > 1s should fail")
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	if !run(durations.RuleGreaterThanOrEqual(time.Second), time.Second) {
		t.Error("1s >= 1s should pass")
	}
	if !run(durations.RuleGreaterThanOrEqual(time.Second), 2*time.Second) {
		t.Error("2s >= 1s should pass")
	}
	if run(durations.RuleGreaterThanOrEqual(time.Second), 0) {
		t.Error("0 >= 1s should fail")
	}
}

func TestLesserThan(t *testing.T) {
	if !run(durations.RuleLesserThan(time.Second), time.Millisecond) {
		t.Error("1ms < 1s should pass")
	}
	if run(durations.RuleLesserThan(time.Second), time.Second) {
		t.Error("1s < 1s should fail")
	}
	if run(durations.RuleLesserThan(time.Second), 2*time.Second) {
		t.Error("2s < 1s should fail")
	}
}

func TestLesserThanOrEqual(t *testing.T) {
	if !run(durations.RuleLesserThanOrEqual(time.Second), time.Second) {
		t.Error("1s <= 1s should pass")
	}
	if !run(durations.RuleLesserThanOrEqual(time.Second), time.Millisecond) {
		t.Error("1ms <= 1s should pass")
	}
	if run(durations.RuleLesserThanOrEqual(time.Second), 2*time.Second) {
		t.Error("2s <= 1s should fail")
	}
}

func TestInRange(t *testing.T) {
	if !run(durations.RuleInRange(time.Second, time.Minute), 30*time.Second) {
		t.Error("30s in [1s,1m] should pass")
	}
	if !run(durations.RuleInRange(time.Second, time.Minute), time.Second) {
		t.Error("1s == min should pass (inclusive)")
	}
	if !run(durations.RuleInRange(time.Second, time.Minute), time.Minute) {
		t.Error("1m == max should pass (inclusive)")
	}
	if run(durations.RuleInRange(time.Second, time.Minute), time.Millisecond) {
		t.Error("1ms below range should fail")
	}
	if run(durations.RuleInRange(time.Second, time.Minute), 2*time.Minute) {
		t.Error("2m above range should fail")
	}
}

func TestPositive(t *testing.T) {
	if !run(durations.RulePositive(), time.Nanosecond) {
		t.Error("1ns should be positive")
	}
	if run(durations.RulePositive(), 0) {
		t.Error("0 should not be positive")
	}
	if run(durations.RulePositive(), -time.Second) {
		t.Error("negative duration should not be positive")
	}
}

func TestPositiveOrZero(t *testing.T) {
	if !run(durations.RulePositiveOrZero(), 0) {
		t.Error("0 should pass PositiveOrZero")
	}
	if !run(durations.RulePositiveOrZero(), time.Second) {
		t.Error("1s should pass PositiveOrZero")
	}
	if run(durations.RulePositiveOrZero(), -time.Second) {
		t.Error("negative duration should fail PositiveOrZero")
	}
}

func TestOneOf(t *testing.T) {
	if !run(durations.RuleOneOf(time.Second, time.Minute), time.Second) {
		t.Error("1s is in list, should pass")
	}
	if !run(durations.RuleOneOf(time.Second, time.Minute), time.Minute) {
		t.Error("1m is in list, should pass")
	}
	if run(durations.RuleOneOf(time.Second, time.Minute), time.Hour) {
		t.Error("1h is not in list, should fail")
	}
}

func TestNotOneOf(t *testing.T) {
	if !run(durations.RuleNotOneOf(time.Second, time.Minute), time.Hour) {
		t.Error("1h is not excluded, should pass")
	}
	if run(durations.RuleNotOneOf(time.Second, time.Minute), time.Second) {
		t.Error("1s is excluded, should fail")
	}
}

func TestNegative(t *testing.T) {
	if !run(durations.RuleNegative(), -time.Second) {
		t.Error("-1s should pass Negative")
	}
	if run(durations.RuleNegative(), 0) {
		t.Error("0 should fail Negative")
	}
	if run(durations.RuleNegative(), time.Second) {
		t.Error("1s should fail Negative")
	}
}

func TestNegativeOrZero(t *testing.T) {
	if !run(durations.RuleNegativeOrZero(), -time.Second) {
		t.Error("-1s should pass NegativeOrZero")
	}
	if !run(durations.RuleNegativeOrZero(), 0) {
		t.Error("0 should pass NegativeOrZero")
	}
	if run(durations.RuleNegativeOrZero(), time.Second) {
		t.Error("1s should fail NegativeOrZero")
	}
}

func TestMultipleOf(t *testing.T) {
	if !run(durations.RuleMultipleOf(time.Minute), 5*time.Minute) {
		t.Error("5m should be a multiple of 1m")
	}
	if !run(durations.RuleMultipleOf(time.Minute), 0) {
		t.Error("0 should be a multiple of 1m")
	}
	if run(durations.RuleMultipleOf(time.Minute), 90*time.Second) {
		t.Error("90s should not be a multiple of 1m")
	}
	// negative multiples
	if !run(durations.RuleMultipleOf(time.Minute), -2*time.Minute) {
		t.Error("-2m should be a multiple of 1m")
	}
}

// MultipleOf(0) is a no-op — every value passes, matching the convention
// used by RuleClamp for nonsensical arguments.
func TestMultipleOf_ZeroIsNoOp(t *testing.T) {
	for _, v := range []time.Duration{0, time.Second, -time.Second, time.Hour} {
		if !run(durations.RuleMultipleOf(0), v) {
			t.Errorf("MultipleOf(0) should pass for %v", v)
		}
	}
}

func TestClamp(t *testing.T) {
	v := 30 * time.Second
	durations.RuleClamp(time.Second, time.Minute).Fn(&v)
	if v != 30*time.Second {
		t.Errorf("in-range should be unchanged, got %v", v)
	}

	v = time.Millisecond
	durations.RuleClamp(time.Second, time.Minute).Fn(&v)
	if v != time.Second {
		t.Errorf("below min should clamp to 1s, got %v", v)
	}

	v = time.Hour
	durations.RuleClamp(time.Second, time.Minute).Fn(&v)
	if v != time.Minute {
		t.Errorf("above max should clamp to 1m, got %v", v)
	}

	// invalid range no-op
	v = time.Hour
	durations.RuleClamp(time.Minute, time.Second).Fn(&v)
	if v != time.Hour {
		t.Errorf("invalid range should be a no-op, got %v", v)
	}
}

func TestClampMin(t *testing.T) {
	v := time.Minute
	durations.RuleClampMin(time.Second).Fn(&v)
	if v != time.Minute {
		t.Errorf("above min should be unchanged, got %v", v)
	}
	v = time.Millisecond
	durations.RuleClampMin(time.Second).Fn(&v)
	if v != time.Second {
		t.Errorf("below min should be raised to 1s, got %v", v)
	}
}

func TestClampMax(t *testing.T) {
	v := time.Second
	durations.RuleClampMax(time.Minute).Fn(&v)
	if v != time.Second {
		t.Errorf("below max should be unchanged, got %v", v)
	}
	v = time.Hour
	durations.RuleClampMax(time.Minute).Fn(&v)
	if v != time.Minute {
		t.Errorf("above max should be lowered to 1m, got %v", v)
	}
}

func TestAbs(t *testing.T) {
	v := -time.Second
	durations.RuleAbs().Fn(&v)
	if v != time.Second {
		t.Errorf("expected 1s, got %v", v)
	}
	v = time.Second
	durations.RuleAbs().Fn(&v)
	if v != time.Second {
		t.Errorf("positive unchanged, got %v", v)
	}
	v = 0
	durations.RuleAbs().Fn(&v)
	if v != 0 {
		t.Errorf("zero unchanged, got %v", v)
	}
	// MinInt64 nanoseconds — no positive counterpart
	v = time.Duration(-1 << 63)
	durations.RuleAbs().Fn(&v)
	if v != time.Duration(-1<<63) {
		t.Errorf("MinInt64 ns should be left as-is, got %v", v)
	}
}

func TestRound(t *testing.T) {
	v := 1500 * time.Millisecond
	durations.RuleRound(time.Second).Fn(&v)
	if v != 2*time.Second {
		t.Errorf("expected 2s (half-away-from-zero), got %v", v)
	}
	v = 1400 * time.Millisecond
	durations.RuleRound(time.Second).Fn(&v)
	if v != time.Second {
		t.Errorf("expected 1s, got %v", v)
	}
	// m <= 0 is no-op
	v = 1500 * time.Millisecond
	durations.RuleRound(0).Fn(&v)
	if v != 1500*time.Millisecond {
		t.Errorf("m=0 should be no-op, got %v", v)
	}
}

func TestTruncate(t *testing.T) {
	v := 1900 * time.Millisecond
	durations.RuleTruncate(time.Second).Fn(&v)
	if v != time.Second {
		t.Errorf("expected 1s (toward zero), got %v", v)
	}
	v = 2 * time.Second
	durations.RuleTruncate(time.Second).Fn(&v)
	if v != 2*time.Second {
		t.Errorf("exact multiple should be unchanged, got %v", v)
	}
	// m <= 0 is no-op
	v = 1900 * time.Millisecond
	durations.RuleTruncate(0).Fn(&v)
	if v != 1900*time.Millisecond {
		t.Errorf("m=0 should be no-op, got %v", v)
	}
}
