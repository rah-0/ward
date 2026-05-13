package ward

import "errors"

var (
	ErrFieldRequiredWithOptional = errors.New("ward: FieldPolicy: Required and Optional cannot both be true")
	ErrFieldRequiredWithAllowNil = errors.New("ward: FieldPolicy: Required and AllowNil cannot both be true")
)
