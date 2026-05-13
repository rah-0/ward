package ward

// FieldPolicy controls validation behaviour for a single field.
// Required and Optional are mutually exclusive.
// StopOnFail halts remaining rules for this field on the first failure.
type FieldPolicy struct {
	Required     bool
	Optional     bool
	AllowDefault bool
	AllowNil     bool
	StopOnFail   bool
}

// Validate checks for contradictory policy combinations.
// Called at the start of Field.Validate().
func (x *FieldPolicy) Validate() error {
	if x.Required && x.Optional {
		return ErrFieldRequiredWithOptional
	}
	if x.Required && x.AllowNil {
		return ErrFieldRequiredWithAllowNil
	}
	return nil
}
