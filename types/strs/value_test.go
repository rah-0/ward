package strs_test

import (
	"testing"

	"github.com/rah-0/ward/types/strs"
)

func TestNew_SetsName(t *testing.T) {
	s := "hello"
	f := strs.New("Email", &s, strs.RuleNotEmpty())
	if f.Name != "Email" {
		t.Errorf("expected Name %q, got %q", "Email", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	s := "hello"
	f := strs.New("Email", &s, strs.RuleNotEmpty())
	if f.Value != &s {
		t.Error("expected Value to point to s")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	s := "hello"
	f := strs.New("Email", &s, strs.RuleNotEmpty())
	if f.TypeID != strs.TypeID {
		t.Errorf("expected TypeID %d, got %d", strs.TypeID, f.TypeID)
	}
}

func TestValidate_PassThroughPointer(t *testing.T) {
	s := "hello@example.com"
	f := strs.New("Email", &s, strs.RuleIsEmail())
	results := f.Validate()
	if len(results) != 0 {
		t.Fatalf("expected no failures, got %d", len(results))
	}
}

func TestValidate_PicksUpPointerChange(t *testing.T) {
	s := "hello@example.com"
	f := strs.New("Email", &s, strs.RuleIsEmail())
	f.Validate()

	s = "not-an-email"
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure after value change, got %d", len(results))
	}
}

func TestValidate_SanitizerWritesBack(t *testing.T) {
	s := "  hello  "
	f := strs.New("Name", &s, strs.RuleTrim())
	f.Validate()
	if s != "hello" {
		t.Errorf("expected s to be trimmed to %q, got %q", "hello", s)
	}
}

func TestValidate_InjectsFieldName(t *testing.T) {
	s := ""
	f := strs.New("Username", &s, strs.RuleNotEmpty())
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].FieldName != "Username" {
		t.Errorf("expected FieldName %q, got %q", "Username", results[0].FieldName)
	}
}

func TestValidate_InjectsTypeIDAndRuleID(t *testing.T) {
	s := ""
	f := strs.New("Field", &s, strs.RuleNotEmpty())
	results := f.Validate()
	if results[0].TypeID != strs.TypeID {
		t.Errorf("expected TypeID %d, got %d", strs.TypeID, results[0].TypeID)
	}
	if results[0].RuleID != strs.IDNotEmpty {
		t.Errorf("expected RuleID %d, got %d", strs.IDNotEmpty, results[0].RuleID)
	}
}

func TestValidate_FailingRule_SetsArg1(t *testing.T) {
	s := "ab"
	f := strs.New("Field", &s, strs.RuleLengthMin(5))
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].Arg1 != 5 {
		t.Errorf("expected Arg1=5, got %v", results[0].Arg1)
	}
}

func TestValidate_FailingRule_SetsArg1AndArg2(t *testing.T) {
	s := "a"
	f := strs.New("Field", &s, strs.RuleLengthBetween(3, 10))
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].Arg1 != 3 {
		t.Errorf("expected Arg1=3, got %v", results[0].Arg1)
	}
	if results[0].Arg2 != 10 {
		t.Errorf("expected Arg2=10, got %v", results[0].Arg2)
	}
}

func TestValidate_FailingRule_SetsErr(t *testing.T) {
	s := "not-an-email"
	f := strs.New("Email", &s, strs.RuleIsEmail())
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].Err == nil {
		t.Error("expected Err to be set for IsEmail failure")
	}
}

func TestValidate_PolicyError_BlocksRules(t *testing.T) {
	s := "hello"
	f := strs.New("Field", &s, strs.RuleNotEmpty())
	f.Policy.Required = true
	f.Policy.Optional = true
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 policy error result, got %d", len(results))
	}
	if results[0].Err == nil {
		t.Error("expected Err to be set on policy error")
	}
	if results[0].RuleID != 0 {
		t.Error("expected RuleID to be zero — rule should not have run")
	}
}

func TestValidate_StopOnFail_FieldLevel(t *testing.T) {
	s := "a"
	f := strs.New("Field", &s, strs.RuleLengthMin(5), strs.RuleLengthMin(10), strs.RuleLengthMin(20))
	f.Policy.StopOnFail = true
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure with StopOnFail, got %d", len(results))
	}
}

func TestValidate_MultipleFailingRules(t *testing.T) {
	s := ""
	f := strs.New("Field", &s, strs.RuleNotEmpty(), strs.RuleLengthMin(5), strs.RuleIsEmail())
	results := f.Validate()
	if len(results) != 3 {
		t.Fatalf("expected 3 failures, got %d", len(results))
	}
}

func TestValidate_NoRules(t *testing.T) {
	s := "anything"
	f := strs.New("Field", &s)
	results := f.Validate()
	if len(results) != 0 {
		t.Fatalf("expected no failures with no rules, got %d", len(results))
	}
}

func TestValidate_EmptyName(t *testing.T) {
	s := ""
	f := strs.New("", &s, strs.RuleNotEmpty())
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].FieldName != "" {
		t.Errorf("expected empty FieldName, got %q", results[0].FieldName)
	}
}
