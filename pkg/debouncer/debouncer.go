package debouncer

import (
	"sync"
	"time"
)

type Debouncer struct {
	interval time.Duration
	mutex    sync.Mutex
	timer    *time.Timer
}

func NewDebouncer(duration string) (*Debouncer, error) {
	interval, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}

	d := new(interval)
	return &d, nil
}

func new(interval time.Duration) Debouncer {
	return Debouncer{
		interval: interval,
	}
}

func (d *Debouncer) Call(fn func()) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.interval, fn)
}
