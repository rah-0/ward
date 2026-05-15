package ward

// Rule associates a type-package ID and rule ID with its validation function.
// TypeID must be set to the owning package's TypeID constant — it is stamped
// into every Result the rule produces, so the caller (and the frontend) can
// identify which type package the failure belongs to.
// Fn receives a pointer to the value so sanitizers can mutate it in place.
// Fn returns nil on pass and a non-nil Result only on failure.
type Rule[T any] struct {
	TypeID uint32
	ID     uint32
	Fn     func(*T) *Result
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
