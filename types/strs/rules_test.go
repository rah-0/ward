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

func TestStartsWith(t *testing.T) {
	if !run(strs.RuleStartsWith("foo"), "foobar") {
		t.Error("foobar starts with foo, should pass")
	}
	if !run(strs.RuleStartsWith(""), "anything") {
		t.Error("empty prefix should always pass")
	}
	if run(strs.RuleStartsWith("foo"), "barfoo") {
		t.Error("barfoo does not start with foo, should fail")
	}
	if run(strs.RuleStartsWith("foo"), "") {
		t.Error("empty string does not start with foo, should fail")
	}
}

func TestEndsWith(t *testing.T) {
	if !run(strs.RuleEndsWith("bar"), "foobar") {
		t.Error("foobar ends with bar, should pass")
	}
	if !run(strs.RuleEndsWith(""), "anything") {
		t.Error("empty suffix should always pass")
	}
	if run(strs.RuleEndsWith("bar"), "barfoo") {
		t.Error("barfoo does not end with bar, should fail")
	}
	if run(strs.RuleEndsWith("bar"), "") {
		t.Error("empty string does not end with bar, should fail")
	}
}

func TestIsIP(t *testing.T) {
	if !run(strs.RuleIsIP(), "192.168.1.1") {
		t.Error("IPv4 should pass IsIP")
	}
	if !run(strs.RuleIsIP(), "::1") {
		t.Error("IPv6 loopback should pass IsIP")
	}
	if !run(strs.RuleIsIP(), "2001:db8::1") {
		t.Error("IPv6 should pass IsIP")
	}
	if run(strs.RuleIsIP(), "not-an-ip") {
		t.Error("non-IP should fail")
	}
	if run(strs.RuleIsIP(), "999.999.999.999") {
		t.Error("invalid octets should fail")
	}
	if run(strs.RuleIsIP(), "") {
		t.Error("empty should fail")
	}
}

func TestIsIPv4(t *testing.T) {
	if !run(strs.RuleIsIPv4(), "192.168.1.1") {
		t.Error("dotted-decimal IPv4 should pass")
	}
	if !run(strs.RuleIsIPv4(), "0.0.0.0") {
		t.Error("0.0.0.0 should pass")
	}
	if run(strs.RuleIsIPv4(), "::1") {
		t.Error("IPv6 should fail IsIPv4")
	}
	if run(strs.RuleIsIPv4(), "::ffff:1.2.3.4") {
		t.Error("IPv4-mapped IPv6 should fail IsIPv4 (it contains ':')")
	}
	if run(strs.RuleIsIPv4(), "999.999.999.999") {
		t.Error("invalid octets should fail")
	}
}

func TestIsIPv6(t *testing.T) {
	if !run(strs.RuleIsIPv6(), "::1") {
		t.Error("IPv6 loopback should pass")
	}
	if !run(strs.RuleIsIPv6(), "2001:db8::1") {
		t.Error("IPv6 should pass")
	}
	if !run(strs.RuleIsIPv6(), "::ffff:1.2.3.4") {
		t.Error("IPv4-mapped IPv6 should pass IsIPv6")
	}
	if run(strs.RuleIsIPv6(), "192.168.1.1") {
		t.Error("plain IPv4 should fail IsIPv6")
	}
	if run(strs.RuleIsIPv6(), "not-an-ip") {
		t.Error("non-IP should fail")
	}
}

func TestIsAlpha(t *testing.T) {
	if !run(strs.RuleIsAlpha(), "Hello") {
		t.Error("letters only should pass")
	}
	if run(strs.RuleIsAlpha(), "Hello123") {
		t.Error("mixed letters and digits should fail")
	}
	if run(strs.RuleIsAlpha(), "Hello!") {
		t.Error("letters with punctuation should fail")
	}
	if run(strs.RuleIsAlpha(), "") {
		t.Error("empty string should fail")
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	if !run(strs.RuleIsAlphaNumeric(), "Hello123") {
		t.Error("letters and digits should pass")
	}
	if !run(strs.RuleIsAlphaNumeric(), "abc") {
		t.Error("letters only should pass")
	}
	if !run(strs.RuleIsAlphaNumeric(), "123") {
		t.Error("digits only should pass")
	}
	if run(strs.RuleIsAlphaNumeric(), "hello world") {
		t.Error("space should fail")
	}
	if run(strs.RuleIsAlphaNumeric(), "abc!") {
		t.Error("punctuation should fail")
	}
	if run(strs.RuleIsAlphaNumeric(), "") {
		t.Error("empty string should fail")
	}
}

func TestIsASCII(t *testing.T) {
	if !run(strs.RuleIsASCII(), "Hello, world!") {
		t.Error("plain ASCII should pass")
	}
	if !run(strs.RuleIsASCII(), "") {
		t.Error("empty string should pass")
	}
	if !run(strs.RuleIsASCII(), "\x00\x7f") {
		t.Error("0x00 and 0x7f boundary should pass")
	}
	if run(strs.RuleIsASCII(), "héllo") {
		t.Error("non-ASCII letter should fail")
	}
	if run(strs.RuleIsASCII(), "日本語") {
		t.Error("multi-byte chars should fail")
	}
}

func TestIsBase64(t *testing.T) {
	if !run(strs.RuleIsBase64(), "aGVsbG8=") {
		t.Error("valid base64 should pass")
	}
	if !run(strs.RuleIsBase64(), "") {
		t.Error("empty string is valid base64 (decodes to empty)")
	}
	if run(strs.RuleIsBase64(), "aGVsbG8") {
		t.Error("base64 missing padding should fail std encoding")
	}
	if run(strs.RuleIsBase64(), "not base64!") {
		t.Error("invalid base64 should fail")
	}
}

func TestIsBase64URL(t *testing.T) {
	if !run(strs.RuleIsBase64URL(), "aGVsbG8=") {
		t.Error("valid URL-safe base64 should pass")
	}
	if !run(strs.RuleIsBase64URL(), "a-_b") {
		// 4 chars, '-' and '_' are URL-safe variants of '+' and '/'
		// "a-_b" -> 3 bytes, no padding needed (4 chars = 3 bytes)
		t.Error("URL-safe characters should pass")
	}
	if run(strs.RuleIsBase64URL(), "a+/b=") {
		t.Error("std base64 chars + and / should fail URL encoding")
	}
}

func TestIsJSON(t *testing.T) {
	for _, v := range []string{
		`{"a":1}`,
		`[1,2,3]`,
		`"hello"`,
		`123`,
		`true`,
		`false`,
		`null`,
	} {
		if !run(strs.RuleIsJSON(), v) {
			t.Errorf("%q should be valid JSON", v)
		}
	}
	for _, v := range []string{
		``,
		`{`,
		`{"a":}`,
		`undefined`,
		`{a:1}`,
	} {
		if run(strs.RuleIsJSON(), v) {
			t.Errorf("%q should not be valid JSON", v)
		}
	}
}

func TestIsLowercase(t *testing.T) {
	if !run(strs.RuleIsLowercase(), "hello") {
		t.Error("all lowercase should pass")
	}
	if !run(strs.RuleIsLowercase(), "hello world 123!") {
		t.Error("lowercase with non-letter chars should pass")
	}
	if !run(strs.RuleIsLowercase(), "") {
		t.Error("empty string should pass")
	}
	if run(strs.RuleIsLowercase(), "Hello") {
		t.Error("mixed case should fail")
	}
	if run(strs.RuleIsLowercase(), "HELLO") {
		t.Error("uppercase should fail")
	}
}

func TestIsUppercase(t *testing.T) {
	if !run(strs.RuleIsUppercase(), "HELLO") {
		t.Error("all uppercase should pass")
	}
	if !run(strs.RuleIsUppercase(), "HELLO WORLD 123!") {
		t.Error("uppercase with non-letter chars should pass")
	}
	if !run(strs.RuleIsUppercase(), "") {
		t.Error("empty string should pass")
	}
	if run(strs.RuleIsUppercase(), "Hello") {
		t.Error("mixed case should fail")
	}
	if run(strs.RuleIsUppercase(), "hello") {
		t.Error("lowercase should fail")
	}
}

func TestToLower(t *testing.T) {
	s := "Hello World"
	strs.RuleToLower().Fn(&s)
	if s != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", s)
	}
}

func TestToUpper(t *testing.T) {
	s := "Hello World"
	strs.RuleToUpper().Fn(&s)
	if s != "HELLO WORLD" {
		t.Errorf("expected %q, got %q", "HELLO WORLD", s)
	}
}

func TestStripHTMLTags(t *testing.T) {
	s := "<p>hello <b>world</b></p>"
	strs.RuleStripHTMLTags().Fn(&s)
	if s != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", s)
	}

	s = "no tags here"
	strs.RuleStripHTMLTags().Fn(&s)
	if s != "no tags here" {
		t.Errorf("plain text should be unchanged, got %q", s)
	}

	s = "<script>alert(1)</script>"
	strs.RuleStripHTMLTags().Fn(&s)
	if s != "alert(1)" {
		t.Errorf("expected %q, got %q", "alert(1)", s)
	}
}
