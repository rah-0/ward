package strs_test

import (
	"regexp"
	"testing"

	"github.com/rah-0/ward/types/strs"
)

func run(rule strs.Rule, value string) bool {
	return rule.Fn(&value) == nil
}

func TestNotEmpty(t *testing.T) {
	if !run(strs.RuleNotEmpty(), "a") {
		t.Error("non-empty string should pass")
	}
	if run(strs.RuleNotEmpty(), "") {
		t.Error("empty string should fail")
	}
}

func TestLengthMin(t *testing.T) {
	if !run(strs.RuleLengthMin(3), "abc") {
		t.Error("exact min length should pass")
	}
	if run(strs.RuleLengthMin(3), "ab") {
		t.Error("below min length should fail")
	}
}

func TestLengthMax(t *testing.T) {
	if !run(strs.RuleLengthMax(3), "abc") {
		t.Error("exact max length should pass")
	}
	if run(strs.RuleLengthMax(3), "abcd") {
		t.Error("above max length should fail")
	}
}

func TestLengthExact(t *testing.T) {
	if !run(strs.RuleLengthExact(3), "abc") {
		t.Error("exact length should pass")
	}
	if run(strs.RuleLengthExact(3), "ab") {
		t.Error("wrong length should fail")
	}
}

func TestLengthBetween(t *testing.T) {
	if !run(strs.RuleLengthBetween(2, 4), "abc") {
		t.Error("within range should pass")
	}
	if run(strs.RuleLengthBetween(2, 4), "a") {
		t.Error("below range should fail")
	}
	if run(strs.RuleLengthBetween(2, 4), "abcde") {
		t.Error("above range should fail")
	}
}

func TestContains(t *testing.T) {
	if !run(strs.RuleContains("foo"), "foobar") {
		t.Error("should pass when contains substring")
	}
	if run(strs.RuleContains("foo"), "bar") {
		t.Error("should fail when missing substring")
	}
}

func TestNotContains(t *testing.T) {
	if !run(strs.RuleNotContains("foo"), "bar") {
		t.Error("should pass when substring absent")
	}
	if run(strs.RuleNotContains("foo"), "foobar") {
		t.Error("should fail when substring present")
	}
}

func TestMatchesRegex(t *testing.T) {
	re := regexp.MustCompile(`^\d+$`)
	if !run(strs.RuleMatchesRegex(re), "123") {
		t.Error("digits-only should pass")
	}
	if run(strs.RuleMatchesRegex(re), "abc") {
		t.Error("non-digits should fail")
	}
}

func TestIsEmail(t *testing.T) {
	if !run(strs.RuleIsEmail(), "user@example.com") {
		t.Error("valid email should pass")
	}
	if run(strs.RuleIsEmail(), "not-an-email") {
		t.Error("invalid email should fail")
	}
}

func TestIsSha512(t *testing.T) {
	valid := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
	if !run(strs.RuleIsSha512(), valid) {
		t.Error("valid 128-char sha512 should pass")
	}
	if run(strs.RuleIsSha512(), valid[:127]) {
		t.Error("127-char hex should fail")
	}
	if run(strs.RuleIsSha512(), "notahash") {
		t.Error("non-hex should fail")
	}
}

func TestHasLowercase(t *testing.T) {
	if !run(strs.RuleHasLowercase(), "Hello") {
		t.Error("string with lowercase should pass")
	}
	if run(strs.RuleHasLowercase(), "HELLO") {
		t.Error("all uppercase should fail")
	}
}

func TestHasUppercase(t *testing.T) {
	if !run(strs.RuleHasUppercase(), "Hello") {
		t.Error("string with uppercase should pass")
	}
	if run(strs.RuleHasUppercase(), "hello") {
		t.Error("all lowercase should fail")
	}
}

func TestIsDigitsOnly(t *testing.T) {
	if !run(strs.RuleIsDigitsOnly(), "12345") {
		t.Error("digits-only should pass")
	}
	if run(strs.RuleIsDigitsOnly(), "123a5") {
		t.Error("mixed should fail")
	}
}

func TestIsURL(t *testing.T) {
	if !run(strs.RuleIsURL(), "https://example.com") {
		t.Error("valid URL should pass")
	}
	if run(strs.RuleIsURL(), "not-a-url") {
		t.Error("invalid URL should fail")
	}
}

func TestIsNotURL(t *testing.T) {
	if !run(strs.RuleIsNotURL(), "not-a-url") {
		t.Error("non-URL should pass")
	}
	if run(strs.RuleIsNotURL(), "https://example.com") {
		t.Error("valid URL should fail")
	}
}

func TestHasDigit(t *testing.T) {
	if !run(strs.RuleHasDigit(), "abc1") {
		t.Error("string with digit should pass")
	}
	if run(strs.RuleHasDigit(), "abc") {
		t.Error("string without digit should fail")
	}
}

func TestHasSpecialChar(t *testing.T) {
	if !run(strs.RuleHasSpecialChar(), "abc!") {
		t.Error("string with special char should pass")
	}
	if run(strs.RuleHasSpecialChar(), "abc") {
		t.Error("string without special char should fail")
	}
}

func TestIsPasswordChars(t *testing.T) {
	if !run(strs.RuleIsPasswordChars(), "Abcd1!") {
		t.Error("password with all char types should pass")
	}
	if run(strs.RuleIsPasswordChars(), "alllower") {
		t.Error("no uppercase/digit/special should fail")
	}
}

func TestIsUsernameChars(t *testing.T) {
	if !run(strs.RuleIsUsernameChars(), "john_doe") {
		t.Error("valid username chars should pass")
	}
	if run(strs.RuleIsUsernameChars(), "john doe") {
		t.Error("space in username should fail")
	}
}

func TestIsBoolString(t *testing.T) {
	for _, v := range []string{"true", "false", "1", "0", "TRUE", "  True  "} {
		if !run(strs.RuleIsBoolString(), v) {
			t.Errorf("%q should pass", v)
		}
	}
	if run(strs.RuleIsBoolString(), "yes") {
		t.Error(`"yes" should fail`)
	}
}

func TestIsNonNegativeInt(t *testing.T) {
	if !run(strs.RuleIsNonNegativeInt(), "0") {
		t.Error("0 should pass")
	}
	if !run(strs.RuleIsNonNegativeInt(), "42") {
		t.Error("42 should pass")
	}
	if run(strs.RuleIsNonNegativeInt(), "-1") {
		t.Error("negative should fail")
	}
	if run(strs.RuleIsNonNegativeInt(), "abc") {
		t.Error("non-numeric should fail")
	}
}

func TestTrim(t *testing.T) {
	s := "  hello  "
	strs.RuleTrim().Fn(&s)
	if s != "hello" {
		t.Errorf("expected %q, got %q", "hello", s)
	}
}

func TestEscapeHTML(t *testing.T) {
	s := "<script>"
	strs.RuleEscapeHTML().Fn(&s)
	if s != "&lt;script&gt;" {
		t.Errorf("expected escaped HTML, got %q", s)
	}
}

func TestUnescapeURL(t *testing.T) {
	s := "hello%20world"
	result := strs.RuleUnescapeURL().Fn(&s)
	if result != nil {
		t.Errorf("expected nil result, got error: %v", result.Err)
	}
	if s != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", s)
	}
}

func TestUnescapeURL_Invalid(t *testing.T) {
	s := "%zz"
	result := strs.RuleUnescapeURL().Fn(&s)
	if result == nil {
		t.Error("expected non-nil result for invalid URL encoding")
	}
	if s != "" {
		t.Errorf("expected empty string on error, got %q", s)
	}
}

func TestNormalizeEmail(t *testing.T) {
	// plain address — passes, no change
	s := "user@example.com"
	if strs.RuleNormalizeEmail().Fn(&s) != nil {
		t.Error("plain address should pass")
	}
	if s != "user@example.com" {
		t.Errorf("plain address should be unchanged, got %q", s)
	}

	// display name stripped
	s = "John Doe <john@example.com>"
	if strs.RuleNormalizeEmail().Fn(&s) != nil {
		t.Error("display name format should pass")
	}
	if s != "john@example.com" {
		t.Errorf("expected display name stripped, got %q", s)
	}
}

func TestOneOf(t *testing.T) {
	if !run(strs.RuleOneOf("active", "inactive"), "active") {
		t.Error("active should pass")
	}
	if !run(strs.RuleOneOf("active", "inactive"), "inactive") {
		t.Error("inactive should pass")
	}
	if run(strs.RuleOneOf("active", "inactive"), "deleted") {
		t.Error("deleted should fail")
	}
	if run(strs.RuleOneOf("active", "inactive"), "") {
		t.Error("empty should fail")
	}
}

func TestNotOneOf(t *testing.T) {
	if !run(strs.RuleNotOneOf("admin", "root"), "user") {
		t.Error("user is not excluded, should pass")
	}
	if run(strs.RuleNotOneOf("admin", "root"), "admin") {
		t.Error("admin is excluded, should fail")
	}
	if run(strs.RuleNotOneOf("admin", "root"), "root") {
		t.Error("root is excluded, should fail")
	}
}
