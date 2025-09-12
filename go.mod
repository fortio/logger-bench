module github.com/fortio/logger-bench

go 1.21

require (
	fortio.org/cli v1.11.0
	fortio.org/log v1.17.2
	fortio.org/scli v1.16.1
	fortio.org/sets v1.3.0
	go.uber.org/zap v1.27.0
)

// replace fortio.org/log => ../log
// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.8.1 // indirect
	fortio.org/struct2env v0.4.2 // indirect
	fortio.org/version v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/kortschak/goroutine v1.1.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto/x509roots/fallback v0.0.0-20250203165127-fa5273e46196 // indirect
	golang.org/x/sys v0.30.0 // indirect
)
