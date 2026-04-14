# Exercise 07: Errors — Custom Types

## Learning Objectives
- Implement the `error` interface on a custom struct
- Return typed errors so callers can use `errors.As`
- Design errors that carry structured context

## Why This Matters (DevOps context)
Terraform providers, AWS SDK, and Kubernetes client-go all return typed errors. `errors.As` lets you extract the type: `var awsErr *awserr.Error; errors.As(err, &awsErr)` to get the AWS error code.

## Concept
Any type with `Error() string` satisfies the `error` interface. Struct errors carry context (field name, resource ID) that callers can inspect without string parsing.

## Your Task
1. Implement `ValidationError.Error()` — format field + message
2. Implement `NotFoundError.Error()` — format resource + id
3. `ValidateAge` — return `*ValidationError` for invalid age
4. `FindUser` — return `*NotFoundError` for missing user

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/07-errors-custom/
```

## Hints
<details><summary>Hint 1</summary>`fmt.Sprintf("field %s: %s", e.Field, e.Message)`</details>
<details><summary>Hint 2</summary>Return the pointer `&ValidationError{...}` so `errors.As` can unwrap it.</details>
<details><summary>Hint 3</summary>For FindUser: `name, ok := users[id]; if !ok { return "", &NotFoundError{...} }`</details>

## References
- [Go Blog — Working with errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [pkg.go.dev — errors.As](https://pkg.go.dev/errors#As)
