package testutil

import "log"

// LogFatalIfErr logs the err during function call
func LogFatalIfErr(f func() error) {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}
