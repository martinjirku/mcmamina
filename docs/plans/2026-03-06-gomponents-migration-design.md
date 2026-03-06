# Gomponents Migration Design (Big Switch)

Date: 2026-03-06
Status: Approved
Owner: @martinjirku

## 1) Context

The repository currently renders server-side HTML using `html/template` and embedded/parsing flows from `templates/**/*.tmpl`. The goal is to migrate rendering to gomponents with a single big switch.

## 2) Scope and Goals

### In scope
- One-shot migration from `html/template` to gomponents.
- Conversion of **all shared pieces** and page rendering to gomponents:
  - layout shell
  - menu/submenu
  - footer
  - icons
  - card wrappers/content
  - animated logo
  - two-week calendar component
  - page-level content components
- Keep handlers/routes/services/data contracts functionally equivalent.
- Preserve strict HTML parity as acceptance target.

### Out of scope
- UI redesign.
- Behavior changes in business logic/services.
- Routing/middleware/auth flow changes.

## 3) Constraints and Acceptance Criteria

- Migration style: **big switch** (no long-lived dual renderer).
- Dependency: adding `gomponents` to `go.mod` is allowed.
- Acceptance bar: **strict HTML parity** with existing output (snapshot-based verification).

## 4) Considered Approaches

1. **Direct rewrite (selected)**
   - Replace template parsing/execution with gomponents rendering in one branch.
   - Pros: minimal temporary complexity, aligns with big switch.
   - Cons: large diff; requires disciplined parity checks.

2. Compatibility shim (not selected)
   - Temporary abstraction for old/new renderer coexistence.
   - Rejected: adds complexity contrary to big-switch objective.

3. Codegen-assisted conversion (not selected)
   - Script conversion to component skeletons.
   - Rejected: setup/cleanup overhead and parity risk.

## 5) Target Architecture

- Remove template runtime flow (`template.ParseFS`, `Clone`, `ExecuteTemplate` paths).
- Introduce Go component packages (illustrative):
  - `pkg/view/layout`
  - `pkg/view/components`
  - `pkg/view/pages`
- Route handlers keep current model construction, then render component tree directly to `http.ResponseWriter`.
- CSS manifest integration remains; resolved CSS path is passed as explicit component data.

## 6) Component/Data Flow Design

- Shared/template fragments are re-authored as gomponents components.
- Every route page becomes a page component composed from shared components.
- Existing handler data shape is preserved to minimize non-rendering risk.
- Rendering pipeline:
  1. handler prepares model
  2. construct page component
  3. render HTML to response writer

## 7) Error Handling

- If gomponents render fails:
  - log structured error
  - return existing equivalent fallback error response behavior
- Error page is also migrated to gomponents so no template fallback remains.

## 8) Testing and Verification Strategy

- Keep snapshot tests as primary parity guardrail.
- Adapt snapshot harness to gomponents output.
- Re-baseline snapshots once as part of migration, then treat future diffs as regressions unless intentional.
- Add focused tests for high-risk shared components (layout/nav/footer/icons).
- Run full test suite (`go test ./...`) before completion.

## 9) Rollout Plan

- Single PR containing complete renderer migration.
- Remove template assets/loading code and now-dead helper functions.
- Update docs to reflect gomponents-based rendering model.
- Verify no stale template execution code paths remain.

## 10) Risks and Mitigations

- Risk: subtle markup drift in shared wrappers.
  - Mitigation: snapshot coverage across all pages + component-focused tests.
- Risk: runtime render error handling differences.
  - Mitigation: preserve equivalent handler error semantics and assert in tests.
- Risk: large change-set reviewability.
  - Mitigation: structure commits logically (infra/components/pages/cleanup/tests) within one branch.

## 11) Decision

Proceed with **Approach 1 (direct rewrite)** using a **single big switch**, migrating all shared and page components to gomponents with **strict HTML parity** as the acceptance bar.
