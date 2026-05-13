package uuids_test

import (
	"testing"

	"github.com/rah-0/ward/types/uuids"
)

func run(rule uuids.Rule, value string) bool {
	return rule.Fn(&value) == nil
}

const (
	validUUID1 = "550e8400-e29b-41d4-a716-446655440000"
	validUUID2 = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	validUUID3 = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
)

func TestIsValid(t *testing.T) {
	if !run(uuids.RuleIsValid(), validUUID1) {
		t.Error("valid UUID should pass")
	}
	if !run(uuids.RuleIsValid(), "6BA7B810-9DAD-11D1-80B4-00C04FD430C8") {
		t.Error("uppercase UUID should pass")
	}
	if run(uuids.RuleIsValid(), "not-a-uuid") {
		t.Error("non-UUID string should fail")
	}
	if run(uuids.RuleIsValid(), "") {
		t.Error("empty string should fail")
	}
	if run(uuids.RuleIsValid(), "550e8400-e29b-41d4-a716-44665544000") {
		t.Error("UUID with wrong length should fail")
	}
	if run(uuids.RuleIsValid(), "550e8400e29b41d4a716446655440000") {
		t.Error("UUID without hyphens should fail")
	}
}

func TestOneOf(t *testing.T) {
	if !run(uuids.RuleOneOf(validUUID1, validUUID2), validUUID1) {
		t.Error("validUUID1 is in list, should pass")
	}
	if !run(uuids.RuleOneOf(validUUID1, validUUID2), validUUID2) {
		t.Error("validUUID2 is in list, should pass")
	}
	if run(uuids.RuleOneOf(validUUID1, validUUID2), validUUID3) {
		t.Error("validUUID3 is not in list, should fail")
	}
	if run(uuids.RuleOneOf(validUUID1, validUUID2), "") {
		t.Error("empty string is not in list, should fail")
	}
}

func TestNotOneOf(t *testing.T) {
	if !run(uuids.RuleNotOneOf(validUUID1, validUUID2), validUUID3) {
		t.Error("validUUID3 is not excluded, should pass")
	}
	if run(uuids.RuleNotOneOf(validUUID1, validUUID2), validUUID1) {
		t.Error("validUUID1 is excluded, should fail")
	}
	if run(uuids.RuleNotOneOf(validUUID1, validUUID2), validUUID2) {
		t.Error("validUUID2 is excluded, should fail")
	}
}
