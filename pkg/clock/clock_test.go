package clock

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Run("create current date", func(t *testing.T) {
		c := NewClock()
		assert.IsType(t, time.Time{}, c.Now())
	})
}
