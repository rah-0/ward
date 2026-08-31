// Package uuids provides UUID validation rules for ward.
// The underlying type is string. UUIDs are parsed and inspected using
// the standard library uuid package; no regex is involved.
//
// Rules in this package are version-tagged. Each version-specific rule
// fully asserts the relevant UUID version, so callers can chain them
// without relying on a separate "is valid" precursor — though composing
// with RuleIsValidV4 is encouraged for explicit intent.
package uuids

import (
	"uuid"

	"github.com/rah-0/ward"
)

const (
	IDIsValidV4         uint32 = 2
	IDIsNotNilV4        uint32 = 3
	IDOneOfV4           uint32 = 4
	IDNotOneOfV4        uint32 = 5
	IDOneOf             uint32 = 6
	IDNotOneOf          uint32 = 7
	IDIsValid           uint32 = 8
	IDIsNil             uint32 = 9
	IDIsNotNil          uint32 = 10
	IDIsVersion         uint32 = 11
	IDIsValidV1         uint32 = 12
	IDIsValidV3         uint32 = 13
	IDIsValidV5         uint32 = 14
	IDIsValidV6         uint32 = 15
	IDIsValidV7         uint32 = 16
	IDIsCanonicalFormat uint32 = 17
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDIsValidV4:         "IsValidV4",
	IDIsNotNilV4:        "IsNotNilV4",
	IDOneOfV4:           "OneOfV4",
	IDNotOneOfV4:        "NotOneOfV4",
	IDOneOf:             "OneOf",
	IDNotOneOf:          "NotOneOf",
	IDIsValid:           "IsValid",
	IDIsNil:             "IsNil",
	IDIsNotNil:          "IsNotNil",
	IDIsVersion:         "IsVersion",
	IDIsValidV1:         "IsValidV1",
	IDIsValidV3:         "IsValidV3",
	IDIsValidV5:         "IsValidV5",
	IDIsValidV6:         "IsValidV6",
	IDIsValidV7:         "IsValidV7",
	IDIsCanonicalFormat: "IsCanonicalFormat",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

// -----------------------------------------------------------------------------
// UUIDv4
// -----------------------------------------------------------------------------

// RuleIsValidV4 passes when *s parses as a UUID and its version is 4.
// The nil UUID has version 0 and fails this rule.
//
// Parsing is delegated to uuid.Parse, which accepts the canonical hyphenated
// form as well as raw hex (no hyphens), the urn:uuid: prefix, and Microsoft-
// style braces. Callers that require strictly canonical input must add their
// own format check.
func RuleIsValidV4() Rule {
	return Rule{ID: IDIsValidV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 4 {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsNotNilV4 passes when *s parses as a UUIDv4 and is not the nil UUID.
// Because the nil UUID is version 0, requiring V4 already excludes it; this
// rule is provided for explicit composition where the intent is to spell
// out "must not be nil" alongside other V4 checks.
func RuleIsNotNilV4() Rule {
	return Rule{ID: IDIsNotNilV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if id == uuid.Nil() {
			return &Result{}
		}
		if uuidVersion(id) != 4 {
			return &Result{}
		}
		return nil
	}}
}

// RuleOneOfV4 passes when *s parses as a UUIDv4 and equals one of the
// allowed values. Allowed UUIDs are parsed once at rule construction;
// any entry that is not a valid UUID is silently skipped, so it can
// never match an input. Comparison is canonical (case-insensitive,
// formatting-tolerant) thanks to uuid.Parse.
func RuleOneOfV4(allowed ...string) Rule {
	parsed := parseUUIDsV4(allowed)
	return Rule{ID: IDOneOfV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err, Arg1: allowed}
		}
		if uuidVersion(id) != 4 {
			return &Result{Arg1: allowed}
		}
		for _, a := range parsed {
			if id == a {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOfV4 passes when *s parses as a UUIDv4 and is not equal to
// any of the excluded values. Excluded UUIDs are parsed once at rule
// construction; invalid entries are silently dropped.
func RuleNotOneOfV4(excluded ...string) Rule {
	parsed := parseUUIDsV4(excluded)
	return Rule{ID: IDNotOneOfV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 4 {
			return &Result{}
		}
		for _, e := range parsed {
			if id == e {
				return &Result{Arg1: excluded}
			}
		}
		return nil
	}}
}

// parseUUIDsV4 parses a slice of UUID strings, keeping only those that
// successfully parse and are version 4. Used at rule construction time
// so the per-call hot path only parses the input once.
func parseUUIDsV4(ss []string) []uuid.UUID {
	out := make([]uuid.UUID, 0, len(ss))
	for _, s := range ss {
		id, err := uuid.Parse(s)
		if err != nil || uuidVersion(id) != 4 {
			continue
		}
		out = append(out, id)
	}
	return out
}

// -----------------------------------------------------------------------------
// Version-agnostic
// -----------------------------------------------------------------------------

// RuleOneOf passes when *s parses as any UUID and equals one of the allowed
// values. Unlike RuleOneOfV4, no version check is performed — v1, v4, v7, and
// any other parseable form are all eligible. Allowed UUIDs are parsed once at
// rule construction; entries that fail to parse are silently skipped.
func RuleOneOf(allowed ...string) Rule {
	parsed := parseUUIDs(allowed)
	return Rule{ID: IDOneOf, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err, Arg1: allowed}
		}
		for _, a := range parsed {
			if id == a {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOf passes when *s parses as any UUID and is not equal to any of
// the excluded values. No version check is performed. Excluded UUIDs are
// parsed once at rule construction; entries that fail to parse are silently
// skipped.
func RuleNotOneOf(excluded ...string) Rule {
	parsed := parseUUIDs(excluded)
	return Rule{ID: IDNotOneOf, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		for _, e := range parsed {
			if id == e {
				return &Result{Arg1: excluded}
			}
		}
		return nil
	}}
}

// parseUUIDs parses a slice of UUID strings, keeping any that successfully
// parse regardless of version.
func parseUUIDs(ss []string) []uuid.UUID {
	out := make([]uuid.UUID, 0, len(ss))
	for _, s := range ss {
		id, err := uuid.Parse(s)
		if err != nil {
			continue
		}
		out = append(out, id)
	}
	return out
}

func uuidVersion(id uuid.UUID) byte {
	return id[6] >> 4
}

// -----------------------------------------------------------------------------
// Version-agnostic identity rules
// -----------------------------------------------------------------------------

// RuleIsValid passes when *s parses as any UUID, regardless of version.
// Parsing is delegated to uuid.Parse, which accepts the canonical hyphenated
// form, raw hex, the urn:uuid: prefix, and Microsoft braces. The nil UUID
// passes this rule — use RuleIsNotNil to exclude it.
func RuleIsValid() Rule {
	return Rule{ID: IDIsValid, Fn: func(s *string) *Result {
		if _, err := uuid.Parse(*s); err != nil {
			return &Result{Err: err}
		}
		return nil
	}}
}

// RuleIsNil passes when *s parses as the nil UUID
// ("00000000-0000-0000-0000-000000000000"). Malformed input fails.
func RuleIsNil() Rule {
	return Rule{ID: IDIsNil, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if id != uuid.Nil() {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsNotNil passes when *s parses as any UUID and is not the nil UUID.
// Version is not checked — use RuleIsNotNilV4 for the v4-specific variant.
func RuleIsNotNil() Rule {
	return Rule{ID: IDIsNotNil, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if id == uuid.Nil() {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsVersion passes when *s parses as a UUID whose version equals n.
// The nil UUID has version 0 and only matches n == 0.
func RuleIsVersion(n int) Rule {
	return Rule{ID: IDIsVersion, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err, Arg1: n}
		}
		if int(uuidVersion(id)) != n {
			return &Result{Arg1: n}
		}
		return nil
	}}
}

// -----------------------------------------------------------------------------
// Version-specific rules (v1, v3, v5, v6, v7)
// -----------------------------------------------------------------------------

// RuleIsValidV1 passes when *s parses as a UUID and its version is 1.
func RuleIsValidV1() Rule {
	return Rule{ID: IDIsValidV1, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 1 {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsValidV3 passes when *s parses as a UUID and its version is 3
// (name-based, MD5).
func RuleIsValidV3() Rule {
	return Rule{ID: IDIsValidV3, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 3 {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsValidV5 passes when *s parses as a UUID and its version is 5
// (name-based, SHA-1).
func RuleIsValidV5() Rule {
	return Rule{ID: IDIsValidV5, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 5 {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsValidV6 passes when *s parses as a UUID and its version is 6
// (reordered time, RFC 9562).
func RuleIsValidV6() Rule {
	return Rule{ID: IDIsValidV6, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 6 {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsValidV7 passes when *s parses as a UUID and its version is 7
// (Unix-time-ordered, RFC 9562). v7 is increasingly preferred over
// v4 for new sortable identifiers.
func RuleIsValidV7() Rule {
	return Rule{ID: IDIsValidV7, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if uuidVersion(id) != 7 {
			return &Result{}
		}
		return nil
	}}
}

// RuleIsCanonicalFormat passes when *s is in the strict canonical hyphenated
// form 8-4-4-4-12 lowercase- or uppercase-hex. Unlike the other rules in this
// package, it rejects the alternate forms uuid.Parse accepts (raw hex,
// urn:uuid: prefix, braces). It performs no version check.
func RuleIsCanonicalFormat() Rule {
	return Rule{ID: IDIsCanonicalFormat, Fn: func(s *string) *Result {
		if !isCanonicalUUID(*s) {
			return &Result{}
		}
		return nil
	}}
}

// isCanonicalUUID reports whether s matches 8-4-4-4-12 hex with hyphens,
// without allocating a regexp. Lowercase and uppercase hex digits are
// accepted; mixed case is allowed.
func isCanonicalUUID(s string) bool {
	if len(s) != 36 {
		return false
	}
	for i := 0; i < 36; i++ {
		c := s[i]
		switch i {
		case 8, 13, 18, 23:
			if c != '-' {
				return false
			}
		default:
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return false
			}
		}
	}
	return true
}
