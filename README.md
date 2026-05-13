# ward

A typed, reflection-free, tag-free validation and sanitization library for Go.

- No struct tags
- No reflection
- No global registry
- Stdlib only
- Generic: works with any type

## Install

```
go get github.com/rah-0/ward
```

## Quick start

See [`examples/loginform`](examples/loginform/) for a complete working example covering basic validation, structured error responses, and field-level failure inspection.

The core pattern for concurrent use (e.g. an HTTP handler):

```go
func handle(r *http.Request) {
    var form LoginForm
    // populate form from request...

    v := ward.New().Add(
        strs.New("Email",    &form.Email,    strs.RuleNotEmpty(), strs.RuleIsEmail()),
        strs.New("Password", &form.Password, strs.RuleNotEmpty(), strs.RuleLengthMin(8)),
    ).Run()

    if v.HasFailures() {
        // inspect v.Failures()
    }
}
```

`Validate`, `Field`, and the form struct must all be per-request. Sharing any of them across goroutines is a data race — `Field` holds a `*T` to the source value, and `Validate` accumulates results in a mutable slice.

## Sanitizers

Sanitizers are rules that mutate the value in place. They run in the same rule chain and write back to the source pointer directly.

```go
name := "  alice  "
v := ward.New().Add(
    strs.New("Name", &name, strs.RuleTrim(), strs.RuleNotEmpty()),
).Run()

fmt.Println(name)    // "alice" — source updated in place
_ = v.HasFailures()
```

Sanitizers write back through the pointer, so the source variable reflects the sanitized value immediately after `Run()`. Callers that need to preserve the original should copy it before calling `Run()`.

## Structured failure responses

`ward.As[T]` maps failures to any type, making it straightforward to produce a JSON-serialisable response:

```go
type ValidationError struct {
    Field string `json:"field"`
    Rule  uint32 `json:"rule"`
    Arg1  any    `json:"arg1,omitempty"`
    Arg2  any    `json:"arg2,omitempty"`
}

errs := ward.As(v.Failures(), func(r *ward.Result) ValidationError {
    return ValidationError{
        Field: r.FieldName,
        Rule:  r.RuleID,
        Arg1:  r.Arg1,
        Arg2:  r.Arg2,
    }
})

// json.Marshal(errs) →
// [{"field":"Password","rule":3,"arg1":8},{"field":"Email","rule":10}]
```

`As` never touches the original `[]*Result` slice — it projects into whatever shape your API layer needs.

## Frontend integration

### Arg1 and Arg2

Parametrized rules carry their values back in `Arg1` and `Arg2`. The frontend receives the exact constraint the backend enforced — no need to duplicate ID's and configuration in client code.

```
// backend: strs.RuleLengthMin(8) fails → Result{RuleID: 3, Arg1: 8}
// frontend receives: {"field":"Password","rule":3,"arg1":8}
// frontend renders:  "Password must be at least 8 characters"
```

Similarly, `RuleLengthBetween(5, 50)` returns `Arg1=5, Arg2=50`, and `RuleContains("@")` returns `Arg1="@"`.

### Exposing available rules

Every type package exports an `IDs` slice listing all rule IDs and a `TypeID` constant identifying the package. These can be served from a single endpoint so the frontend always knows what validations exist:

```go
// GET /api/validation-rules
func GetValidationRules(w http.ResponseWriter, r *http.Request) {
    rules := map[uint32][]uint32{
        strs.TypeID: strs.IDs,
        // add further type packages here as the API grows
    }
    json.NewEncoder(w).Encode(rules)
}

// response:
// {"2":[2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25]}
```

When a failure arrives at the frontend with `TypeID=2, RuleID=3`, it looks up TypeID 2 → strs, RuleID 3 → `LengthMin`, and can display the right message using `Arg1` as the actual minimum value. The frontend never hardcodes validation logic — it derives everything from what the backend exposes.

## StopOnFail

Stop at the first failing field across the whole validator:

```go
v.Policy.StopOnFail = true
```

Stop at the first failing rule within a single field:

```go
fieldEmail.Policy.StopOnFail = true
```

## Custom types

`ward.Rule[T]` and `ward.Field[T]` are generic over any type `T` — a struct, a primitive, a type alias. Implementing a custom type package requires only a TypeID, two type aliases, and a `New()` function.

See [`examples/`](examples/) for the full implementation guide and the following working examples:

| Example | T | Demonstrates |
|---|---|---|
| [loginform](examples/loginform/) | `string` | Basic usage, `As[T]`, structured error responses |
| [phonenumber](examples/phonenumber/) | `struct` | Multi-field struct, parametrized rules |
| [percentage](examples/percentage/) | `float64` | Primitive type, numeric range rules |

## Benchmarks

Full comparison against `go-playground/validator` and `ozzo-validation`:
[github.com/rah-0/benchmarks/tree/master/validator](https://github.com/rah-0/benchmarks/tree/master/validator#readme)

## ☕ Support

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/default-orange.png)](https://www.buymeacoffee.com/rah.0)
