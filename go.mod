module github.com/fortio/logger-bench

go 1.20

require (
	fortio.org/cli v1.4.2
	fortio.org/log v1.11.0
	fortio.org/scli v1.12.1
	fortio.org/sets v1.0.3
	go.uber.org/zap v1.25.0
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

// replace fortio.org/log => ../log

// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.7.0 // indirect
	fortio.org/version v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
