
# Small arg for overview testing
# 150 calls to logger (3 go routines * 10 iterations * (4 invisible + 1 visible)) so with 3*40 of them not logged unless -loglevel debug is passed:
ARGS:=-n 10 -e 4 -r 3


manual-check: manual-check-fortio manual-check-zap manual-check-slog

manual-check-fortio:
	$(MAKE) manual-check-param PARAM=fortio

manual-check-zap:
	$(MAKE) manual-check-param PARAM=zap

manual-check-slog:
	$(MAKE) manual-check-param PARAM=slog

manual-check-param:
	@echo "--- Manual eyeball test for $(PARAM), should have 10 log entries (out of 50 made) ---"
	go run . $(PARAM) $(ARGS) 2>&1
	@echo "--- Manual eyeball test, should have all 50 entries - in color and with goroutine just for fun ---"
	GOMAXPROCS=8 go run -race . $(PARAM) $(ARGS) -loglevel debug -logger-force-color=true -logger-goroutine=true
	@echo "--- end of $(PARAM) manual check ---"
