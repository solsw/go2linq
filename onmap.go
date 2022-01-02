//go:build go1.18

package go2linq

import (
	"sync"
)

func NewOnMapImmediate[Key comparable, Element any](m map[Key]Element) Enumerator[KeyElement[Key, Element]] {
	r := make([]KeyElement[Key, Element], 0, len(m))
	for k, e := range m {
		r = append(r, KeyElement[Key, Element]{k, e})
	}
	return NewOnSlice(r...)
}

// OnMap is an Enumerator implementation based on map[Key]Element.
type OnMap[Key comparable, Element any] struct {
	mp       map[Key]Element
	closed   bool
	closedCh chan struct{}
	closeMtx sync.Mutex
	crrntCh  chan KeyElement[Key, Element]
	crrnt    KeyElement[Key, Element]
}

func rangeOnMap[Key comparable, Element any](om *OnMap[Key, Element]) {
	for k, e := range om.mp {
		if om.closed {
			return
		}
		select {
		case <-om.closedCh:
			return
		case om.crrntCh <- KeyElement[Key, Element]{key: k, element: e}:
		}
	}
	om.Close()
}

func reInitOnMap[Key comparable, Element any](om *OnMap[Key, Element]) {
	om.closed = false
	om.closedCh = make(chan struct{}, 1)
	om.crrntCh = make(chan KeyElement[Key, Element])
	go rangeOnMap(om)
}

// NewOnMap creates a new OnMap based on the provided map.
func NewOnMap[Key comparable, Element any](m map[Key]Element) *OnMap[Key, Element] {
	onMap := OnMap[Key, Element]{mp: m}
	reInitOnMap(&onMap)
	return &onMap
}

// MoveNext implements the Enumerator.MoveNext method.
func (en *OnMap[Key, Element]) MoveNext() bool {
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
func (en *OnMap[Key, Element]) Current() KeyElement[Key, Element] {
	return en.crrnt
}

// Reset implements the Enumerator.Reset method.
func (en *OnMap[Key, Element]) Reset() {
	if !en.closed {
		en.Close()
	}
	reInitOnMap(en)
}

// Close implements the io.Closer interface.
func (en *OnMap[Key, Element]) Close() error {
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
