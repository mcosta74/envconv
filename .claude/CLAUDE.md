# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test ./... -v

# Run a single test function
go test -run TestGetBool ./...

# Lint
go fmt ./...
go vet ./...
```

## Architecture

This is a single-file Go library (`envconv.go`) with no external dependencies. It provides typed getters for environment variables, each following the same pattern: look up the env var, attempt to parse it, and fall back to a `defaultValue` on any error or when the variable is unset.

**Public API:**
- `GetBool`, `GetInt`, `GetString`, `GetDuration`, `GetFloat64` — typed getters using standard library parsers
- `GetTextUnmarshaler[T, TPtr]` — generic getter for any type whose pointer implements `encoding.TextUnmarshaler`; this is the extension point for custom types (e.g. `slog.Level`, `net.IP`)
- `GetSlogLevel` — deprecated wrapper around `GetTextUnmarshaler`

The generic `GetTextUnmarshaler` uses a two-type-parameter constraint (`T any, TPtr interface{ *T; encoding.TextUnmarshaler }`) to allow value-type return while still calling the pointer-receiver `UnmarshalText`. New typed getters for standard types should follow the existing pattern. New types that implement `encoding.TextUnmarshaler` do not need a dedicated getter — `GetTextUnmarshaler` handles them directly.

Tests live in `envconv_test.go` (external test package `envconv_test`) and use table-driven tests with `os.Setenv`/`os.Unsetenv` per case.
