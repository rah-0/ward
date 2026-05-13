package result

type Result struct {
	TypeID    uint32
	RuleID    uint32
	FieldName string
	Arg1      any
	Arg2      any
	Err       error
}

type Check interface {
	Validate() []*Result
}

func As[T any](results []*Result, fn func(*Result) T) []T {
	out := make([]T, len(results))
	for i, r := range results {
		out[i] = fn(r)
	}
	return out
}
