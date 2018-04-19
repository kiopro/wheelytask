package main

import (
	"runtime"

	app "./app"
	"github.com/rs/zerolog"
)

// *****************************************************************************
// * Application start
// *****************************************************************************

func init() {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = ""

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	go app.StartHTTPServer()
	app.StartGRPCServer()
}
