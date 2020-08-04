package retry

import (
	"time"

	"github.com/pkg/errors"
)

// Func is a function that could be retried.
type Func func() (retry bool, err error)

// Option is a function that mutates config.
type Option func(*config)

type logger interface {
	Debugf(string, ...interface{})
}

type config struct {
	log      logger
	interval time.Duration
	backoff  float64
	maxRetry int
}

func getConfig(opts []Option) config {
	c := config{}
	for _, o := range opts {
		o(&c)
	}
	return c
}

// WithLog adds additional logging in Retry.
func WithLog(log logger) Option {
	return func(c *config) {
		c.log = log
	}
}

// WithInterval adds additional logging in Retry.
func WithInterval(interval time.Duration) Option {
	return func(c *config) {
		c.interval = interval
	}
}

// WithBackoff controls whether the interval is doubled between retries
func WithBackoff(backoff float64) Option {
	return func(c *config) {
		c.backoff = backoff
	}
}

// WithMaxRetry limit the number of retries, so the total number of run is maxRetry + 1
// maxRetry smaller that 1 has no effect, retry.DO function keeps retrying until succeed or meeting no-retry error
func WithMaxRetry(maxRetry int) Option {
	return func(c *config) {
		if maxRetry >= 1 {
			c.maxRetry = maxRetry
		}
	}
}

// Do runs function f in loop until the function returns retry == false.
func Do(f Func, opts ...Option) error {
	c := getConfig(opts)

	for round := 0; ; round++ {
		retry, err := f()

		if !retry {
			return err
		}

		// maximum retry limit check
		if c.maxRetry >= 1 && round >= c.maxRetry {
			return errors.Wrap(err, "maximum retry reached")
		}

		// logging check
		if log := c.log; log != nil {
			log.Debugf("Retry %d, error was: %v", round, err)
		}

		// sleep - zero and negative interval value lead to time.Sleep return immediately
		time.Sleep(c.interval)

		// note that backoff factor less than 1 is also accepted here
		if c.backoff > 0 {
			c.interval = time.Duration(float64(c.interval) * c.backoff)
		}
	}
}
