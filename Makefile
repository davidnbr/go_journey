.PHONY: test-all test-phase check progress

test-all:
	go test ./...

test-race:
	go test -race ./...

test-phase:
	go test ./phases/$(PHASE)/...

check:
	go test ./phases/$(PHASE)/exercises/$(EX)/...

progress:
	@go test ./... 2>&1 | grep -E "^(ok|FAIL|---)" | sort
