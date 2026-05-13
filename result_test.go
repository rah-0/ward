package ward_test

import (
	"testing"

	ward "github.com/rah-0/ward"
)

func TestAs(t *testing.T) {
	results := []*ward.Result{
		{FieldName: "email", RuleID: 1},
		{FieldName: "username", RuleID: 2},
	}

	names := ward.As(results, func(r *ward.Result) string { return r.FieldName })

	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d", len(names))
	}
	if names[0] != "email" {
		t.Errorf("expected %q, got %q", "email", names[0])
	}
	if names[1] != "username" {
		t.Errorf("expected %q, got %q", "username", names[1])
	}
}

func TestAs_Empty(t *testing.T) {
	names := ward.As([]*ward.Result{}, func(r *ward.Result) string { return r.FieldName })
	if len(names) != 0 {
		t.Fatalf("expected empty slice, got %d elements", len(names))
	}
}
