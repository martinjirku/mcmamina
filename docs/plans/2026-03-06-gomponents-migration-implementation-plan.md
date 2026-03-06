# Gomponents Big-Switch Migration Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Replace all `html/template` rendering with gomponents in one switch while preserving strict HTML parity across routes.

**Architecture:** Introduce a dedicated Go view layer (`pkg/view/...`) for shared/layout/page components and switch handlers to render component trees directly to `http.ResponseWriter`. Keep business logic, models, and routing unchanged; only rendering pipeline changes. Snapshot tests remain the primary parity gate.

**Tech Stack:** Go 1.25, gomponents, Gorilla Mux, existing handler/services layer, snapshot tests in `handlers/pages_snapshot_test.go`.

---

### Task 1: Add gomponents dependency + rendering primitive

**Files:**
- Modify: `go.mod`
- Modify: `go.sum`
- Create: `pkg/view/render.go`
- Create: `pkg/view/render_test.go`

**Step 1: Write the failing test**

```go
// pkg/view/render_test.go
func TestRenderHTML_WritesDOCTYPEAndBody(t *testing.T) {
    rr := httptest.NewRecorder()
    err := RenderHTML(rr, g.Text("ok"))
    if err != nil { t.Fatalf("RenderHTML err: %v", err) }
    out := rr.Body.String()
    if !strings.Contains(out, "<!doctype html>") { t.Fatalf("missing doctype") }
    if !strings.Contains(out, "ok") { t.Fatalf("missing body") }
}
```

**Step 2: Run test to verify it fails**

Run: `go test ./pkg/view -run TestRenderHTML_WritesDOCTYPEAndBody -v`
Expected: FAIL (missing function/package).

**Step 3: Write minimal implementation**

```go
// pkg/view/render.go
func RenderHTML(w io.Writer, nodes ...g.Node) error {
    _, _ = io.WriteString(w, "<!doctype html>")
    return g.Render(w, g.Group(nodes))
}
```

**Step 4: Run test to verify it passes**

Run: `go test ./pkg/view -run TestRenderHTML_WritesDOCTYPEAndBody -v`
Expected: PASS.

**Step 5: Add dependency and tidy**

Run: `go get github.com/maragudk/gomponents@latest && go mod tidy`
Expected: `go.mod` and `go.sum` include gomponents.

**Step 6: Commit**

```bash
git add go.mod go.sum pkg/view/render.go pkg/view/render_test.go
git commit -m "feat(view): add gomponents render primitive"
```

---

### Task 2: Build shared atomic components (icons/cards/animated logo)

**Files:**
- Create: `pkg/view/components/icons.go`
- Create: `pkg/view/components/cards.go`
- Create: `pkg/view/components/animated_logo.go`
- Create: `pkg/view/components/components_test.go`

**Step 1: Write failing component tests**

- Assert each component emits expected root selectors/classes currently used in snapshots.
- Assert icon components preserve key SVG attributes (`viewBox`, class passthrough).

**Step 2: Run tests to verify failure**

Run: `go test ./pkg/view/components -run Test -v`
Expected: FAIL.

**Step 3: Implement minimal component code**

- Port HTML from:
  - `templates/components/icons.tmpl`
  - `templates/components/cards.tmpl`
  - `templates/components/animatedlogo.tmpl`
- Keep class names and structural wrappers unchanged.

**Step 4: Re-run tests**

Run: `go test ./pkg/view/components -run Test -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add pkg/view/components/icons.go pkg/view/components/cards.go pkg/view/components/animated_logo.go pkg/view/components/components_test.go
git commit -m "feat(view): port shared atomic components to gomponents"
```

---

### Task 3: Build layout-level components (page shell, nav/footer, submenu)

**Files:**
- Create: `pkg/view/layout/page.go`
- Create: `pkg/view/layout/layout.go`
- Create: `pkg/view/layout/submenu.go`
- Create: `pkg/view/layout/layout_test.go`

**Step 1: Write failing tests for layout skeleton**

- Validate generated markup contains `<head>`, CSS link with provided path, nav container, footer container.
- Validate submenu highlights current section exactly like existing templates.

**Step 2: Run tests to verify failure**

Run: `go test ./pkg/view/layout -run Test -v`
Expected: FAIL.

**Step 3: Implement layout components**

Port from:
- `templates/layout/page.tmpl`
- `templates/layout/layout.tmpl`
- `templates/layout/submenu.tmpl`

**Step 4: Re-run tests**

Run: `go test ./pkg/view/layout -run Test -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add pkg/view/layout/page.go pkg/view/layout/layout.go pkg/view/layout/submenu.go pkg/view/layout/layout_test.go
git commit -m "feat(view): port layout shell and submenu to gomponents"
```

---

### Task 4: Port static page components (about/activity/support family)

**Files:**
- Create: `pkg/view/pages/about_us.go`
- Create: `pkg/view/pages/activity.go`
- Create: `pkg/view/pages/activity_babydelivery.go`
- Create: `pkg/view/pages/activity_marketplace.go`
- Create: `pkg/view/pages/activity_supportgroups.go`
- Create: `pkg/view/pages/activity_calendar.go`
- Create: `pkg/view/pages/support.go`
- Create: `pkg/view/pages/support_taxbonus.go`
- Create: `pkg/view/pages/support_volunteers.go`
- Create: `pkg/view/pages/pages_static_test.go`

**Step 1: Write failing tests for representative pages**

- For each page, assert critical headings/sections/classes present.
- Keep assertions focused on structure present in existing snapshots.

**Step 2: Run tests to verify failure**

Run: `go test ./pkg/view/pages -run TestStatic -v`
Expected: FAIL.

**Step 3: Implement minimal page components**

Port from matching `templates/pages/*.tmpl` files listed above.

**Step 4: Re-run tests**

Run: `go test ./pkg/view/pages -run TestStatic -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add pkg/view/pages/about_us.go pkg/view/pages/activity*.go pkg/view/pages/support*.go pkg/view/pages/pages_static_test.go
git commit -m "feat(view): port static page templates to gomponents"
```

---

### Task 5: Port dynamic pages/components (index, login, admin, error, calendar widget)

**Files:**
- Create: `pkg/view/components/twoweeks_calendar.go`
- Create: `pkg/view/pages/index.go`
- Create: `pkg/view/pages/login.go`
- Create: `pkg/view/pages/admin.go`
- Create: `pkg/view/pages/error.go`
- Create: `pkg/view/pages/pages_dynamic_test.go`

**Step 1: Write failing tests**

- Index test: events/activities blocks render with empty and non-empty data.
- Login test: GET and POST-invalid states include expected messages/recaptcha markup.
- Error/Admin tests: page-specific headline and action links exist.

**Step 2: Run tests to verify failure**

Run: `go test ./pkg/view/pages -run TestDynamic -v`
Expected: FAIL.

**Step 3: Implement dynamic components**

Port from:
- `templates/components/twoweekscal.tmpl`
- `templates/pages/index.tmpl`
- `templates/pages/login.tmpl`
- `templates/pages/admin.tmpl`
- `templates/pages/error.tmpl`

**Step 4: Re-run tests**

Run: `go test ./pkg/view/pages -run TestDynamic -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add pkg/view/components/twoweeks_calendar.go pkg/view/pages/index.go pkg/view/pages/login.go pkg/view/pages/admin.go pkg/view/pages/error.go pkg/view/pages/pages_dynamic_test.go
git commit -m "feat(view): port dynamic page rendering to gomponents"
```

---

### Task 6: Switch handlers from template execution to gomponents render

**Files:**
- Modify: `handlers/Index.go`
- Modify: `handlers/Pages.go`
- Modify: `handlers/SupportedUs.go`
- Modify: `handlers/Login.go`
- Modify: `handlers/Admin.go`

**Step 1: Write failing handler tests (or extend snapshot harness stubs first)**

- Create/extend tests verifying handlers return 200 and HTML body using gomponents path (no template parse dependencies).

**Step 2: Run tests to verify failure**

Run: `go test ./handlers -run TestPageSnapshots -v`
Expected: FAIL (constructor signatures/render path mismatch).

**Step 3: Implement minimal handler refactor**

- Remove `*template.Template` dependencies from handler constructors/structs.
- Build page component from model and call `view.RenderHTML`.
- Preserve status codes and error logging semantics.

**Step 4: Re-run tests**

Run: `go test ./handlers -run TestPageSnapshots -v`
Expected: closer; still may fail until router/bootstrap updated.

**Step 5: Commit**

```bash
git add handlers/Index.go handlers/Pages.go handlers/SupportedUs.go handlers/Login.go handlers/Admin.go
git commit -m "refactor(handlers): render pages with gomponents"
```

---

### Task 7: Update app bootstrap to remove template engine wiring

**Files:**
- Modify: `main.go`

**Step 1: Write failing integration compile check**

Run: `go test ./...`
Expected: FAIL from removed template parameters not yet aligned.

**Step 2: Implement minimal bootstrap changes**

- Remove `html/template` and `sprout` parsing setup.
- Update route wiring to new handler signatures.
- Adjust embed directive as needed (drop `templates/**/*.tmpl` when no longer required at runtime).

**Step 3: Re-run compile/tests**

Run: `go test ./...`
Expected: handlers compile; snapshot test likely still needs harness migration.

**Step 4: Commit**

```bash
git add main.go
git commit -m "refactor(main): remove html/template bootstrap wiring"
```

---

### Task 8: Migrate snapshot suite to gomponents rendering and enforce parity

**Files:**
- Modify: `handlers/pages_snapshot_test.go`
- Update: `handlers/testdata/snapshots/pages/*.snap.html`

**Step 1: Write failing test adjustments**

- Remove template parsing helper usage.
- Instantiate router using new handler constructors.
- Keep route matrix unchanged.

**Step 2: Run snapshot verify mode (expect failures first)**

Run: `go test ./handlers -run TestPageSnapshots -v`
Expected: FAIL with diff output.

**Step 3: Fix markup parity gaps**

- Iterate in `pkg/view/...` until diffs are only intentional/known.

**Step 4: Update snapshots once**

Run: `UPDATE_SNAPSHOTS=1 go test ./handlers -run TestPageSnapshots -v`
Expected: PASS and snapshot files updated.

**Step 5: Verify again without update mode**

Run: `go test ./handlers -run TestPageSnapshots -v`
Expected: PASS.

**Step 6: Commit**

```bash
git add handlers/pages_snapshot_test.go handlers/testdata/snapshots/pages/*.snap.html
git commit -m "test: migrate snapshots to gomponents output"
```

---

### Task 9: Remove legacy template assets and dead code

**Files:**
- Delete: `templates/**/*.tmpl`
- Modify: `tailwind.config.ts` (remove stale template globs if no longer needed)
- Modify: `README.md`
- Modify: `CLAUDE.md`
- Modify: `AGENTS.md` (repo section mentioning templates, if applicable)

**Step 1: Write failing guard test/check**

- Add a small check in tests/docs lint notes (or rely on grep command) to ensure no `ExecuteTemplate`/`ParseFS` usage remains.

**Step 2: Run check to verify failure before cleanup**

Run: `rg -n "ExecuteTemplate|ParseFS\(|html/template|templates/.*\.tmpl" main.go handlers pkg`
Expected: matches found before full cleanup.

**Step 3: Apply cleanup**

- Delete template files.
- Update docs/config references to gomponents.

**Step 4: Re-run guard check**

Run: `rg -n "ExecuteTemplate|ParseFS\(|html/template" main.go handlers pkg`
Expected: no matches for legacy render flow.

**Step 5: Commit**

```bash
git add -A
git commit -m "chore: remove legacy templates and update docs for gomponents"
```

---

### Task 10: Full verification before completion

**Files:**
- No new files (verification only)

**Step 1: Run focused tests**

Run:
- `go test ./pkg/view/... -v`
- `go test ./handlers -run TestPageSnapshots -v`

Expected: PASS.

**Step 2: Run full test suite**

Run: `go test ./...`
Expected: PASS.

**Step 3: Optional runtime smoke (if environment ready)**

Run: `task watch` and manually verify key routes:
- `/`
- `/o-nas`
- `/aktivity/*`
- `/podpora/*`
- `/prihlasenie`
- `/admin` (with auth fixture/session)

**Step 4: Final commit (if verification touched files)**

```bash
git add -A
git commit -m "test: finalize gomponents migration verification"
```

(Only if there are unstaged verification-driven edits.)

