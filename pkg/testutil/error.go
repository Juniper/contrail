package testutil

import "github.com/Juniper/contrail/pkg/log"

// LogFatalIfError executes given function and calls log.Fatal() if it returned an error.
func LogFatalIfError(f func() error) {
	if err := f(); err != nil {
		log.FatalWithStackTrace(err)
	}
}
