package time

import "time"

type Throttler struct {
	lastTime time.Time
	interval time.Duration
}

func NewThrottler(interval time.Duration) *Throttler {
	throttler := &Throttler{
		lastTime: time.Now(),
		interval: interval,
	}
	return throttler
}

func (self *Throttler) Sync() {
	diff := time.Since(self.lastTime)
	if diff < self.interval {
		time.Sleep(self.interval - diff)
	}
	self.lastTime = time.Now()
}
