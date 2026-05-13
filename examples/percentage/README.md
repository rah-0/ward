# Example — percentage

Demonstrates a custom ward type package where **T is a primitive** (`float64`).

Rules receive a `*float64` and operate on the value directly. This shows that `T` does not need to be a struct — any type works.

## Type

```go
// No custom struct needed — T is float64 directly.
const TypeID uint32 = 101

type Rule  = ward.Rule[float64]
type Field = ward.Field[float64]
```

## Usage

```go
var score float64

v := ward.New()
score = 75
v.Add(
    percentage.New("Score", &score,
        percentage.RuleInRange(0, 100),
        percentage.RuleIsWhole(),
    ),
)
v.Run()

if v.HasFailures() {
    for _, f := range v.Failures() {
        fmt.Println(f.FieldName, f.RuleID, f.Arg1, f.Arg2)
    }
}
v.Reset()
```

## Rules

| ID | Rule | Description |
|---|---|---|
| 2 | `RuleInRange(min, max)` | Value must be within [min, max]. Arg1=min, Arg2=max on failure |
| 3 | `RuleIsPositive()` | Value must be > 0 |
| 4 | `RuleIsWhole()` | Value must have no fractional part. Arg1=value on failure |
