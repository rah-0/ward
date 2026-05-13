// Package strs provides string validation and sanitization rules for ward.
package strs

import (
	"html"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
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
)

// IDs lists all rule IDs in this package. Use this to enumerate available
// rules for a frontend or documentation — iterate the slice rather than
// maintaining a separate list alongside the constants.
var IDs = []uint32{
	IDNotEmpty, IDLengthMin, IDLengthMax, IDLengthExact, IDLengthBetween,
	IDContains, IDNotContains, IDMatchesRegex,
	IDIsEmail, IDIsSha512,
	IDHasLowercase, IDHasUppercase, IDIsDigitsOnly,
	IDIsURL, IDIsNotURL,
	IDHasDigit, IDHasSpecialChar, IDIsPasswordChars, IDIsUsernameChars,
	IDIsBoolString, IDIsNonNegativeInt,
	IDTrim, IDEscapeHTML, IDUnescapeURL,
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
