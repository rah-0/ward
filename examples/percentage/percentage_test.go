package percentage_test

import (
	"testing"

	"github.com/rah-0/ward"
	"github.com/rah-0/ward/examples/percentage"
)

func newScoreField(value *float64) *percentage.Field {
	return percentage.New("Score", value,
		percentage.RuleInRange(0, 100),
		percentage.RuleIsWhole(),
	)
}

func TestPercentage_Valid(t *testing.T) {
	value := 75.0
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	if v.HasFailures() {
		t.Fatalf("expected no failures, got %d", len(v.Failures()))
	}
}

func TestPercentage_Zero(t *testing.T) {
	value := 0.0
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	if v.HasFailures() {
		t.Fatalf("expected no failures for 0, got %d", len(v.Failures()))
	}
}

func TestPercentage_Hundred(t *testing.T) {
	value := 100.0
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	if v.HasFailures() {
		t.Fatalf("expected no failures for 100, got %d", len(v.Failures()))
	}
}

func TestPercentage_BelowRange(t *testing.T) {
	value := -1.0
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for value below range")
	}
	if v.Failures()[0].RuleID != percentage.IDInRange {
		t.Errorf("expected RuleID %d, got %d", percentage.IDInRange, v.Failures()[0].RuleID)
	}
	if v.Failures()[0].Arg1 != 0.0 || v.Failures()[0].Arg2 != 100.0 {
		t.Errorf("expected Arg1=0 Arg2=100, got Arg1=%v Arg2=%v", v.Failures()[0].Arg1, v.Failures()[0].Arg2)
	}
}

func TestPercentage_AboveRange(t *testing.T) {
	value := 101.0
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for value above range")
	}
	if v.Failures()[0].RuleID != percentage.IDInRange {
		t.Errorf("expected RuleID %d, got %d", percentage.IDInRange, v.Failures()[0].RuleID)
	}
}

func TestPercentage_NotWhole(t *testing.T) {
	value := 75.5
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for non-whole value")
	}
	if v.Failures()[0].RuleID != percentage.IDIsWhole {
		t.Errorf("expected RuleID %d, got %d", percentage.IDIsWhole, v.Failures()[0].RuleID)
	}
}

func TestPercentage_IsPositive_Zero(t *testing.T) {
	value := 0.0
	v := ward.New()
	v.Add(percentage.New("Score", &value, percentage.RuleIsPositive()))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for 0 with IsPositive")
	}
}

func TestPercentage_TypeID(t *testing.T) {
	value := -5.0
	v := ward.New()
	v.Add(newScoreField(&value))
	v.Run()
	for _, f := range v.Failures() {
		if f.TypeID != percentage.TypeID {
			t.Errorf("expected TypeID %d, got %d", percentage.TypeID, f.TypeID)
		}
	}
}
