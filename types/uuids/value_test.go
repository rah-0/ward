package uuids_test

import (
	"testing"

	"github.com/rah-0/ward/types/uuids"
)

func TestNew_SetsName(t *testing.T) {
	s := validV4_1
	f := uuids.New("UserID", &s, uuids.RuleIsValidV4())
	if f.Name != "UserID" {
		t.Errorf("expected Name %q, got %q", "UserID", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	s := validV4_1
	f := uuids.New("UserID", &s, uuids.RuleIsValidV4())
	if f.Value != &s {
		t.Error("expected Value to point to s")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	s := validV4_1
	f := uuids.New("UserID", &s, uuids.RuleIsValidV4())
	if f.TypeID != uuids.TypeID {
		t.Errorf("expected TypeID %d, got %d", uuids.TypeID, f.TypeID)
	}
}

func TestValidate_ValidV4(t *testing.T) {
	s := validV4_1
	f := uuids.New("UserID", &s, uuids.RuleIsValidV4())
	results := f.Validate()
	if len(results) != 0 {
		t.Fatalf("expected no failures, got %d", len(results))
	}
}

func TestValidate_InvalidV4(t *testing.T) {
	s := "not-a-uuid"
	f := uuids.New("UserID", &s, uuids.RuleIsValidV4())
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].TypeID != uuids.TypeID {
		t.Errorf("expected TypeID %d, got %d", uuids.TypeID, results[0].TypeID)
	}
	if results[0].RuleID != uuids.IDIsValidV4 {
		t.Errorf("expected RuleID %d, got %d", uuids.IDIsValidV4, results[0].RuleID)
	}
}
