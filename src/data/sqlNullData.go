package data

import (
    "database/sql"
    "encoding/json"
    "bytes"
)

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

func NewNullInt64(val int64) (NullInt64) {
	return NullInt64{sql.NullInt64{Int64: val, Valid: true}}
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

func NewNullBool(val bool) (NullBool) {
	return NullBool{sql.NullBool{Bool: val, Valid: true}}
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		nb.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return err
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

func NewNullFloat64(val float64) (NullFloat64) {
	return NullFloat64{sql.NullFloat64{Float64: val, Valid: true}}
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		nf.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = true
	return err
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

func NewNullString(val string) (NullString) {
	return NullString{sql.NullString{String: val, Valid: true}}
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		ns.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = true
	return err
}