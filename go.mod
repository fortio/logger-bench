module github.com/fortio/logger-bench

go 1.21

require (
	fortio.org/cli v1.9.2
	fortio.org/log v1.17.1
	fortio.org/scli v1.15.3
	fortio.org/sets v1.2.0
	go.uber.org/zap v1.27.0
)

// replace fortio.org/log => ../log
// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.7.3 // indirect
	fortio.org/struct2env v0.4.1 // indirect
	fortio.org/version v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/kortschak/goroutine v1.1.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto/x509roots/fallback v0.0.0-20240806160748-b2d3a6a4b4d3 // indirect
	golang.org/x/exp v0.0.0-20240808152545-0cdaa3abc0fa // indirect
	golang.org/x/sys v0.25.0 // indirect
)
