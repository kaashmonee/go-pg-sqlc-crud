# go-pg-sqlc-crud

goal is to build an AST and generate postgres CRUD statements for each table based on the pg_dump of the database

## current progress so far

1. builds the AST and pretty prints the Postgres AST in color

### note

* to be completely transparent and up-front about attribution, this code is written with the help of AI tools, including but not limited to:
  * continue.dev
  * mistral codestral
  * claude 3.5 sonnet

* the correctness of this tool is only guaranteed to the extent that the code is unit tested
* but please note: NO code is _generated_ with the help of any AI tools: it's consistently deterministically generated
  by building an AST and parsing it. as a result, the code should be syntactically correct and compile as needed
* the code is generated as `schema.crud.sql`
