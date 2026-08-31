# ward

A typed, reflection-free, tag-free validation and sanitization library for Go.

- No struct tags
- No reflection
- No global registry — type packages export a mutable `IDs` map for app-level extensions
- Stdlib only
- Generic: works with any type

## Install

Ward requires Go 1.27 or newer.

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

`Validate` is mutable and must be created per request. A `Field` holds a `*T` to caller-owned data; share a field or its source across goroutines only when every rule is read-only and no goroutine writes the value. Sanitizers necessarily require exclusive access.

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

`(*Validate).FailuresAs[T]` maps failures from the last run to an application-owned type, making it straightforward to shape a response for Go 1.27's `encoding/json/v2`:

```go
import json "encoding/json/v2"

type ValidationError struct {
    Field string `json:"field"`
    Rule  uint32 `json:"rule"`
    Arg1  any    `json:"arg1,omitzero"`
    Arg2  any    `json:"arg2,omitzero"`
}

errs := v.FailuresAs(func(r *ward.Result) ValidationError {
    return ValidationError{
        Field: r.FieldName,
        Rule:  r.RuleID,
        Arg1:  r.Arg1,
        Arg2:  r.Arg2,
    }
})

// json.Marshal(errs) →
// [{"field":"Email","rule":10},{"field":"Password","rule":3,"arg1":8}]
```

The `omitzero` tags omit only nil argument interfaces. JSON v2 has no default representation for `time.Duration`, so applications exposing duration-rule arguments should convert them to their API representation in the `FailuresAs` callback or register a v2 marshaler.

`FailuresAs` preserves the stored slice and failure order while projecting into whatever shape your API layer needs. Its callback receives the original `*Result` pointers, so treat them as read-only if the stored failures must remain unchanged.

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

Every type package exports an `IDs` map (`map[uint32]string`) associating each rule ID with its name, and a `TypeID` constant identifying the package. These can be served from a single endpoint so the frontend always knows what validations exist:

```go
import json "encoding/json/v2"

// GET /api/validation-rules
func GetValidationRules(w http.ResponseWriter, r *http.Request) {
    rules := map[uint32]map[uint32]string{
        strs.TypeID: strs.IDs,
        // add further type packages here as the API grows
    }
    if err := json.MarshalWrite(w, rules, json.Deterministic(true)); err != nil {
        return
    }
}

// response:
// {"2":{"2":"NotEmpty","3":"LengthMin","4":"LengthMax",...}}
```

When a failure arrives at the frontend with `TypeID=2, RuleID=3`, it looks up TypeID 2 → strs, RuleID 3 → `LengthMin`, and can use `Arg1` as the actual minimum value. This keeps numeric rule IDs and configured constraint values synchronized with the backend; the frontend still owns its message templates.

Applications can register custom rules in two ways:

**Automatic ID assignment** — `IDsAdd` picks an ID one greater than the map's current maximum:

```go
var (
    idPasswordsMatch    = strs.IDsAdd("PasswordsMatch")
    idUsernameAvailable = strs.IDsAdd("UsernameAvailable")
)
```

**Manual ID assignment** — write directly to the map when you need a specific ID:

```go
func init() {
    strs.IDs[1000] = "PasswordsMatch"
    strs.IDs[1001] = "UsernameAvailable"
}
```

Register custom IDs during application initialization. The exported `IDs` maps are mutable and are not safe to write while other goroutines read or modify them. Manually assigned IDs must not overwrite existing entries.

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

`New()` stamps `TypeID` on every rule automatically — rule constructors only need `ID` and `Fn`.

See [`examples/`](examples/) for the full implementation guide and the following working examples:

| Example | T | Demonstrates |
|---|---|---|
| [loginform](examples/loginform/) | `string` | Basic usage, `FailuresAs[T]`, structured error responses |
| [phonenumber](examples/phonenumber/) | `struct` | Multi-field struct, parametrized rules |
| [percentage](examples/percentage/) | `float64` | Primitive type, numeric range rules |

## Benchmarks

Historical comparison against `go-playground/validator` and `ozzo-validation`:
[github.com/rah-0/benchmarks/tree/master/validator](https://github.com/rah-0/benchmarks/tree/master/validator#readme)

## ☕ Support
Enjoying ward?
If it saved you time or brought value to your project, feel free to show some support. Every bit is appreciated 🙂

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/default-orange.png)](https://www.buymeacoffee.com/rah.0)
