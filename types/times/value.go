package times

import (
	"time"

	"github.com/rah-0/ward"
)

const (
	TypeID uint32 = 6
)

type Rule = ward.Rule[time.Time]
type Field = ward.Field[time.Time]
type Result = ward.Result

// New constructs a Field pointing directly to fieldPtr.
func New(name string, fieldPtr *time.Time, rules ...Rule) *Field {
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
