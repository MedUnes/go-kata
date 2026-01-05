# 🥋 Go Katas 🥋

> "I fear not the man who has practiced 10,000 kicks once, but I fear the man who has practiced one kick 10,000 times."
>
> (Bruce Lee)

## What should it be?

- Go is simple to learn, but nuanced to master. The difference between "working code" and "idiomatic code" often lies in details such as safety, memory efficiency, and concurrency control.

- This repository is a collection of **Daily Katas**: small, standalone coding challenges designed to drill specific Go patterns into your muscle memory.

## What should it NOT be?

- This is not intended to teach coding, nor to use Go as a general-purpose learning vehicle. It is not intended to teach Go **in general**.

- The focus should be, as much as possible, on challenging oneself to solve common software engineering problems **the Go way**.

- Several seasoned developers spend years learning and applying best practices in production-grade contexts. When they decide to switch to Go, they often face two challenges:

  - Is there a way to transfer knowledge so that I don’t have to throw away years of experience and start from zero?

  - If yes, which parts should I focus on to recognize mismatches and use them the expected way in the Go ecosystem?

## How to Use This Repo

1. **Pick a Kata:** Navigate to any `XX-kata-yy` folder.
2. **Read the Challenge:** Open the `README.md` inside that folder. It defines the goal, the constraints, and the "idiomatic patterns" you must use.
3. **Solve It:** Initialize a module inside the folder and write your solution.
4. **Reflect:** Compare your solution with the provided "Reference Implementation" (if available) or the core patterns listed.

## Contribution Guidelines

 Please refer to the [CONTRIBUTING](CONTRIBUTING.md) file

## Katas Index (Grouped)

### 01) Context, Cancellation, and Fail-Fast Concurrency

Real-world concurrency patterns that prevent leaks, enforce backpressure, and fail fast under cancellation.

- [01 - The Fail-Fast Data Aggregator](./01-context-cancellation-concurrency/01-concurrent-aggregator)
- [03 - Graceful Shutdown Server](./01-context-cancellation-concurrency/03-graceful-shutdown-server)
- [05 - Context-Aware Error Propagator](./01-context-cancellation-concurrency/05-context-aware-error-propagator)
- [07 - The Rate-Limited Fan-Out Client](./01-context-cancellation-concurrency/07-rate-limited-fanout)
- [09 - The Cache Stampede Shield (singleflight TTL)](./01-context-cancellation-concurrency/09-single-flight-ttl-cache)
- [10 - Worker Pool with Backpressure and errors.Join](./01-context-cancellation-concurrency/10-worker-pool-errors-join)
- [14 - The Leak-Free Scheduler](./01-context-cancellation-concurrency/14-leak-free-scheduler)
- [17 - Context-Aware Channel Sender (No Leaked Producers)](./01-context-cancellation-concurrency/17-context-aware-channel-sender)

---

### 02) Performance, Allocation, and Throughput

Drills focused on memory efficiency, allocation control, and high-throughput data paths.

- [02 - Concurrent Map with Sharded Locks](./02-performance-allocation/02-concurrent-map-with-sharded-locks)
- [04 - Zero-Allocation JSON Parser](./02-performance-allocation/04-zero-allocation-json-parser)
- [11 - NDJSON Stream Reader (Long Lines)](./02-performance-allocation/11-ndjson-stream-reader)
- [12 - sync.Pool Buffer Middleware](./02-performance-allocation/12-sync-pool-buffer-middleware)

---

### 03) HTTP and Middleware Engineering

Idiomatic HTTP client/server patterns, middleware composition, and production hygiene.

- [06 - Interface-Based Middleware Chain](./03-http-middleware/06-interface-based-middleware-chain)
- [16 - HTTP Client Hygiene Wrapper](./03-http-middleware/16-http-client-hygiene)

---

### 04) Errors: Semantics, Wrapping, and Edge Cases

Modern Go error handling: retries, cleanup, wrapping, and infamous pitfalls.

- [08 - Retry Policy That Respects Context](./04-errors-semantics/08-retry-backoff-policy)
- [19 - The Cleanup Chain (defer + LIFO + Error Preservation)](./04-errors-semantics/19-defer-cleanup-chain)
- [20 - The “nil != nil” Interface Trap (Typed nil Errors)](./04-errors-semantics/20-nil-interface-gotcha)

---

### 05) Filesystems, Packaging, and Deployment Ergonomics

Portable binaries, testable filesystem code, and dev/prod parity.

- [13 - Filesystem-Agnostic Config Loader (io/fs)](./05-filesystems-packaging/13-iofs-config-loader)
- [18 - embed.FS Dev/Prod Switch](./05-filesystems-packaging/18-embedfs-dev-prod-switch)

---

### 06) Testing and Quality Gates

Idiomatic Go testing: table-driven tests, parallelism, and fuzzing.

- [15 - Go Test Harness (Subtests, Parallel, Fuzz)](./06-testing-quality/15-testing-parallel-fuzz-harness)
