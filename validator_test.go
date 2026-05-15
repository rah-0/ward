package ward_test

import (
	"testing"

	"github.com/rah-0/ward"
)

func makeField(value string, rules ...ward.Rule[string]) *ward.Field[string] {
	return &ward.Field[string]{
		TypeID: testTypeID,
		Name:   "field",
		Value:  &value,
		Rules:  rules,
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
	v := ward.New()
	v.Policy.StopOnFail = true
	v.Add(makeField("bad", failingRule))
	v.Add(makeField("bad", failingRule))
	v.Add(makeField("bad", failingRule))
	v.Run()
	if len(v.Failures()) != 1 {
		t.Fatalf("expected 1 failure with StopOnFail, got %d", len(v.Failures()))
	}
}
