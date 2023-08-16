package main

import (
	"fortio.org/log"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func SetupZapLogger() {
	cfg := zap.NewProductionConfig()
	// Important: turn off sampling!
	cfg.Sampling = nil
	if log.GetLogLevel() == log.Debug {
		cfg.Level.SetLevel(zap.DebugLevel)
	}
	zlog, _ = cfg.Build(zap.WithCaller(false))
}

func ZapLog1(id string, numLogged int64, numExtraNotLogged int) {
	// iterate numCalls time
	for i := int64(1); i <= numLogged; i++ {
		// Not optimized version - otherwise we can mutate the KeyVals
		zlog.Info("A visible log entry with 3 attributes",
			zap.Int64("iteration", i),
			zap.String("id", id),
			zap.String("logger", "zaplog"), // same byte lentgh as `fortio`
		)
		for j := 1; j <= numExtraNotLogged; j++ {
			// Not optimized version - otherwise we'd check
			// if log.LogDebug() {...}
			zlog.Debug("Not visible entry with 4 attributes",
				zap.String("id", id),
				zap.Int64("iteration", i),
				zap.Int("sub-iteration", j),
				zap.String("logger", "zaplog"),
			)
		}
	}
}
