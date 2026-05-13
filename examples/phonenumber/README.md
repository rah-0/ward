# Example — phonenumber

Demonstrates a custom ward type package where **T is a struct**.

The validation target is a struct with multiple fields. Rules receive a `*PhoneNumber` and access any field directly — fully type-safe, no casting.

## Type

```go
type PhoneNumber struct {
    CountryCode string
    Number      string
}
```

## TypeID

```go
const TypeID uint32 = 100
```

## Usage

```go
var phone phonenumber.PhoneNumber

v := ward.New()
phone = phonenumber.PhoneNumber{CountryCode: "+1", Number: "8005550100"}
v.Add(
    phonenumber.New("Phone", &phone,
        phonenumber.RuleHasCountryCode(),
        phonenumber.RuleHasNumber(),
        phonenumber.RuleValidLength(7, 15),
        phonenumber.RuleDigitsOnly(),
    ),
)
v.Run()

if v.HasFailures() {
    for _, f := range v.Failures() {
        fmt.Println(f.FieldName, f.RuleID)
    }
}
v.Reset()
```

## Rules

| ID | Rule | Description |
|---|---|---|
| 2 | `RuleHasCountryCode()` | CountryCode must not be blank |
| 3 | `RuleHasNumber()` | Number must not be blank |
| 4 | `RuleValidLength(min, max)` | Number rune count within range. Arg1=min, Arg2=max on failure |
| 5 | `RuleDigitsOnly()` | Number must contain only digit characters |
