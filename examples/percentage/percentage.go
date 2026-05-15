// Package percentage demonstrates a custom ward type package where T is a primitive (float64).
package percentage

import "github.com/rah-0/ward"

const TypeID uint32 = 101

type Rule = ward.Rule[float64]
type Field = ward.Field[float64]

func New(name string, ptr *float64, rules ...Rule) *Field {
	return &Field{
		TypeID: TypeID,
		Name:   name,
		Value:  ptr,
		Rules:  rules,
	}
}
