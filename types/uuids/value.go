package uuids

import "github.com/rah-0/ward"

const (
	TypeID uint32 = 8
)

type Rule = ward.Rule[string]
type Field = ward.Field[string]
type Result = ward.Result

// New constructs a Field pointing directly to fieldPtr.
// The underlying type is string; a UUID is validated as a formatted string.
func New(name string, fieldPtr *string, rules ...Rule) *Field {
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
