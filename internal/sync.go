// +build muka

package impl

import (
	"fmt"
	"sync"
	"time"
)

// SyncTimeout provides syncronized calls with timeout
type SyncTimeout struct {
	m sync.Mutex
}

// NewSt structure
func NewSt() *SyncTimeout {
	st := SyncTimeout{m: sync.Mutex{}}
	return &st
}

// Callback function
type Callback func(chan interface{}, chan error)

// Call callback
func (sm *SyncTimeout) Call(timeout time.Duration, callback Callback) (interface{}, error) {
	dc := make(chan interface{}, 1)
	ec := make(chan error, 1)

	go func() {
		sm.m.Lock()
		defer sm.m.Unlock()
		callback(dc, ec)
	}()

	select {
	case data := <-dc:
		return data, nil
	case err := <-ec:
		return nil, err
	case <-time.After(timeout):
		return nil, fmt.Errorf("calltimeout %.2f sec", timeout.Seconds())
	}
}
