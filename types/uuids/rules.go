// Package uuids provides UUID validation rules for ward.
// The underlying type is string. UUIDs are parsed and inspected using
// github.com/google/uuid; no regex is involved.
//
// Rules in this package are version-tagged. Each version-specific rule
// fully asserts the relevant UUID version, so callers can chain them
// without relying on a separate "is valid" precursor — though composing
// with RuleIsValidV4 is encouraged for explicit intent.
package uuids

import (
	"github.com/google/uuid"

	"github.com/rah-0/ward"
)

const (
	IDIsValidV4  uint32 = 2
	IDIsNotNilV4 uint32 = 3
	IDOneOfV4    uint32 = 4
	IDNotOneOfV4 uint32 = 5
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDIsValidV4:  "IsValidV4",
	IDIsNotNilV4: "IsNotNilV4",
	IDOneOfV4:    "OneOfV4",
	IDNotOneOfV4: "NotOneOfV4",
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
	return Rule{TypeID: TypeID, ID: IDIsValidV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if id.Version() != 4 {
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
	return Rule{TypeID: TypeID, ID: IDIsNotNilV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if id == uuid.Nil {
			return &Result{}
		}
		if id.Version() != 4 {
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
	return Rule{TypeID: TypeID, ID: IDOneOfV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err, Arg1: allowed}
		}
		if id.Version() != 4 {
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
	return Rule{TypeID: TypeID, ID: IDNotOneOfV4, Fn: func(s *string) *Result {
		id, err := uuid.Parse(*s)
		if err != nil {
			return &Result{Err: err}
		}
		if id.Version() != 4 {
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
		if err != nil || id.Version() != 4 {
			continue
		}
		out = append(out, id)
	}
	return out
}
