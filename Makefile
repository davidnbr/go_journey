.PHONY: test-all test-phase check progress

test-all:
	go test ./...

test-race:
	go test -race ./...

test-phase:
	go test ./phases/$(PHASE)/...

check:
	@if echo "$(EX)" | grep -q '\.'; then \
		P=$$(echo "$(EX)" | cut -d. -f1); \
		N=$$(echo "$(EX)" | cut -d. -f2); \
		DIR=$$(ls -d phases/phase$$P-*/exercises/$$(printf "%02d" $$N)-* 2>/dev/null | head -1); \
		if [ -z "$$DIR" ]; then echo "Exercise $(EX) not found"; exit 1; fi; \
		go test ./$$DIR/...; \
	else \
		go test ./phases/$(PHASE)/exercises/$(EX)/...; \
	fi

progress:
	@go test ./... 2>&1 | grep -E "^(ok|FAIL|---)" | sort
