package result

type Result struct {
	TypeID    uint32
	RuleID    uint32
	FieldName string
	Valid     bool
	Arg1      any
	Arg2      any
	Err       error
}

type Check interface {
	Validate() []*Result
}
