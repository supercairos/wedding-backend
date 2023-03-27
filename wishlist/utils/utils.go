package utils

import (
	"strconv"

	"gopkg.in/guregu/null.v4"
)

// validID checks if the given string is a valid id.
func ValidID(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
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
