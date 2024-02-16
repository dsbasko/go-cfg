// Package gocfg is a Go library for reading configuration data from various sources.
// It provides a unified interface for reading configuration data from environment variables,
// command-line flags, and configuration files.
//
// # Installation
//
// To install the library, use the go get command:
//
//	go get github.com/dsbasko/go-cfg
//
// # Usage
//
// The library provides several functions for reading configuration data:
//
//	ReadEnv(cfg any) error
//	    Reads environment variables into the provided cfg structure. Each field in the cfg structure represents an environment variable.
//
//	MustReadEnv(cfg any)
//	    Similar to ReadEnv but panics if the reading process fails.
//
//	ReadFlag(cfg any) error
//	    Reads command-line flags into the provided cfg structure. Each field in the cfg structure represents a command-line flag.
//
//	MustReadFlag(cfg any)
//	    Similar to ReadFlag but panics if the reading process fails.
//
//	ReadFile(path string, cfg any) error
//	    Reads configuration from a file into the provided cfg structure. The path parameter is the path to the configuration file. Each field in the cfg structure represents a configuration option. Supported file formats include JSON and YAML.
//
//	MustReadFile(path string, cfg any)
//	    Similar to ReadFile but panics if the reading process fails.
//
// Here is an example of how to use the library:
package gocfg
