package strs

import "github.com/rah-0/ward"

const (
	TypeID uint32 = 2
)

type Rule = ward.Rule[string]
type Field = ward.Field[string]
type Result = ward.Result

// New constructs a Field pointing directly to fieldPtr.
// Sanitizers mutate *fieldPtr in place, so the source struct field reflects
// the sanitized value after Run(). Callers that need to preserve the original
// should copy it before calling Run().
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
