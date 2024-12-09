package main

import (
	"fmt"
	"os"

	"github.com/kaashmonee/go-pg-sqlc-crud/util"
	pgquery "github.com/pganalyze/pg_query_go/v6"
)

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
	// takes in postgres schema dump and generates sqlc crud stubs
	// uses an ast to parse the schema dump and generate the stubs

	fmt.Println("building the AST")
	tree, err := pgquery.ParseToJSON(examplePgSchemaDump)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	prettyJson, err := util.JsonStringPrettyPrint([]byte(tree))
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("pretty printed JSON:")
	fmt.Println(string(prettyJson))
}
