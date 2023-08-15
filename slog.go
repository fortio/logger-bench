package main

import (
	"os"

	"fortio.org/log"
	"golang.org/x/exp/slog"
)

var slogger *slog.Logger

func SetupSlogLogger() {
	level := new(slog.LevelVar) // Info by default
	if log.GetLogLevel() == log.Debug {
		level.Set(slog.LevelDebug)
	}
	// Ouch... slog jsonhandler doesn't have a hook or config for configuring how is the timestamp being logged.
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if len(groups) == 0 && a.Key == slog.TimeKey {
			return slog.Attr{
				Key:   slog.TimeKey,
				Value: slog.Float64Value(log.TimeToTS(a.Value.Time())),
			}
		}
		return a
	}
	slogger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level, ReplaceAttr: replace}))
}

func SlogLog1(id string, numLogged int64, numExtraNotLogged int) {
	// iterate numCalls time
	for i := int64(1); i <= numLogged; i++ {
		// Not optimized version - otherwise we can mutate the KeyVals
		slogger.Info("A visible log entry with 3 attributes",
			slog.Any("iteration", i),
			slog.String("id", id),
			slog.String("logger", "sloggr"), // same byte lentgh as `fortio`
		)
		for j := 1; j <= numExtraNotLogged; j++ {
			// Not optimized version - otherwise we'd check
			// if log.LogDebug() {...}
			slogger.Debug("Not visible entry with 4 attributes",
				slog.String("id", id),
				slog.Int64("iteration", i),
				slog.Int("sub-iteration", j),
				slog.String("logger", "sloggr"),
			)
		}
	}
}
