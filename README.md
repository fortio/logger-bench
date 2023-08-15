# logger-bench

Realistic scenario e2e golang logger test (slog, zap, fortio, zerolog etc)

Or at least 1 semi relistic scenario:

- 10 go routines,
- Doing 1M "logger on" structured logs (ie level that does get logged)
- And 9M "logger off" structured logs (e.g debug level)
- To actual stderr captured to a file (ideally on a fast FS nvme or ram disk - or maybe `> /dev/null`)

Measure total time taken, and resource utilization (cpu seconds, peak memory)

Obviously results are only valid on same hardware/comparing different logger (and interleaving tests etc)


Currently the output looks like this for each logger:
```json
{"level":"info","ts":1692131277.761052,"msg":"A visible log entry with 3 attributes","iteration":1,"id":"R1","logger":"zaplog"}
{"ts":1692131243.359859,"level":"INFO","msg":"A visible log entry with 3 attributes","iteration":1,"id":"R3","logger":"sloggr"}
{"ts":1692131314.582912,"level":"info","msg":"A visible log entry with 3 attributes","iteration":"1","id":"R3","logger":"fortio"}
```
