package ward

// ValidatorPolicy controls global validation behaviour across all fields.
// StopOnFail halts remaining fields on the first field that produces any failure.
type ValidatorPolicy struct {
	StopOnFail bool
}
