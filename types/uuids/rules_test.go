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

// TestIsValidV4_AcceptsAlternateForms documents the non-canonical encodings
// accepted by uuid.Parse. These all decode to the same underlying v4 UUID, so
// IsValidV4 accepts them.
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

func TestIsValidV4_RejectsUnsupportedAlternateForms(t *testing.T) {
	for _, v := range []string{
		"URN:UUID:550e8400-e29b-41d4-a716-446655440000", // prefix must be lowercase
		"[550e8400-e29b-41d4-a716-446655440000]",        // only braces are supported
	} {
		if run(uuids.RuleIsValidV4(), v) {
			t.Errorf("%q should fail IsValidV4 (unsupported alternate form)", v)
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

// -----------------------------------------------------------------------------
// RuleOneOf (version-agnostic)
// -----------------------------------------------------------------------------

func TestOneOf_AcceptsMemberAnyVersion(t *testing.T) {
	if !run(uuids.RuleOneOf(validV1, validV4_1, validV7), validV1) {
		t.Error("v1 member should pass OneOf")
	}
	if !run(uuids.RuleOneOf(validV1, validV4_1, validV7), validV4_1) {
		t.Error("v4 member should pass OneOf")
	}
	if !run(uuids.RuleOneOf(validV1, validV4_1, validV7), validV7) {
		t.Error("v7 member should pass OneOf")
	}
}

func TestOneOf_RejectsNonMember(t *testing.T) {
	if run(uuids.RuleOneOf(validV4_1, validV4_2), validV4_3) {
		t.Error("non-member should fail OneOf")
	}
}

func TestOneOf_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleOneOf(validV4_1), "not-a-uuid") {
		t.Error("malformed input should fail OneOf")
	}
}

func TestOneOf_AllowsV1InAllowedList(t *testing.T) {
	// In OneOfV4, a v1 in the allowed list is dropped; in OneOf, it's kept.
	if !run(uuids.RuleOneOf(validV1), validV1) {
		t.Error("v1 should match a v1 entry under OneOf")
	}
}

// -----------------------------------------------------------------------------
// RuleNotOneOf (version-agnostic)
// -----------------------------------------------------------------------------

func TestNotOneOf_AcceptsNonMemberAnyVersion(t *testing.T) {
	if !run(uuids.RuleNotOneOf(validV4_1, validV4_2), validV1) {
		t.Error("v1 non-member should pass NotOneOf")
	}
	if !run(uuids.RuleNotOneOf(validV4_1, validV4_2), validV7) {
		t.Error("v7 non-member should pass NotOneOf")
	}
}

func TestNotOneOf_RejectsMember(t *testing.T) {
	if run(uuids.RuleNotOneOf(validV4_1, validV4_2), validV4_1) {
		t.Error("member should fail NotOneOf")
	}
}

func TestNotOneOf_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleNotOneOf(validV4_1), "not-a-uuid") {
		t.Error("malformed input should fail NotOneOf")
	}
}

// Additional UUIDs used by version-specific tests below.
const (
	// v3: MD5 hash of "www.widgets.com" in the DNS namespace.
	validV3 = "3d813cbb-47fb-32ba-91df-831e1593ac29"

	// v5: SHA-1 hash of "www.widgets.com" in the DNS namespace.
	validV5 = "21f7f8de-8051-5b89-8680-0195ef798b6a"

	// v6: time-ordered, RFC 9562.
	validV6 = "1ec9414c-232a-6b00-b3c8-9e6bdeced846"
)

// -----------------------------------------------------------------------------
// RuleIsValid (version-agnostic)
// -----------------------------------------------------------------------------

func TestIsValid_AcceptsAnyVersion(t *testing.T) {
	for _, v := range []string{validV1, validV3, validV4_1, validV5, validV6, validV7, nilUUID} {
		if !run(uuids.RuleIsValid(), v) {
			t.Errorf("%q should pass IsValid", v)
		}
	}
}

func TestIsValid_RejectsMalformed(t *testing.T) {
	for _, v := range []string{"", "not-a-uuid", "550e8400-e29b-41d4-a716-44665544000"} {
		if run(uuids.RuleIsValid(), v) {
			t.Errorf("%q should fail IsValid", v)
		}
	}
}

// -----------------------------------------------------------------------------
// RuleIsNil / RuleIsNotNil
// -----------------------------------------------------------------------------

func TestIsNil_AcceptsNilUUID(t *testing.T) {
	if !run(uuids.RuleIsNil(), nilUUID) {
		t.Error("nil UUID should pass IsNil")
	}
}

func TestIsNil_RejectsNonNil(t *testing.T) {
	for _, v := range []string{validV1, validV3, validV4_1, validV5, validV6, validV7} {
		if run(uuids.RuleIsNil(), v) {
			t.Errorf("%q should fail IsNil", v)
		}
	}
}

func TestIsNil_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleIsNil(), "not-a-uuid") {
		t.Error("malformed input should fail IsNil")
	}
}

func TestIsNotNil_AcceptsAnyNonNil(t *testing.T) {
	for _, v := range []string{validV1, validV3, validV4_1, validV5, validV6, validV7} {
		if !run(uuids.RuleIsNotNil(), v) {
			t.Errorf("%q should pass IsNotNil", v)
		}
	}
}

func TestIsNotNil_RejectsNil(t *testing.T) {
	if run(uuids.RuleIsNotNil(), nilUUID) {
		t.Error("nil UUID should fail IsNotNil")
	}
}

func TestIsNotNil_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleIsNotNil(), "not-a-uuid") {
		t.Error("malformed input should fail IsNotNil")
	}
}

// -----------------------------------------------------------------------------
// RuleIsVersion
// -----------------------------------------------------------------------------

func TestIsVersion_MatchesExact(t *testing.T) {
	cases := []struct {
		s string
		v int
	}{
		{validV1, 1}, {validV3, 3}, {validV4_1, 4},
		{validV5, 5}, {validV6, 6}, {validV7, 7},
		{nilUUID, 0},
	}
	for _, c := range cases {
		if !run(uuids.RuleIsVersion(c.v), c.s) {
			t.Errorf("%q should pass IsVersion(%d)", c.s, c.v)
		}
	}
}

func TestIsVersion_RejectsMismatch(t *testing.T) {
	if run(uuids.RuleIsVersion(4), validV1) {
		t.Error("v1 should fail IsVersion(4)")
	}
	if run(uuids.RuleIsVersion(7), validV4_1) {
		t.Error("v4 should fail IsVersion(7)")
	}
}

func TestIsVersion_RejectsMalformed(t *testing.T) {
	if run(uuids.RuleIsVersion(4), "not-a-uuid") {
		t.Error("malformed input should fail IsVersion")
	}
}

// -----------------------------------------------------------------------------
// Version-specific rules
// -----------------------------------------------------------------------------

func TestIsValidV1(t *testing.T) {
	if !run(uuids.RuleIsValidV1(), validV1) {
		t.Error("v1 should pass IsValidV1")
	}
	for _, v := range []string{validV3, validV4_1, validV5, validV6, validV7, nilUUID, "not-a-uuid"} {
		if run(uuids.RuleIsValidV1(), v) {
			t.Errorf("%q should fail IsValidV1", v)
		}
	}
}

func TestIsValidV3(t *testing.T) {
	if !run(uuids.RuleIsValidV3(), validV3) {
		t.Error("v3 should pass IsValidV3")
	}
	for _, v := range []string{validV1, validV4_1, validV5, validV6, validV7, nilUUID, "not-a-uuid"} {
		if run(uuids.RuleIsValidV3(), v) {
			t.Errorf("%q should fail IsValidV3", v)
		}
	}
}

func TestIsValidV5(t *testing.T) {
	if !run(uuids.RuleIsValidV5(), validV5) {
		t.Error("v5 should pass IsValidV5")
	}
	for _, v := range []string{validV1, validV3, validV4_1, validV6, validV7, nilUUID, "not-a-uuid"} {
		if run(uuids.RuleIsValidV5(), v) {
			t.Errorf("%q should fail IsValidV5", v)
		}
	}
}

func TestIsValidV6(t *testing.T) {
	if !run(uuids.RuleIsValidV6(), validV6) {
		t.Error("v6 should pass IsValidV6")
	}
	for _, v := range []string{validV1, validV3, validV4_1, validV5, validV7, nilUUID, "not-a-uuid"} {
		if run(uuids.RuleIsValidV6(), v) {
			t.Errorf("%q should fail IsValidV6", v)
		}
	}
}

func TestIsValidV7(t *testing.T) {
	if !run(uuids.RuleIsValidV7(), validV7) {
		t.Error("v7 should pass IsValidV7")
	}
	for _, v := range []string{validV1, validV3, validV4_1, validV5, validV6, nilUUID, "not-a-uuid"} {
		if run(uuids.RuleIsValidV7(), v) {
			t.Errorf("%q should fail IsValidV7", v)
		}
	}
}

// -----------------------------------------------------------------------------
// RuleIsCanonicalFormat
// -----------------------------------------------------------------------------

func TestIsCanonicalFormat_AcceptsCanonical(t *testing.T) {
	for _, v := range []string{
		validV1, validV3, validV4_1, validV5, validV6, validV7, nilUUID,
		"550E8400-E29B-41D4-A716-446655440000", // uppercase
		"550e8400-E29B-41d4-A716-446655440000", // mixed case
	} {
		if !run(uuids.RuleIsCanonicalFormat(), v) {
			t.Errorf("%q should pass IsCanonicalFormat", v)
		}
	}
}

func TestIsCanonicalFormat_RejectsAlternateForms(t *testing.T) {
	// These all parse via uuid.Parse but are not in canonical form.
	for _, v := range []string{
		"550e8400e29b41d4a716446655440000",              // raw hex, no hyphens
		"urn:uuid:550e8400-e29b-41d4-a716-446655440000", // URN form
		"{550e8400-e29b-41d4-a716-446655440000}",        // Microsoft braces
	} {
		if run(uuids.RuleIsCanonicalFormat(), v) {
			t.Errorf("%q should fail IsCanonicalFormat (alternate form)", v)
		}
	}
}

func TestIsCanonicalFormat_RejectsMalformed(t *testing.T) {
	for _, v := range []string{
		"",
		"not-a-uuid",
		"550e8400-e29b-41d4-a716-44665544000",   // too short
		"550e8400-e29b-41d4-a716-4466554400000", // too long
		"550e8400_e29b-41d4-a716-446655440000",  // wrong separator
		"zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz",  // non-hex
	} {
		if run(uuids.RuleIsCanonicalFormat(), v) {
			t.Errorf("%q should fail IsCanonicalFormat", v)
		}
	}
}
