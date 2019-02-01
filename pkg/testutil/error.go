package testutil

import "github.com/Juniper/contrail/pkg/logutil"

// LogFatalIfError executes given function and calls log.Fatal() if it returned an error.
func LogFatalIfError(f func() error) {
	if err := f(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
