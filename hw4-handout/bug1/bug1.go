package bug1

import (
	"sync"
)
// Counter stores a count.
type Counter struct {
	n int64
	l sync.Mutex
}

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	c.l.Lock()
	c.n++
	c.l.Unlock()
}
