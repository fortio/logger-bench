module github.com/fortio/logger-bench

go 1.21

require (
	fortio.org/cli v1.5.2
	fortio.org/log v1.12.1
	fortio.org/scli v1.14.2
	fortio.org/sets v1.0.4
	go.uber.org/zap v1.27.0
)

// replace fortio.org/log => ../log
// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.7.1 // indirect
	fortio.org/struct2env v0.4.0 // indirect
	fortio.org/version v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20240318143956-a85f2c67cd81 // indirect
	golang.org/x/sys v0.18.0 // indirect
)
