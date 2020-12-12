// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Cockroach is a postgres migration driver.
type Cockroach struct {
	Postgres
	dialect.Driver
	version string
}

// init loads the Cockroach version from the database for later use in the migration process.
func (d *Cockroach) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW server_version", []interface{}{}, rows); err != nil {
		return fmt.Errorf("querying server version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return fmt.Errorf("server_version variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("scanning version: %v", err)
	}

	d.version = version
	if compareVersions(d.version, "9.0.0") == -1 {
		return fmt.Errorf("unsupported cockroach version: %s", d.version)
	}
	return nil
}

/*
// scanColumn scans the information a column from column description.
func (d *Postgres) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		nullable sql.NullString
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &defaults); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	switch c.typ {
	case "boolean":
		c.Type = field.TypeBool
	case "smallint":
		c.Type = field.TypeInt16
	case "integer":
		c.Type = field.TypeInt32
	case "bigint":
		c.Type = field.TypeInt64
	case "real":
		c.Type = field.TypeFloat32
	case "numeric", "decimal", "double precision":
		c.Type = field.TypeFloat64
	case "text":
		c.Type = field.TypeString
		c.Size = maxCharSize + 1
	case "character", "character varying":
		c.Type = field.TypeString
	case "date", "time", "timestamp", "timestamp with time zone", "timestamp without time zone":
		c.Type = field.TypeTime
	case "bytea":
		c.Type = field.TypeBytes
	case "jsonb":
		c.Type = field.TypeJSON
	case "uuid":
		c.Type = field.TypeUUID
	case "cidr", "inet", "macaddr", "macaddr8":
		c.Type = field.TypeOther
	}
	switch {
	case !defaults.Valid || c.Type == field.TypeTime || seqfunc(defaults.String):
		return nil
	case strings.Contains(defaults.String, "::"):
		parts := strings.Split(defaults.String, "::")
		defaults.String = strings.Trim(parts[0], "'")
		fallthrough
	default:
		return c.ScanDefault(defaults.String)
	}
}
*/

/*
// tBuilder returns the TableBuilder for the given table.
func (d *Postgres) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.Dialect(dialect.Postgres).
		CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(d.addColumn(c))
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	return b
}
*/

// addColumn returns the ColumnBuilder for adding the given column to a table.
func (d *Cockroach) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Dialect(dialect.Postgres).
		Column(c.Name).Type(d.cType(c)).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("SERIAL")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}
