package main

import (
	"fmt"
	"os"

	pgquery "github.com/pganalyze/pg_query_go/v6"

	"github.com/kaashmonee/go-pg-sqlc-crud/crud"
)

// Replace this here with your own schema or modify this code to read from a schema dump file
const examplePgSchemaDump = `
CREATE TABLE public.users (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    email text NOT NULL,
    encrypted_password text NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE TABLE public.todos (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    title text NOT NULL,
    description text NOT NULL,
    is_done boolean NOT NULL DEFAULT false,
    user_id uuid NOT NULL,
    CONSTRAINT todos_pkey PRIMARY KEY (id),
    CONSTRAINT todos_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);

CREATE TABLE public.sessions (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    user_id uuid NOT NULL,
    refresh_token text NOT NULL,
    user_agent text NOT NULL,
    client_ip text NOT NULL,
    is_blocked boolean NOT NULL DEFAULT false,
    expires_at timestamp with time zone NOT NULL,
    CONSTRAINT sessions_pkey PRIMARY KEY (id),
	CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);
`

func main() {
	fmt.Println("building the AST")
	tree, err := pgquery.ParseToJSON(examplePgSchemaDump)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// Debug: Print the raw tree
	fmt.Println("Raw AST:")
	fmt.Println(tree)

	crud, err := crud.GenerateCRUD(tree)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// Debug: Print the CRUD content
	fmt.Println("Generated CRUD content:")
	fmt.Println(crud)

	// Create the generated directory if it doesn't exist
	if err := os.MkdirAll("generated", 0755); err != nil {
		fmt.Println("error creating directory:", err)
		os.Exit(1)
	}

	// Write the CRUD to file
	if err := os.WriteFile("generated/schema.crud.sql", []byte(crud), 0644); err != nil {
		fmt.Println("error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated CRUD operations in generated/schema.crud.sql")
}
