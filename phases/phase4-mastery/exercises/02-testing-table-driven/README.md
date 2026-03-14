# Exercise 02: Testing — Table-Driven Tests

## Learning Objectives
- Write table-driven tests (the idiomatic Go testing pattern)
- Use named test cases for clear failure messages
- Test both success and error paths

## Why This Matters (DevOps context)
Table-driven tests are the standard in Go open source projects (Go standard library uses them extensively). Adding a new test case = adding one struct literal. Named cases make CI failure messages immediately actionable.

## Concept
Define a slice of test cases with inputs, expected outputs, and a `wantErr bool`. Range over them with `t.Run(tt.name, ...)`. This produces `TestFoo/positive`, `TestFoo/zero_divisor` subtest names in output.

## Your Task
The test file is already written — implement `ParseInt` and `SafeDivide` to make all cases pass.

## Run Tests
```bash
go test -v ./phases/phase4-mastery/exercises/02-testing-table-driven/
```

## Hints
<details><summary>Hint 1</summary>ParseInt: `v, err := strconv.Atoi(s); return v, err`</details>
<details><summary>Hint 2</summary>SafeDivide: `if b == 0 { return 0, fmt.Errorf("division by zero") }; return a/b, nil`</details>

## References
- [Go Wiki — TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
