package ward_test

import (
	"testing"

	"github.com/rah-0/ward"
)

func TestValidateFailuresAs(t *testing.T) {
	v := ward.New().Add(
		newStringField("email", "bad", failingRule),
		newStringField("username", "bad", failingRule),
	).Run()

	names := v.FailuresAs(func(r *ward.Result) string { return r.FieldName })

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

func TestValidateFailuresAs_Empty(t *testing.T) {
	names := ward.New().FailuresAs(func(r *ward.Result) string { return r.FieldName })
	if len(names) != 0 {
		t.Fatalf("expected empty slice, got %d elements", len(names))
	}
}
