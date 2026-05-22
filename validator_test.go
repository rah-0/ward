package ward_test

import (
	"testing"

	"github.com/rah-0/ward"
)

func makeField(value string, rules ...ward.Rule[string]) *ward.Field[string] {
	stamped := make([]ward.Rule[string], len(rules))
	copy(stamped, rules)
	for i := range stamped {
		stamped[i].TypeID = testTypeID
	}
	return &ward.Field[string]{
		TypeID: testTypeID,
		Name:   "field",
		Value:  &value,
		Rules:  stamped,
	}
}

func TestValidatorNew(t *testing.T) {
	v := ward.New()
	if v == nil {
		t.Fatal("expected non-nil Validate")
	}
}

func TestValidatorRun_NoFailures(t *testing.T) {
	v := ward.New().Add(makeField("ok", passingRule)).Run()
	if v.HasFailures() {
		t.Fatalf("expected no failures, got %d", len(v.Failures()))
	}
}

func TestValidatorRun_WithFailures(t *testing.T) {
	v := ward.New().Add(makeField("bad", failingRule)).Run()
	if !v.HasFailures() {
		t.Fatal("expected HasFailures true")
	}
	if len(v.Failures()) != 1 {
		t.Fatalf("expected Failures len 1, got %d", len(v.Failures()))
	}
}

func TestValidatorRun_MultipleFields(t *testing.T) {
	v := ward.New()
	v.Add(makeField("ok", passingRule))
	v.Add(makeField("bad", failingRule))
	v.Add(makeField("bad", failingRule))
	v.Run()
	if len(v.Failures()) != 2 {
		t.Fatalf("expected 2 failures, got %d", len(v.Failures()))
	}
}

func TestValidatorRun_StopOnFail(t *testing.T) {
	var subsequentCalls int
	countingRule := ward.Rule[string]{ID: 4, Fn: func(s *string) *ward.Result {
		subsequentCalls++
		return &ward.Result{}
	}}
	v := ward.New()
	v.Policy.StopOnFail = true
	v.Add(makeField("bad", failingRule))
	v.Add(makeField("bad", countingRule))
	v.Add(makeField("bad", countingRule))
	v.Run()
	if len(v.Failures()) != 1 {
		t.Fatalf("expected 1 failure with StopOnFail, got %d", len(v.Failures()))
	}
	if subsequentCalls != 0 {
		t.Fatalf("expected 0 subsequent rule invocations after StopOnFail, got %d", subsequentCalls)
	}
}
