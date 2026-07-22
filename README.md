# Taskman – CLI Task Manager

> 🎓 **Learning Project** — This repository exists purely to familiarise myself with **GORM**, **Cobra**, and **Viper** in Go.

A lightweight CLI task manager that stores tasks in a local SQLite database. Add, list, complete, and delete tasks right from the terminal.

---

## Tech Stack

| Library | Role |
|---------|------|
| [Cobra](https://github.com/spf13/cobra) | CLI command framework (subcommands, flags, argument validation) |
| [Viper](https://github.com/spf13/viper) | Configuration management (reads `taskman.yaml`) |
| [GORM](https://gorm.io/) | ORM for database operations (SQLite via `glebarez/sqlite`) |

---

## Project Structure

```
Cli_App/
├── main.go                     # Application entrypoint
├── taskman.yaml                # App configuration (Viper)
├── go.mod / go.sum             # Go module files
│
├── allvar/
│   └── variables.go            # Shared global variables (flags)
│
├── cmd/
│   ├── root.go                 # Root Cobra command & flag registration
│   └── process/
│       ├── add_proc.go         # `add` subcommand
│       ├── list_proc.go        # `list` subcommand
│       ├── complete_proc.go    # `complete` subcommand
│       └── delete_proc.go      # `delete` subcommand
│
├── config/
│   ├── app_config.go           # Viper config setup
│   └── database.go             # GORM database connection
│
└── database/
    ├── model/
    │   ├── task_model.go       # Task GORM model
    │   └── user_model.go       # User GORM model
    └── taskman.sqlite3         # SQLite database file
```

---

## Configuration

The app reads its settings from `taskman.yaml` via Viper:

```yaml
storage:
  driver: sqlite
  path: ./database/taskman.sqlite3
  auto_migrate: true

behavior:
  hide_done_after_day: 4        # Hide completed tasks after N days
  task_limit: 10                # Max number of tasks allowed

logging:
  gorm_log_level: silent        # silent | error | warn | info
```

---

## Usage

### Build & Run

```bash
go build -o taskman .
./taskman <command> [flags]
```

### Commands

#### `add` — Add a new task

```bash
./taskman add "Write unit tests" -p high
./taskman add "Read GORM docs"              # defaults to medium priority
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--priority` | `-p` | `medium` | Task priority (`high`, `medium`, `low`) |

#### `list` — List tasks

```bash
./taskman list                  # shows pending tasks
./taskman list -a               # includes completed tasks
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--all` | `-a` | `false` | Show completed tasks too |

#### `complete` — Mark a task as done

```bash
./taskman complete 3            # marks task #3 as completed
```

#### `delete` — Delete tasks

```bash
./taskman delete 1 2 5                  # delete by ID(s)
./taskman delete --status SELESAI       # delete all completed tasks
./taskman delete --status PENDING       # delete all pending tasks
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--status` | `-s` | `""` | Delete by status (`SELESAI` or `PENDING`) |

---

## Database Models

### Task

| Column | Type | Notes |
|--------|------|-------|
| `ID` | `integer` | Primary key, auto-increment |
| `Description` | `text` | Not null |
| `Priority` | `text` | `high`, `medium`, or `low` |
| `Completed` | `numeric` | `0` = pending, `1` = done |
| `CreatedAt` | `integer` | Unix timestamp |
| `CompletedAt` | `integer` | Unix timestamp (set on completion) |
| `UserID` | `byte` | Foreign key → `User` |

### User

| Column | Type | Notes |
|--------|------|-------|
| `ID` | `byte` | Primary key |
| `Name` | `text` | Not null |
| `Email` | `string` | Nullable |
| `Tasks` | `[]Task` | Has-many relationship |

---

## Prerequisites

- **Go** ≥ 1.26

```bash
go mod tidy
```

---

## What I Learned

- Structuring a CLI app with **Cobra** (root command, subcommands, flags, argument validation).
- Loading and reading YAML config with **Viper** (`SetConfigFile`, `ReadInConfig`, `GetString`, `GetInt`).
- Defining models and performing CRUD operations with **GORM** (`Create`, `Find`, `Updates`, `Delete`, `Where`).
- Using **SQLite** as a lightweight embedded database via `glebarez/sqlite`.
- Organising a Go project with a clean package layout (`cmd/`, `config/`, `database/model/`, `allvar/`).