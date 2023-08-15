package main

import (
	"fortio.org/log"
	"go.uber.org/zap"
)

func ZapLog1(id string, numLogged int64, numExtraNotLogged int) {
	cfg := zap.NewProductionConfig()
	if log.GetLogLevel() == log.Debug {
		cfg.Level.SetLevel(zap.DebugLevel)
	}
	logger, _ := cfg.Build(zap.WithCaller(false))
	defer logger.Sync()
	// iterate numCalls time
	for i := int64(1); i <= numLogged; i++ {
		// Not optimized version - otherwise we can mutate the KeyVals
		logger.Info("A visible log entry with 3 attributes",
			zap.Int64("iteration", i),
			zap.String("id", id),
			zap.String("logger", "zaplog"), // same byte lentgh as `fortio`
		)
		for j := 1; j <= numExtraNotLogged; j++ {
			// Not optimized version - otherwise we'd check
			// if log.LogDebug() {...}
			logger.Debug("Not visible entry with 4 attributes",
				zap.String("id", id),
				zap.Int64("iteration", i),
				zap.Int("sub-iteration", j),
				zap.String("logger", "fortio"),
			)
		}
	}
}
