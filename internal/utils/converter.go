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
			return sql.NullString{
				String: time.Time{}.Format(models.DateFormat),
				Valid:  true,
			}
		} else {
			return sql.NullString{
				String: value.Format(models.DateFormat),
				Valid:  true,
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
	case sql.NullString:
		if value.Valid {
			date, err := time.Parse(models.DateFormat, value.String)
			if err != nil {
				return value.String
			}
			return models.CustomDate{Time: date}
		} else {
			return ""
		}
	case sql.NullInt64:
		if value.Int64 == 0 {
			return false
		} else {
			return true
		}
	default:
		return ""
	}
}
