package durations

import (
	"time"

	"github.com/rah-0/ward"
)

const (
	TypeID uint32 = 7
)

type Rule = ward.Rule[time.Duration]
type Field = ward.Field[time.Duration]
type Result = ward.Result

// New constructs a Field pointing directly to fieldPtr.
func New(name string, fieldPtr *time.Duration, rules ...Rule) *Field {
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
