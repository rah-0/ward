// Package strs provides string validation and sanitization rules for ward.
package strs

import (
	"encoding/base64"
	"encoding/json"
	"html"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/rah-0/ward"
)

const (
	IDNotEmpty         uint32 = 2
	IDLengthMin        uint32 = 3
	IDLengthMax        uint32 = 4
	IDLengthExact      uint32 = 5
	IDLengthBetween    uint32 = 6
	IDContains         uint32 = 7
	IDNotContains      uint32 = 8
	IDMatchesRegex     uint32 = 9
	IDIsEmail          uint32 = 10
	IDIsSha512         uint32 = 11
	IDHasLowercase     uint32 = 12
	IDHasUppercase     uint32 = 13
	IDIsDigitsOnly     uint32 = 14
	IDIsURL            uint32 = 15
	IDIsNotURL         uint32 = 16
	IDHasDigit         uint32 = 17
	IDHasSpecialChar   uint32 = 18
	IDIsPasswordChars  uint32 = 19
	IDIsUsernameChars  uint32 = 20
	IDIsBoolString     uint32 = 21
	IDIsNonNegativeInt uint32 = 22
	IDTrim             uint32 = 23
	IDEscapeHTML       uint32 = 24
	IDUnescapeURL      uint32 = 25
	IDNormalizeEmail   uint32 = 26
	IDOneOf            uint32 = 27
	IDNotOneOf         uint32 = 28
	IDStartsWith       uint32 = 29
	IDEndsWith         uint32 = 30
	IDIsIP             uint32 = 31
	IDIsIPv4           uint32 = 32
	IDIsIPv6           uint32 = 33
	IDIsAlpha          uint32 = 34
	IDIsAlphaNumeric   uint32 = 35
	IDIsASCII          uint32 = 36
	IDIsBase64         uint32 = 37
	IDIsBase64URL      uint32 = 38
	IDIsJSON           uint32 = 39
	IDIsLowercase      uint32 = 40
	IDIsUppercase      uint32 = 41
	IDToLower          uint32 = 42
	IDToUpper          uint32 = 43
	IDStripHTMLTags    uint32 = 44
	IDIsUUID           uint32 = 45
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDNotEmpty:         "NotEmpty",
	IDLengthMin:        "LengthMin",
	IDLengthMax:        "LengthMax",
	IDLengthExact:      "LengthExact",
	IDLengthBetween:    "LengthBetween",
	IDContains:         "Contains",
	IDNotContains:      "NotContains",
	IDMatchesRegex:     "MatchesRegex",
	IDIsEmail:          "IsEmail",
	IDIsSha512:         "IsSha512",
	IDHasLowercase:     "HasLowercase",
	IDHasUppercase:     "HasUppercase",
	IDIsDigitsOnly:     "IsDigitsOnly",
	IDIsURL:            "IsURL",
	IDIsNotURL:         "IsNotURL",
	IDHasDigit:         "HasDigit",
	IDHasSpecialChar:   "HasSpecialChar",
	IDIsPasswordChars:  "IsPasswordChars",
	IDIsUsernameChars:  "IsUsernameChars",
	IDIsBoolString:     "IsBoolString",
	IDIsNonNegativeInt: "IsNonNegativeInt",
	IDTrim:             "Trim",
	IDEscapeHTML:       "EscapeHTML",
	IDUnescapeURL:      "UnescapeURL",
	IDNormalizeEmail:   "NormalizeEmail",
	IDOneOf:            "OneOf",
	IDNotOneOf:         "NotOneOf",
	IDStartsWith:       "StartsWith",
	IDEndsWith:         "EndsWith",
	IDIsIP:             "IsIP",
	IDIsIPv4:           "IsIPv4",
	IDIsIPv6:           "IsIPv6",
	IDIsAlpha:          "IsAlpha",
	IDIsAlphaNumeric:   "IsAlphaNumeric",
	IDIsASCII:          "IsASCII",
	IDIsBase64:         "IsBase64",
	IDIsBase64URL:      "IsBase64URL",
	IDIsJSON:           "IsJSON",
	IDIsLowercase:      "IsLowercase",
	IDIsUppercase:      "IsUppercase",
	IDToLower:          "ToLower",
	IDToUpper:          "ToUpper",
	IDStripHTMLTags:    "StripHTMLTags",
	IDIsUUID:           "IsUUID",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

func RuleNotEmpty() Rule {
	return Rule{ID: IDNotEmpty, Fn: func(s *string) *Result {
		if utf8.RuneCountInString(*s) > 0 {
			return nil
		}
		return &Result{}
	}}
}

func RuleLengthMin(min int) Rule {
	return Rule{ID: IDLengthMin, Fn: func(s *string) *Result {
		if utf8.RuneCountInString(*s) >= min {
			return nil
		}
		return &Result{
			Arg1: min,
		}
	}}
}

func RuleLengthMax(max int) Rule {
	return Rule{ID: IDLengthMax, Fn: func(s *string) *Result {
		if utf8.RuneCountInString(*s) <= max {
			return nil
		}
		return &Result{
			Arg1: max,
		}
	}}
}

func RuleLengthExact(length int) Rule {
	return Rule{ID: IDLengthExact, Fn: func(s *string) *Result {
		if utf8.RuneCountInString(*s) == length {
			return nil
		}
		return &Result{
			Arg1: length,
		}
	}}
}

func RuleLengthBetween(min, max int) Rule {
	return Rule{ID: IDLengthBetween, Fn: func(s *string) *Result {
		l := utf8.RuneCountInString(*s)
		if l >= min && l <= max {
			return nil
		}
		return &Result{
			Arg1: min,
			Arg2: max,
		}
	}}
}

func RuleContains(sub string) Rule {
	return Rule{ID: IDContains, Fn: func(s *string) *Result {
		if strings.Contains(*s, sub) {
			return nil
		}
		return &Result{
			Arg1: sub,
		}
	}}
}

func RuleNotContains(sub string) Rule {
	return Rule{ID: IDNotContains, Fn: func(s *string) *Result {
		if !strings.Contains(*s, sub) {
			return nil
		}
		return &Result{
			Arg1: sub,
		}
	}}
}

func RuleMatchesRegex(pattern *regexp.Regexp) Rule {
	return Rule{ID: IDMatchesRegex, Fn: func(s *string) *Result {
		if pattern.MatchString(*s) {
			return nil
		}
		return &Result{
			Arg1: pattern.String(),
		}
	}}
}

func RuleIsEmail() Rule {
	return Rule{ID: IDIsEmail, Fn: func(s *string) *Result {
		_, err := mail.ParseAddress(*s)
		if err == nil {
			return nil
		}
		return &Result{
			Err: err,
		}
	}}
}

func RuleIsSha512() Rule {
	return Rule{ID: IDIsSha512, Fn: func(s *string) *Result {
		if len(*s) == 128 && RegexpSha512.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleHasLowercase() Rule {
	return Rule{ID: IDHasLowercase, Fn: func(s *string) *Result {
		if RegexpHasLowercase.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleHasUppercase() Rule {
	return Rule{ID: IDHasUppercase, Fn: func(s *string) *Result {
		if RegexpHasUppercase.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleIsDigitsOnly() Rule {
	return Rule{ID: IDIsDigitsOnly, Fn: func(s *string) *Result {
		if RegexpDigitsOnly.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleIsURL() Rule {
	return Rule{ID: IDIsURL, Fn: func(s *string) *Result {
		u, err := url.ParseRequestURI(*s)
		if err == nil && u.Host != "" && (u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "ftp" || u.Scheme == "ftps") {
			return nil
		}
		return &Result{
			Err: err,
		}
	}}
}

func RuleIsNotURL() Rule {
	return Rule{ID: IDIsNotURL, Fn: func(s *string) *Result {
		u, err := url.ParseRequestURI(*s)
		if !(err == nil && u.Host != "" && (u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "ftp" || u.Scheme == "ftps")) {
			return nil
		}
		return &Result{
			Err: err,
		}
	}}
}

func RuleHasDigit() Rule {
	return Rule{ID: IDHasDigit, Fn: func(s *string) *Result {
		if RegexpHasDigit.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleHasSpecialChar() Rule {
	return Rule{ID: IDHasSpecialChar, Fn: func(s *string) *Result {
		if RegexpHasSpecialChar.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleIsPasswordChars() Rule {
	return Rule{ID: IDIsPasswordChars, Fn: func(s *string) *Result {
		if RegexpHasLowercase.MatchString(*s) &&
			RegexpHasUppercase.MatchString(*s) &&
			RegexpHasDigit.MatchString(*s) &&
			RegexpHasSpecialChar.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleIsUsernameChars() Rule {
	return Rule{ID: IDIsUsernameChars, Fn: func(s *string) *Result {
		if RegexpUsernameChars.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

func RuleIsBoolString() Rule {
	return Rule{ID: IDIsBoolString, Fn: func(s *string) *Result {
		v := strings.ToLower(strings.TrimSpace(*s))
		if v == "1" || v == "0" || v == "true" || v == "false" {
			return nil
		}
		return &Result{}
	}}
}

func RuleIsNonNegativeInt() Rule {
	return Rule{ID: IDIsNonNegativeInt, Fn: func(s *string) *Result {
		if RegexpNonNegativeInt.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

// -----------------------------------------------------------------------------
// Sanitizers — the following rules mutate *s
// -----------------------------------------------------------------------------

func RuleTrim() Rule {
	return Rule{ID: IDTrim, Fn: func(s *string) *Result {
		*s = strings.TrimSpace(*s)
		return nil
	}}
}

func RuleEscapeHTML() Rule {
	return Rule{ID: IDEscapeHTML, Fn: func(s *string) *Result {
		*s = html.EscapeString(*s)
		return nil
	}}
}

func RuleUnescapeURL() Rule {
	return Rule{ID: IDUnescapeURL, Fn: func(s *string) *Result {
		decoded, err := url.QueryUnescape(*s)
		if err != nil {
			*s = ""
			return &Result{
				Err: err,
			}
		}
		*s = decoded
		return nil
	}}
}

// RuleNormalizeEmail is a sanitizer that parses the address, strips any display
// name (e.g. "John Doe <john@example.com>" → "john@example.com"), and writes
// the canonical address back to *s. Returns a failure if the value cannot be
// parsed as a valid email at all.
func RuleNormalizeEmail() Rule {
	return Rule{ID: IDNormalizeEmail, Fn: func(s *string) *Result {
		addr, err := mail.ParseAddress(*s)
		if err != nil {
			return &Result{Err: err}
		}
		*s = addr.Address
		return nil
	}}
}

// RuleOneOf passes when *s equals one of the allowed values.
func RuleOneOf(allowed ...string) Rule {
	return Rule{ID: IDOneOf, Fn: func(s *string) *Result {
		for _, a := range allowed {
			if *s == a {
				return nil
			}
		}
		return &Result{Arg1: allowed}
	}}
}

// RuleNotOneOf passes when *s does not equal any of the excluded values.
func RuleNotOneOf(excluded ...string) Rule {
	return Rule{ID: IDNotOneOf, Fn: func(s *string) *Result {
		for _, e := range excluded {
			if *s == e {
				return &Result{Arg1: excluded}
			}
		}
		return nil
	}}
}

// RuleStartsWith passes when *s begins with prefix.
func RuleStartsWith(prefix string) Rule {
	return Rule{ID: IDStartsWith, Fn: func(s *string) *Result {
		if strings.HasPrefix(*s, prefix) {
			return nil
		}
		return &Result{Arg1: prefix}
	}}
}

// RuleEndsWith passes when *s ends with suffix.
func RuleEndsWith(suffix string) Rule {
	return Rule{ID: IDEndsWith, Fn: func(s *string) *Result {
		if strings.HasSuffix(*s, suffix) {
			return nil
		}
		return &Result{Arg1: suffix}
	}}
}

// RuleIsIP passes when *s parses as either an IPv4 or IPv6 address.
func RuleIsIP() Rule {
	return Rule{ID: IDIsIP, Fn: func(s *string) *Result {
		if net.ParseIP(*s) != nil {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsIPv4 passes when *s is a dotted-decimal IPv4 address.
// The input form must contain a "." and no ":", so IPv4-mapped IPv6
// addresses (e.g. "::ffff:1.2.3.4") are rejected here and accepted by
// RuleIsIPv6 instead.
func RuleIsIPv4() Rule {
	return Rule{ID: IDIsIPv4, Fn: func(s *string) *Result {
		if net.ParseIP(*s) != nil && strings.Contains(*s, ".") && !strings.Contains(*s, ":") {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsIPv6 passes when *s is an IPv6 address (any form containing ":").
// IPv4-mapped IPv6 addresses such as "::ffff:1.2.3.4" pass here.
func RuleIsIPv6() Rule {
	return Rule{ID: IDIsIPv6, Fn: func(s *string) *Result {
		if net.ParseIP(*s) != nil && strings.Contains(*s, ":") {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsAlpha passes when *s contains only ASCII letters (a-z, A-Z) and is non-empty.
func RuleIsAlpha() Rule {
	return Rule{ID: IDIsAlpha, Fn: func(s *string) *Result {
		if RegexpAlpha.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsAlphaNumeric passes when *s contains only ASCII letters and digits and is non-empty.
func RuleIsAlphaNumeric() Rule {
	return Rule{ID: IDIsAlphaNumeric, Fn: func(s *string) *Result {
		if RegexpAlphaNumeric.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsASCII passes when every byte of *s is in the 7-bit ASCII range (0-127).
// The empty string passes trivially.
func RuleIsASCII() Rule {
	return Rule{ID: IDIsASCII, Fn: func(s *string) *Result {
		for i := 0; i < len(*s); i++ {
			if (*s)[i] > 127 {
				return &Result{}
			}
		}
		return nil
	}}
}

// RuleIsBase64 passes when *s is valid standard base64 (with padding).
func RuleIsBase64() Rule {
	return Rule{ID: IDIsBase64, Fn: func(s *string) *Result {
		if _, err := base64.StdEncoding.DecodeString(*s); err == nil {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsBase64URL passes when *s is valid URL-safe base64 (with padding).
func RuleIsBase64URL() Rule {
	return Rule{ID: IDIsBase64URL, Fn: func(s *string) *Result {
		if _, err := base64.URLEncoding.DecodeString(*s); err == nil {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsJSON passes when *s is a syntactically valid JSON value
// (object, array, string, number, true, false, or null).
func RuleIsJSON() Rule {
	return Rule{ID: IDIsJSON, Fn: func(s *string) *Result {
		if json.Valid([]byte(*s)) {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsLowercase passes when *s is unchanged by strings.ToLower.
// Characters without case (digits, punctuation, spaces) do not cause failure.
// The empty string passes.
func RuleIsLowercase() Rule {
	return Rule{ID: IDIsLowercase, Fn: func(s *string) *Result {
		if strings.ToLower(*s) == *s {
			return nil
		}
		return &Result{}
	}}
}

// RuleIsUppercase passes when *s is unchanged by strings.ToUpper.
// Characters without case (digits, punctuation, spaces) do not cause failure.
// The empty string passes.
func RuleIsUppercase() Rule {
	return Rule{ID: IDIsUppercase, Fn: func(s *string) *Result {
		if strings.ToUpper(*s) == *s {
			return nil
		}
		return &Result{}
	}}
}

// -----------------------------------------------------------------------------
// Sanitizers — the following rules mutate *s
// -----------------------------------------------------------------------------

// RuleToLower is a sanitizer that lowercases *s.
func RuleToLower() Rule {
	return Rule{ID: IDToLower, Fn: func(s *string) *Result {
		*s = strings.ToLower(*s)
		return nil
	}}
}

// RuleToUpper is a sanitizer that uppercases *s.
func RuleToUpper() Rule {
	return Rule{ID: IDToUpper, Fn: func(s *string) *Result {
		*s = strings.ToUpper(*s)
		return nil
	}}
}

// RuleStripHTMLTags is a sanitizer that removes anything matching <...> from *s.
// This is not a security feature — use RuleEscapeHTML for output-escaping.
// It is intended for simple content cleanup (e.g. stripping markup from a
// pasted snippet) and does not attempt full HTML parsing.
func RuleStripHTMLTags() Rule {
	return Rule{ID: IDStripHTMLTags, Fn: func(s *string) *Result {
		*s = RegexpHTMLTag.ReplaceAllString(*s, "")
		return nil
	}}
}

// RuleIsUUID passes when *s matches the canonical UUID format
// (8-4-4-4-12 hex digits with hyphens). It does not enforce a specific
// UUID version — any version (1–5) and the nil UUID all pass as long as
// the format is correct.
func RuleIsUUID() Rule {
	return Rule{ID: IDIsUUID, Fn: func(s *string) *Result {
		if !RegexpUUID.MatchString(*s) {
			return &Result{}
		}
		return nil
	}}
}
