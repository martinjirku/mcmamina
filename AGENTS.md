# Repository Guidelines

## Project Structure & Module Organization
- `main.go` bootstraps routes, middleware, templates, and services.
- `handlers/` contains HTTP handlers.
- `pkg/` contains business logic (`calendar`, `mail`, `middleware`, `models`, `services`).
- `templates/` stores Go templates and page TS entry files.
- `assets/` contains frontend modules/styles built to `dist/` by Vite.
- `migrations/` holds PostgreSQL migrations.
- Tests live beside code as `*_test.go`.

## Build, Test, and Development Commands
- `task install` — install Air + pnpm packages.
- `task watch` — run Vite watch and Go hot reload (Air).
- `pnpm build` — production frontend build to `dist/`.
- `go test ./...` — run all Go tests.
- `task db:start` / `task db:stop` — start/stop local Postgres.
- `task db:up` / `task db:down` — apply/rollback migrations.

## Coding Style & Naming Conventions
- Go: format with `gofmt`; use idiomatic naming (exported `PascalCase`, internal `camelCase`).
- TypeScript: keep strict-safe code (`noImplicitAny` enabled), ES modules, small focused files.
- Keep handlers thin; move reusable logic into `pkg/`.

## Testing Guidelines
- Use Go’s `testing` package.
- Test files: `*_test.go`; test names: `TestXxx`.
- Prefer table-driven tests for service/business logic.
- Run `go test ./...` before opening a PR.

### Snapshot Tests (server-rendered pages)
- Suite location: `handlers/pages_snapshot_test.go`.
- Shared helper: `pkg/testutil/snapshot.go`.
- Snapshot files: `handlers/testdata/snapshots/pages/`.
- Naming: `<METHOD>_<normalized-route>.snap.html` (example: `GET__aktivity__kalendar.snap.html`).
- Verify mode (default): `go test ./handlers -run TestPageSnapshots`.
- Update mode: `UPDATE_SNAPSHOTS=1 go test ./handlers -run TestPageSnapshots`.
- Keep skipped routes documented with TODO + reason in the test table.

## Commit & Pull Request Guidelines
- Use short, imperative commit messages (e.g., `fix: phone number`).
- Keep commits focused; avoid mixing unrelated concerns.
- PRs should include summary, linked issue/task, screenshots for UI changes, and notes on env vars/migrations.

## Security & Configuration Tips
- Never commit secrets from `.env`.
- Required vars include Google OAuth/Calendar keys, DB credentials, and `SESSION_KEY` (32 bytes).
- Review auth/callback URL changes carefully before deploy.
