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


## Results

For 1M logs logged and 10M logs issued:

### Fortio
```bash
$ time go run . fortio -n 100000 2>&1 | head -2
```json
{"ts":1700637382.217323,"level":"info","msg":"Starting","command":"logger-bench","version":"dev  go1.21.4 arm64 darwin","go-max-procs":10}
{"ts":1700637382.217334,"level":"info","msg":"Testing","logger":"fortio","num-calls":100000,"num-extra":9,"num-goroutines":10,"gomaxprocs":10,"profile":""}
```
```bash
real	0m0.515s
user	0m0.357s
sys	0m0.285s
```

### Slog

```bash
$ time go run . slog -n 100000 2>&1 | head -2
```
```json
{"ts":1700637390.850496,"level":"info","msg":"Starting","command":"logger-bench","version":"dev  go1.21.4 arm64 darwin","go-max-procs":10}
{"ts":1700637390.850510,"level":"info","msg":"Testing","logger":"slog","num-calls":100000,"num-extra":9,"num-goroutines":10,"gomaxprocs":10,"profile":""}
```
```bash
real	0m0.548s
user	0m0.361s
sys	0m0.295s
```

### Zap
```bash
$ time go run . zap -n 100000 2>&1 | head -2
```json
{"ts":1700637401.684503,"level":"info","msg":"Starting","command":"logger-bench","version":"dev  go1.21.4 arm64 darwin","go-max-procs":10}
{"ts":1700637401.684517,"level":"info","msg":"Testing","logger":"zap","num-calls":100000,"num-extra":9,"num-goroutines":10,"gomaxprocs":10,"profile":""}
```
```bash
real	0m0.513s
user	0m0.358s
sys	0m0.295s
```

## Conclusion

All 3 are about the same performance, with zap marginally about 1% faster than fortio and slog slowest at about 6% slower than fortio.
