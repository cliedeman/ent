// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent/dialect/sql/schema/internal"
	"text/template"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -mode=420 -modtime=1 ./template/...

func initTemplates() *template.Template {
	t := template.New("schema")

	schemaTemplate := internal.MustAsset("template/schema.tmpl")
	fieldTemplate := internal.MustAsset("template/field.tmpl")

	t = template.Must(t.Parse(string(schemaTemplate)))
	t = template.Must(t.Parse(string(fieldTemplate)))

	return t
}
