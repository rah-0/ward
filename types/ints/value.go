package ints

import ward "github.com/rah-0/ward"

const (
	TypeID uint32 = 3
)

type Rule = ward.Rule[int64]
type Field = ward.Field[int64]
type Result = ward.Result

// New constructs a Field pointing directly to fieldPtr.
// Sanitizers mutate *fieldPtr in place, so the source struct field reflects
// the sanitized value after Run(). Callers that need to preserve the original
// should copy it before calling Run().
func New(name string, fieldPtr *int64, rules ...Rule) *Field {
	return &Field{
		TypeID: TypeID,
		Name:   name,
		Value:  fieldPtr,
		Rules:  rules,
	}
}
