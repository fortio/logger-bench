package main

import "fortio.org/log"

func FortioLog(numLogged int64, numExtraNotLogged int) {
	// iterate numCalls time
	for i := int64(0); i < numLogged; i++ {
		log.S(log.Info, "visible log", log.Attr("iteration", i), log.Str("logger", "fortio"))
		for j := 0; j < numExtraNotLogged; j++ {
			log.S(log.Debug, "not visible log", log.Attr("iteration", i), log.Attr("sub-iteration", j), log.Str("logger", "fortio"))
		}
	}
}
