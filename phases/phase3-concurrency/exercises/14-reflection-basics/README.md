# Exercise 14: Reflection — Basics

## Learning Objectives
- Use `reflect.TypeOf` and `reflect.ValueOf`
- Iterate struct fields with `NumField`/`Field`
- Set struct fields dynamically with `FieldByName` + `Set`

## Why This Matters (DevOps context)
Reflection powers JSON/YAML marshalers, configuration libraries (viper, envconfig), ORM frameworks, and test assertion libraries. Understanding the basics helps you debug and extend these tools.

## Concept
`reflect.TypeOf(v)` returns the type. `reflect.ValueOf(v)` returns the value. For structs: `.NumField()`, `.Field(i)`. To mutate: pass a pointer, call `.Elem()`, then `.FieldByName()`, then `.Set(reflect.ValueOf(newVal))`.

## Your Task
1. `TypeName` — return type string
2. `FieldNames` — iterate exported struct fields
3. `SetField` — dynamically set a struct field by name

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/14-reflection-basics/
```

## Hints
<details><summary>Hint 1</summary>`reflect.TypeOf(v).String()`</details>
<details><summary>Hint 2</summary>FieldNames: `t := reflect.TypeOf(v); if t.Kind() != reflect.Struct { return nil }; for i := 0; i < t.NumField(); i++ { names = append(names, t.Field(i).Name) }`</details>
<details><summary>Hint 3</summary>SetField: `v := reflect.ValueOf(ptr).Elem(); f := v.FieldByName(fieldName); if !f.IsValid() { return error }; f.Set(reflect.ValueOf(value))`</details>

## References
- [Go Blog — The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
- [pkg.go.dev — reflect](https://pkg.go.dev/reflect)
