package strs

import (
	"github.com/rah-0/ward/policy"
	"github.com/rah-0/ward/result"
)

type Value struct {
	Current  string
	Original *string
}

var _ result.Check = (*Field)(nil)

type Field struct {
	FieldName string
	Policy    policy.Field
	Value     Value
	Rules     []Rule
}

func (f *Field) RulesAddFromIDs(ids ...uint32) error {
	for _, id := range ids {
		r, err := RuleGet(id)
		if err != nil {
			return err
		}
		f.Rules = append(f.Rules, r)
	}
	return nil
}

func (f *Field) Validate() []*result.Result {
	var results []*result.Result
	for _, rule := range f.Rules {
		r := rule.Fn(&f.Value)
		r.TypeID = TypeID
		r.RuleID = rule.ID
		r.FieldName = f.FieldName
		results = append(results, r)
		if !r.Valid && f.Policy.StopOnFail {
			break
		}
	}
	return results
}
