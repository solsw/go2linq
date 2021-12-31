//go:build go1.18

package go2linq

import (
	"sync"
)

// OnMap is an Enumerator implementation based on map[Key]Value.
type OnMap[Key comparable, Value any] struct {
	mp       map[Key]Value
	closed   bool
	closedCh chan struct{}
	closeMtx sync.Mutex
	crrntCh  chan KeyValue[Key, Value]
	crrnt    KeyValue[Key, Value]
}

func rangeOnMap[Key comparable, Value any](om *OnMap[Key, Value]) {
	for k, v := range om.mp {
		if om.closed {
			return
		}
		select {
		case <-om.closedCh:
			return
		case om.crrntCh <- KeyValue[Key, Value]{key: k, value: v}:
		}
	}
	om.Close()
}

func reInitOnMap[Key comparable, Value any](om *OnMap[Key, Value]) {
	om.closed = false
	om.closedCh = make(chan struct{}, 1)
	om.crrntCh = make(chan KeyValue[Key, Value])
	go rangeOnMap(om)
}

// NewOnMap creates a new OnMap based on the provided map.
func NewOnMap[Key comparable, Value any](m map[Key]Value) *OnMap[Key, Value] {
	onMap := OnMap[Key, Value]{mp: m}
	reInitOnMap(&onMap)
	return &onMap
}

// MoveNext implements the Enumerator.MoveNext method.
func (en *OnMap[Key, Value]) MoveNext() bool {
	if en.closed {
		return false
	}
	var ok bool
	en.crrnt, ok = <-en.crrntCh
	if !ok {
		return false
	}
	return true
}

// Current implements the Enumerator.Current method.
func (en *OnMap[Key, Value]) Current() KeyValue[Key, Value] {
	return en.crrnt
}

// Reset implements the Enumerator.Reset method.
func (en *OnMap[Key, Value]) Reset() {
	if !en.closed {
		en.Close()
	}
	reInitOnMap(en)
}

// Close implements the io.Closer interface.
func (en *OnMap[Key, Value]) Close() error {
	en.closeMtx.Lock()
	defer en.closeMtx.Unlock()
	if en.closed {
		return nil
	}
	en.closed = true
	en.closedCh <- struct{}{}
	close(en.crrntCh)
	return nil
}
