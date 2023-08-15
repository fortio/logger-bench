package main

import (
	"flag"

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
	numCallsFlag := flag.Int64("n", 100, "Number of calls logged (info level)")
	numExtraFlag := flag.Int("e", 9, "Number of extra debug calls (not logged for each `n` iteration), total call will be n*(e+1)")
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
	log.S(log.Info, "Testing", log.Str("logger", cli.Command), log.Attr("num-calls", numCalls), log.Attr("num-extra", numExtra))

	switch cli.Command {
	case "fortio":
		FortioLog(numCalls, numExtra)
	case "zap":
		userErrorf("Zap test/bench not implemented yet")
		// ZapLog()
	case "slog":
		// SLog()
		userErrorf("slog test/bench not implemented yet")
	}
}
