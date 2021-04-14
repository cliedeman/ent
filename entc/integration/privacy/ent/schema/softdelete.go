// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/privacy/ent/privacy"
	"entgo.io/ent/entc/integration/privacy/rule"
	"entgo.io/ent/schema/field"
)

// SoftDelete defines the schema of a SoftDelete.
type SoftDelete struct {
	ent.Schema
}

// Mixin list of schemas to the SoftDelete.
func (SoftDelete) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the SoftDelete.
func (SoftDelete) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Immutable().
			NotEmpty().
			Unique(),
		field.Bool("active").
			Default(true),
	}
}

// Policy of the SoftDelete.
func (SoftDelete) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterInactiveRule(),
		},
	}
}
