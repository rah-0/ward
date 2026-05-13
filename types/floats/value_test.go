package floats_test

import (
	"testing"

	"github.com/rah-0/ward/types/floats"
)

func TestNew_SetsName(t *testing.T) {
	v := 3.14
	f := floats.New("Price", &v, floats.RulePositive())
	if f.Name != "Price" {
		t.Errorf("expected Name %q, got %q", "Price", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	v := 3.14
	f := floats.New("Price", &v, floats.RulePositive())
	if f.Value != &v {
		t.Error("expected Value to point to v")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	v := 3.14
	f := floats.New("Price", &v)
	if f.TypeID != floats.TypeID {
		t.Errorf("expected TypeID %d, got %d", floats.TypeID, f.TypeID)
	}
}
