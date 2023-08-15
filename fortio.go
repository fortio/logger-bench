package main

import (
	"fortio.org/log"
)

func FortioLog1(id string, numLogged int64, numExtraNotLogged int) {
	// iterate numCalls time
	for i := int64(1); i <= numLogged; i++ {
		// Not optimized version - otherwise we can mutate the KeyVals
		log.S(log.Info, "A visible log entry with 3 attributes",
			log.Attr("iteration", i),
			log.Str("id", id),
			log.Str("logger", "fortio"),
		)
		for j := 1; j <= numExtraNotLogged; j++ {
			// Not optimized version - otherwise we'd check
			// if log.LogDebug() {...}
			log.S(log.Debug, "Not visible entry with 4 attributes",
				log.Str("id", id),
				log.Attr("iteration", i),
				log.Attr("sub-iteration", j),
				log.Str("logger", "fortio"),
			)
		}
	}
}
