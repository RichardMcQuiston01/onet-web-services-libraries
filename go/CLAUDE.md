# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go client library for the [O*NET Web Services API](https://services.onetcenter.org/reference/start/overview). This is a library module, not a runnable application.

## Commands

```sh
go build ./...          # Build all packages
go test ./...           # Run all tests
go test -run TestName   # Run a single test
go vet ./...            # Static analysis
go mod tidy             # Sync go.mod/go.sum
```

## Design Constraints

These are hard requirements — do not deviate from them:

- **Standard library only** — no third-party dependencies for HTTP requests or JSON parsing. Use `net/http` and `encoding/json`.
- **Thread-safe** — no package-level mutable state. All state lives on the client struct.
- **Consistent surface** — request types, response types, and error types follow the same shape across all endpoints.
- **Error fidelity** — HTTP status codes and error messages must match the O*NET API specification exactly.

## Module

`github.com/RichardMcQuiston01/onet-web-services-go` (Go 1.21, zero external dependencies)

## Architecture

- **`Client`** (`client.go`) — holds `apiKey`, `baseURL`, and `*http.Client`. All state is on the struct; no globals. `do(ctx, path, query, out)` is the single internal HTTP helper: sets `X-API-Key`, decodes JSON, and maps non-2xx responses to `*APIError`.
- **`APIError`** (`error.go`) — `{ StatusCode int; Message string }`. HTTP 429 is returned as-is; callers handle backoff.
- **`PageParams` / `PageMeta`** (`page.go`) — shared across all paginated endpoints.
- **Shared types** (`types.go`) — `Occupation`, `Career`, `Element`, `Score`, `NamedItem`.
- **Endpoint files** — one file per API service group, all methods on `*Client`:
  - `about.go` — `About()`
  - `online.go` — search, list, single occupation
  - `online_summary.go` / `online_details.go` — 16 topic endpoints each; share `occupationResource(ctx, code, kind, topic, out)`
  - `online_browse.go` — BrightOutlook, CareerClusters, Industries, JobFamilies, JobZones, STEM, HotTechnologies
  - `online_crosswalk.go` — 7 crosswalk methods via internal `crosswalk(ctx, name, keyword, page)`
  - `online_tools.go` — JobDuties, Associations, RelatedActivities, SoftSkills, Technology
  - `online_onetdata.go` — 8 O*NET data attributes × 2 (list + detail) via `onetDataList` / `onetDataDetail`
  - `mnm.go` — MNM (English) + MPP (Spanish) + Veterans; shared helpers with prefix param
  - `taxonomy.go` — 6 crosswalk methods via `taxonomyCrosswalk(ctx, from, to, code)`
  - `database.go` — `DatabaseTables`, `DatabaseTableInfo`, `DatabaseRows` with `FilterParam` / `SortParam`

## Code Style

- Standard `gofmt` formatting — run before committing.
- Exported symbols get Go doc comments (`// TypeName does ...`).
- Table-driven tests using `t.Run` subtests.
- Test files live alongside source files (`*_test.go`).
