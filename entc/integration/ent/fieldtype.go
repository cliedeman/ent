// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/ent/fieldtype"
	"github.com/facebook/ent/entc/integration/ent/role"
	"github.com/facebook/ent/entc/integration/ent/schema"
	"github.com/google/uuid"
)

// FieldType is the model entity for the FieldType schema.
type FieldType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Int holds the value of the "int" field.
	Int int `json:"int,omitempty"`
	// Int8 holds the value of the "int8" field.
	Int8 int8 `json:"int8,omitempty"`
	// Int16 holds the value of the "int16" field.
	Int16 int16 `json:"int16,omitempty"`
	// Int32 holds the value of the "int32" field.
	Int32 int32 `json:"int32,omitempty"`
	// Int64 holds the value of the "int64" field.
	Int64 int64 `json:"int64,omitempty"`
	// OptionalInt holds the value of the "optional_int" field.
	OptionalInt int `json:"optional_int,omitempty"`
	// OptionalInt8 holds the value of the "optional_int8" field.
	OptionalInt8 int8 `json:"optional_int8,omitempty"`
	// OptionalInt16 holds the value of the "optional_int16" field.
	OptionalInt16 int16 `json:"optional_int16,omitempty"`
	// OptionalInt32 holds the value of the "optional_int32" field.
	OptionalInt32 int32 `json:"optional_int32,omitempty"`
	// OptionalInt64 holds the value of the "optional_int64" field.
	OptionalInt64 int64 `json:"optional_int64,omitempty"`
	// NillableInt holds the value of the "nillable_int" field.
	NillableInt *int `json:"nillable_int,omitempty"`
	// NillableInt8 holds the value of the "nillable_int8" field.
	NillableInt8 *int8 `json:"nillable_int8,omitempty"`
	// NillableInt16 holds the value of the "nillable_int16" field.
	NillableInt16 *int16 `json:"nillable_int16,omitempty"`
	// NillableInt32 holds the value of the "nillable_int32" field.
	NillableInt32 *int32 `json:"nillable_int32,omitempty"`
	// NillableInt64 holds the value of the "nillable_int64" field.
	NillableInt64 *int64 `json:"nillable_int64,omitempty"`
	// ValidateOptionalInt32 holds the value of the "validate_optional_int32" field.
	ValidateOptionalInt32 int32 `json:"validate_optional_int32,omitempty"`
	// OptionalUint holds the value of the "optional_uint" field.
	OptionalUint uint `json:"optional_uint,omitempty"`
	// OptionalUint8 holds the value of the "optional_uint8" field.
	OptionalUint8 uint8 `json:"optional_uint8,omitempty"`
	// OptionalUint16 holds the value of the "optional_uint16" field.
	OptionalUint16 uint16 `json:"optional_uint16,omitempty"`
	// OptionalUint32 holds the value of the "optional_uint32" field.
	OptionalUint32 uint32 `json:"optional_uint32,omitempty"`
	// OptionalUint64 holds the value of the "optional_uint64" field.
	OptionalUint64 uint64 `json:"optional_uint64,omitempty"`
	// State holds the value of the "state" field.
	State fieldtype.State `json:"state,omitempty"`
	// OptionalFloat holds the value of the "optional_float" field.
	OptionalFloat float64 `json:"optional_float,omitempty"`
	// OptionalFloat32 holds the value of the "optional_float32" field.
	OptionalFloat32 float32 `json:"optional_float32,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Decimal holds the value of the "decimal" field.
	Decimal float64 `json:"decimal,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir http.Dir `json:"dir,omitempty"`
	// Ndir holds the value of the "ndir" field.
	Ndir *http.Dir `json:"ndir,omitempty"`
	// Str holds the value of the "str" field.
	Str sql.NullString `json:"str,omitempty"`
	// NullStr holds the value of the "null_str" field.
	NullStr *sql.NullString `json:"null_str,omitempty"`
	// Link holds the value of the "link" field.
	Link schema.Link `json:"link,omitempty"`
	// NullLink holds the value of the "null_link" field.
	NullLink *schema.Link `json:"null_link,omitempty"`
	// Active holds the value of the "active" field.
	Active schema.Status `json:"active,omitempty"`
	// NullActive holds the value of the "null_active" field.
	NullActive *schema.Status `json:"null_active,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted sql.NullBool `json:"deleted,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
	// IP holds the value of the "ip" field.
	IP net.IP `json:"ip,omitempty"`
	// NullInt64 holds the value of the "null_int64" field.
	NullInt64 sql.NullInt64 `json:"null_int64,omitempty"`
	// SchemaInt holds the value of the "schema_int" field.
	SchemaInt schema.Int `json:"schema_int,omitempty"`
	// SchemaInt8 holds the value of the "schema_int8" field.
	SchemaInt8 schema.Int8 `json:"schema_int8,omitempty"`
	// SchemaInt64 holds the value of the "schema_int64" field.
	SchemaInt64 schema.Int64 `json:"schema_int64,omitempty"`
	// SchemaFloat holds the value of the "schema_float" field.
	SchemaFloat schema.Float64 `json:"schema_float,omitempty"`
	// SchemaFloat32 holds the value of the "schema_float32" field.
	SchemaFloat32 schema.Float32 `json:"schema_float32,omitempty"`
	// NullFloat holds the value of the "null_float" field.
	NullFloat sql.NullFloat64 `json:"null_float,omitempty"`
	// Role holds the value of the "role" field.
	Role role.Role `json:"role,omitempty"`
	// MAC holds the value of the "mac" field.
	MAC schema.MAC `json:"mac,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID       uuid.UUID `json:"uuid,omitempty"`
	file_field *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FieldType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullInt64{},   // int
		&sql.NullInt64{},   // int8
		&sql.NullInt64{},   // int16
		&sql.NullInt64{},   // int32
		&sql.NullInt64{},   // int64
		&sql.NullInt64{},   // optional_int
		&sql.NullInt64{},   // optional_int8
		&sql.NullInt64{},   // optional_int16
		&sql.NullInt64{},   // optional_int32
		&sql.NullInt64{},   // optional_int64
		&sql.NullInt64{},   // nillable_int
		&sql.NullInt64{},   // nillable_int8
		&sql.NullInt64{},   // nillable_int16
		&sql.NullInt64{},   // nillable_int32
		&sql.NullInt64{},   // nillable_int64
		&sql.NullInt64{},   // validate_optional_int32
		&sql.NullInt64{},   // optional_uint
		&sql.NullInt64{},   // optional_uint8
		&sql.NullInt64{},   // optional_uint16
		&sql.NullInt64{},   // optional_uint32
		&sql.NullInt64{},   // optional_uint64
		&sql.NullString{},  // state
		&sql.NullFloat64{}, // optional_float
		&sql.NullFloat64{}, // optional_float32
		&sql.NullTime{},    // datetime
		&sql.NullFloat64{}, // decimal
		&sql.NullString{},  // dir
		&sql.NullString{},  // ndir
		&sql.NullString{},  // str
		&sql.NullString{},  // null_str
		&schema.Link{},     // link
		&schema.Link{},     // null_link
		&sql.NullBool{},    // active
		&sql.NullBool{},    // null_active
		&sql.NullBool{},    // deleted
		&sql.NullTime{},    // deleted_at
		&[]byte{},          // ip
		&sql.NullInt64{},   // null_int64
		&sql.NullInt64{},   // schema_int
		&sql.NullInt64{},   // schema_int8
		&sql.NullInt64{},   // schema_int64
		&sql.NullFloat64{}, // schema_float
		&sql.NullFloat64{}, // schema_float32
		&sql.NullFloat64{}, // null_float
		&sql.NullString{},  // role
		&schema.MAC{},      // mac
		&uuid.UUID{},       // uuid
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*FieldType) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // file_field
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FieldType fields.
func (ft *FieldType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(fieldtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ft.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field int", values[0])
	} else if value.Valid {
		ft.Int = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field int8", values[1])
	} else if value.Valid {
		ft.Int8 = int8(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field int16", values[2])
	} else if value.Valid {
		ft.Int16 = int16(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field int32", values[3])
	} else if value.Valid {
		ft.Int32 = int32(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field int64", values[4])
	} else if value.Valid {
		ft.Int64 = value.Int64
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_int", values[5])
	} else if value.Valid {
		ft.OptionalInt = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_int8", values[6])
	} else if value.Valid {
		ft.OptionalInt8 = int8(value.Int64)
	}
	if value, ok := values[7].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_int16", values[7])
	} else if value.Valid {
		ft.OptionalInt16 = int16(value.Int64)
	}
	if value, ok := values[8].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_int32", values[8])
	} else if value.Valid {
		ft.OptionalInt32 = int32(value.Int64)
	}
	if value, ok := values[9].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_int64", values[9])
	} else if value.Valid {
		ft.OptionalInt64 = value.Int64
	}
	if value, ok := values[10].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field nillable_int", values[10])
	} else if value.Valid {
		ft.NillableInt = new(int)
		*ft.NillableInt = int(value.Int64)
	}
	if value, ok := values[11].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field nillable_int8", values[11])
	} else if value.Valid {
		ft.NillableInt8 = new(int8)
		*ft.NillableInt8 = int8(value.Int64)
	}
	if value, ok := values[12].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field nillable_int16", values[12])
	} else if value.Valid {
		ft.NillableInt16 = new(int16)
		*ft.NillableInt16 = int16(value.Int64)
	}
	if value, ok := values[13].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field nillable_int32", values[13])
	} else if value.Valid {
		ft.NillableInt32 = new(int32)
		*ft.NillableInt32 = int32(value.Int64)
	}
	if value, ok := values[14].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field nillable_int64", values[14])
	} else if value.Valid {
		ft.NillableInt64 = new(int64)
		*ft.NillableInt64 = value.Int64
	}
	if value, ok := values[15].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field validate_optional_int32", values[15])
	} else if value.Valid {
		ft.ValidateOptionalInt32 = int32(value.Int64)
	}
	if value, ok := values[16].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_uint", values[16])
	} else if value.Valid {
		ft.OptionalUint = uint(value.Int64)
	}
	if value, ok := values[17].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_uint8", values[17])
	} else if value.Valid {
		ft.OptionalUint8 = uint8(value.Int64)
	}
	if value, ok := values[18].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_uint16", values[18])
	} else if value.Valid {
		ft.OptionalUint16 = uint16(value.Int64)
	}
	if value, ok := values[19].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_uint32", values[19])
	} else if value.Valid {
		ft.OptionalUint32 = uint32(value.Int64)
	}
	if value, ok := values[20].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_uint64", values[20])
	} else if value.Valid {
		ft.OptionalUint64 = uint64(value.Int64)
	}
	if value, ok := values[21].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[21])
	} else if value.Valid {
		ft.State = fieldtype.State(value.String)
	}
	if value, ok := values[22].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_float", values[22])
	} else if value.Valid {
		ft.OptionalFloat = value.Float64
	}
	if value, ok := values[23].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field optional_float32", values[23])
	} else if value.Valid {
		ft.OptionalFloat32 = float32(value.Float64)
	}
	if value, ok := values[24].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field datetime", values[24])
	} else if value.Valid {
		ft.Datetime = value.Time
	}
	if value, ok := values[25].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field decimal", values[25])
	} else if value.Valid {
		ft.Decimal = value.Float64
	}
	if value, ok := values[26].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field dir", values[26])
	} else if value.Valid {
		ft.Dir = http.Dir(value.String)
	}
	if value, ok := values[27].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ndir", values[27])
	} else if value.Valid {
		ft.Ndir = new(http.Dir)
		*ft.Ndir = http.Dir(value.String)
	}
	if value, ok := values[28].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field str", values[28])
	} else if value != nil {
		ft.Str = *value
	}
	if value, ok := values[29].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field null_str", values[29])
	} else if value != nil {
		ft.NullStr = value
	}
	if value, ok := values[30].(*schema.Link); !ok {
		return fmt.Errorf("unexpected type %T for field link", values[30])
	} else if value != nil {
		ft.Link = *value
	}
	if value, ok := values[31].(*schema.Link); !ok {
		return fmt.Errorf("unexpected type %T for field null_link", values[31])
	} else if value != nil {
		ft.NullLink = value
	}
	if value, ok := values[32].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field active", values[32])
	} else if value.Valid {
		ft.Active = schema.Status(value.Bool)
	}
	if value, ok := values[33].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field null_active", values[33])
	} else if value.Valid {
		ft.NullActive = new(schema.Status)
		*ft.NullActive = schema.Status(value.Bool)
	}
	if value, ok := values[34].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field deleted", values[34])
	} else if value != nil {
		ft.Deleted = *value
	}
	if value, ok := values[35].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[35])
	} else if value != nil {
		ft.DeletedAt = *value
	}
	if value, ok := values[36].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field ip", values[36])
	} else if value != nil {
		ft.IP = *value
	}
	if value, ok := values[37].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field null_int64", values[37])
	} else if value != nil {
		ft.NullInt64 = *value
	}
	if value, ok := values[38].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field schema_int", values[38])
	} else if value.Valid {
		ft.SchemaInt = schema.Int(value.Int64)
	}
	if value, ok := values[39].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field schema_int8", values[39])
	} else if value.Valid {
		ft.SchemaInt8 = schema.Int8(value.Int64)
	}
	if value, ok := values[40].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field schema_int64", values[40])
	} else if value.Valid {
		ft.SchemaInt64 = schema.Int64(value.Int64)
	}
	if value, ok := values[41].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field schema_float", values[41])
	} else if value.Valid {
		ft.SchemaFloat = schema.Float64(value.Float64)
	}
	if value, ok := values[42].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field schema_float32", values[42])
	} else if value.Valid {
		ft.SchemaFloat32 = schema.Float32(value.Float64)
	}
	if value, ok := values[43].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field null_float", values[43])
	} else if value != nil {
		ft.NullFloat = *value
	}
	if value, ok := values[44].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field role", values[44])
	} else if value.Valid {
		ft.Role = role.Role(value.String)
	}
	if value, ok := values[45].(*schema.MAC); !ok {
		return fmt.Errorf("unexpected type %T for field mac", values[45])
	} else if value != nil {
		ft.MAC = *value
	}
	if value, ok := values[46].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field uuid", values[46])
	} else if value != nil {
		ft.UUID = *value
	}
	values = values[47:]
	if len(values) == len(fieldtype.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field file_field", value)
		} else if value.Valid {
			ft.file_field = new(int)
			*ft.file_field = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this FieldType.
// Note that, you need to call FieldType.Unwrap() before calling this method, if this FieldType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FieldType) Update() *FieldTypeUpdateOne {
	return (&FieldTypeClient{config: ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ft *FieldType) Unwrap() *FieldType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FieldType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FieldType) String() string {
	var builder strings.Builder
	builder.WriteString("FieldType(")
	builder.WriteString(fmt.Sprintf("id=%v", ft.ID))
	builder.WriteString(", int=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int))
	builder.WriteString(", int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int8))
	builder.WriteString(", int16=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int16))
	builder.WriteString(", int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int32))
	builder.WriteString(", int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int64))
	builder.WriteString(", optional_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt))
	builder.WriteString(", optional_int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt8))
	builder.WriteString(", optional_int16=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt16))
	builder.WriteString(", optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt32))
	builder.WriteString(", optional_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt64))
	if v := ft.NillableInt; v != nil {
		builder.WriteString(", nillable_int=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt8; v != nil {
		builder.WriteString(", nillable_int8=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt16; v != nil {
		builder.WriteString(", nillable_int16=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt32; v != nil {
		builder.WriteString(", nillable_int32=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt64; v != nil {
		builder.WriteString(", nillable_int64=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", validate_optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.ValidateOptionalInt32))
	builder.WriteString(", optional_uint=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint))
	builder.WriteString(", optional_uint8=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint8))
	builder.WriteString(", optional_uint16=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint16))
	builder.WriteString(", optional_uint32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint32))
	builder.WriteString(", optional_uint64=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint64))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", ft.State))
	builder.WriteString(", optional_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalFloat))
	builder.WriteString(", optional_float32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalFloat32))
	builder.WriteString(", datetime=")
	builder.WriteString(ft.Datetime.Format(time.ANSIC))
	builder.WriteString(", decimal=")
	builder.WriteString(fmt.Sprintf("%v", ft.Decimal))
	builder.WriteString(", dir=")
	builder.WriteString(fmt.Sprintf("%v", ft.Dir))
	if v := ft.Ndir; v != nil {
		builder.WriteString(", ndir=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", str=")
	builder.WriteString(fmt.Sprintf("%v", ft.Str))
	if v := ft.NullStr; v != nil {
		builder.WriteString(", null_str=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", link=")
	builder.WriteString(fmt.Sprintf("%v", ft.Link))
	if v := ft.NullLink; v != nil {
		builder.WriteString(", null_link=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", ft.Active))
	if v := ft.NullActive; v != nil {
		builder.WriteString(", null_active=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", deleted=")
	builder.WriteString(fmt.Sprintf("%v", ft.Deleted))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ft.DeletedAt))
	builder.WriteString(", ip=")
	builder.WriteString(fmt.Sprintf("%v", ft.IP))
	builder.WriteString(", null_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.NullInt64))
	builder.WriteString(", schema_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt))
	builder.WriteString(", schema_int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt8))
	builder.WriteString(", schema_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt64))
	builder.WriteString(", schema_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaFloat))
	builder.WriteString(", schema_float32=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaFloat32))
	builder.WriteString(", null_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.NullFloat))
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", ft.Role))
	builder.WriteString(", mac=")
	builder.WriteString(fmt.Sprintf("%v", ft.MAC))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", ft.UUID))
	builder.WriteByte(')')
	return builder.String()
}

// FieldTypes is a parsable slice of FieldType.
type FieldTypes []*FieldType

func (ft FieldTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
