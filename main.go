package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
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
	profileFlag := flag.String("profile", "",
		"Write a cpu and memory profile to using the given file `prefix` (will add logger name .cpu and .mem)")
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
	log.SetOutput(os.Stdout) // we use stderr for the logger tests, and so stdout for this
	scli.ServerMain()
	vSet := sets.FromSlice(valid)
	if !vSet.Has(cli.Command) {
		userErrorf("Invalid sub-command %q, valid ones are %v", cli.Command, valid)
	}
	numCalls := *numCallsFlag
	numExtra := *numExtraFlag
	numThrds := *numGoroutineFlag
	profile := ""
	if *profileFlag != "" {
		profile = *profileFlag + "-" + cli.Command
	}
	log.S(log.Info, "Testing",
		log.Str("logger", cli.Command),
		log.Attr("num-calls", numCalls),
		log.Attr("num-extra", numExtra),
		log.Attr("num-goroutines", numThrds),
		log.Attr("gomaxprocs", runtime.GOMAXPROCS(0)),
		log.Attr("profile", profile),
	)
	switch cli.Command {
	case "fortio":
		Drive(profile, FortioLog1, numThrds, numCalls, numExtra)
	case "zap":
		SetupZapLogger()
		Drive(profile, ZapLog1, numThrds, numCalls, numExtra)
		_ = zlog.Sync()
	case "slog":
		SetupSlogLogger()
		Drive(profile, SlogLog1, numThrds, numCalls, numExtra)
	}
}

func PrintMemoryStats(start, end *runtime.MemStats) {
	log.S(log.Info, "Memory stats",
		log.Attr("alloc", end.Alloc-start.Alloc),
		log.Attr("total-alloc", end.TotalAlloc-start.TotalAlloc),
		log.Attr("num-gc", end.NumGC-start.NumGC),
	)
}

// Drive the given (iteration logging) function from multiple goroutines.
func Drive(profile string, fn func(string, int64, int), numGoroutines int, numLogged int64, numExtraNotLogged int) {
	// wait group
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)
	var mStart runtime.MemStats
	var mEnd runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&mStart)
	var fc *os.File
	if profile != "" {
		var err error
		fc, err = os.Create(profile + ".cpu")
		if err != nil {
			log.Fatalf("Unable to create .cpu profile: %v", err)
		}
		if err = pprof.StartCPUProfile(fc); err != nil {
			log.Critf("Unable to start cpu profile: %v", err)
		}
	}
	log.SetOutput(os.Stderr)
	for i := 1; i <= numGoroutines; i++ {
		go func(c int) {
			fn(fmt.Sprintf("R%d", c), numLogged, numExtraNotLogged)
			wg.Done()
		}(i)
	}
	wg.Wait()
	runtime.ReadMemStats(&mEnd)
	log.SetOutput(os.Stdout)
	if profile != "" {
		pprof.StopCPUProfile()
		fm, err := os.Create(profile + ".mem")
		if err != nil {
			log.Fatalf("Unable to create .mem profile: %v", err)
		}
		runtime.GC() // get up-to-date statistics
		if err = pprof.WriteHeapProfile(fm); err != nil {
			log.Critf("Unable to write heap profile: %v", err)
		}
		fm.Close()
		log.S(log.Info, "Wrote profile data", log.Str("cpu", profile+".cpu"), log.Str("mem", profile+".mem"))
	}
	PrintMemoryStats(&mStart, &mEnd)
}
