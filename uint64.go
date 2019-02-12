package nullx

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

type NullUint64 struct {
	Uint64 uint64
	Valid  bool // Valid is true if Uint64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint64) Scan(value interface{}) error {
	if value == nil {
		n.Uint64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	s := asString(value)
	var err error
	n.Uint64, err = strconv.ParseUint(s, 10, 64)
	return err
}

func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	return fmt.Sprintf("%v", src)
}

// Value implements the driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	//return int64(n.Uint64), nil
	return n.Uint64, nil
}

type Uint64 struct {
	NullUint64
}

// Uint64From creates a new Uint64 that will never be blank.
func Uint64From(u uint64) Uint64 {
	return NewUint64(u, true)
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (s Uint64) ValueOrZero() uint64 {
	if !s.Valid {
		return 0
	}
	return s.Uint64
}

// NewUint64 creates a new Uint64
func NewUint64(u uint64, valid bool) Uint64 {
	return Uint64{
		NullUint64: NullUint64{
			Uint64: u,
			Valid:  valid,
		},
	}
}
