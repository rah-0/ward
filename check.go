package ward

// Check is the only contract between the validator and any type package.
// Any value that implements Validate() can be passed to Validate.Add().
type Check interface {
	Validate() []*Result
}
