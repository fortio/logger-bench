package main

import (
	"fmt"
	"sync"

	"fortio.org/log"
)

func FortioLog1(id string, numLogged int64, numExtraNotLogged int) {
	// iterate numCalls time
	for i := int64(1); i <= numLogged; i++ {
		// Not optimized version - otherwise we can mutate the KeyVals
		log.S(log.Info, "visible log",
			log.Attr("iteration", i),
			log.Str("id", id),
			log.Str("logger", "fortio"),
		)
		for j := 1; j <= numExtraNotLogged; j++ {
			// Not optimized version - otherwise we'd check
			// if log.LogDebug() {...}
			log.S(log.Debug, "not visible log",
				log.Str("id", id),
				log.Attr("iteration", i),
				log.Attr("sub-iteration", j),
				log.Str("logger", "fortio"),
			)
		}
	}
}

func FortioLog(numGoroutines int, numLogged int64, numExtraNotLogged int) {
	// wait group
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		go func(c int) {
			FortioLog1(fmt.Sprintf("R%d", c), numLogged, numExtraNotLogged)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
