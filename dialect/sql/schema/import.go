package schema

import (
	"bytes"
	"context"
	"fmt"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"
	_ "github.com/lib/pq"
	"os"
	"path"
	"strings"
)

// Importer generates field config from an existing database
type Importer struct {
	sqlImporter
	ignoredTables []string
	pkgName       string
	path          string
	// TODO discovery mixins?
}

func NewImporter(d, connectionString, pkgName, path string, opts ...ImporterOption) (*Importer, error) {
	i := &Importer{
		pkgName: pkgName,
		path:    path,
	}
	for _, opt := range opts {
		opt(i)
	}
	switch d {
	case dialect.MySQL:
		// i.sqlImporter = &MySQL{Driver: d}
		return nil, fmt.Errorf("MySQL is not currently impleted")
	case dialect.SQLite:
		// i.sqlImporter = &SQLite{Driver: d}
		return nil, fmt.Errorf("SQLite is not currently impleted")
	case dialect.Postgres:
		driver, err := sql.Open(dialect.Postgres, connectionString)
		if err != nil {
			return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
		}
		i.sqlImporter = &Postgres{Driver: driver}
	default:
		return nil, fmt.Errorf("sql/schema: unsupported dialect %q", d)
	}
	return i, nil
}

func (i *Importer) Close() error {
	return i.sqlImporter.Close()
}

func (i *Importer) Import(ctx context.Context) error {
	tmpl := initTemplates()

	tx, err := i.sqlImporter.Tx(ctx)
	if err != nil {
		return err
	}
	tables, err := i.allTables(ctx, tx, i.ignoredTables)
	if err != nil {
		return rollback(tx, err)
	}

	schemas := make([]importSchema, 0)

	for _, table := range tables {
		schemas = append(schemas, tableToSchema(table, i.pkgName))
	}

	files := make(map[string]bytes.Buffer)

	for _, schema := range schemas {
		var buf bytes.Buffer
		err := tmpl.Execute(&buf, schema)

		if err != nil {
			return fmt.Errorf("failed to generate schema for %s: %v", schema.Name, err)
		}

		files[schema.Filename] = buf
	}

	for name, buf := range files {
		filename := path.Join(i.path, i.pkgName, fmt.Sprintf("%s.go", name))
		f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("failed to open destination file: %s for writting: %w", filename, err)
		}

		_, err = f.Write(buf.Bytes())

		if err != nil {
			return fmt.Errorf("failed to write schema: %v", err)
		}
	}

	// TODO fmt and imports

	_ = tx.Rollback()
	return nil
}

func schemaFilename(name string) string {
	return strings.Replace(strings.ToLower(name), "_", "", -1)
}

func schemaName(name string) string {
	return strings.Replace(strings.ToLower(name), "_", "", -1)
}

// ImporterOption allows for managing importer configuration using functional options.
type ImporterOption func(*Importer)

func WithIgnoredTables(ignoredTables []string) ImporterOption {
	return func(i *Importer) {
		i.ignoredTables = ignoredTables
	}
}

type MixinMatcher interface {
	Match([]Table)
}

type importSchema struct {
	Name        string
	PackageName string
	TableName   string
	Filename    string
	Fields      []importField
}

type importField struct {
	*Column
	Name       string
	StorageKey string
	Optional   bool
}

func (f importField) IsPrimary() bool {
	return f.Column.Key == "PRI"
}

func (f importField) SchemaType() string {
	return fmt.Sprintf("map[string]{dialect.Postgres: %s}", f.Column.SchemaType[dialect.Postgres])
}

func (f importField) HasSchemaType() bool {
	_, ok := f.Column.SchemaType[dialect.Postgres]
	return ok
}

func (f importField) FieldType() string {
	switch f.Column.Type {
	case field.TypeBool:
		return "Bool"
	case field.TypeTime:
		return "Time"
	case field.TypeJSON:
		return "JSON"
	case field.TypeUUID:
		return "UUID"
	case field.TypeBytes:
		return "Bytes"
	case field.TypeEnum:
		return "Enum"
	case field.TypeString:
		return "String"
	case field.TypeInt8:
		return "Int8"
	case field.TypeInt16:
		return "Int16"
	case field.TypeInt32:
		return "Int32"
	case field.TypeInt:
		return "Int"
	case field.TypeInt64:
		return "Int64"
	case field.TypeUint8:
		return "Uint8"
	case field.TypeUint16:
		return "Uint16"
	case field.TypeUint32:
		return "Uint32"
	case field.TypeUint:
		return "Uint"
	case field.TypeUint64:
		return "Uint64"
	case field.TypeFloat32:
		return "Float32"
	case field.TypeFloat64:
		return "Float"
	default:
		return "Invalid"
	}
}

func tableToSchema(table *Table, pkgName string) importSchema {
	fields := make([]importField, 0)

	for _, col := range table.Columns {
		fields = append(fields, columnToField(col))
	}

	return importSchema{
		Name:        strings.Title(camel(table.Name)),
		PackageName: pkgName,
		TableName:   table.Name,
		Filename:    schemaFilename(table.Name),
		Fields:      fields,
	}
}

func columnToField(col *Column) importField {
	// TODO: check if storage key differs
	return importField{
		Column:     col,
		Name:       col.Name,
		StorageKey: "",
		Optional:   false,
	}
}

func camel(s string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range s {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '_' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}
