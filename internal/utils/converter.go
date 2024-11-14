package utils

import (
	"database/sql"
	"time"
	"todo/internal/models"
)

func ToNullType[T any](v T) interface{} {
	switch value := any(v).(type) {
	case string:
		if value == "" {
			return sql.NullString{
				String: "",
				Valid:  false,
			}
		} else {
			return sql.NullString{
				String: value,
				Valid:  true,
			}
		}
	case models.CustomDate:
		if value.IsZero() {
			return sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			}
		} else {
			return sql.NullTime{
				Time:  value.Time,
				Valid: true,
			}
		}
	case bool:
		if value == false {
			return sql.NullInt64{
				Int64: 0,
				Valid: true,
			}
		} else {
			return sql.NullInt64{
				Int64: 1,
				Valid: true,
			}
		}
	default:
		return nil
	}
}

func ToNormalType[T any](v T) interface{} {
	switch value := any(v).(type) {
	case sql.NullTime:
		if value.Valid {
			return models.CustomDate{Time: value.Time}
		} else {
			return models.CustomDate{}
		}
	case sql.NullString:
		if value.Valid {
			return value.String
		} else {
			return ""
		}
	case sql.NullInt64:
		return value.Int64 != 0
	default:
		return nil
	}
}
