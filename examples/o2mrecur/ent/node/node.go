// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package node

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"

	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"

	// Table holds the table name of the node in the database.
	Table = "nodes"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "nodes"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "node_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "nodes"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "node_children"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Node type.
var ForeignKeys = []string{
	"node_children",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
