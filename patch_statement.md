# Codebase Change Statement: yamux

**Current Codebase:** `zhfreal/yamux` (`master-patched`)  
**Original Upstream Baseline:** `MetaCubeX/yamux` (Branch `master`, Commit `dd5f17c0be4982ec3da70ef7bce61c4641bf7665`)  
**Repository Location:** `./yamux-mine`  
**Status:** Production-Ready, Race-Verified (`go test -race ./...` passing 100%)

---

## 1. Executive Summary

This document records the delta between `zhfreal/yamux @ master-patched` and the upstream baseline `MetaCubeX/yamux @ master` (`dd5f17c0`).

The modifications cover buffer pool bounds safety, `sendLoop()` buffer drain cleanup on session shutdown, benchmark modernization, and online dependency alignment with `zhfreal/sing`.

---

## 2. Bug Fixes & Memory Remediations

### 2.1 Pool Bounds & Capacity Guards (`pool.go`)
* Guarded `poolGet(size)` for `size <= 0` (returning `nil`) and `poolPut(buf)` for `cap(buf) == 0 || cap(buf) > 65536`, preventing invalid or oversized slices from entering `sing/common/buf`'s pool.

### 2.2 Shutdown Channel Buffer Reclamation (`session.go`)
* Updated `Session.sendLoop()`'s `defer` cleanup to drain all queued buffers from `s.sendCh` via `poolPut(b)` after closing `s.sendDoneCh`, eliminating pooled buffer leaks when a session shuts down with pending frames.

---

## 3. Modified Files Matrix

| File Path | Status | Change Category | Summary of Delta Relative to Upstream `master` (`dd5f17c0`) |
| :--- | :--- | :--- | :--- |
| `pool.go` | **Modified** | **Safety / Bugfix** | Added `size <= 0` and `cap(buf) == 0 \|\| cap(buf) > 65536` guards in `poolGet`/`poolPut`. |
| `session.go` | **Modified** | **Memory Leak Fix** | Drains `s.sendCh` through `poolPut` inside `sendLoop()`'s `defer` cleanup. |
| `bench_test.go` | **Modified** | **Testing** | Updated benchmark loops to `for b.Loop()`. |
| `go.mod` / `go.sum` | **Modified** | **Dependencies** | Updated `go 1.26.0`, `golang.org/x/sys v0.50.0`, and replaced `github.com/metacubex/sing` with `github.com/zhfreal/sing v0.5.10-0.20261125005639-d6d53d30c901`. |
