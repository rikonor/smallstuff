package monitor

import (
	"sync"
	"time"
)

type MonitorFunc func()

// Monitor inactivity
// it calls a given function if its not pinged for a certain amount of time
type Monitor struct {
	mu    *sync.Mutex
	dur   time.Duration
	timer *time.Timer
}

// NewMonitor creates a new Monitor
func NewMonitor(dur time.Duration, fn MonitorFunc) *Monitor {
	m := &Monitor{
		mu:    &sync.Mutex{},
		dur:   dur,
		timer: time.NewTimer(dur),
	}

	go func() {
		select {
		case <-m.timer.C:
			fn()
		}
	}()

	return m
}

// Heartbeat resets the timer of the monitor (meaning we are still active)
func (m *Monitor) Heartbeat() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.timer.Reset(m.dur)
}

// Stop the monitor from tracking inactivity
func (m *Monitor) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.timer.Stop()
}
