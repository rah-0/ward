package ward_test

import (
	"errors"
	"testing"

	"github.com/rah-0/ward"
)

func TestFieldPolicyValidate_RequiredAndOptional(t *testing.T) {
	p := ward.FieldPolicy{Required: true, Optional: true}
	if err := p.Validate(); !errors.Is(err, ward.ErrFieldRequiredWithOptional) {
		t.Fatalf("expected ErrFieldRequiredWithOptional, got %v", err)
	}
}

func TestFieldPolicyValidate_RequiredAndAllowNil(t *testing.T) {
	p := ward.FieldPolicy{Required: true, AllowNil: true}
	if err := p.Validate(); !errors.Is(err, ward.ErrFieldRequiredWithAllowNil) {
		t.Fatalf("expected ErrFieldRequiredWithAllowNil, got %v", err)
	}
}

func TestFieldPolicyValidate_Valid(t *testing.T) {
	cases := []ward.FieldPolicy{
		{},
		{Required: true},
		{Optional: true},
		{AllowDefault: true},
		{AllowNil: true},
		{StopOnFail: true},
		{Required: true, StopOnFail: true},
	}
	for _, p := range cases {
		if err := p.Validate(); err != nil {
			t.Fatalf("expected nil, got %v for policy %+v", err, p)
		}
	}
}
