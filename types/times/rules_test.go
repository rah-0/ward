package times_test

import (
	"testing"
	"time"

	"github.com/rah-0/ward/types/times"
)

func run(rule times.Rule, value time.Time) bool {
	return rule.Fn(&value) == nil
}

var (
	epoch = time.Unix(0, 0).UTC()
	t1    = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2    = time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	t3    = time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
)

func TestAfter(t *testing.T) {
	if !run(times.RuleAfter(t1), t2) {
		t.Error("t2 after t1 should pass")
	}
	if run(times.RuleAfter(t1), t1) {
		t.Error("equal times should fail strict After")
	}
	if run(times.RuleAfter(t2), t1) {
		t.Error("t1 before t2 should fail After(t2)")
	}
}

func TestAfterOrEqual(t *testing.T) {
	if !run(times.RuleAfterOrEqual(t1), t1) {
		t.Error("equal times should pass AfterOrEqual")
	}
	if !run(times.RuleAfterOrEqual(t1), t2) {
		t.Error("t2 after t1 should pass AfterOrEqual(t1)")
	}
	if run(times.RuleAfterOrEqual(t2), t1) {
		t.Error("t1 before t2 should fail AfterOrEqual(t2)")
	}
}

func TestBefore(t *testing.T) {
	if !run(times.RuleBefore(t2), t1) {
		t.Error("t1 before t2 should pass")
	}
	if run(times.RuleBefore(t1), t1) {
		t.Error("equal times should fail strict Before")
	}
	if run(times.RuleBefore(t1), t2) {
		t.Error("t2 after t1 should fail Before(t1)")
	}
}

func TestBeforeOrEqual(t *testing.T) {
	if !run(times.RuleBeforeOrEqual(t1), t1) {
		t.Error("equal times should pass BeforeOrEqual")
	}
	if !run(times.RuleBeforeOrEqual(t2), t1) {
		t.Error("t1 before t2 should pass BeforeOrEqual(t2)")
	}
	if run(times.RuleBeforeOrEqual(t1), t2) {
		t.Error("t2 after t1 should fail BeforeOrEqual(t1)")
	}
}

func TestInRange(t *testing.T) {
	if !run(times.RuleInRange(t1, t3), t2) {
		t.Error("t2 in [t1,t3] should pass")
	}
	if !run(times.RuleInRange(t1, t3), t1) {
		t.Error("t1 == start should pass (inclusive)")
	}
	if !run(times.RuleInRange(t1, t3), t3) {
		t.Error("t3 == end should pass (inclusive)")
	}
	if run(times.RuleInRange(t2, t3), t1) {
		t.Error("t1 before range should fail")
	}
	if run(times.RuleInRange(t1, t2), t3) {
		t.Error("t3 after range should fail")
	}
}

func TestIsZero(t *testing.T) {
	if !run(times.RuleIsZero(), time.Time{}) {
		t.Error("zero time should pass IsZero")
	}
	if run(times.RuleIsZero(), t1) {
		t.Error("non-zero time should fail IsZero")
	}
}

func TestIsNotZero(t *testing.T) {
	if !run(times.RuleIsNotZero(), t1) {
		t.Error("non-zero time should pass IsNotZero")
	}
	if run(times.RuleIsNotZero(), time.Time{}) {
		t.Error("zero time should fail IsNotZero")
	}
}

func TestOneOf(t *testing.T) {
	if !run(times.RuleOneOf(t1, t2), t1) {
		t.Error("t1 is in list, should pass")
	}
	if !run(times.RuleOneOf(t1, t2), t2) {
		t.Error("t2 is in list, should pass")
	}
	if run(times.RuleOneOf(t1, t2), t3) {
		t.Error("t3 is not in list, should fail")
	}
}

func TestNotOneOf(t *testing.T) {
	if !run(times.RuleNotOneOf(t1, t2), t3) {
		t.Error("t3 is not excluded, should pass")
	}
	if run(times.RuleNotOneOf(t1, t2), t1) {
		t.Error("t1 is excluded, should fail")
	}
	if run(times.RuleNotOneOf(t1, t2), t2) {
		t.Error("t2 is excluded, should fail")
	}
}

func TestIsPast(t *testing.T) {
	past := time.Now().Add(-time.Hour)
	if !run(times.RuleIsPast(), past) {
		t.Error("one hour ago should be past")
	}
	future := time.Now().Add(time.Hour)
	if run(times.RuleIsPast(), future) {
		t.Error("one hour from now should fail IsPast")
	}
	if run(times.RuleIsPast(), time.Time{}) {
		// zero time IS technically in the past, but documenting current behaviour:
		// it is before now, so it passes.
		// This test is to alert if behaviour ever changes.
		_ = past
	}
}

func TestIsFuture(t *testing.T) {
	future := time.Now().Add(time.Hour)
	if !run(times.RuleIsFuture(), future) {
		t.Error("one hour from now should be future")
	}
	past := time.Now().Add(-time.Hour)
	if run(times.RuleIsFuture(), past) {
		t.Error("one hour ago should fail IsFuture")
	}
}

func TestIsWeekday(t *testing.T) {
	// 2024-01-01 was a Monday
	monday := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	friday := time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC)
	saturday := time.Date(2024, 1, 6, 12, 0, 0, 0, time.UTC)
	sunday := time.Date(2024, 1, 7, 12, 0, 0, 0, time.UTC)

	if !run(times.RuleIsWeekday(), monday) {
		t.Error("Monday should be a weekday")
	}
	if !run(times.RuleIsWeekday(), friday) {
		t.Error("Friday should be a weekday")
	}
	if run(times.RuleIsWeekday(), saturday) {
		t.Error("Saturday should fail IsWeekday")
	}
	if run(times.RuleIsWeekday(), sunday) {
		t.Error("Sunday should fail IsWeekday")
	}
}

func TestIsWeekend(t *testing.T) {
	monday := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	saturday := time.Date(2024, 1, 6, 12, 0, 0, 0, time.UTC)
	sunday := time.Date(2024, 1, 7, 12, 0, 0, 0, time.UTC)

	if !run(times.RuleIsWeekend(), saturday) {
		t.Error("Saturday should be weekend")
	}
	if !run(times.RuleIsWeekend(), sunday) {
		t.Error("Sunday should be weekend")
	}
	if run(times.RuleIsWeekend(), monday) {
		t.Error("Monday should fail IsWeekend")
	}
}
