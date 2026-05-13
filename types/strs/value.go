package strs

import (
	"github.com/rah-0/ward/config"
	"github.com/rah-0/ward/policy"
	"github.com/rah-0/ward/result"
)

type Value struct {
	Current  string
	Original *string
}

type Rule struct {
	ID uint32
	Fn func(*Value) *result.Result
}

type Field struct {
	FieldName string
	Policy    policy.Field
	Value     Value
	Rules     []Rule
}

var _ result.Check = (*Field)(nil)

func (f *Field) RulesAddFromIDs(ids ...uint32) error {
	for _, id := range ids {
		raw, err := config.RuleGet(TypeID, id)
		if err != nil {
			return err
		}
		f.Rules = append(f.Rules, raw.(Rule))
	}
	return nil
}

func (f *Field) Validate() []*result.Result {
	if err := f.Policy.Validate(); err != nil {
		return []*result.Result{{FieldName: f.FieldName, Err: err}}
	}
	var results []*result.Result
	for _, rule := range f.Rules {
		r := rule.Fn(&f.Value)
		if r == nil {
			continue
		}
		r.TypeID = TypeID
		r.RuleID = rule.ID
		r.FieldName = f.FieldName
		results = append(results, r)
		if f.Policy.StopOnFail {
			break
		}
	}
	return results
}
