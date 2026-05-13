// Package uuids provides UUID validation rules for ward.
// The underlying type is string. A UUID is represented as a
// hyphenated hex string: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
package uuids

const (
	IDIsValid  uint32 = 2
	IDOneOf    uint32 = 3
	IDNotOneOf uint32 = 4
)

// IDs lists all rule IDs in this package.
var IDs = []uint32{
	IDIsValid,
	IDOneOf, IDNotOneOf,
}

// RuleIsValid passes when *s is a valid UUID string.
func RuleIsValid() Rule {
	return Rule{ID: IDIsValid, Fn: func(s *string) *Result {
		if RegexpUUID.MatchString(*s) {
			return nil
		}
		return &Result{}
	}}
}

// RuleOneOf passes when *s equals one of the allowed UUIDs.
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

// RuleNotOneOf passes when *s does not equal any of the excluded UUIDs.
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
