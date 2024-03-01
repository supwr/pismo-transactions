//go:generate mockgen -destination=mock/interface.go -source=clock.go -package=mock
package clock

import "time"

type Clock interface {
	Now() time.Time
}

type clock struct{}

func NewClock() Clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now()
}
