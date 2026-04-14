# Exercise 04: Testing — Mocks via Interfaces

## Learning Objectives
- Use interfaces to make code testable without real dependencies
- Build a manual mock (test double) that records calls
- Test error paths by controlling mock behavior

## Why This Matters (DevOps context)
Real email sending, AWS API calls, and database writes must not happen in unit tests. Interfaces let you swap real implementations for mocks. This is how Terraform providers, AWS SDK clients, and HTTP handlers are tested.

## Concept
Design around interfaces: `EmailSender` instead of `*SMTPClient`. Tests pass a `MockEmailSender`. Production passes the real SMTP sender. The service doesn't know or care which it gets.

## Your Task
Implement `RegisterUser`:
1. Validate name and email are non-empty
2. Call `s.sender.Send(email, subject, body)` with a welcome message
3. Return any send error

## Run Tests
```bash
go test -v ./phases/phase4-mastery/exercises/04-testing-mocks/
```

## Hints
<details><summary>Hint 1</summary>Validation: `if name == "" || email == "" { return fmt.Errorf("...") }`</details>
<details><summary>Hint 2</summary>Send: `return s.sender.Send(email, "Welcome "+name, "Welcome to our service!")`</details>

## References
- [Go Blog — Dependency Injection](https://go.dev/blog/wire)
- [pkg.go.dev — testing](https://pkg.go.dev/testing)
