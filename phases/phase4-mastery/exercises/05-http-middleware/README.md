# Exercise 05: HTTP Middleware Pattern

## Learning Objectives
- Implement the middleware pattern for `net/http`
- Chain middlewares: each wraps the next handler
- Implement logging, panic recovery, and auth as middleware

## Why This Matters (DevOps context)
Every production HTTP service uses middleware: logging (structured logs per request), recovery (prevent one panic from killing the server), auth (JWT/bearer token validation). This pattern is used by Gorilla, Chi, and Gin.

## Concept
A middleware is `func(http.Handler) http.Handler`. `Chain` applies them right-to-left so the first middleware in the list is the outermost wrapper. Inside each middleware: call `next.ServeHTTP(w, r)` to pass control.

## Your Task
1. `Chain` — apply middlewares in order (first = outermost)
2. `Logger` — wrap handler, log method/path/duration
3. `Recovery` — defer recover(), return 500 on panic
4. `Auth` — check "Authorization: Bearer <secret>" header

## Run Tests
```bash
go test ./phases/phase4-mastery/exercises/05-http-middleware/
```

## Hints
<details><summary>Hint 1</summary>Chain (right-to-left): `for i := len(middlewares)-1; i >= 0; i-- { h = middlewares[i](h) }; return h`</details>
<details><summary>Hint 2</summary>Recovery: `defer func() { if r := recover(); r != nil { http.Error(w, "Internal Server Error", 500) } }()`</details>
<details><summary>Hint 3</summary>Auth: `token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "); if token != secret { http.Error(w, "Unauthorized", 401); return }`</details>

## References
- [pkg.go.dev — net/http](https://pkg.go.dev/net/http)
- [Go Blog — Writing HTTP middleware](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246333)
