package strs

import (
	"html"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/rah-0/ward/config"
	"github.com/rah-0/ward/result"
)

const (
	TypeID uint32 = 2

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

var IDs = []uint32{
	IDNotEmpty, IDLengthMin, IDLengthMax, IDLengthExact, IDLengthBetween,
	IDContains, IDNotContains, IDMatchesRegex, IDIsEmail, IDIsSha512,
	IDHasLowercase, IDHasUppercase, IDIsDigitsOnly, IDIsURL, IDIsNotURL,
	IDHasDigit, IDHasSpecialChar, IDIsPasswordChars, IDIsUsernameChars,
	IDIsBoolString, IDIsNonNegativeInt,
	IDTrim, IDEscapeHTML, IDUnescapeURL,
}

func NotEmpty() Rule {
	return Rule{ID: IDNotEmpty, Fn: func(v *Value) *result.Result {
		if utf8.RuneCountInString(v.Current) > 0 {
			return nil
		}
		return &result.Result{}
	}}
}

func LengthMin(min int) Rule {
	return Rule{ID: IDLengthMin, Fn: func(v *Value) *result.Result {
		if utf8.RuneCountInString(v.Current) >= min {
			return nil
		}
		return &result.Result{Arg1: min}
	}}
}

func LengthMax(max int) Rule {
	return Rule{ID: IDLengthMax, Fn: func(v *Value) *result.Result {
		if utf8.RuneCountInString(v.Current) <= max {
			return nil
		}
		return &result.Result{Arg1: max}
	}}
}

func LengthExact(length int) Rule {
	return Rule{ID: IDLengthExact, Fn: func(v *Value) *result.Result {
		if utf8.RuneCountInString(v.Current) == length {
			return nil
		}
		return &result.Result{Arg1: length}
	}}
}

func LengthBetween(min, max int) Rule {
	return Rule{ID: IDLengthBetween, Fn: func(v *Value) *result.Result {
		l := utf8.RuneCountInString(v.Current)
		if l >= min && l <= max {
			return nil
		}
		return &result.Result{Arg1: min, Arg2: max}
	}}
}

func Contains(sub string) Rule {
	return Rule{ID: IDContains, Fn: func(v *Value) *result.Result {
		if strings.Contains(v.Current, sub) {
			return nil
		}
		return &result.Result{Arg1: sub}
	}}
}

func NotContains(sub string) Rule {
	return Rule{ID: IDNotContains, Fn: func(v *Value) *result.Result {
		if !strings.Contains(v.Current, sub) {
			return nil
		}
		return &result.Result{Arg1: sub}
	}}
}

func MatchesRegex(pattern *regexp.Regexp) Rule {
	return Rule{ID: IDMatchesRegex, Fn: func(v *Value) *result.Result {
		if pattern.MatchString(v.Current) {
			return nil
		}
		return &result.Result{Arg1: pattern.String()}
	}}
}

func IsEmail() Rule {
	return Rule{ID: IDIsEmail, Fn: func(v *Value) *result.Result {
		_, err := mail.ParseAddress(v.Current)
		if err == nil {
			return nil
		}
		return &result.Result{Err: err}
	}}
}

func IsSha512() Rule {
	return Rule{ID: IDIsSha512, Fn: func(v *Value) *result.Result {
		if len(v.Current) == 128 && config.RegexpSha512.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func HasLowercase() Rule {
	return Rule{ID: IDHasLowercase, Fn: func(v *Value) *result.Result {
		if config.RegexpHasLowercase.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func HasUppercase() Rule {
	return Rule{ID: IDHasUppercase, Fn: func(v *Value) *result.Result {
		if config.RegexpHasUppercase.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func IsDigitsOnly() Rule {
	return Rule{ID: IDIsDigitsOnly, Fn: func(v *Value) *result.Result {
		if config.RegexpDigitsOnly.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func IsURL() Rule {
	return Rule{ID: IDIsURL, Fn: func(v *Value) *result.Result {
		u, err := url.ParseRequestURI(v.Current)
		if err == nil && u.Host != "" && (u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "ftp" || u.Scheme == "ftps") {
			return nil
		}
		return &result.Result{Err: err}
	}}
}

func IsNotURL() Rule {
	return Rule{ID: IDIsNotURL, Fn: func(v *Value) *result.Result {
		u, err := url.ParseRequestURI(v.Current)
		if !(err == nil && u.Host != "" && (u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "ftp" || u.Scheme == "ftps")) {
			return nil
		}
		return &result.Result{Err: err}
	}}
}

func HasDigit() Rule {
	return Rule{ID: IDHasDigit, Fn: func(v *Value) *result.Result {
		if config.RegexpHasDigit.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func HasSpecialChar() Rule {
	return Rule{ID: IDHasSpecialChar, Fn: func(v *Value) *result.Result {
		if config.RegexpHasSpecialChar.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func IsPasswordChars() Rule {
	return Rule{ID: IDIsPasswordChars, Fn: func(v *Value) *result.Result {
		if config.RegexpHasLowercase.MatchString(v.Current) &&
			config.RegexpHasUppercase.MatchString(v.Current) &&
			config.RegexpHasDigit.MatchString(v.Current) &&
			config.RegexpHasSpecialChar.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func IsUsernameChars() Rule {
	return Rule{ID: IDIsUsernameChars, Fn: func(v *Value) *result.Result {
		if config.RegexpUsernameChars.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

func IsBoolString() Rule {
	return Rule{ID: IDIsBoolString, Fn: func(v *Value) *result.Result {
		s := strings.ToLower(strings.TrimSpace(v.Current))
		if s == "1" || s == "0" || s == "true" || s == "false" {
			return nil
		}
		return &result.Result{}
	}}
}

func IsNonNegativeInt() Rule {
	return Rule{ID: IDIsNonNegativeInt, Fn: func(v *Value) *result.Result {
		if config.RegexpNonNegativeInt.MatchString(v.Current) {
			return nil
		}
		return &result.Result{}
	}}
}

// -----------------------------------------------------------------------------
// Sanitizers — the following rules mutate v.Current
// -----------------------------------------------------------------------------

func Trim() Rule {
	return Rule{ID: IDTrim, Fn: func(v *Value) *result.Result {
		v.Current = strings.TrimSpace(v.Current)
		return nil
	}}
}

func EscapeHTML() Rule {
	return Rule{ID: IDEscapeHTML, Fn: func(v *Value) *result.Result {
		v.Current = html.EscapeString(v.Current)
		return nil
	}}
}

func UnescapeURL() Rule {
	return Rule{ID: IDUnescapeURL, Fn: func(v *Value) *result.Result {
		decoded, err := url.QueryUnescape(v.Current)
		if err != nil {
			v.Current = ""
			return &result.Result{Err: err}
		}
		v.Current = decoded
		return nil
	}}
}
