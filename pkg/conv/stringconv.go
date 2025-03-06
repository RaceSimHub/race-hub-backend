package conv

import (
	"database/sql"
	"strconv"
)

type StringConv struct{}

func (StringConv) StringToNullInt(value string) sql.NullInt32 {
	if value == "" {
		return sql.NullInt32{}
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{Int32: int32(number), Valid: true}
}
