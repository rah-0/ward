package ward

// Field holds a pointer to the value being validated, its rules, and its policy.
// TypeID identifies which type package created the field; it is set by each
// package's New() function and can be read for informational purposes, but
// Validate() reads the TypeID from each Rule, not from the Field.
// Sanitizers mutate *Value in place — callers that need to preserve the original
// should copy it before calling Run().
type Field[T any] struct {
	TypeID uint32
	Name   string
	Value  *T
	Rules  []Rule[T]
	Policy FieldPolicy
}

var _ Check = (*Field[any])(nil)

func (f *Field[T]) Validate() []*Result {
	if err := f.Policy.Validate(); err != nil {
		return []*Result{
			{
				FieldName: f.Name,
				Err:       err,
			},
		}
	}

	var results []*Result
	for _, r := range f.Rules {
		res := r.Fn(f.Value)
		if res == nil {
			continue
		}
		res.TypeID = r.TypeID
		res.RuleID = r.ID
		res.FieldName = f.Name
		results = append(results, res)
		if f.Policy.StopOnFail {
			break
		}
	}
	return results
}
