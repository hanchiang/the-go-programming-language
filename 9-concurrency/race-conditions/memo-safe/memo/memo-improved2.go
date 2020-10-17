/**
Same as "memo-improved.go", but without using mutex
**/

package memo

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}
type Memo3 struct{ requests chan request }

// New returns a memoization of f. Clients must subsequently call Close.
func New3(f Func) *Memo3 {
	memo := &Memo3{requests: make(chan request)}
	go memo.server(f)
	return memo
}
func (memo *Memo3) Get(key string) (interface{}, error) {
	response := make(chan result)
	// send to requests channel
	memo.requests <- request{key, response}
	// receive from response channel
	res := <-response
	return res.value, res.err
}

func (memo *Memo3) Close() { close(memo.requests) }

func (memo *Memo3) server(f Func) {
	cache := make(map[string]*entry)
	// receive from requests channel
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}
func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
