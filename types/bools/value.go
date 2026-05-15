package bools

import "github.com/rah-0/ward"

const (
	TypeID uint32 = 5
)

type Rule = ward.Rule[bool]
type Field = ward.Field[bool]
type Result = ward.Result

// New constructs a Field pointing directly to fieldPtr.
func New(name string, fieldPtr *bool, rules ...Rule) *Field {
	for i := range rules {
		rules[i].TypeID = TypeID
	}
	return &Field{
		TypeID: TypeID,
		Name:   name,
		Value:  fieldPtr,
		Rules:  rules,
	}
}
