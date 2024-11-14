package utils

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo/internal/models"
)

func TestToNullType(t *testing.T) {
	t.Run("Empty string to sql.NullString", func(t *testing.T) {
		expected := sql.NullString{String: "", Valid: false}
		result := ToNullType("")
		assert.Equal(t, expected, result)
	})

	t.Run("Non-empty string to sql.NullString", func(t *testing.T) {
		expected := sql.NullString{String: "test", Valid: true}
		result := ToNullType("test")
		assert.Equal(t, expected, result)
	})

	t.Run("Empty CustomDate to sql.NullString", func(t *testing.T) {
		expected := sql.NullString{String: time.Time{}.Format(models.DateFormat), Valid: true}
		result := ToNullType(models.CustomDate{})
		assert.Equal(t, expected, result)
	})

	t.Run("Non-empty CustomDate to sql.NullString", func(t *testing.T) {
		expected := sql.NullString{String: "2024-01-01 00:00:00", Valid: true}
		result := ToNullType(models.CustomDate{
			Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		})
		assert.Equal(t, expected, result)
	})

	t.Run("True bool to sql.NullInt64", func(t *testing.T) {
		expected := sql.NullInt64{Int64: 1, Valid: true}
		result := ToNullType(true)
		assert.Equal(t, expected, result)
	})

	t.Run("False bool to sql.NullInt64", func(t *testing.T) {
		expected := sql.NullInt64{Int64: 0, Valid: true}
		result := ToNullType(false)
		assert.Equal(t, expected, result)
	})
}

func TestToNormalType(t *testing.T) {
	t.Run("sql.NullString with valid:false to string", func(t *testing.T) {
		expected := ""
		result := ToNormalType(sql.NullString{Valid: false})
		assert.Equal(t, expected, result)
	})

	t.Run("sql.NullString with valid:true to string", func(t *testing.T) {
		expected := "test"
		result := ToNormalType(sql.NullString{String: "test", Valid: true})
		assert.Equal(t, expected, result)
	})

	t.Run("sql.NullString with valid:true to CustomDate", func(t *testing.T) {
		expected := models.CustomDate{
			Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		result := ToNormalType(sql.NullString{String: "2024-01-01 00:00:00", Valid: true})
		assert.Equal(t, expected, result)
	})

	t.Run("sql.NullInt64 0 to false", func(t *testing.T) {
		result := ToNormalType(sql.NullInt64{Int64: 0, Valid: true})
		assert.Equal(t, false, result)
	})

	t.Run("sql.NullInt64 1 to true", func(t *testing.T) {
		result := ToNormalType(sql.NullInt64{Int64: 1, Valid: true})
		assert.Equal(t, true, result)
	})
}
