package percentage

import ward "github.com/rah-0/ward"

const (
	IDInRange    uint32 = 2
	IDIsPositive uint32 = 3
	IDIsWhole    uint32 = 4
)

var IDs = []uint32{
	IDInRange,
	IDIsPositive,
	IDIsWhole,
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
