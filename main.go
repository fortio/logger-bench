package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"

	"fortio.org/cli"
	"fortio.org/log"
	"fortio.org/scli"
	"fortio.org/sets"
)

func userErrorf(format string, args ...interface{}) {
	// Revert to color mode instead of json as this is a user error.
	log.Config.ConsoleColor = true
	log.SetColorMode()
	cli.ErrUsage(format, args...)
}

func main() {
	numCallsFlag := flag.Int64("n", 100, "Number of calls logged (info level), per goroutine `r`")
	numExtraFlag := flag.Int("e", 9, "Number of extra debug calls (not logged for each `n` iteration), total call will be n*(e+1)")
	numGoroutineFlag := flag.Int("r", 10, "Number of go routines to use (multiplies the other numbers)")
	// Force JSON output even on console and disable expensive debug file/line logging
	// as well as goroutine id logging which most other loggers don't have.
	cli.BeforeFlagParseHook = func() {
		for _, f := range [][]string{
			{"false", "logger-file-line", "logger-goroutine"},
			{"true", "logger-no-color"},
		} {
			err := cli.ChangeFlagsDefault(f[0], f[1:]...)
			if err != nil {
				log.Fatalf("Bug in setting flag new default: %v", err)
			}
		}
	}
	valid := []string{"fortio", "zap", "slog"}
	cli.CommandBeforeFlags = true
	cli.CommandHelp = "{" + cli.ColorJoin(log.Colors.Purple, valid...) + "}"
	scli.ServerMain()
	vSet := sets.FromSlice(valid)
	if !vSet.Has(cli.Command) {
		userErrorf("Invalid sub-command %q, valid ones are %v", cli.Command, valid)
	}
	numCalls := *numCallsFlag
	numExtra := *numExtraFlag
	numThrds := *numGoroutineFlag
	log.S(log.Info, "Testing",
		log.Str("logger", cli.Command),
		log.Attr("num-calls", numCalls),
		log.Attr("num-extra", numExtra),
		log.Attr("num-goroutines", numThrds),
		log.Attr("gomaxprocs", runtime.GOMAXPROCS(0)),
	)
	switch cli.Command {
	case "fortio":
		Drive(FortioLog1, numThrds, numCalls, numExtra)
	case "zap":
		SetupZapLogger()
		Drive(ZapLog1, numThrds, numCalls, numExtra)
		_ = zlog.Sync()
	case "slog":
		// Drive(Slog1, numThrds, numCalls, numExtra)
		userErrorf("slog test/bench not implemented yet")
	}
}

// Drive the given (iteration logging) function from multiple goroutines.
func Drive(fn func(string, int64, int), numGoroutines int, numLogged int64, numExtraNotLogged int) {
	// wait group
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		go func(c int) {
			fn(fmt.Sprintf("R%d", c), numLogged, numExtraNotLogged)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
