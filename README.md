# envconv
[![Go Reference](https://pkg.go.dev/badge/mcosta74/envconv.svg)](https://pkg.go.dev/mcosta74/envconv)

Utility functions to parse environment variables

One scenario where these functions could be useful is when they are used in combination with the `flag` package

```go
var (
    logLevel slog.Level
    timeout time.Duration
    ...
)

func main() {
    flag.TextVar(&logLevel, "log.level", envconv.GetSlogLevel("APP_LOG_LEVEL", slog.LevelInfo), "application log level")
    flag.TextVar(&timeout, "timeout", envconv.GetDuration("APP_TIMEOUT", 30*time.Second), "application timeout")
    ...
}
```