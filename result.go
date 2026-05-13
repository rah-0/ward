package ward

// Result is returned by a rule only on failure. A nil return means the rule passed.
// TypeID and RuleID identify the type package and specific rule that failed.
// FieldName is injected by Field.Validate() from the field's configured name.
// Arg1 and Arg2 carry rule parameters back to the caller (e.g. min/max for length rules).
// Err carries the underlying error for rules that parse or decode (IsEmail, IsURL, UnescapeURL).
type Result struct {
	TypeID    uint32
	RuleID    uint32
	FieldName string
	Arg1      any
	Arg2      any
	Err       error
}

// As maps a slice of results to any type T using fn, allowing callers to
// project failures into their own wire format without coupling to Result.
func As[T any](results []*Result, fn func(*Result) T) []T {
	out := make([]T, len(results))
	for i, r := range results {
		out[i] = fn(r)
	}
	return out
}
