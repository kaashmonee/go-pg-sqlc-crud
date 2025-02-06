# go-pg-sqlc-crud

A CLI tool to generate [SQLC](https://sqlc.dev/)-compatible CRUD operations from PostgreSQL schemas.

## Built With

Standing on the shoulders of giants - most of the code was written with the assistance of Claude 3.5 Sonnet. This tool is made possible by these excellent projects:

### Core Dependencies

- [sqlc](https://github.com/sqlc-dev/sqlc) - The SQL compiler that makes this useful
- [pg_query_go](https://github.com/pganalyze/pg_query_go) - PostgreSQL parser from pganalyze
- [Go 1.23](https://go.dev/) - The Go programming language

### Development Tools

- [Claude 3.5 Sonnet](https://www.anthropic.com/claude) - Primary coding assistant
- [VSCode](https://code.visualstudio.com/) - Code editor
- [continue.dev](https://continue.dev/) - AI coding assistant

Special thanks to all the teams and contributors behind these tools.

## Installation

```bash
go install github.com/kaashmonee/go-pg-sqlc-crud@latest
```

## Usage

```bash
go-pg-sqlc-crud generate -schema ./schema.sql -output ./generated/schema.crud.sql
```

## Flags

- `-schema`: Path to the PostgreSQL schema dump file (required)
- `-output`: Path where the generated CRUD file should be written (required)

## Try it yourself

1. Clone the repo
2. `cd example`
3. `docker compose up`
4. `pg_dump ...`
5. Run `go-pg-sqlc-crud generate -schema ./schema.sql -output ./generated/crud.sql`
6. Copy the `crud.sql` over to the directory you've bootstrapped `sqlc` into
7. Make sure it's included in `sqlc.yaml`
8. Run `sqlc generate`
9. If this runs and succeeds, then congrats, the tool works
10. Please LMK or create a PR if you find bugs! TY

## How It Works

1. Takes a PostgreSQL schema dump as input
2. Uses pg_query_go to parse the schema into an AST
3. Generates SQLC-compatible CRUD queries
4. Writes the queries to the specified output file

## License

MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
