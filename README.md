# mcmamina

Go web app for the MCMamina site (server-rendered templates + Vite assets).

## Development

- Install deps: `task install`
- Run dev watchers: `task watch`
- Build frontend: `pnpm build`

## Testing

- Run all tests: `go test ./...`

### Page snapshot tests

Snapshot tests for server-rendered pages are in:

- test suite: `handlers/pages_snapshot_test.go`
- helper: `pkg/testutil/snapshot.go`
- snapshots: `handlers/testdata/snapshots/pages/`

Usage:

- Verify snapshots (default): `go test ./handlers -run TestPageSnapshots`
- Regenerate snapshots: `UPDATE_SNAPSHOTS=1 go test ./handlers -run TestPageSnapshots`

Update snapshots only when intentional HTML/template output changes are made.
