// package envconv provides utilities for converting environment variable values
package envconv

import (
	"encoding"
	"log/slog"
	"os"
	"strconv"
	"time"
)

// GetBool retrieves the boolean value of the environment variable named by the key.
//
// The conversion follows [strconv.ParseBool] rules. If the variable is not set or cannot be
// converted to a boolean, it returns the provided defaultValue.
func GetBool(name string, defaultValue bool) bool {
	val, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}

	if boolValue, err := strconv.ParseBool(val); err == nil {
		return boolValue
	}
	return defaultValue
}

// GetInt retrieves the integer value of the environment variable named by the key.
//
// The conversion follows [strconv.Atoi] rules. If the variable is not set or cannot be
// converted to an integer, it returns the provided defaultValue.
func GetInt(name string, defaultValue int) int {
	val, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}

	if intValue, err := strconv.Atoi(val); err == nil {
		return intValue
	}
	return defaultValue
}

// GetDuration retrieves the time.Duration value of the environment variable named by the key.
//
// The conversion follows [time.ParseDuration] rules. If the variable is not set or cannot be
// converted to a time.Duration, it returns the provided defaultValue.
func GetDuration(name string, defaultValue time.Duration) time.Duration {
	val, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}

	if durationValue, err := time.ParseDuration(val); err == nil {
		return durationValue
	}
	return defaultValue
}

// GetString retrieves the string value of the environment variable named by the key.
//
// If the variable is not set, it returns the provided defaultValue.
func GetString(name, defaultValue string) string {
	val, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}
	return val
}

// GetSlogLevel retrieves the slog.Level value of the environment variable named by the key.
//
// The conversion follows [slog.Level.UnmarshalText] rules. If the variable is not set or cannot be
// converted to a slog.Level, it returns the provided defaultValue.
//
// Deprecated: use GetTextUnmarshaler instead.
func GetSlogLevel(name string, defaultValue slog.Level) slog.Level {
	return GetTextUnmarshaler(name, defaultValue)
}

// GetTextUnmarshaler retrieves a value of type T that implements encoding.TextUnmarshaler
// from the environment variable named by the key.
//
// If the variable is not set or cannot be converted, it returns the provided defaultValue.
// The type TPtr is a pointer to T that implements encoding.TextUnmarshaler.
//
// Example usage:
//
//	var slogLevel slog.Level = envconv.GetTextUnmarshaler("MY_LOG_LEVEL", slog.LevelInfo)
//	var netIP net.IP = envconv.GetTextUnmarshaler("MY_IP", net.IPv4(192.168, 0, 1))
func GetTextUnmarshaler[T any, TPtr interface {
	*T
	encoding.TextUnmarshaler
}](name string, defaultValue T) T {
	val, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}

	var result T
	ptr := TPtr(&result)

	if err := ptr.UnmarshalText([]byte(val)); err != nil {
		return defaultValue
	}
	return result
}
