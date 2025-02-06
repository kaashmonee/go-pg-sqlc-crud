package crud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

type AST struct {
	Version int       `json:"version"`
	Stmts   []ASTNode `json:"stmts"`
}

type ASTNode struct {
	Stmt         StmtContent `json:"stmt"`
	StmtLen      int         `json:"stmt_len,omitempty"`
	StmtLocation int         `json:"stmt_location,omitempty"`
}

type StmtContent struct {
	CreateStmt *CreateTableStmt `json:"CreateStmt,omitempty"`
}

type CreateTableStmt struct {
	Relation struct {
		Schemaname     string `json:"schemaname"`
		Relname        string `json:"relname"`
		Inh            bool   `json:"inh"`
		Relpersistence string `json:"relpersistence"`
		Location       int    `json:"location"`
	} `json:"relation"`
	TableElts []TableElt `json:"tableElts"`
}

type TableElt struct {
	ColumnDef *struct {
		Colname  string `json:"colname"`
		TypeName struct {
			Names []struct {
				String struct {
					Str string `json:"sval"`
				} `json:"String"`
			} `json:"names"`
		} `json:"typeName"`
	} `json:"ColumnDef,omitempty"`
}

type TableInfo struct {
	Schema  string
	Name    string
	Columns []string
}

func GenerateCRUD(tree string) (string, error) {
	var ast AST
	if err := json.Unmarshal([]byte(tree), &ast); err != nil {
		return "", fmt.Errorf("failed to parse AST: %v", err)
	}

	tables := make([]TableInfo, 0)
	for _, node := range ast.Stmts {
		if node.Stmt.CreateStmt == nil {
			continue
		}

		table := TableInfo{
			Schema:  node.Stmt.CreateStmt.Relation.Schemaname,
			Name:    node.Stmt.CreateStmt.Relation.Relname,
			Columns: make([]string, 0),
		}

		for _, elt := range node.Stmt.CreateStmt.TableElts {
			if elt.ColumnDef != nil {
				table.Columns = append(table.Columns, elt.ColumnDef.Colname)
			}
		}

		tables = append(tables, table)
	}

	var result bytes.Buffer

	funcMap := template.FuncMap{
		"title": strings.Title,
		"add": func(a, b int) int {
			return a + b
		},
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
	}
	tmpl := template.Must(template.New("crud").Funcs(funcMap).Parse(`
  -- {{ .Schema }}.{{ .Name }} CRUD Operations
  
  -- name: Get{{ .Name | title }}ByID :one
  SELECT {{ range $i, $c := .Columns }}{{ if $i }}, {{ end }}{{ $c }}{{ end }}
  FROM {{ .Schema }}.{{ .Name }}
  WHERE id = $1;
  
  -- name: List{{ .Name | title }}s :many
  SELECT {{ range $i, $c := .Columns }}{{ if $i }}, {{ end }}{{ $c }}{{ end }}
  FROM {{ .Schema }}.{{ .Name }}
  ORDER BY created_at DESC
  LIMIT $1 OFFSET $2;
  
  -- name: Create{{ .Name | title }} :one
  INSERT INTO {{ .Schema }}.{{ .Name }} (
	{{- range $i, $c := .Columns }}
	{{- if ne $c "id" }}
	{{- if ne $c "created_at" }}
	{{- if ne $c "updated_at" }}
	{{ $c }}{{ if not (last $i $.Columns) }},{{ end }}
	{{- end }}
	{{- end }}
	{{- end }}
	{{- end }}
  )
  VALUES (
	{{- $count := 1 }}
	{{- range $i, $c := .Columns }}
	{{- if ne $c "id" }}
	{{- if ne $c "created_at" }}
	{{- if ne $c "updated_at" }}
	${{ $count }}{{ if not (last $i $.Columns) }},{{ end }}
	{{- $count = add $count 1 }}
	{{- end }}
	{{- end }}
	{{- end }}
	{{- end }}
  )
  RETURNING {{ range $i, $c := .Columns }}{{ if $i }}, {{ end }}{{ $c }}{{ end }};
  
  -- name: Update{{ .Name | title }} :one
  UPDATE {{ .Schema }}.{{ .Name }}
  SET
	{{- range $i, $c := .Columns }}
	{{- if ne $c "id" }}
	{{- if ne $c "created_at" }}
	{{ $c }} = ${{ add $i 1 }}{{ if not (last $i $.Columns) }},{{ end }}
	{{- end }}
	{{- end }}
	{{- end }}
  WHERE id = $1
  RETURNING {{ range $i, $c := .Columns }}{{ if $i }}, {{ end }}{{ $c }}{{ end }};
  
  -- name: Delete{{ .Name | title }} :exec
  DELETE FROM {{ .Schema }}.{{ .Name }}
  WHERE id = $1;
  `))

	for _, table := range tables {
		if result.Len() > 0 {
			result.WriteString("\n\n")
		}

		if err := tmpl.Execute(&result, table); err != nil {
			return "", fmt.Errorf("failed to execute template for table %s: %v", table.Name, err)
		}
	}

	return result.String(), nil
}
