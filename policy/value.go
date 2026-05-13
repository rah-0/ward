package policy

type Field struct {
	Required     bool
	Optional     bool
	AllowDefault bool
	AllowNil     bool
	StopOnFail   bool
}

func (x *Field) Validate() error {
	if x.Required && x.Optional {
		return ErrFieldRequiredWithOptional
	}
	if x.Required && x.AllowNil {
		return ErrFieldRequiredWithAllowNil
	}
	return nil
}
