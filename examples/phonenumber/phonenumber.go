// Package phonenumber demonstrates a custom ward type package where T is a struct.
package phonenumber

import ward "github.com/rah-0/ward"

const TypeID uint32 = 100

// PhoneNumber is a custom struct type used as the validation target.
// T in ward.Rule[T] and ward.Field[T] can be any type — struct, primitive, or otherwise.
type PhoneNumber struct {
	CountryCode string
	Number      string
}

type Rule = ward.Rule[PhoneNumber]
type Field = ward.Field[PhoneNumber]

func New(name string, ptr *PhoneNumber, rules ...Rule) *Field {
	return &Field{
		TypeID: TypeID,
		Name:   name,
		Value:  ptr,
		Rules:  rules,
	}
}
