# Agents Guide — Crusade Tracker API

This file provides context for AI agents working on this codebase.

---

## Project Overview

**Crusade Tracker API** is a RESTful backend service for managing Warhammer 40k Crusade campaigns. It is written in Go using the Gin web framework and backed by a PostgreSQL database.

The API serves a companion app (frontend not in this repo) that helps players run narrative Warhammer 40k Crusade campaigns. A Crusade is a persistent, evolving campaign format where armies gain experience, earn battle honours and scars, unlock relics, and change over time based on game results. The API needs to track all of this state across multiple players, armies, and ongoing campaigns.

### Domain Concepts

Understanding these Warhammer 40k Crusade terms is important when working on this codebase:

| Term | Description |
|------|-------------|
| **Crusade** | A campaign shared by a group of players. Has a supply limit and tracks overall crusade points. |
| **Army** | A player's collection of units participating in a Crusade. Has a supply limit (PL or points) and a crusade points total. |
| **Faction** | The Warhammer 40k faction the army belongs to (e.g. Space Marines, Aeldari, Chaos Space Marines). |
| **Roster** | The full list of units available to an army in a Crusade. Units on the roster have Crusade Cards. |
| **Crusade Card** | A record attached to each unit tracking its XP, rank, battle honours, battle scars, and equipment. |
| **Unit** | An individual model or squad on the roster. Gains XP and advances through ranks. |
| **Battle Report** | A record of a completed game: participants, result, agenda achieved, and casualties. |
| **Experience (XP)** | Points earned by units through battle. Accumulate to unlock rank advancements. |
| **Rank** | Unit advancement tier based on XP: Blooded → Proven → Crusader → Respected Hero → Legendary Hero. |
| **Battle Honour** | A bonus ability or stat improvement earned when a unit ranks up. |
| **Battle Scar** | A penalty applied to a unit when it is devastated in battle. Can be removed with requisition. |
| **Requisition (RP)** | Campaign currency spent to reinforce armies, remove scars, add units, or purchase upgrades. |
| **Supply Limit** | The maximum points or power level an army may field. Can be increased with requisition. |
| **Agenda** | An optional in-game objective that awards XP or other bonuses when completed. |

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go |
| HTTP Framework | Gin |
| Database | PostgreSQL |
| Migrations | Custom bash runner (`database/migrate.sh`) |

---

## Project Structure

```
crusadetrackerapi/
├── cmd/
│   └── main.go              # Entry point — initialises DB connection and starts the server
├── internal/
│   ├── router/
│   │   └── router.go        # Registers all route groups; add new modules here
│   ├── users/
│   ├── armies/
│   ├── factions/
│   └── rosters/
├── database/
│   ├── migrate.sh           # Migration runner
│   └── sql/                 # Numbered SQL migration files (001_*.sql, 002_*.sql, ...)
├── go.mod
└── go.sum
```

### Module Pattern

Every domain module under `internal/` follows this layout:

```
internal/<module>/
├── types.go     # Structs, interfaces, and constants for this domain
├── routes.go    # Registers the module's routes onto a Gin RouterGroup
└── service.go   # Handler functions and business logic
```

When adding a new module, create this directory structure and register the route group in `internal/router/router.go`.

---

## Conventions

- **Base path:** All routes are versioned under `/api/v1/`.
- **IDs:** Use UUIDs for all primary keys.
- **Migrations:** Add new SQL files to `database/sql/` with the next sequential prefix (e.g. `004_create_units.sql`). Never edit a migration that has already been applied — add a new one instead.
- **Error responses:** Return JSON with a consistent shape: `{ "error": "message" }`.
- **No global state:** Pass dependencies (e.g. DB connection) explicitly via constructors or closures — do not use package-level vars for shared resources.

---

## Testing

Every piece of meaningful logic must have a corresponding test. Do not implement a feature without writing tests for it.

### What to test

- **Service/handler logic** — test each handler function against expected inputs and outputs. Use `net/http/httptest` and Gin's test utilities to fire requests and assert on the response status code and JSON body.
- **Business logic** — any non-trivial calculation (XP thresholds, rank advancement, requisition cost, supply limit checks) should be unit tested in isolation.
- **Edge cases** — invalid UUIDs, missing required fields, out-of-range values, attempts to exceed supply limits, etc.

### Test file placement

Tests live alongside the code they cover, following Go conventions:

```
internal/<module>/
├── types.go
├── routes.go
├── service.go
└── service_test.go   # tests for this module
```

For shared helpers or middleware, place tests in the same package as the code under test.

### Running tests

```bash
# Run all tests
go test ./...

# Run tests for a single module with verbose output
go test -v ./internal/armies/...

# Run a specific test by name
go test -v -run TestGetAllArmies ./internal/armies/...
```

### Guidelines

- Use the standard library `testing` package. Avoid introducing a test framework unless there is a clear need.
- Keep tests independent — each test should set up and tear down its own state. Do not rely on test execution order.
- Mock the database in unit tests. Depend on an interface, not a concrete `*sql.DB`, so tests can swap in a fake implementation without a running Postgres instance.
- Integration tests that require a real database should be gated behind a build tag (e.g. `//go:build integration`) so they do not run in standard `go test ./...`.
- Test names should read as sentences: `TestCreateArmy_ExceedsSupplyLimit`, `TestGetUnit_NotFound`.

---

## Keeping Docs Up to Date

`agents.md` and `README.md` are living documents. Update them as part of the same work that changes the code — not as an afterthought.

### When to update `agents.md`

- **New module added** — add it to the project structure tree and describe its responsibility.
- **New domain concept introduced** — add it to the Domain Concepts table with a clear description.
- **Convention changes** — if the team agrees to change a pattern (error shape, ID type, route structure), update the Conventions section immediately so agents don't follow the old pattern.
- **Current State** — tick off completed items and add newly planned work to "Up next" whenever a significant milestone is reached.
- **Testing rules change** — if a new test helper, build tag convention, or tool is adopted, reflect it in the Testing section.

### When to update `README.md`

- **New endpoints** — add them to the API table.
- **New environment variables** — add them to the setup env block and keep `.env.example` in sync.
- **New planned features** — add them to the Planned Features list when they are committed to.
- **Features shipped** — remove items from Planned Features once they are live and documented in the API table instead.
- **Dependency or tooling changes** — update the Tech Stack and any `go run` / `go build` commands that change.

### General rule

If a code change would confuse someone reading the docs, update the docs in the same commit. Stale documentation is worse than no documentation.

---

## Current State

The project is in early scaffolding. The routing infrastructure and module layout are in place, but most handlers return placeholder responses and no data models or SQL migrations have been written yet.

**Done:**
- Gin server with versioned route groups
- Module scaffolding for users, armies, factions, and rosters
- PostgreSQL migration runner with checksum drift detection

**Up next:**
- Define structs in each module's `types.go`
- Write SQL migrations to create the schema
- Implement handlers with real database queries
- Add authentication middleware
- Add a crusades module
- Add units and crusade card tracking
- Add battle reports
