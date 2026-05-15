package phonenumber_test

import (
	"testing"

	"github.com/rah-0/ward"
	"github.com/rah-0/ward/examples/phonenumber"
)

func newField(p *phonenumber.PhoneNumber) *phonenumber.Field {
	return phonenumber.New("Phone", p,
		phonenumber.RuleHasCountryCode(),
		phonenumber.RuleHasNumber(),
		phonenumber.RuleValidLength(7, 15),
		phonenumber.RuleDigitsOnly(),
	)
}

func TestPhoneNumber_Valid(t *testing.T) {
	phone := phonenumber.PhoneNumber{CountryCode: "+1", Number: "8005550100"}
	v := ward.New()
	v.Add(newField(&phone))
	v.Run()
	if v.HasFailures() {
		t.Fatalf("expected no failures, got %d", len(v.Failures()))
	}
}

func TestPhoneNumber_MissingCountryCode(t *testing.T) {
	phone := phonenumber.PhoneNumber{CountryCode: "", Number: "8005550100"}
	v := ward.New()
	v.Add(newField(&phone))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for missing country code")
	}
	if v.Failures()[0].RuleID != phonenumber.IDHasCountryCode {
		t.Errorf("expected RuleID %d, got %d", phonenumber.IDHasCountryCode, v.Failures()[0].RuleID)
	}
}

func TestPhoneNumber_MissingNumber(t *testing.T) {
	phone := phonenumber.PhoneNumber{CountryCode: "+1", Number: ""}
	v := ward.New()
	v.Add(newField(&phone))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for missing number")
	}
	found := false
	for _, f := range v.Failures() {
		if f.RuleID == phonenumber.IDHasNumber {
			found = true
		}
	}
	if !found {
		t.Errorf("expected failure with RuleID %d", phonenumber.IDHasNumber)
	}
}

func TestPhoneNumber_NumberTooShort(t *testing.T) {
	phone := phonenumber.PhoneNumber{CountryCode: "+1", Number: "123"}
	v := ward.New()
	v.Add(newField(&phone))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for number too short")
	}
	found := false
	for _, f := range v.Failures() {
		if f.RuleID == phonenumber.IDValidLength {
			found = true
			if f.Arg1 != 7 || f.Arg2 != 15 {
				t.Errorf("expected Arg1=7 Arg2=15, got Arg1=%v Arg2=%v", f.Arg1, f.Arg2)
			}
		}
	}
	if !found {
		t.Errorf("expected failure with RuleID %d", phonenumber.IDValidLength)
	}
}

func TestPhoneNumber_NonDigits(t *testing.T) {
	phone := phonenumber.PhoneNumber{CountryCode: "+1", Number: "800-555-0100"}
	v := ward.New()
	v.Add(newField(&phone))
	v.Run()
	if !v.HasFailures() {
		t.Fatal("expected failure for non-digit characters")
	}
	found := false
	for _, f := range v.Failures() {
		if f.RuleID == phonenumber.IDDigitsOnly {
			found = true
		}
	}
	if !found {
		t.Errorf("expected failure with RuleID %d", phonenumber.IDDigitsOnly)
	}
}

func TestPhoneNumber_TypeID(t *testing.T) {
	phone := phonenumber.PhoneNumber{CountryCode: "", Number: ""}
	v := ward.New()
	v.Add(newField(&phone))
	v.Run()
	for _, f := range v.Failures() {
		if f.TypeID != phonenumber.TypeID {
			t.Errorf("expected TypeID %d, got %d", phonenumber.TypeID, f.TypeID)
		}
	}
}
