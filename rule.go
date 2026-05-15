package ward

// Rule associates a rule ID with its validation function for a given value type T.
// Fn receives a pointer to the value so sanitizers can mutate it in place.
// Fn returns nil on pass and a non-nil Result only on failure.
type Rule[T any] struct {
	ID uint32
	Fn func(*T) *Result
}

// IDsAdd registers a custom rule name into ids and returns its automatically
// assigned ID. The ID is one greater than the current maximum key in the map.
func IDsAdd(ids map[uint32]string, name string) uint32 {
	var maxID uint32
	for id := range ids {
		if id > maxID {
			maxID = id
		}
	}
	maxID++
	ids[maxID] = name
	return maxID
}
