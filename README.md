# sqlc-crud

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
sqlc-crud generate -schema ./schema.sql -output ./generated/schema.crud.sql
```

## Flags

- `-schema`: Path to the PostgreSQL schema dump file (required)
- `-output`: Path where the generated CRUD file should be written (required)

## Example

```bash
# Generate CRUD operations from a schema file
sqlc-crud generate -schema ./schema.sql -output ./generated/crud.sql
```

## How It Works

1. Takes a PostgreSQL schema dump as input
2. Uses pg_query_go to parse the schema into an AST
3. Generates SQLC-compatible CRUD queries
4. Writes the queries to the specified output file

## License

MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
