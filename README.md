# go_journey

A structured, test-driven Go course — novice to production-ready. Built on learning science: spaced repetition, active recall, and Bloom's Taxonomy scaffolding.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/davidnbr/go_journey)
[![CI](https://github.com/davidnbr/go_journey/actions/workflows/ci.yml/badge.svg)](https://github.com/davidnbr/go_journey/actions/workflows/ci.yml)

---

## How It Works

Each exercise contains:
- **`exercise.go`** — a stub with `TODO` markers for you to fill in
- **`exercise_test.go`** — failing tests; your goal is to make them pass
- **`README.md`** — objectives, concept explanation, hints (spoiler-gated), and references

**The workflow**: read the exercise README → attempt the stub → run tests → reveal hints only if stuck.

---

## Quick Start

**Option 1 — GitHub Codespaces (zero setup)**

Click the badge above. Everything is pre-configured.

**Option 2 — Local**

```bash
git clone https://github.com/davidnbr/go_journey.git
cd go_journey
go mod download
```

Requires Go 1.22+.

---

## Run Tests

```bash
# All exercises
make test-all

# One phase
make test-phase PHASE=phase1-foundations

# One exercise (shorthand: phase.exercise)
make check EX=1.1

# One exercise (explicit)
make check PHASE=phase1-foundations EX=01-hello-world

# Check overall progress
make progress
```

---

## Course Structure

| Phase | Focus | Bloom Level | Exercises | Projects |
|-------|-------|-------------|-----------|---------|
| [Phase 1 — Foundations](phases/phase1-foundations/README.md) | Syntax, types, control flow, errors | Remember → Apply | 16 | 2 |
| [Phase 2 — Core Patterns](phases/phase2-core-patterns/README.md) | Structs, interfaces, pointers, I/O | Apply → Analyze | 18 | 3 |
| [Phase 3 — Concurrency](phases/phase3-concurrency/README.md) | Goroutines, channels, sync, generics | Analyze → Evaluate | 15 | 3 |
| [Phase 4 — Mastery](phases/phase4-mastery/README.md) | Testing, profiling, HTTP, databases | Evaluate → Create | 10 | 2 |

---

## Solutions

Solutions live on the `solutions` branch to avoid spoilers. Switch only after you've made a genuine attempt.

```bash
git fetch origin solutions
git show origin/solutions:phases/phase1-foundations/exercises/01-hello-world/exercise.go
```

---

## DevOps Context

Every exercise includes a "Why This Matters" section connecting the concept to real infrastructure work: CLI tools, HTTP services, configuration parsing, concurrency patterns for monitoring agents, and more.
