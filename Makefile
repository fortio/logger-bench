
# Small arg for overview testing
# 150 calls to logger (3 go routines * 10 iterations * (4 invisible + 1 visible)) so with 3*40 of them not logged unless -loglevel debug is passed:
ARGS:=-n 10 -e 4 -r 3


manual-check:
	@echo "--- Manual eyeball test, should have 10 log entries (out of 50 made) ---"
	go run . fortio $(ARGS) 2>&1 | jq -c
	@echo "--- Manual eyeball test, should have all 50 entries - in color and with goroutine just for fun ---"
	GOMAXPROCS=8 go run -race . fortio $(ARGS) -loglevel debug -logger-force-color=true -logger-goroutine=true
