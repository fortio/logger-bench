# logger-bench

Realistic scenario e2e golang logger test (slog, zap, fortio, zerolog etc)

Or at least 1 semi relistic scenario:

- 10 go routines,
- Doing 1M "logger on" structured logs (ie level that does get logged)
- And 9M "logger off" structured logs (e.g debug level)
- To actual stderr captured to a file (ideally on a fast FS nvme or ram disk - or maybe `> /dev/null`)

Measure total time taken, and resource utilization (cpu seconds, peak memory)

Obviously results are only valid on same hardware/comparing different logger (and interleaving tests etc)
