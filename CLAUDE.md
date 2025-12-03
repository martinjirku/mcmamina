# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**mcmamina** is a community center website for "Materské Centrum Mamina" (Mother's Center) - a Slovak organization providing activities, support groups, and services for mothers and children. The site features an activity calendar (Google Calendar integration), information about children's activities, and an admin portal with Google OAuth authentication.

## Technology Stack

- **Backend**: Go 1.23 with Gorilla Mux routing, html/template templating, PostgreSQL database
- **Frontend**: TypeScript/Vite 5.4 with TailwindCSS 3.4, HTMX for dynamic interactions
- **Authentication**: Google OAuth2 with session-based auth (Gorilla sessions), CSRF protection (nosurf)
- **Package Manager**: pnpm 9.12
- **Development**: Air for hot reload, Task for task running
- **Deployment**: Docker multi-stage builds, Fly.io configuration

## Development Commands

### Primary Development Workflow
```bash
task install    # Install all dependencies (Go modules + pnpm packages)
task watch      # Start development environment (Vite watch + Air hot reload)
task db:start   # Start PostgreSQL container
task db:up      # Run database migrations up
task db:down    # Rollback database migrations
```

### Frontend Build
```bash
pnpm build      # Production build with Vite → dist/ directory
```

### Testing
```bash
go test ./...                    # Run all Go tests
go test ./pkg/calendar -v        # Run specific package tests
```

### Development Flow
1. Run `task watch` which starts:
   - `watch:ts`: Vite in watch mode (rebuilds TypeScript on changes)
   - `watch:server`: Air hot reload (rebuilds Go server on changes)
2. Docker Compose `develop.watch` syncs handler/template/pkg file changes
3. Air watches Go files and templates, auto-rebuilds on modification

## Architecture

### Hybrid SSR + Progressive Enhancement Pattern

The application uses server-side rendering with progressive enhancement:
- **Go Templates**: Primary rendering via `html/template` with Sprout functions
- **TypeScript Modules**: Per-page interactivity (loaded as ES modules)
- **Vite Integration**: Builds TypeScript to `dist/`, which is then embedded in Go binary

### Embedded Asset Distribution

Critical for understanding production builds:
1. Vite builds frontend assets to `dist/` directory
2. Go binary embeds `dist/` via `//go:embed dist` directive (in `main.go`)
3. Production Docker image uses scratch base (minimal, self-contained)
4. CSS paths dynamically resolved via Vite's `manifest.json` for cache-busting

**Implementation**: `main.go` has `CSSPathFromManifest()` function that reads `.vite/manifest.json` to find hashed CSS filenames.

### Service-Oriented Handler Design

Clean separation of concerns:
```
HTTP Handlers (handlers/) → Services (pkg/) → External APIs/Database
```

- **Handlers**: HTTP request/response, template rendering
- **Services**: Business logic (e.g., `pkg/calendar/service.go` fetches/caches Google Calendar events)
- **Interfaces**: Dependency injection via interfaces (EventsGetter, CSSPathGetter, RecaptchaValidator)

**Middleware Stack** (applied in `main.go:191`):
```
RedirectFromWWWSubdomain → Recover → RequestID → Logger → CSRF → Authentication
```

### Caching Strategy

**Google Calendar Events**:
- 30-minute cache with atomic operations (`sync/atomic.Value`)
- Concurrent event fetching with goroutines (20 parallel max)
- Recurring events expanded within date range

**CSS Manifest**:
- Loaded once on startup, cached in memory
- Required for production builds to find hashed asset filenames

### Dual Template Loading

Templates are loaded differently based on environment:
- **Production**: Parsed from `embed.FS` (templates embedded in binary)
- **Development**: Parsed from filesystem (when `--public ./dist` flag is passed to Air)

This enables hot reload in development while maintaining single-binary distribution.

## Critical Configuration

### Environment Variables (.env)

Required for application to function:
- `GOOGLE_API_KEY` - Google Calendar API access
- `GOOGLE_CALENDAR_ID` - Calendar to fetch events from
- `GOOGLE_CAPTCHA_SITE` / `GOOGLE_CAPTCHA_SECRET` - ReCAPTCHA validation
- `GOOGLE_AUTH_CLIENT_ID` / `GOOGLE_AUTH_CLIENT_SECRET` - OAuth credentials
- `GOOGLE_AUTH_REDIRECT_PATH` - OAuth callback URL (e.g., `/auth/google/callback`)
- `POSTGRES_*` - Database connection (HOST, PORT, USER, PASSWORD, DATABASE)
- `SESSION_KEY` - Cookie encryption key (CRITICAL for security, must be 32 bytes)

### Routes

- **Public**: `/`, `/o-nas`, `/aktivity/*`, `/podpora/*`
  - `/aktivity/predporodny-kurz` - Pre-birth course page (email: predporodnykurz.mcmamina@gmail.com)
  - `/aktivity/podporne-skupiny` - Support groups
  - `/aktivity/burzy` - Marketplace
  - `/aktivity/kalendar` - Calendar view
- **Auth**: `/prihlasenie` (login), `/auth/google/callback` (OAuth)
- **Admin**: `/admin` (requires authentication)
- **Utility**: `/healthcheck` (Fly.io health checks), `/panorama` (reverse proxy)

### Database Schema

Migrations in `migrations/` directory (use golang-migrate CLI):
- `google_users`: OAuth user profiles (email, name, picture, google_id)
- `oauth_tokens`: Token storage (currently unused in auth flow)

## Non-Obvious Implementation Details

### CSS Path Resolution

The build process creates hashed CSS filenames (e.g., `style-D6JX9lqD.css`) for cache-busting. To reference these in templates:

1. Vite generates `.vite/manifest.json` with original → hashed filename mappings
2. `CSSPathFromManifest()` in `main.go` reads manifest at startup
3. Templates receive CSS path via `cssPath` function injected into template data

### Recurring Event Expansion

`pkg/calendar/service.go` fetches Google Calendar events and expands recurring instances:
- Uses goroutines to fetch multiple event instances in parallel (max 20 concurrent)
- Caches all expanded instances for 30 minutes
- Atomic swap pattern prevents race conditions during cache updates

### Activity Carousel

`assets/ts/modules/activityManager.ts` uses TypeScript Proxy pattern:
- Manages active activity state across multiple activity cards
- Auto-rotates every 30 seconds
- Ensures only one activity is active at a time

### Air Configuration Quirk

`.air.toml` includes post-command to kill orphaned processes:
```toml
post_cmd = ["lsof -ti:8080 | xargs kill -9 || true"]
```
This prevents "address already in use" errors after rebuilds.

### WWW Subdomain Redirect

`pkg/middleware/redirect.go` implements automatic redirect from www subdomain:
- Middleware checks if `r.Host == "www.mcmamina.sk"`
- Redirects with HTTP 301 (permanent) to `https://mcmamina.sk`
- Applied as first middleware in the stack (see `main.go:191`)
- Important for SEO and canonical URL consistency

### Docker Multi-Stage Build

`Dockerfile` has multiple targets:
- `watch-base`: Development dependencies (Go + Node + Task + Air)
- `watch`: Development with hot reload
- `be-builder`: Production build stage
- Final stage: `FROM scratch` (minimal 20MB image with only binary + embedded assets)

## Deployment

The project uses GitHub Actions for CI/CD (`.github/workflows/deploy.yml`):

- **Main branch** → Production deployment to `mcmamina` app on Fly.io (Warsaw region)
- **Dev branch** → Staging deployment to `mcmamina-staging` app on Fly.io
- Both deploy on push to respective branches
- Uses remote-only builds with commit SHA as build arg

**Production URL**: https://mcmamina.sk (hosted on Fly.io)

## Testing Notes

- Unit tests exist for calendar and mail services (`pkg/calendar/service_test.go`, `pkg/mail/mail_test.go`)
- No test commands in Taskfile, use standard `go test` commands
- Mock interfaces used for dependency injection (e.g., `MockEventsGetter` in tests)
