package validator

import (
	"github.com/rah-0/ward/policy"
	"github.com/rah-0/ward/result"
)

type Validate struct {
	Policy policy.Validator
	checks []result.Check
}

func (v *Validate) Add(c result.Check) {
	v.checks = append(v.checks, c)
}

func (v *Validate) Run() []*result.Result {
	var results []*result.Result
	for _, check := range v.checks {
		fieldResults := check.Validate()
		results = append(results, fieldResults...)
		if v.Policy.StopOnFail {
			for _, r := range fieldResults {
				if !r.Valid {
					return results
				}
			}
		}
	}
	return results
}
