package phonenumber

import (
	"strings"
	"unicode/utf8"

	"github.com/rah-0/ward"
)

const (
	IDHasCountryCode uint32 = 2
	IDHasNumber      uint32 = 3
	IDValidLength    uint32 = 4
	IDDigitsOnly     uint32 = 5
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDHasCountryCode: "HasCountryCode",
	IDHasNumber:      "HasNumber",
	IDValidLength:    "ValidLength",
	IDDigitsOnly:     "DigitsOnly",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

func RuleHasCountryCode() Rule {
	return Rule{ID: IDHasCountryCode, Fn: func(p *PhoneNumber) *ward.Result {
		if strings.TrimSpace(p.CountryCode) != "" {
			return nil
		}
		return &ward.Result{}
	}}
}

func RuleHasNumber() Rule {
	return Rule{ID: IDHasNumber, Fn: func(p *PhoneNumber) *ward.Result {
		if strings.TrimSpace(p.Number) != "" {
			return nil
		}
		return &ward.Result{}
	}}
}

func RuleValidLength(min, max int) Rule {
	return Rule{ID: IDValidLength, Fn: func(p *PhoneNumber) *ward.Result {
		n := utf8.RuneCountInString(p.Number)
		if n >= min && n <= max {
			return nil
		}
		return &ward.Result{
			Arg1: min,
			Arg2: max,
		}
	}}
}

func RuleDigitsOnly() Rule {
	return Rule{ID: IDDigitsOnly, Fn: func(p *PhoneNumber) *ward.Result {
		for _, c := range p.Number {
			if c < '0' || c > '9' {
				return &ward.Result{}
			}
		}
		return nil
	}}
}
