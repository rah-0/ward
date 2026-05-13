package bools_test

import (
	"testing"

	"github.com/rah-0/ward/types/bools"
)

func TestNew_SetsName(t *testing.T) {
	v := true
	f := bools.New("Enabled", &v, bools.RuleIsTrue())
	if f.Name != "Enabled" {
		t.Errorf("expected Name %q, got %q", "Enabled", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	v := true
	f := bools.New("Enabled", &v)
	if f.Value != &v {
		t.Error("expected Value to point to v")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	v := false
	f := bools.New("Enabled", &v)
	if f.TypeID != bools.TypeID {
		t.Errorf("expected TypeID %d, got %d", bools.TypeID, f.TypeID)
	}
}
