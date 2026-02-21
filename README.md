# Crusade Tracker API

A RESTful API for tracking Warhammer 40k Crusade campaigns. Manage crusades, armies, rosters, battle reports, experience points, and everything needed to run a narrative Crusade campaign.

## Tech Stack

- **Language:** Go
- **Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** PostgreSQL
- **Migrations:** Custom bash migration runner with checksum drift detection

## Project Structure

```
crusadetrackerapi/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── router/
│   │   └── router.go        # Route registration and server startup
│   ├── users/               # User management
│   ├── armies/              # Army tracking
│   ├── factions/            # Faction data
│   └── rosters/             # Roster management
├── database/
│   ├── migrate.sh           # Migration runner script
│   └── sql/                 # SQL migration files
├── go.mod
└── go.sum
```

Each domain module follows a consistent structure:
- `types.go` — data structs and interfaces
- `routes.go` — route registration
- `service.go` — request handlers and business logic

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL

### Setup

1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd crusadetrackerapi
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up your environment variables (copy and edit as needed):
   ```bash
   cp .env.example .env
   ```

   Required variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=crusadetracker
   PORT=8080
   ```

4. Run database migrations:
   ```bash
   ./database/migrate.sh
   ```

   Or with explicit connection options:
   ```bash
   ./database/migrate.sh --host localhost --port 5432 --user postgres --password secret --dbname crusadetracker
   ```

5. Start the server:
   ```bash
   go run cmd/main.go
   ```

## API

All endpoints are versioned under `/api/v1/`.

### Users
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users/` | List all users |

### Armies
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/armies/` | List all armies |

### Factions
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/factions/` | List all factions |

### Rosters
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/rosters/` | List all rosters |

## Database Migrations

Migration files live in `database/sql/` and are named sequentially:

```
database/sql/
├── 001_create_users.sql
├── 002_create_factions.sql
├── 003_create_armies.sql
└── ...
```

The migration runner tracks applied migrations in a `schema_migrations` table and uses SHA256 checksums to detect drift on already-applied migrations.

## Planned Features

- **Crusades** — create and manage Crusade campaigns with participant tracking
- **Armies** — track army supply limits, crusade points, and requisition
- **Rosters** — manage unit rosters with crusade cards
- **Units** — track individual units with experience, battle honours, battle scars, and relics
- **Battle Reports** — log game results, objectives achieved, and casualties
- **Experience & Advancement** — handle XP gain, rank advancement, and stratagems
- **Requisition** — manage requisition point spending and army upgrades

## Development

```bash
# Run the server with live reload (requires air)
air

# Run tests
go test ./...

# Build binary
go build -o crusadetracker cmd/main.go
```
