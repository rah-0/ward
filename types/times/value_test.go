package times_test

import (
	"testing"
	"time"

	"github.com/rah-0/ward/types/times"
)

func TestNew_SetsName(t *testing.T) {
	v := time.Now()
	f := times.New("CreatedAt", &v, times.RuleIsNotZero())
	if f.Name != "CreatedAt" {
		t.Errorf("expected Name %q, got %q", "CreatedAt", f.Name)
	}
}

func TestNew_BindsPointer(t *testing.T) {
	v := time.Now()
	f := times.New("CreatedAt", &v)
	if f.Value != &v {
		t.Error("expected Value to point to v")
	}
}

func TestNew_SetsTypeID(t *testing.T) {
	v := time.Now()
	f := times.New("CreatedAt", &v)
	if f.TypeID != times.TypeID {
		t.Errorf("expected TypeID %d, got %d", times.TypeID, f.TypeID)
	}
}
