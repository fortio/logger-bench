package main

import (
	"fortio.org/cli"
	"fortio.org/log"
	"fortio.org/scli"
	"fortio.org/sets"
)

func main() {
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
		// Revert to color mode instead of json as this is a user error.
		log.Config.ConsoleColor = true
		log.SetColorMode()
		cli.ErrUsage("Invalid sub-command %q, valid ones are %v", cli.Command, valid)
	}
}
