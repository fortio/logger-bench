module github.com/fortio/logger-bench

go 1.21

require (
	fortio.org/cli v1.5.1
	fortio.org/log v1.12.0
	fortio.org/scli v1.14.1
	fortio.org/sets v1.0.3
	go.uber.org/zap v1.26.0
)

// replace fortio.org/log => ../log
// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.7.0 // indirect
	fortio.org/struct2env v0.4.0 // indirect
	fortio.org/version v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/sys v0.17.0 // indirect
)
