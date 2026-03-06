# Snapshot Testing Plan

## Objective

Implement snapshot (golden) tests for every server-rendered HTML page route in the Go web app, with support for verify mode (default) and update mode via `UPDATE_SNAPSHOTS=1`.

## Constraints

- Use Go standard tooling (`go test ./...`).
- Avoid new third-party dependencies unless absolutely necessary.
- Keep implementation incremental and idiomatic.
- Store snapshots under `testdata/snapshots/pages/`.
- One deterministic snapshot file per page route.

## Snapshot Naming Convention

- Format: `<METHOD>_<normalized-route>.snap.html`
- Examples:
  - `GET__calendar.snap.html`
  - `GET__clients__new.snap.html`

## Implementation Steps

### 1) Add Shared Snapshot Helper

Create a shared helper (e.g., `pkg/testutil/snapshot.go` or test helper file) that provides:

- update mode detection using `os.Getenv("UPDATE_SNAPSHOTS") == "1"`
- HTML normalization (at minimum trim surrounding whitespace)
- snapshot read/write utilities
- compare behavior with actionable mismatch output (route, snapshot path, expected/actual snippets)

### 2) Create Page Snapshot Test Suite

Add a dedicated snapshot test file (e.g., `handlers/pages_snapshot_test.go`) that:

- initializes router/app in test mode
- uses table-driven tests for each HTML page route
- sends requests using `httptest`
- validates status code per route
- captures HTML response body
- calls shared snapshot helper to verify or update snapshot

Suggested test table fields:

- `name`
- `method`
- `route`
- `snapshotFile`
- optional request setup for auth/session/query specifics

### 3) Ensure Route Coverage

Include all current server-rendered page routes.
For routes requiring complex setup/auth:

- add fixtures/setup where practical, or
- explicitly skip with TODO + reason documented in test file

### 4) Add Developer Documentation

Document usage in `README.md` (or top-of-test-file comments):

- verify: `go test ./...`
- regenerate: `UPDATE_SNAPSHOTS=1 go test ./...`
- when snapshot updates are appropriate

### 5) Generate Initial Snapshots

Run update mode once to create baseline snapshots:

- `UPDATE_SNAPSHOTS=1 go test ./...`

### 6) Validate in Verify Mode

Run:

- `go test ./...`
  Ensure mismatches fail with clear, actionable output.

## Acceptance Criteria

- Default test run verifies snapshots and fails on real mismatches.
- Update mode rewrites snapshots and passes.
- Snapshot files exist under `testdata/snapshots/pages/`.
- Every existing HTML page has a corresponding snapshot test (or explicit documented skip).
- Failure messages identify route and snapshot path with useful mismatch context.

## Deliverables

- Snapshot helper utility
- Page snapshot test suite
- Generated snapshot files
- Documentation updates
- Summary of changed files, commands run, and any skipped routes with reasons
