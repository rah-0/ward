package ints_test

import (
	"testing"

	"github.com/rah-0/ward/types/ints"
)

func TestNew_SetsName(t *testing.T) {
	v := int64(42)
	f := ints.New("Age", &v, ints.RulePositive())
	if f.Name != "Age" {
		t.Errorf("expected Name %q, got %q", "Age", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	v := int64(42)
	f := ints.New("Age", &v, ints.RulePositive())
	if f.Value != &v {
		t.Error("expected Value to point to v")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	v := int64(42)
	f := ints.New("Age", &v)
	if f.TypeID != ints.TypeID {
		t.Errorf("expected TypeID %d, got %d", ints.TypeID, f.TypeID)
	}
}
