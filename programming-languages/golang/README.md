# Mastering Golang project

This is a project to master Golang by building a distributed, fault-tolerant key-value store (like a mini etcd/Redis).

## Stage 1: Naive in-memory store

Build a correct, single-threaded in-memory KV store with a clean API.

### Milestones

- [x] Define a Store interface with Get, Set, Delete and Keys methods
- [x] Implement the Store with a plain `map[string][]byte`
- [x] Write table-driven tests with testing package
- [x] Add a CLI client using cobra or flag package
- [ ] Define and return typed errors

## Stage 2: Concurrent access with sync primitives

Add concurrent read/write safety. Understand Go's memory model, when to use mutexes vs channels, and how to test concurrent code reliably.

### Milestones

- [ ] Wrap map access with `sync.RWMutex` (RLock for reads, Lock for writes)
- [ ] Run tests with `go test -race` and make them pass
- [ ] Add a TTL/expirty feature using a background goroutine and `time.Ticker`
- [ ] Implement a `BatchGet` using goroutines + `sync.WaitGroup`
- [ ] Benchmark with `go test -bench` and profile with `pprof`

### References

- [The Go memory model](https://go.dev/ref/mem)
- [sync package docs](https://pkg.go.dev/sync)
- [go test -race flag](https://go.dev/doc/articles/race_detector)

## Stage 3: Network layer - gRPC + channels

Build a real network protocol. Learn how channels replace callbacks and how context carries deadlines and cancellation across goroutine boundaries.

### Milestones

- [ ] Define a `.proto` schema; generate Go stubs with `protoc`
- [ ] Implement a gRPC server wrapping the store
- [ ] Add a streaming Watch RPC using channels internally
- [ ] Propagate `context.Context` through every call for cancellation
- [ ] Build a HTTP/JSON gateway with `net/http` for human-readable access
- [ ] Write integration tests with an in-process server

### References

- [Go concurrency patterns](https://go.dev/blog/pipelines)
- [gRPC Go quickstart](https://grpc.io/docs/languages/go/quickstart/)
- [context package explainer](https://go.dev/blog/context)

## Stage 4: Persistence -- WAL + SSTable

Make writes durable. Implement a write-ahead log and a basic compacting SSTable engine. Learn Go's file I/O patterns and how real storage engine works.

- [ ] Write a WAL: append-only binary log with CRC checksums
- [ ] Implement crash recovery: replay WAL on startup to rebuild in-memory state
- [ ] Add MemTable -> SSTable flush when MemTable exceeds threshold
- [ ] Implement a basic compaction strategy (merge overlapping SSTables)
- [ ] Use `bufio.Writer` for batched I/O; measure fsync cost
- [ ] Add snapshot/checkpoint to bound WAL replay time

### References

- [Log-Structured Merge Trees](https://www.cs.umb.edu/~poneil/lsmtree.pdf)
- [Building a storage engine in Go](https://arvenil.github.io/blog/storage-engine/)
- [io package deep dive](https://pkg.go.dev/io)

## Stage 5: Raft consensus -- leader election & log replication

Implement the core Raft algorithm. By now channels, goroutines, mutexes and contexts are muscle memory -- Raft's complexity becomes about the algorithm, not Go.

### Milestones

- [ ] Implement Raft state machine: Follower -> Candidate -> Leader transitions
- [ ] Leader election with randomized election timeouts (`time.Timer` + reset)
- [ ] `AppendEntries` RPC for log replication; track `nextIndex` and `matchIndex`
- [ ] Apply committed entries to the KV store state machine
- [ ] Implement log compaction via Raft snapshots (`SnapshotInstall` RPC)
- [ ] Chaos test. kill random nodes, partition network, verify linearizability

### References

- [Raft paper](https://raft.github.io/raft.pdf)
- [The Raft website](https://raft.github.io/)
- [MIT 6.824 Lab 2 (Raft)](https://pdos.csail.mit.edu/6.824/labs/lab-raft.html)
- [etcd raft library](https://github.com/etcd-io/raft)

## Stage 6: Production hardening

Close the gap to a real system. Add dynamic cluster membership, client-side session guarantees, structured observability and operational tooling.

### Milestones

- [ ] Prometheus metrics: op latency histograms, raft log lag, snapshot size
- [ ] Structured logging with `log/slog`
- [ ] Raft joint-consensus for safe cluster membership charges
- [ ] Client sessions with fencing tokens for exactly-once semantics
- [ ] Read leases/lease-based reads to serve reads without Raft round-trip
- [ ] Admin API: cluster status, leader transfer, force-compact
- [ ] Docker Compose multi-node setup; Jepsen-style failure injection script 

### References

- [Design Data-Intensive Applications](https://dataintensive.net/)
- [Jepsen -- testing distributed systems](https://jepsen.io/)
- [log/slog docs](https://pkg.go.dev/log/slog)
- [Promethous Goclient](https://github.com/prometheus/client_golang)
