# Webka Content Refresh Design

## Goal
Refresh the site content based on `webka.docx`, fix the currently incorrect activities page structure, and keep the changes scoped to existing templates/routes unless a small structural adjustment is needed.

## Approved approach
Approach 2: rename and restructure the existing activities subsection while keeping the current route.

## Scope
- Update the homepage contribution pricing.
- Update the footer contact for `Akcie` to `akcie.mcmamina@gmail.com`.
- Replace the `Predpôrodný kurz` page copy with the new five-session content, fee, instructors, dates, and Wednesday schedule.
- Replace the wrong `Podporné skupiny` page content with a new `Krúžky a cvičenia` weekly-program structure.
- Keep the existing route `/aktivity/podporne-skupiny` for now, but rename the submenu label to `Krúžky a cvičenia`.
- Fix the handler/template mismatch that currently causes `/aktivity/podporne-skupiny` to render marketplace content.

## Current issues found
- `handlers/Pages.go` renders `activity.marketplace.tmpl` in `SupportGroups`, so `/aktivity/podporne-skupiny` currently shows the wrong page.
- Footer contacts are duplicated in both `templates/layout/layout.tmpl` and `pkg/view/layout/layout.go`, so both need to stay aligned.
- The DOCX weekly-program content fits a schedule page, not the current support-groups copy.

## Information architecture
### Activities submenu
- Predpôrodný kurz
- Krúžky a cvičenia
- Burzy
- Kalendár

### Pages
- `Predpôrodný kurz`: updated static informational page with dates and contribution details.
- `Krúžky a cvičenia`: weekday-structured schedule page with activities for Monday through Thursday.
- `Burzy`: keep existing marketplace/burza content.
- `Kalendár`: unchanged.

## Layout/content design
### Homepage
Keep the existing layout and only update values:
- `1 rok a viac` → `4€`
- `10-vstupová pernamentka` → `30€`

### Predpôrodný kurz page
Keep the current card-based layout, but update content structure to:
- intro paragraph with registration email
- course contents paragraph
- contribution paragraph
- start dates list
- time/location paragraph

### Krúžky a cvičenia page
Repurpose `templates/pages/activity.supportgroups.tmpl` into a weekly-program page.
Use a simple, readable structure:
- page title
- repeated sections for `Pondelok`, `Utorok`, `Streda`, `Štvrtok`
- for each item: activity name with time, then a short description below

No new frontend components are required. Existing headings, dividers, and spacing utilities are sufficient.

## Error handling / edge cases
- Since content is static, the main risk is rendering the wrong template or leaving labels inconsistent.
- Keep the route stable to avoid accidental broken links.
- Update snapshots so content regressions are visible in tests.

## Testing strategy
- Run focused page snapshot tests for the affected pages.
- Update snapshots for:
  - homepage
  - `/aktivity/predporodny-kurz`
  - `/aktivity/podporne-skupiny`
  - any other pages affected by footer contact changes
- Optionally run the full handler snapshot suite if quick enough.
