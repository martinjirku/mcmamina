# Webka Content Refresh Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Update homepage/footer/activity content from `webka.docx`, fix the wrong activities page/template wiring, and restructure the old support-groups section into a weekly `Krúžky a cvičenia` page.

**Architecture:** This change stays within the existing server-rendered page structure. It updates static templates, fixes one handler/template mismatch in `handlers/Pages.go`, and refreshes snapshots to reflect the new content and footer contact details.

**Tech Stack:** Go, Go templates, existing page snapshot tests, Task/go test tooling

---

### Task 1: Fix activities submenu label and handler/template wiring

**Files:**
- Modify: `handlers/Pages.go`
- Test: `handlers/pages_snapshot_test.go`
- Snapshot impact: `handlers/testdata/snapshots/pages/GET__aktivity__podporne-skupiny.snap.html`, `handlers/testdata/snapshots/pages/GET__aktivity.snap.html`, related activity pages with submenu output

**Step 1: Write the failing test expectation**

Update snapshot expectations indirectly by ensuring the route `/aktivity/podporne-skupiny` should render its own page content and the submenu label should read `Krúžky a cvičenia`.

**Step 2: Run focused snapshot test to confirm current failure state**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- Current snapshots pass before changes, but inspection confirms `SupportGroups` uses the wrong template and the submenu label is outdated.

**Step 3: Write minimal implementation**

In `handlers/Pages.go`:
- change submenu label from `Podporné skupiny` to `Krúžky a cvičenia`
- change `SupportGroups` to parse `activity.supportgroups.tmpl` instead of `activity.marketplace.tmpl`
- keep route `/aktivity/podporne-skupiny` unchanged

**Step 4: Run focused snapshot test to capture expected diffs**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL with snapshot diffs for affected activity pages

**Step 5: Commit checkpoint**

```bash
git add handlers/Pages.go
git commit -m "fix: wire activities page to correct template"
```

### Task 2: Update Predpôrodný kurz page content

**Files:**
- Modify: `templates/pages/activity.babydelivery.tmpl`
- Snapshot impact: `handlers/testdata/snapshots/pages/GET__aktivity__predporodny-kurz.snap.html`

**Step 1: Write the failing test expectation**

Use snapshot expectations for:
- updated registration email text
- five-session course description
- 65 € contribution
- dates `11.2.2026`, `8.4.2026`, `13.5.2026`
- Wednesday 17:00–19:00 schedule

**Step 2: Run focused snapshot test to verify current content differs**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL after content edits until snapshots are updated

**Step 3: Write minimal implementation**

In `templates/pages/activity.babydelivery.tmpl`:
- replace the current intro, curriculum, fee, dates, and time/location copy with the approved text
- use the visible email text `predporodnykurz.mcmamina@gmail.com`
- preserve existing template structure/classes unless a small list/paragraph cleanup is needed for readability

**Step 4: Run focused snapshot test**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL only on snapshot diffs caused by the new page content

**Step 5: Commit checkpoint**

```bash
git add templates/pages/activity.babydelivery.tmpl
git commit -m "feat: refresh predporodny kurz content"
```

### Task 3: Replace old support-groups page with weekly `Krúžky a cvičenia`

**Files:**
- Modify: `templates/pages/activity.supportgroups.tmpl`
- Snapshot impact: `handlers/testdata/snapshots/pages/GET__aktivity__podporne-skupiny.snap.html`

**Step 1: Write the failing test expectation**

Expect the page to render:
- title `Krúžky a cvičenia`
- weekday sections for Monday through Thursday
- the new activity descriptions from the DOCX
- no old support-group text
- no stray `Burzy` page content on this route

**Step 2: Run focused snapshot test to verify current content is wrong**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL with snapshot diffs on the weekly-program page

**Step 3: Write minimal implementation**

In `templates/pages/activity.supportgroups.tmpl`:
- replace existing support-group copy entirely
- add structured sections for:
  - `Pondelok`
  - `Utorok`
  - `Streda`
  - `Štvrtok`
- render each activity with a heading line (name + time) and a short paragraph below
- omit the text explicitly marked for deletion in the DOCX

**Step 4: Run focused snapshot test**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL only on snapshot diffs caused by the new page content

**Step 5: Commit checkpoint**

```bash
git add templates/pages/activity.supportgroups.tmpl
git commit -m "feat: replace support groups with weekly activities"
```

### Task 4: Update homepage contribution pricing

**Files:**
- Modify: `templates/pages/index.tmpl`
- Snapshot impact: `handlers/testdata/snapshots/pages/GET__prihlasenie.snap.html`? no
- Snapshot impact: `handlers/testdata/snapshots/pages/GET__*.snap.html` for pages rendering shared homepage only if applicable
- Primary snapshot impact: homepage snapshot if present in suite; verify exact coverage in `handlers/pages_snapshot_test.go`

**Step 1: Write the failing test expectation**

Expect homepage pricing text to show:
- `1 rok a viac` → `4€`
- `10-vstupová pernamentka` → `30€`

**Step 2: Run relevant snapshot test to verify current content differs**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL on homepage snapshot after template edit until snapshots are updated

**Step 3: Write minimal implementation**

In `templates/pages/index.tmpl`:
- change `3€` to `4€`
- change `25€` to `30€`
- leave the rest of the pricing layout unchanged

**Step 4: Run focused snapshot test**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL only on snapshot diffs caused by homepage pricing changes

**Step 5: Commit checkpoint**

```bash
git add templates/pages/index.tmpl
git commit -m "fix: update homepage contribution pricing"
```

### Task 5: Update footer contact email in both template implementations

**Files:**
- Modify: `templates/layout/layout.tmpl`
- Modify: `pkg/view/layout/layout.go`
- Snapshot impact: many page snapshots that include footer contacts

**Step 1: Write the failing test expectation**

Expect footer contact row `Akcie` to render:
- `mailto:akcie.mcmamina@gmail.com`
- visible text `akcie.mcmamina@gmail.com`

**Step 2: Run snapshot test to verify broad shared-layout impact**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL on all snapshots that include the footer

**Step 3: Write minimal implementation**

Update both:
- `templates/layout/layout.tmpl`
- `pkg/view/layout/layout.go`

Change the `Akcie` contact email from `akcie@mcmamina.sk` to `akcie.mcmamina@gmail.com`.

**Step 4: Run snapshot test**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- FAIL only on snapshot diffs reflecting the new footer email plus prior intentional content changes

**Step 5: Commit checkpoint**

```bash
git add templates/layout/layout.tmpl pkg/view/layout/layout.go
git commit -m "fix: update akcie contact email"
```

### Task 6: Refresh snapshots and verify the final result

**Files:**
- Modify: `handlers/testdata/snapshots/pages/*.snap.html` (affected files only)
- Test: `handlers/pages_snapshot_test.go`

**Step 1: Update snapshots**

Run:
```bash
UPDATE_SNAPSHOTS=1 go test ./handlers -run TestPageSnapshots
```
Expected:
- PASS and updated snapshots for affected pages

**Step 2: Run handler snapshot suite again in verify mode**

Run:
```bash
go test ./handlers -run TestPageSnapshots
```
Expected:
- PASS

**Step 3: Run broader verification**

Run:
```bash
go test ./...
```
Expected:
- PASS

**Step 4: Review changed files**

Run:
```bash
git status --short
git diff -- handlers/Pages.go templates/pages/activity.babydelivery.tmpl templates/pages/activity.supportgroups.tmpl templates/pages/index.tmpl templates/layout/layout.tmpl pkg/view/layout/layout.go handlers/testdata/snapshots/pages
```
Expected:
- only the planned files and corresponding snapshots changed

**Step 5: Commit final integration**

```bash
git add handlers/Pages.go templates/pages/activity.babydelivery.tmpl templates/pages/activity.supportgroups.tmpl templates/pages/index.tmpl templates/layout/layout.tmpl pkg/view/layout/layout.go handlers/testdata/snapshots/pages docs/plans/2026-03-11-webka-content-refresh-design.md docs/plans/2026-03-11-webka-content-refresh-implementation-plan.md
git commit -m "feat: refresh website activity content"
```
