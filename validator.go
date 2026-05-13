package ward

// Validate holds a set of fields to validate and accumulates their results.
// Create one instance per request — never share across goroutines.
type Validate struct {
	Policy  ValidatorPolicy
	checks  []Check
	results []*Result
}

// New returns a ready-to-use Validate instance. Call once per request.
func New() *Validate {
	return &Validate{}
}

// Add registers one or more fields to be validated on the next Run call.
// Returns the same Validate instance for chaining.
func (v *Validate) Add(checks ...Check) *Validate {
	v.checks = append(v.checks, checks...)
	return v
}

// Run executes all added checks in order and accumulates failures.
// If Policy.StopOnFail is set, it stops as soon as the first field fails.
// Results are only stored for failing rules — passing rules produce no entries.
// Returns the same Validate instance for chaining: ward.New().Add(...).Run()
func (v *Validate) Run() *Validate {
	v.results = v.results[:0]
	for _, check := range v.checks {
		fieldResults := check.Validate()
		v.results = append(v.results, fieldResults...)
		if v.Policy.StopOnFail && len(fieldResults) > 0 {
			return v
		}
	}
	return v
}

// HasFailures reports whether the last Run produced any failures.
func (v *Validate) HasFailures() bool {
	return len(v.results) > 0
}

// Failures returns all failures from the last Run.
func (v *Validate) Failures() []*Result {
	return v.results
}
