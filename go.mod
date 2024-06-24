module github.com/fortio/logger-bench

go 1.21

require (
	fortio.org/cli v1.6.0
	fortio.org/log v1.12.2
	fortio.org/scli v1.15.0
	fortio.org/sets v1.1.1
	go.uber.org/zap v1.27.0
)

// replace fortio.org/log => ../log
// replace fortio.org/cli => ../cli

require (
	fortio.org/dflag v1.7.2 // indirect
	fortio.org/struct2env v0.4.0 // indirect
	fortio.org/version v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto/x509roots/fallback v0.0.0-20240604170348-d4e7c9cb6cb8 // indirect
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8 // indirect
	golang.org/x/sys v0.21.0 // indirect
)
