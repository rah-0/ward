package loginform_test

import (
	"testing"

	"github.com/rah-0/ward/examples/loginform"
	"github.com/rah-0/ward/types/strs"
)

func TestValidate_Valid(t *testing.T) {
	form := &loginform.Form{
		Email:    "user@example.com",
		Password: "Secret1!",
	}
	errs, ok := loginform.Validate(form)
	if !ok {
		t.Fatalf("expected valid form, got %d errors", len(errs))
	}
}

func TestValidate_EmptyFields(t *testing.T) {
	form := &loginform.Form{}
	errs, ok := loginform.Validate(form)
	if ok {
		t.Fatal("expected validation to fail for empty form")
	}
	hasEmailErr := false
	hasPasswordErr := false
	for _, e := range errs {
		if e.Field == "Email" {
			hasEmailErr = true
		}
		if e.Field == "Password" {
			hasPasswordErr = true
		}
	}
	if !hasEmailErr {
		t.Error("expected failure for Email field")
	}
	if !hasPasswordErr {
		t.Error("expected failure for Password field")
	}
}

func TestValidate_InvalidEmail(t *testing.T) {
	form := &loginform.Form{
		Email:    "not-an-email",
		Password: "Secret1!",
	}
	errs, ok := loginform.Validate(form)
	if ok {
		t.Fatal("expected validation to fail for invalid email")
	}
	found := false
	for _, e := range errs {
		if e.Field == "Email" && e.Rule == strs.IDIsEmail {
			found = true
		}
	}
	if !found {
		t.Errorf("expected IsEmail failure on Email field (RuleID %d)", strs.IDIsEmail)
	}
}

func TestValidate_PasswordTooShort(t *testing.T) {
	form := &loginform.Form{
		Email:    "user@example.com",
		Password: "Ab1!",
	}
	errs, ok := loginform.Validate(form)
	if ok {
		t.Fatal("expected validation to fail for short password")
	}
	found := false
	for _, e := range errs {
		if e.Field == "Password" && e.Rule == strs.IDLengthMin {
			found = true
			if e.Arg1 != 8 {
				t.Errorf("expected Arg1=8, got %v", e.Arg1)
			}
		}
	}
	if !found {
		t.Errorf("expected LengthMin failure on Password field (RuleID %d)", strs.IDLengthMin)
	}
}

func TestValidate_WeakPassword(t *testing.T) {
	form := &loginform.Form{
		Email:    "user@example.com",
		Password: "alllowercase",
	}
	errs, ok := loginform.Validate(form)
	if ok {
		t.Fatal("expected validation to fail for weak password")
	}
	found := false
	for _, e := range errs {
		if e.Field == "Password" && e.Rule == strs.IDIsPasswordChars {
			found = true
		}
	}
	if !found {
		t.Errorf("expected IsPasswordChars failure on Password field (RuleID %d)", strs.IDIsPasswordChars)
	}
}
