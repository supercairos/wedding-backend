package utils

import (
	"os"
	"strconv"

	"gopkg.in/guregu/null.v4"
)

// validID checks if the given string is a valid id.
func ValidID(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
}

func ToIntValue(val null.Int) int {
	if val.Valid {
		return int(val.Int64)
	} else {
		return 0
	}
}

func ToFloatValue(val null.Float) float64 {
	if val.Valid {
		return val.Float64
	} else {
		return 0
	}
}

func ToStringValue(val null.String) string {
	if val.Valid {
		return val.String
	} else {
		return ""
	}
}

func ToBoolValue(val null.Bool) bool {
	if val.Valid {
		return val.Bool
	} else {
		return false
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
