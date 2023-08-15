module github.com/fortio/logger-bench

go 1.20

require (
	fortio.org/cli v1.4.2
	fortio.org/log v1.10.0
	fortio.org/scli v1.11.0
	fortio.org/sets v1.0.3
	go.uber.org/zap v1.25.0
)

// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.5.3 // indirect
	fortio.org/version v1.0.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20230801115018-d63ba01acd4b // indirect
	golang.org/x/sys v0.10.0 // indirect
)
