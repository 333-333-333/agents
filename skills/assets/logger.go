func setupLogger(env string) *slog.Logger {
    var handler slog.Handler

    switch env {
    case "production", "staging":
        handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
            Level: slog.LevelInfo,
        })
    case "development":
        handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
            Level:     slog.LevelDebug,
            AddSource: true,
        })
    default: // local
        handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
            Level:     slog.LevelDebug,
            AddSource: true,
        })
    }

    return slog.New(handler)
}
