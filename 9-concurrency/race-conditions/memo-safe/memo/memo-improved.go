package memo

import "sync"

/**
Each entry contains the memoized result of a call to the function f, as before, but it additionally
contains a channel called ready. Just after the entry’s result has been set, this channel
will be closed, to broadcast to any other goroutines that it is now safe for them to read
the result fro m the entry
*/
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New2(f Func) *Memo2 {
	return &Memo2{f: f, cache: make(map[string]*entry)}
}

type Memo2 struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo2) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
	}
	<-e.ready // wait for ready condition
	return e.res.value, e.res.err
}
