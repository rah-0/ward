# Ward — Implementation Plan

A typed, reflection-free, tag-free validation + sanitization library for Go.
Performance-first. Zero coupling between type packages.

- Module: `github.com/rah-0/ward`
- Go: 1.26.2
- Core dependencies: **stdlib only**

## Architecture

### Core principle
Zero coupling between packages. The `result.Check` interface is the only contract.
Any package — first-party or user-written — that implements `Validate() []*result.Result`
works with the validator out of the box, no registration or import needed on ward's side.

### Dependency graph
```
result        ← no deps (pure)
policy        ← no deps (pure)
config        ← no deps (pure, holds regexes + global rule registry)
validator     ← imports result, policy
types/strs    ← imports result, policy, config
types/ints    ← imports result, policy, config  (TODO)
types/bools   ← imports result, policy, config  (TODO)
types/times   ← imports result, policy, config  (TODO)
```
User code imports `validator` + whichever type packages it needs. Ward core never
imports any type package.

### Result
```go
type Result struct {
    TypeID    uint32  // identifies the type package
    RuleID    uint32  // identifies the rule within the package
    FieldName string  // injected by Field.Validate(), set on Field struct
    Valid     bool
    Arg1      any     // rule parameter (e.g. min for LengthMin)
    Arg2      any     // second rule parameter (e.g. max for LengthBetween)
    Err       error   // underlying error for debug (IsEmail, IsURL, UnescapeURL)
}

type Check interface {
    Validate() []*Result
}
```

### Policy
```go
// policy.Field — per-field validation behaviour
type Field struct {
    Required     bool
    Optional     bool
    AllowDefault bool
    AllowNil     bool
    StopOnFail   bool  // stop remaining rules for this field on first failure
}

// policy.Validator — global validation behaviour
type Validator struct {
    StopOnFail bool  // stop remaining fields on first field with any failure
}
```

### Validator
```go
type Validate struct {
    Policy policy.Validator
    checks []result.Check
}
func (v *Validate) Add(c result.Check)
func (v *Validate) Run() []*result.Result
```
Not safe for concurrent use. One instance per request. `checks` is a flat
slice of `result.Check` — any type package's `Field` goes in here.

### Type package pattern (strs as reference)
Every type package follows this exact structure:
```
types/strs/
    value.go   — Value{Current, Original}, Field{FieldName, Policy, Value, Rules},
                 Validate() []*result.Result, RulesAddFromIDs()
    rules.go   — TypeID, ID* constants, IDs []uint32, Rule{ID, Fn},
                 RuleSet, RuleGet, all built-in rule funcs + sanitizers
    errors.go  — package-level errors (ErrRuleNotFound etc.)
```
Each `Field` satisfies `result.Check` enforced by compile-time assertion:
```go
var _ result.Check = (*Field)(nil)
```
`Field.Validate()` injects `TypeID`, `RuleID`, `FieldName` into every result.
`Fn` takes `*Value` so sanitizers can mutate `Current`.

### Global rule registry (config)
```go
func RuleSet(typeID, ruleID uint32, rule any) error  // write-once, returns ErrRuleAlreadyRegistered on duplicate
func RuleGet(typeID, ruleID uint32) (any, error)
func RuleList() map[uint32][]uint32                  // TypeID → []RuleID, for frontend mapping
```
Registry is `map[uint32]map[uint32]any` — two-level: TypeID → RuleID → rule.
Write-once enforced: registering the same TypeID+RuleID twice returns an error.

## What exists

### ✅ ward/result
- `Result` struct with `TypeID`, `RuleID`, `FieldName`, `Valid`, `Arg1`, `Arg2`, `Err`
- `Check` interface

### ✅ ward/policy
- `policy.Field` — `Required`, `Optional`, `AllowDefault`, `AllowNil`, `StopOnFail`
- `policy.Validator` — `StopOnFail`
- `policy/errors.go` — `ErrFieldRequiredWithOptional`, `ErrFieldRequiredWithAllowNil`

### ✅ ward/config
- `config/regexp.go` — precompiled: `RegexpSha512`, `RegexpHasLowercase`, `RegexpHasUppercase`,
  `RegexpDigitsOnly`, `RegexpHasDigit`, `RegexpNonNegativeInt`, `RegexpUsernameChars`, `RegexpHasSpecialChar`
- `config/registry.go` — `RuleSet`, `RuleGet`, `RuleList`
- `config/errors.go` — `ErrRuleNotFound`, `ErrRuleAlreadyRegistered`

### ✅ ward/validator
- `Validate` struct with `Policy policy.Validator`
- `Add(result.Check)`, `Run() []*result.Result`
- `StopOnFail` wired at both global (validator) and field level

### ✅ ward/types/strs (TypeID = 2)
22 validators + 3 sanitizers:

| ID | Rule | Notes |
|----|------|-------|
| 2  | `NotEmpty` | rune count > 0 |
| 3  | `LengthMin(min)` | Arg1=min |
| 4  | `LengthMax(max)` | Arg1=max |
| 5  | `LengthExact(n)` | Arg1=n |
| 6  | `LengthBetween(min,max)` | Arg1=min, Arg2=max |
| 7  | `Contains(sub)` | Arg1=sub |
| 8  | `NotContains(sub)` | Arg1=sub |
| 9  | `MatchesRegex(pattern)` | Arg1=pattern.String() |
| 10 | `IsEmail` | Err set on failure |
| 11 | `IsSha512` | 128 hex chars |
| 12 | `HasLowercase` | at least one [a-z] |
| 13 | `HasUppercase` | at least one [A-Z] |
| 14 | `IsDigitsOnly` | all digits, no spaces |
| 15 | `IsURL` | http/https/ftp/ftps, Err set |
| 16 | `IsNotURL` | inverse of IsURL, Err set |
| 17 | `HasDigit` | at least one digit |
| 18 | `HasSpecialChar` | `@$!%*?&_-=` |
| 19 | `IsPasswordChars` | lower+upper+digit+special |
| 20 | `IsUsernameChars` | `[a-zA-Z0-9_.@-]+` |
| 21 | `IsBoolString` | "true/false/1/0" (trimmed, lowered) |
| 22 | `IsNonNegativeInt` | 0–999,999,999 |
| 23 | `Trim` | sanitizer — TrimSpace |
| 24 | `EscapeHTML` | sanitizer — html.EscapeString |
| 25 | `UnescapeURL` | sanitizer — blanks on error, Err set |

All length rules use `utf8.RuneCountInString` (no allocation).
All `Fn` take `*Value` so sanitizers can mutate `Current`.

## What is missing

### Type packages
- `types/ints` — `int8/16/32/64`, `uint8/16/32/64`, `float32/64` (TypeID TBD)
  Rules: `GT`, `GTE`, `LT`, `LTE`, `Between`, `GTZero`, `GTEZero`
- `types/bools` — `bool` (TypeID TBD)
  Rules: `MustBeTrue`, `MustBeFalse`
- `types/times` — `time.Time` (TypeID TBD)
  Rules: `After`, `AfterOrEqual`, `Before`, `BeforeOrEqual`, `Between`, `NotZero`
- `types/durations` — `time.Duration` (TypeID TBD)
  Rules: `GT`, `GTE`, `LT`, `LTE`, `Between`
- `types/slices` — `[]T` (TypeID TBD)
  Rules: `NotEmpty`, `LengthMin`, `LengthMax`, `LengthBetween`, `Each`

### Validator
- `Reset()` — clear checks and results for instance reuse
- Result filtering helpers — `Failures() []*result.Result` (filter Valid==false)
- Mapping utility — `As[T](results []*result.Result, fn func(*result.Result) T) []T`

### policy.Field
- `Validate()` is implemented but never called during `Field.Validate()` — wire it in
  so invalid policy combinations (Required+Optional etc.) surface as errors

### strs
- Missing string sanitizers: `TrimLeft`, `TrimRight`, `ToLower`, `ToUpper`

### Tests
- No tests written yet for any package
- Target: `go test -race ./...` clean, ≥ 90% coverage

### Benchmarks
- Separate repo `github.com/rah-0/benchmarks` — compare against:
  - `github.com/go-playground/validator/v10`
  - `github.com/go-ozzo/ozzo-validation/v4`
  - legacy `utils/security/validator`
- Scenarios: single field, multi-field DTO, happy path, StopOnFail modes

### README
- Package overview, quick-start example, custom type package guide
