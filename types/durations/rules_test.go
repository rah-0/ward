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

func TestLessThan(t *testing.T) {
	if !run(durations.RuleLessThan(time.Second), time.Millisecond) {
		t.Error("1ms < 1s should pass")
	}
	if run(durations.RuleLessThan(time.Second), time.Second) {
		t.Error("1s < 1s should fail")
	}
	if run(durations.RuleLessThan(time.Second), 2*time.Second) {
		t.Error("2s < 1s should fail")
	}
}

func TestLessThanOrEqual(t *testing.T) {
	if !run(durations.RuleLessThanOrEqual(time.Second), time.Second) {
		t.Error("1s <= 1s should pass")
	}
	if !run(durations.RuleLessThanOrEqual(time.Second), time.Millisecond) {
		t.Error("1ms <= 1s should pass")
	}
	if run(durations.RuleLessThanOrEqual(time.Second), 2*time.Second) {
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

func TestNonNegative(t *testing.T) {
	if !run(durations.RuleNonNegative(), 0) {
		t.Error("0 should be non-negative")
	}
	if !run(durations.RuleNonNegative(), time.Second) {
		t.Error("1s should be non-negative")
	}
	if run(durations.RuleNonNegative(), -time.Second) {
		t.Error("negative duration should fail non-negative")
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
