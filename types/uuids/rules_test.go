package uuids_test

import (
	"testing"

	"github.com/rah-0/ward/types/uuids"
)

func run(rule uuids.Rule, value string) bool {
	return rule.Fn(&value) == nil
}

const (
	// v4 UUIDs
	validV4_1 = "550e8400-e29b-41d4-a716-446655440000"
	validV4_2 = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	validV4_3 = "9a72e3ed-4b6f-4adc-9c0a-7f0e74b5b6a0"

	// v1 UUID
	validV1 = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

	// v7 UUID (time-ordered)
	validV7 = "018f6c4e-79ed-7c5e-a716-446655440000"

	nilUUID = "00000000-0000-0000-0000-000000000000"
)

// -----------------------------------------------------------------------------
// RuleIsValidV4
// -----------------------------------------------------------------------------

func TestIsValidV4_AcceptsV4(t *testing.T) {
	if !run(uuids.RuleIsValidV4(), validV4_1) {
		t.Error("valid v4 UUID should pass")
	}
}

func TestIsValidV4_AcceptsUppercase(t *testing.T) {
	if !run(uuids.RuleIsValidV4(), "550E8400-E29B-41D4-A716-446655440000") {
		t.Error("uppercase v4 UUID should pass")
	}
}

func TestIsValidV4_RejectsV1(t *testing.T) {
	if run(uuids.RuleIsValidV4(), validV1) {
		t.Error("v1 UUID should fail IsValidV4")
	}
}

func TestIsValidV4_RejectsV7(t *testing.T) {
	if run(uuids.RuleIsValidV4(), validV7) {
		t.Error("v7 UUID should fail IsValidV4")
	}
}

func TestIsValidV4_RejectsNil(t *testing.T) {
	if run(uuids.RuleIsValidV4(), nilUUID) {
		t.Error("nil UUID should fail IsValidV4 (it is version 0)")
	}
}

func TestIsValidV4_RejectsMalformed(t *testing.T) {
	for _, v := range []string{
		"",
		"not-a-uuid",
		"550e8400-e29b-41d4-a716-44665544000",  // too short
		"zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz", // non-hex
	} {
		if run(uuids.RuleIsValidV4(), v) {
			t.Errorf("%q should fail IsValidV4", v)
		}
	}
}

// TestIsValidV4_AcceptsAlternateForms documents that google/uuid.Parse accepts
// several non-canonical encodings besides the hyphenated form. These all decode
// to the same underlying v4 UUID, so IsValidV4 accepts them.
func TestIsValidV4_AcceptsAlternateForms(t *testing.T) {
	for _, v := range []string{
		"550e8400e29b41d4a716446655440000",              // raw hex, no hyphens
		"urn:uuid:550e8400-e29b-41d4-a716-446655440000", // URN form
		"{550e8400-e29b-41d4-a716-446655440000}",        // Microsoft braces
	} {
		if !run(uuids.RuleIsValidV4(), v) {
			t.Errorf("%q should pass IsValidV4 (accepted by uuid.Parse)", v)
		}
	}
}

// -----------------------------------------------------------------------------
// RuleIsNotNilV4
// -----------------------------------------------------------------------------

func TestIsNotNilV4_AcceptsV4(t *testing.T) {
	if !run(uuids.RuleIsNotNilV4(), validV4_1) {
		t.Error("non-nil v4 UUID should pass IsNotNilV4")
	}
}

func TestIsNotNilV4_RejectsNil(t *testing.T) {
	if run(uuids.RuleIsNotNilV4(), nilUUID) {
		t.Error("nil UUID should fail IsNotNilV4")
	}
}

func TestIsNotNilV4_RejectsOtherVersions(t *testing.T) {
	if run(uuids.RuleIsNotNilV4(), validV1) {
		t.Error("v1 UUID should fail IsNotNilV4")
	}
	if run(uuids.RuleIsNotNilV4(), validV7) {
		t.Error("v7 UUID should fail IsNotNilV4")
	}
}

func TestIsNotNilV4_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleIsNotNilV4(), "not-a-uuid") {
		t.Error("non-UUID should fail IsNotNilV4")
	}
}

// -----------------------------------------------------------------------------
// RuleOneOfV4
// -----------------------------------------------------------------------------

func TestOneOfV4_AcceptsMember(t *testing.T) {
	if !run(uuids.RuleOneOfV4(validV4_1, validV4_2), validV4_1) {
		t.Error("member of allowed list should pass")
	}
	if !run(uuids.RuleOneOfV4(validV4_1, validV4_2), validV4_2) {
		t.Error("member of allowed list should pass")
	}
}

func TestOneOfV4_CanonicalComparison(t *testing.T) {
	// uppercase input should match lowercase entry in allowed list
	if !run(uuids.RuleOneOfV4(validV4_1), "550E8400-E29B-41D4-A716-446655440000") {
		t.Error("comparison should be case-insensitive via uuid.Parse")
	}
}

func TestOneOfV4_RejectsNonMember(t *testing.T) {
	if run(uuids.RuleOneOfV4(validV4_1, validV4_2), validV4_3) {
		t.Error("non-member should fail")
	}
}

func TestOneOfV4_RejectsWrongVersion(t *testing.T) {
	if run(uuids.RuleOneOfV4(validV4_1), validV1) {
		t.Error("v1 input should fail OneOfV4 even before list lookup")
	}
}

func TestOneOfV4_SilentlySkipsInvalidAllowed(t *testing.T) {
	// "garbage" is dropped at construction; "validV4_1" still matches itself.
	if !run(uuids.RuleOneOfV4("garbage", validV4_1), validV4_1) {
		t.Error("invalid allowed entries should be silently skipped, not cause failure")
	}
	// A non-V4 in the allowed list cannot match a V4 input.
	if run(uuids.RuleOneOfV4(validV1), validV4_1) {
		t.Error("v1 in allowed list should not match a v4 input")
	}
}

// -----------------------------------------------------------------------------
// RuleNotOneOfV4
// -----------------------------------------------------------------------------

func TestNotOneOfV4_AcceptsNonMember(t *testing.T) {
	if !run(uuids.RuleNotOneOfV4(validV4_1, validV4_2), validV4_3) {
		t.Error("non-member v4 should pass NotOneOfV4")
	}
}

func TestNotOneOfV4_RejectsMember(t *testing.T) {
	if run(uuids.RuleNotOneOfV4(validV4_1, validV4_2), validV4_1) {
		t.Error("member should fail NotOneOfV4")
	}
}

func TestNotOneOfV4_RejectsWrongVersion(t *testing.T) {
	if run(uuids.RuleNotOneOfV4(validV4_1), validV1) {
		t.Error("v1 input should fail NotOneOfV4 (rule requires v4)")
	}
}

func TestNotOneOfV4_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleNotOneOfV4(validV4_1), "not-a-uuid") {
		t.Error("malformed input should fail NotOneOfV4")
	}
}
