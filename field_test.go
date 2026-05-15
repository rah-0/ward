package ward_test

import (
	"testing"

	"github.com/rah-0/ward"
)

const testTypeID uint32 = 99

var (
	passingRule = ward.Rule[string]{ID: 1, Fn: func(s *string) *ward.Result { return nil }}
	failingRule = ward.Rule[string]{ID: 2, Fn: func(s *string) *ward.Result { return &ward.Result{Arg1: *s} }}
)

func newStringField(name string, value string, rules ...ward.Rule[string]) *ward.Field[string] {
	f := &ward.Field[string]{
		TypeID: testTypeID,
		Name:   name,
		Value:  &value,
		Rules:  rules,
	}
	return f
}

func TestFieldValidate_PassingRule(t *testing.T) {
	f := newStringField("email", "test@example.com", passingRule)
	results := f.Validate()
	if len(results) != 0 {
		t.Fatalf("expected no failures, got %d", len(results))
	}
}

func TestFieldValidate_FailingRule(t *testing.T) {
	f := newStringField("email", "bad", failingRule)
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
}

func TestFieldValidate_InjectsIDs(t *testing.T) {
	f := newStringField("email", "bad", failingRule)
	results := f.Validate()
	r := results[0]
	if r.TypeID != testTypeID {
		t.Errorf("expected TypeID %d, got %d", testTypeID, r.TypeID)
	}
	if r.RuleID != failingRule.ID {
		t.Errorf("expected RuleID %d, got %d", failingRule.ID, r.RuleID)
	}
	if r.FieldName != "email" {
		t.Errorf("expected FieldName %q, got %q", "email", r.FieldName)
	}
}

func TestFieldValidate_StopOnFail(t *testing.T) {
	f := newStringField("email", "bad", failingRule, failingRule, failingRule)
	f.Policy.StopOnFail = true
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure with StopOnFail, got %d", len(results))
	}
}

func TestFieldValidate_MultipleRules_AllFail(t *testing.T) {
	f := newStringField("email", "bad", failingRule, failingRule, failingRule)
	results := f.Validate()
	if len(results) != 3 {
		t.Fatalf("expected 3 failures, got %d", len(results))
	}
}

func TestFieldValidate_MixedRules(t *testing.T) {
	f := newStringField("email", "bad", passingRule, failingRule, passingRule)
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
}

func TestFieldValidate_PolicyError(t *testing.T) {
	f := newStringField("email", "test", passingRule)
	f.Policy.Required = true
	f.Policy.Optional = true
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 policy error, got %d", len(results))
	}
	if results[0].Err == nil {
		t.Fatal("expected Err to be set on policy error result")
	}
}
