package percentage

import "github.com/rah-0/ward"

const (
	IDInRange    uint32 = 2
	IDIsPositive uint32 = 3
	IDIsWhole    uint32 = 4
)

// IDs maps every rule ID in this package to its name.
var IDs = map[uint32]string{
	IDInRange:    "InRange",
	IDIsPositive: "IsPositive",
	IDIsWhole:    "IsWhole",
}

// IDsAdd registers a custom rule name and returns its automatically assigned ID.
func IDsAdd(name string) uint32 {
	return ward.IDsAdd(IDs, name)
}

func RuleInRange(min, max float64) Rule {
	return Rule{ID: IDInRange, Fn: func(v *float64) *ward.Result {
		if *v >= min && *v <= max {
			return nil
		}
		return &ward.Result{
			Arg1: min,
			Arg2: max,
		}
	}}
}

func RuleIsPositive() Rule {
	return Rule{ID: IDIsPositive, Fn: func(v *float64) *ward.Result {
		if *v > 0 {
			return nil
		}
		return &ward.Result{}
	}}
}

func RuleIsWhole() Rule {
	return Rule{ID: IDIsWhole, Fn: func(v *float64) *ward.Result {
		if *v == float64(int64(*v)) {
			return nil
		}
		return &ward.Result{Arg1: *v}
	}}
}
