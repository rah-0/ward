# Ward — Custom Type Examples

## Available examples

| Example | T | Demonstrates |
|---|---|---|
| [loginform](loginform/) | `string` | Basic usage, `As[T]`, structured error responses |
| [phonenumber](phonenumber/) | `struct` | Multi-field struct, parametrized rules, accessing individual fields through `*T` |
| [percentage](percentage/) | `float64` | Primitive type, numeric range rules, whole-number check |

---

## How to implement a custom type package

Any type can be validated with ward — a struct, a primitive, a type alias, even an array. All you need is four things.

### 1. Pick a TypeID

Choose a `uint32` that uniquely identifies your type package. Ward's built-ins occupy 1–99. Examples use 100+. Pick anything above that for your own packages.

```go
const TypeID uint32 = 200
```

### 2. Alias Rule and Field

```go
type Rule  = ward.Rule[YourType]
type Field = ward.Field[YourType]
```

These are type aliases (using `=`), not new types. They exist purely for ergonomics — so callers write `Rule` instead of `ward.Rule[YourType]` everywhere.

### 3. Write New()

```go
func New(name string, ptr *YourType, rules ...Rule) *Field {
    return &Field{
        TypeID: TypeID,
        Name:   name,
        Value:  ptr,
        Rules:  rules,
    }
}
```

`ptr` must point to the actual value being validated. Ward reads and writes through it directly — no copy is made. Sanitizers that mutate `*ptr` will update the source immediately.

### 4. Write rule functions

A rule returns `nil` on pass and a non-nil `*ward.Result` only on failure.

```go
func RuleMustBePositive() Rule {
    return Rule{ID: 2, Fn: func(v *YourType) *ward.Result {
        if *v > 0 {
            return nil
        }
        return &ward.Result{}
    }}
}
```

Use `Arg1` and `Arg2` on the result to carry rule parameters back to the caller:

```go
func RuleInRange(min, max YourType) Rule {
    return Rule{ID: 3, Fn: func(v *YourType) *ward.Result {
        if *v >= min && *v <= max {
            return nil
        }
        return &ward.Result{Arg1: min, Arg2: max}
    }}
}
```

Use `Err` for rules that wrap a stdlib parse or decode call:

```go
func RuleIsValid() Rule {
    return Rule{ID: 4, Fn: func(v *YourType) *ward.Result {
        if err := validate(*v); err != nil {
            return &ward.Result{Err: err}
        }
        return nil
    }}
}
```

### Sanitizers

A rule that mutates `*v` and returns `nil` is a sanitizer. It runs in the same chain as validators.

```go
func RuleNormalize() Rule {
    return Rule{ID: 5, Fn: func(v *YourType) *ward.Result {
        *v = normalize(*v)
        return nil
    }}
}
```

### Minimal complete example

```go
package mytype

import ward "github.com/rah-0/ward"

type MyType int

const TypeID uint32 = 200

type Rule  = ward.Rule[MyType]
type Field = ward.Field[MyType]

func New(name string, ptr *MyType, rules ...Rule) *Field {
    return &Field{TypeID: TypeID, Name: name, Value: ptr, Rules: rules}
}

func RuleIsPositive() Rule {
    return Rule{ID: 2, Fn: func(v *MyType) *ward.Result {
        if *v > 0 {
            return nil
        }
        return &ward.Result{}
    }}
}
```

Usage:

```go
var n mytype.MyType
field := mytype.New("Count", &n, mytype.RuleIsPositive())

v := ward.New()
n = -1
v.Add(field)
v.Run()
v.HasFailures() // true
v.Reset()
```
