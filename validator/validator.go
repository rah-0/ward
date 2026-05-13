package validator

import (
	"github.com/rah-0/ward/policy"
	"github.com/rah-0/ward/result"
)

type Validate struct {
	Policy  policy.Validator
	checks  []result.Check
	results []*result.Result
}

func New() *Validate {
	return &Validate{}
}

func (v *Validate) Reset() {
	v.checks = v.checks[:0]
	v.results = v.results[:0]
}

func (v *Validate) Add(c result.Check) {
	v.checks = append(v.checks, c)
}

func (v *Validate) Run() []*result.Result {
	v.results = v.results[:0]
	for _, check := range v.checks {
		fieldResults := check.Validate()
		v.results = append(v.results, fieldResults...)
		if v.Policy.StopOnFail {
			for _, r := range fieldResults {
				if !r.Valid {
					return v.results
				}
			}
		}
	}
	return v.results
}

func (v *Validate) HasFailures() bool {
	for _, r := range v.results {
		if !r.Valid {
			return true
		}
	}
	return false
}

func (v *Validate) Failures() []*result.Result {
	var out []*result.Result
	for _, r := range v.results {
		if !r.Valid {
			out = append(out, r)
		}
	}
	return out
}
