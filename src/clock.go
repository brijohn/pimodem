package main

import (
	"time"
)

type TimerID int

type Timer struct {
	lastTick time.Time
	period   time.Duration
	C        chan time.Time
}

type Clock struct {
	ticker *time.Ticker
	timers map[TimerID]*Timer
	done   chan bool
}

func NewClock(interval time.Duration) *Clock {
	var c Clock
	c.timers = make(map[TimerID]*Timer)
	c.done = make(chan bool)
	c.ticker = time.NewTicker(interval)
	go func() {
		for {
			select {
			case t := <-c.ticker.C:
				for _, timer := range c.timers {
					if timer.period > 0 && time.Since(timer.lastTick) >= timer.period {
						timer.lastTick = time.Now()
						select {
						case timer.C <- t:
						default:
						}
					}
				}
			case <-c.done:
				c.ticker.Stop()
				return
			}
		}
	}()
	return &c
}

func (c *Clock) GetTimer(id TimerID) <-chan time.Time {
	timer, ok := c.timers[id]
	if !ok {
		timer = &Timer{}
		timer.C = make(chan time.Time, 1)
		c.timers[id] = timer
	}
	return timer.C
}

func (c *Clock) SetDuration(id TimerID, period time.Duration) {
	timer, ok := c.timers[id]
	if ok {
		timer.period = period
		timer.lastTick = time.Now()
	}
}

func (c *Clock) Stop() {
	c.done <- true
}
