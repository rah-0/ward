package uuids_test

import (
	"testing"

	"github.com/rah-0/ward/types/uuids"
)

func TestNew_SetsName(t *testing.T) {
	s := validUUID1
	f := uuids.New("UserID", &s, uuids.RuleIsValid())
	if f.Name != "UserID" {
		t.Errorf("expected Name %q, got %q", "UserID", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	s := validUUID1
	f := uuids.New("UserID", &s, uuids.RuleIsValid())
	if f.Value != &s {
		t.Error("expected Value to point to s")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	s := validUUID1
	f := uuids.New("UserID", &s, uuids.RuleIsValid())
	if f.TypeID != uuids.TypeID {
		t.Errorf("expected TypeID %d, got %d", uuids.TypeID, f.TypeID)
	}
}

func TestValidate_ValidUUID(t *testing.T) {
	s := validUUID1
	f := uuids.New("UserID", &s, uuids.RuleIsValid())
	results := f.Validate()
	if len(results) != 0 {
		t.Fatalf("expected no failures, got %d", len(results))
	}
}

func TestValidate_InvalidUUID(t *testing.T) {
	s := "not-a-uuid"
	f := uuids.New("UserID", &s, uuids.RuleIsValid())
	results := f.Validate()
	if len(results) != 1 {
		t.Fatalf("expected 1 failure, got %d", len(results))
	}
	if results[0].TypeID != uuids.TypeID {
		t.Errorf("expected TypeID %d, got %d", uuids.TypeID, results[0].TypeID)
	}
	if results[0].RuleID != uuids.IDIsValid {
		t.Errorf("expected RuleID %d, got %d", uuids.IDIsValid, results[0].RuleID)
	}
}
