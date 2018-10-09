package testutil

import "log"

// LogFatalIfError executes given function and calls log.Fatal() if it returned an error.
func LogFatalIfError(f func() error) {
	if err := f(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
