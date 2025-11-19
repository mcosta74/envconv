# envconv
![CI Test](https://github.com/mcosta74/envconv/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/mcosta74/envconv.svg)](https://pkg.go.dev/github.com/mcosta74/envconv)

Utility functions to parse environment variables

One scenario where these functions could be useful is when they are used in combination with the `flag` package

```go
var (
    logLevel slog.Level
    timeout time.Duration
    ...
)

func main() {
    flag.TextVar(&logLevel, "log.level", envconv.GetTextUnmarshaler("APP_LOG_LEVEL", slog.LevelInfo), "application log level")
    flag.TextVar(&timeout, "timeout", envconv.GetDuration("APP_TIMEOUT", 30*time.Second), "application timeout")
    ...
}
```