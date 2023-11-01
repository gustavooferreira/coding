package ops

import (
	"sync"
	"sync/atomic"
)

// Operation has no synchonization primitives and therefore is not thread-safe.
type Operation struct {
	counter int
}

func (o *Operation) Increment() {
	o.counter += 1
}

func (o *Operation) Get() int {
	return o.counter
}

// OperationAtomicInteger provides synchonization via an atomic integer.
type OperationAtomicInteger struct {
	counter int32
}

func (o *OperationAtomicInteger) Increment() {
	atomic.AddInt32(&o.counter, 1)
}

func (o *OperationAtomicInteger) Get() int32 {
	return atomic.LoadInt32(&o.counter)
}

// OperationMutex provides synchonization via a mutex.
type OperationMutex struct {
	mu      sync.RWMutex
	counter int
}

func (o *OperationMutex) Increment() {
	o.mu.Lock()
	o.counter += 1
	o.mu.Unlock()
}

func (o *OperationMutex) Get() int {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.counter
}

// OperationAtomicValue provides synchonization via an atomic.Value.
type OperationAtomicValue struct {
	// atomicValue stores the counter integer
	atomicValue atomic.Value
}

// Increment isn't really atomically incremented as we can't atomically get the value and increment it when using
// sync.Value.
func (o *OperationAtomicValue) Increment() {
	o.atomicValue.Store(10) // we store 10 every time instead of actually incrementing the value
}

func (o *OperationAtomicValue) Get() int {
	return o.atomicValue.Load().(int)
}
