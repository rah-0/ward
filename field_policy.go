package ward

// FieldPolicy controls validation behavior for a single field.
// StopOnFail halts remaining rules for this field on the first failure.
type FieldPolicy struct {
	StopOnFail bool
}
