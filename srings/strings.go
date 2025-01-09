package lib

import (
	"errors"
	"strconv"
	"strings"
)

// ToFloat32 converts a string to float32.
func ToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// ToFloat64 converts a string to float64.
func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ToInt32 converts a string to int32.
func ToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// ToInt64 converts a string to int64.
func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// ToUint32 converts a string to uint32.
func ToUint32(s string) (uint32, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}

// ToUint64 converts a string to uint64.
func ToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// ToBool converts a string to bool. Supports "true", "false", "1", "0".
func ToBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, errors.New("invalid boolean value")
	}
}

// IsNumeric checks if a string is numeric.
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// TrimSpace removes leading and trailing spaces from a string.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// ToUpperCase converts a string to uppercase.
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ToLowerCase converts a string to lowercase.
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// SplitByDelimiter splits a string by a delimiter.
func SplitByDelimiter(s, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// JoinWithDelimiter joins a slice of strings with a delimiter.
func JoinWithDelimiter(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}

// Contains checks if a substring exists in a string.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// HasPrefix checks if a string starts with a specific prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if a string ends with a specific suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
