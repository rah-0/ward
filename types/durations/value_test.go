package durations_test

import (
	"testing"
	"time"

	"github.com/rah-0/ward/types/durations"
)

func TestNew_SetsName(t *testing.T) {
	v := 5 * time.Second
	f := durations.New("Timeout", &v, durations.RulePositive())
	if f.Name != "Timeout" {
		t.Errorf("expected Name %q, got %q", "Timeout", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	v := 5 * time.Second
	f := durations.New("Timeout", &v)
	if f.Value != &v {
		t.Error("expected Value to point to v")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	v := 5 * time.Second
	f := durations.New("Timeout", &v)
	if f.TypeID != durations.TypeID {
		t.Errorf("expected TypeID %d, got %d", durations.TypeID, f.TypeID)
	}
}
