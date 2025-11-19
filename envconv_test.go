package envconv_test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/mcosta74/envconv"
)

func TestGetBool(t *testing.T) {
	type args struct {
		isSet        bool
		envValue     string
		defaultValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"MissingEnv", args{false, "", true}, true},
		{"InvalidBool", args{true, "aabbcc", true}, true},
		{"Valid_1", args{true, "1", false}, true},
		{"Valid_t", args{true, "t", false}, true},
		{"Valid_T", args{true, "T", false}, true},
		{"Valid_TRUE", args{true, "TRUE", false}, true},
		{"Valid_0", args{true, "0", true}, false},
		{"Valid_f", args{true, "f", true}, false},
		{"Valid_F", args{true, "F", true}, false},
		{"Valid_FALSE", args{true, "FALSE", true}, false},
	}
	for _, tt := range tests {
		envName := "TEST_ENV_INT"
		if tt.args.isSet {
			os.Setenv(envName, tt.args.envValue)
			defer os.Unsetenv(envName)
		} else {
			os.Unsetenv(envName)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := envconv.GetBool(envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		isSet        bool
		envValue     string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"MissingEnv", args{false, "", 42}, 42},
		{"InvalidInt", args{true, "aabbcc", 42}, 42},
		{"ValidInt", args{true, "123", 42}, 123},
	}
	for _, tt := range tests {
		envName := "TEST_ENV_INT"
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.isSet {
				os.Setenv(envName, tt.args.envValue)
				defer os.Unsetenv(envName)
			} else {
				os.Unsetenv(envName)
			}
			if got := envconv.GetInt(envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	type args struct {
		isSet        bool
		envValue     string
		defaultValue time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"MissingEnv", args{false, "", 5 * time.Second}, 5 * time.Second},
		{"InvalidDuration", args{true, "abc", 5 * time.Second}, 5 * time.Second},
		{"ValidDuration", args{true, "2h45m", 5 * time.Second}, 2*time.Hour + 45*time.Minute},
	}
	for _, tt := range tests {
		envName := "TEST_ENV_DURATION"
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.isSet {
				os.Setenv(envName, tt.args.envValue)
				defer os.Unsetenv(envName)
			} else {
				os.Unsetenv(envName)
			}
			if got := envconv.GetDuration(envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		isSet        bool
		envValue     string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"MissingEnv", args{false, "", "HELLO"}, "HELLO"},
		{"EnvSet", args{true, "HI", "HELLO"}, "HI"},
	}
	for _, tt := range tests {
		envName := "TEST_ENV_STRING"
		if tt.args.isSet {
			os.Setenv(envName, tt.args.envValue)
			defer os.Unsetenv(envName)
		} else {
			os.Unsetenv(envName)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := envconv.GetString(envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlogLevel(t *testing.T) {
	type args struct {
		isSet        bool
		envValue     string
		defaultValue slog.Level
	}
	tests := []struct {
		name string
		args args
		want slog.Level
	}{
		{"MissingEnv", args{false, "", slog.LevelInfo}, slog.LevelInfo},
		{"InvalidLevel", args{true, "INVALID", slog.LevelInfo}, slog.LevelInfo},
		{"ValidLevel_Debug", args{true, "DEBUG", slog.LevelInfo}, slog.LevelDebug},
		{"ValidLevel_Info", args{true, "INFO", slog.LevelDebug}, slog.LevelInfo},
		{"ValidLevel_Error", args{true, "ERROR", slog.LevelInfo}, slog.LevelError},
		{"ValidLevel_Warn", args{true, "WARN", slog.LevelInfo}, slog.LevelWarn},
	}
	for _, tt := range tests {
		envName := "TEST_ENV_SLOG_LEVEL"
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.isSet {
				os.Setenv(envName, tt.args.envValue)
				defer os.Unsetenv(envName)
			} else {
				os.Unsetenv(envName)
			}
			if got := envconv.GetSlogLevel(envName, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetSlogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
