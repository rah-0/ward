package policy

import "errors"

var (
	ErrFieldRequiredWithOptional = errors.New("ward: policy field: Required and Optional cannot both be true")
	ErrFieldRequiredWithAllowNil = errors.New("ward: policy field: Required and AllowNil cannot both be true")
)
