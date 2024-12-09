package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kaashmonee/go-pg-sqlc-crud/crud"
	pgquery "github.com/pganalyze/pg_query_go/v6"
)

/*
sqlc-crud - A tool to generate SQLC-compatible CRUD operations from PostgreSQL schemas

Usage:
  sqlc-crud generate [flags]

Flags:
  -schema string   Path to the PostgreSQL schema dump file (required)
                  Example: ./schema.sql or /path/to/schema.sql

  -output string   Path where the generated CRUD file should be written (required)
                  Example: ./generated/crud.sql or /path/to/output/crud.sql

Example:
  sqlc-crud generate -schema ./schema.sql -output ./generated/crud.sql

Notes:
  - The schema file should be a valid PostgreSQL schema dump
  - The tool will create any necessary directories in the output path
  - Existing output files will be overwritten
  - Output files are created with 0644 permissions (rw-r--r--)
*/

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	// Handle subcommands
	switch os.Args[1] {
	case "generate":
		generateCmd(os.Args[2:])
	case "help", "-h", "--help":
		printUsageAndExit()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsageAndExit()
	}
}

func generateCmd(args []string) {
	// Create a new FlagSet for the generate subcommand
	generateFlags := flag.NewFlagSet("generate", flag.ExitOnError)
	schemaPath := generateFlags.String("schema", "", "Path to the PostgreSQL schema dump file")
	outputPath := generateFlags.String("output", "", "Path where the generated CRUD file should be written")

	if err := generateFlags.Parse(args); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *schemaPath == "" || *outputPath == "" {
		fmt.Println("Error: Both -schema and -output flags are required")
		fmt.Println("Usage:")
		generateFlags.PrintDefaults()
		os.Exit(1)
	}

	schemaDump, err := os.ReadFile(*schemaPath)
	if err != nil {
		fmt.Printf("Error reading schema file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Building the AST")
	tree, err := pgquery.ParseToJSON(string(schemaDump))
	if err != nil {
		fmt.Printf("Error parsing schema: %v\n", err)
		os.Exit(1)
	}

	crud, err := crud.GenerateCRUD(tree)
	if err != nil {
		fmt.Printf("Error generating CRUD: %v\n", err)
		os.Exit(1)
	}

	outputDir := filepath.Dir(*outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*outputPath, []byte(crud), 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated CRUD operations in %s\n", *outputPath)
}

func printUsageAndExit() {
	fmt.Printf(`sqlc-crud - A tool to generate SQLC-compatible CRUD operations from PostgreSQL schemas

Usage:
  sqlc-crud generate [flags]

Flags:
  -schema string   Path to the PostgreSQL schema dump file (required)
                  Example: ./schema.sql or /path/to/schema.sql
  
  -output string   Path where the generated CRUD file should be written (required)
                  Example: ./generated/crud.sql or /path/to/output/crud.sql

Example:
  sqlc-crud generate -schema ./schema.sql -output ./generated/crud.sql

Notes:
  - The schema file should be a valid PostgreSQL schema dump
  - The tool will create any necessary directories in the output path
  - Existing output files will be overwritten
  - Output files are created with 0644 permissions (rw-r--r--)
`)
	os.Exit(1)
}
