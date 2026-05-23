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
	// NaN / ±Inf must fail — they have no decimal representation.
	if run(floats.RuleMaxDecimalPlaces(2), math.NaN()) {
		t.Error("NaN should fail MaxDecimalPlaces")
	}
	if run(floats.RuleMaxDecimalPlaces(2), math.Inf(1)) {
		t.Error("+Inf should fail MaxDecimalPlaces")
	}
	if run(floats.RuleMaxDecimalPlaces(2), math.Inf(-1)) {
		t.Error("-Inf should fail MaxDecimalPlaces")
	}
	// Negative n is clamped to 0 — 1.5 has 1 decimal place, fails max=0.
	if run(floats.RuleMaxDecimalPlaces(-1), 1.5) {
		t.Error("MaxDecimalPlaces(-1) on 1.5 should fail (negative n clamped to 0)")
	}
	if !run(floats.RuleMaxDecimalPlaces(-1), 5.0) {
		t.Error("MaxDecimalPlaces(-1) on 5.0 should pass (5.0 has 0 decimals)")
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

	// Large n that overflows 10^n to +Inf must not corrupt v to NaN.
	v = 3.14
	floats.RuleRound(400).Fn(&v)
	if math.IsNaN(v) {
		t.Error("Round(400) must not produce NaN")
	}
	if v != 3.14 {
		t.Errorf("Round(400) should leave v unchanged when shift overflows, got %v", v)
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

func TestIsNumber(t *testing.T) {
	if !run(floats.RuleIsNumber(), 1.0) {
		t.Error("1.0 should pass IsNumber")
	}
	if !run(floats.RuleIsNumber(), math.Inf(1)) {
		t.Error("+Inf should pass IsNumber")
	}
	if !run(floats.RuleIsNumber(), math.Inf(-1)) {
		t.Error("-Inf should pass IsNumber")
	}
	if run(floats.RuleIsNumber(), math.NaN()) {
		t.Error("NaN should fail IsNumber")
	}
}

func TestApproxEqual(t *testing.T) {
	if !run(floats.RuleApproxEqual(1.0, 0.01), 1.005) {
		t.Error("1.005 should be within 0.01 of 1.0")
	}
	if !run(floats.RuleApproxEqual(1.0, 0.0), 1.0) {
		t.Error("exact equality with zero tolerance should pass")
	}
	if run(floats.RuleApproxEqual(1.0, 0.01), 1.1) {
		t.Error("1.1 should not be within 0.01 of 1.0")
	}
	// negative tolerance treated as 0 (exact)
	if !run(floats.RuleApproxEqual(1.0, -1.0), 1.0) {
		t.Error("negative tolerance should be treated as 0 — exact match should pass")
	}
	if run(floats.RuleApproxEqual(1.0, -1.0), 1.0001) {
		t.Error("negative tolerance should be treated as 0 — non-exact should fail")
	}
	// NaN always fails
	if run(floats.RuleApproxEqual(1.0, 0.01), math.NaN()) {
		t.Error("NaN input should fail ApproxEqual")
	}
	if run(floats.RuleApproxEqual(math.NaN(), 0.01), 1.0) {
		t.Error("NaN target should fail ApproxEqual")
	}
}

func TestClampMinFloat(t *testing.T) {
	v := 5.0
	floats.RuleClampMin(0.0).Fn(&v)
	if v != 5.0 {
		t.Errorf("above min should be unchanged, got %v", v)
	}
	v = -10.0
	floats.RuleClampMin(0.0).Fn(&v)
	if v != 0.0 {
		t.Errorf("below min should be raised, got %v", v)
	}
	// NaN unchanged
	v = math.NaN()
	floats.RuleClampMin(0.0).Fn(&v)
	if !math.IsNaN(v) {
		t.Error("NaN should be left unchanged by ClampMin")
	}
}

func TestClampMaxFloat(t *testing.T) {
	v := 5.0
	floats.RuleClampMax(10.0).Fn(&v)
	if v != 5.0 {
		t.Errorf("below max should be unchanged, got %v", v)
	}
	v = 100.0
	floats.RuleClampMax(10.0).Fn(&v)
	if v != 10.0 {
		t.Errorf("above max should be lowered, got %v", v)
	}
	// NaN unchanged
	v = math.NaN()
	floats.RuleClampMax(10.0).Fn(&v)
	if !math.IsNaN(v) {
		t.Error("NaN should be left unchanged by ClampMax")
	}
}

func TestTrunc(t *testing.T) {
	v := 1.7
	floats.RuleTrunc().Fn(&v)
	if v != 1.0 {
		t.Errorf("expected 1.0, got %v", v)
	}
	v = -1.7
	floats.RuleTrunc().Fn(&v)
	if v != -1.0 {
		t.Errorf("expected -1.0 (toward zero), got %v", v)
	}
	v = 0.0
	floats.RuleTrunc().Fn(&v)
	if v != 0.0 {
		t.Errorf("zero unchanged, got %v", v)
	}
}

// -----------------------------------------------------------------------------
// Reusability / closure-state regressions
//
// These tests guard against a class of bug where a rule constructor accepts
// an argument that needs normalization (e.g. negative tolerance, negative n)
// and the closure mutates the captured parameter on the first call. After
// that first call, the rule reports the *normalized* value back in Arg1/Arg2
// rather than the value the user originally supplied — and concurrent use
// becomes a data race. Each test invokes a single Rule value multiple times
// and asserts the reported Args remain equal to the constructor arguments.
// -----------------------------------------------------------------------------

// TestApproxEqual_ReusableAcrossCalls: negative tolerance is normalized to 0
// at construction time. Arg2 must report the *effective* tolerance, and that
// value must be identical on every invocation regardless of which code path
// the call takes.
func TestApproxEqual_ReusableAcrossCalls(t *testing.T) {
	rule := floats.RuleApproxEqual(1.0, -1.0)

	// Failure path with NaN — exercises the early-return arm.
	v := math.NaN()
	r1 := rule.Fn(&v)
	if r1 == nil {
		t.Fatal("NaN should fail ApproxEqual")
	}

	// Successful non-NaN path — under the old bug, this call rewrote the
	// captured tolerance to 0, contaminating subsequent failure reports.
	v = 1.0
	if rule.Fn(&v) != nil {
		t.Fatal("exact match with negative tolerance should pass")
	}

	// Failure path again — Arg2 must equal the value seen on the first call.
	v = math.NaN()
	r3 := rule.Fn(&v)
	if r3 == nil {
		t.Fatal("NaN should fail ApproxEqual on second call")
	}
	if r1.Arg2 != r3.Arg2 {
		t.Errorf("Arg2 differs across calls (closure mutation): %v vs %v", r1.Arg2, r3.Arg2)
	}
	if r3.Arg2 != 0.0 {
		t.Errorf("Arg2 should report the normalized tolerance 0, got %v", r3.Arg2)
	}

	// Failure on the non-NaN compare path — Arg2 must match the others.
	v = 5.0
	r4 := rule.Fn(&v)
	if r4 == nil {
		t.Fatal("5.0 should fail ApproxEqual(target=1.0, tolerance=-1.0)")
	}
	if r4.Arg2 != r1.Arg2 {
		t.Errorf("Arg2 on compare-fail path differs: %v vs %v", r4.Arg2, r1.Arg2)
	}
}

// TestApproxEqual_PositiveTolerancePreserved guards the simple case where
// no normalization happens — Arg1/Arg2 must equal the user-supplied target
// and tolerance exactly, even after many calls.
func TestApproxEqual_PositiveTolerancePreserved(t *testing.T) {
	rule := floats.RuleApproxEqual(1.0, 0.5)

	for i := 0; i < 5; i++ {
		v := 1.2
		if rule.Fn(&v) != nil {
			t.Fatalf("iter %d: 1.2 within 0.5 of 1.0 should pass", i)
		}
		v = 5.0
		r := rule.Fn(&v)
		if r == nil {
			t.Fatalf("iter %d: 5.0 should fail", i)
		}
		if r.Arg1 != 1.0 || r.Arg2 != 0.5 {
			t.Errorf("iter %d: expected Arg1=1.0 Arg2=0.5, got Arg1=%v Arg2=%v", i, r.Arg1, r.Arg2)
		}
	}
}

func TestMaxDecimalPlaces_ReusableAcrossCalls(t *testing.T) {
	rule := floats.RuleMaxDecimalPlaces(-1)

	// Failure on NaN — exercises the early-return arm before normalization.
	v := math.NaN()
	r1 := rule.Fn(&v)
	if r1 == nil {
		t.Fatal("NaN should fail MaxDecimalPlaces")
	}

	// Path that previously wrote n = 0 into the closure (any non-NaN/Inf input
	// reaches the normalization branch).
	v = 5.0
	if rule.Fn(&v) != nil {
		t.Fatal("integer should pass MaxDecimalPlaces(-1)")
	}

	// Now NaN again — Arg1 must still match the original n = -1.
	v = math.NaN()
	r3 := rule.Fn(&v)
	if r3 == nil {
		t.Fatal("NaN should fail MaxDecimalPlaces on second call")
	}
	if r1.Arg1 != r3.Arg1 {
		t.Errorf("Arg1 differs across calls (closure mutation): %v vs %v", r1.Arg1, r3.Arg1)
	}
	if r3.Arg1 != 0 {
		// Negative n is normalized to 0 by spec; both calls should see 0.
		t.Errorf("Arg1 should be normalized to 0, got %v", r3.Arg1)
	}

	// Failure on the fractional path — Arg1 must still match.
	v = 1.5
	r4 := rule.Fn(&v)
	if r4 == nil {
		t.Fatal("1.5 should fail MaxDecimalPlaces(0)")
	}
	if r4.Arg1 != 0 {
		t.Errorf("Arg1 on fractional-fail path should be 0, got %v", r4.Arg1)
	}
}

// TestApproxEqual_ConcurrentReuse runs the same rule from many goroutines
// to flag any data race introduced by closure-state mutation. With -race,
// any concurrent write to a shared closure variable would fail the run.
func TestApproxEqual_ConcurrentReuse(t *testing.T) {
	rule := floats.RuleApproxEqual(1.0, -1.0)

	const goroutines = 32
	const iterations = 1000
	done := make(chan struct{}, goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < iterations; i++ {
				v := 1.0
				rule.Fn(&v)
			}
			done <- struct{}{}
		}()
	}
	for g := 0; g < goroutines; g++ {
		<-done
	}
}

func TestMaxDecimalPlaces_ConcurrentReuse(t *testing.T) {
	rule := floats.RuleMaxDecimalPlaces(-1)

	const goroutines = 32
	const iterations = 1000
	done := make(chan struct{}, goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < iterations; i++ {
				v := 1.5
				rule.Fn(&v)
			}
			done <- struct{}{}
		}()
	}
	for g := 0; g < goroutines; g++ {
		<-done
	}
}
