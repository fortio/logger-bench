
# Small arg for overview testing
# 50 calls to logger (10*(4+1)) with 40 of them not logged unless -loglevel debug is passed:
ARGS:=-n 10 -e 4


manual-check:
	@echo "--- Manual eyeball test, should have 10 log entries (out of 50 made) ---"
	go run . fortio $(ARGS) 2>&1 | jq -c
	@echo "--- Manual eyeball test, should have all 50 entries - in color just for fun ---"
	go run . fortio $(ARGS) -loglevel debug -logger-force-color=true
