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

var IDs = []uint32{
	IDHasCountryCode,
	IDHasNumber,
	IDValidLength,
	IDDigitsOnly,
}

func RuleHasCountryCode() Rule {
	return Rule{TypeID: TypeID, ID: IDHasCountryCode, Fn: func(p *PhoneNumber) *ward.Result {
		if strings.TrimSpace(p.CountryCode) != "" {
			return nil
		}
		return &ward.Result{}
	}}
}

func RuleHasNumber() Rule {
	return Rule{TypeID: TypeID, ID: IDHasNumber, Fn: func(p *PhoneNumber) *ward.Result {
		if strings.TrimSpace(p.Number) != "" {
			return nil
		}
		return &ward.Result{}
	}}
}

func RuleValidLength(min, max int) Rule {
	return Rule{TypeID: TypeID, ID: IDValidLength, Fn: func(p *PhoneNumber) *ward.Result {
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
	return Rule{TypeID: TypeID, ID: IDDigitsOnly, Fn: func(p *PhoneNumber) *ward.Result {
		for _, c := range p.Number {
			if c < '0' || c > '9' {
				return &ward.Result{}
			}
		}
		return nil
	}}
}
