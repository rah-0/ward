package ward

// Rule associates a rule ID with its validation function for a given value type T.
// Fn receives a pointer to the value so sanitizers can mutate it in place.
// Fn returns nil on pass and a non-nil Result only on failure.
type Rule[T any] struct {
	ID uint32
	Fn func(*T) *Result
}
