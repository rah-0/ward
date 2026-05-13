// Package loginform demonstrates basic ward usage with string fields.
package loginform

import (
	ward "github.com/rah-0/ward"
	"github.com/rah-0/ward/types/strs"
)

type Form struct {
	Email    string
	Password string
}

type ValidationError struct {
	Field string
	Rule  uint32
	Arg1  any
	Arg2  any
}

func Validate(form *Form) ([]ValidationError, bool) {
	v := ward.New().Add(
		strs.New("Email", &form.Email, strs.RuleNotEmpty(), strs.RuleIsEmail()),
		strs.New("Password", &form.Password, strs.RuleNotEmpty(), strs.RuleLengthMin(8), strs.RuleIsPasswordChars()),
	).Run()

	if !v.HasFailures() {
		return nil, true
	}

	errs := ward.As(v.Failures(), func(r *ward.Result) ValidationError {
		return ValidationError{
			Field: r.FieldName,
			Rule:  r.RuleID,
			Arg1:  r.Arg1,
			Arg2:  r.Arg2,
		}
	})

	return errs, false
}
