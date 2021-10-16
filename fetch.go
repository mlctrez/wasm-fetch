//go:build wasm
package fetch

import (
	"errors"
	"syscall/js"
)

// Fetch uses the JS Fetch API to make requests over WASM.
func Fetch(url string, opts *Opts) (*Response, error) {
	optsMap, err := mapOpts(opts)
	if err != nil {
		return nil, err
	}

	type fetchResponse struct {
		r *Response
		e error
	}
	ch := make(chan *fetchResponse)
	done := make(chan struct{}, 1)
	if opts.Signal != nil {
		controller := js.Global().Get("AbortController").New()
		signal := controller.Get("signal")
		optsMap["signal"] = signal
		go func() {
			select {
			case <-opts.Signal.Done():
				controller.Call("abort")
			case <-done:
			}
		}()
	}

	js.Global().Call("fetch", url, optsMap).Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var r Response
		resp := args[0]
		headersIt := resp.Get("headers").Call("entries")
		headers := Header{}
		for {
			n := headersIt.Call("next")
			if n.Get("done").Bool() {
				break
			}
			pair := n.Get("value")
			key, value := pair.Index(0).String(), pair.Index(1).String()
			headers.Add(key, value)
		}
		r.Headers = headers
		r.OK = resp.Get("ok").Bool()
		r.Redirected = resp.Get("redirected").Bool()
		r.Status = resp.Get("status").Int()
		r.StatusText = resp.Get("statusText").String()
		r.Type = resp.Get("type").String()
		r.URL = resp.Get("url").String()
		r.BodyUsed = resp.Get("bodyUsed").Bool()

		args[0].Call("text").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			r.Body = []byte(args[0].String())
			done <- struct{}{}
			ch <- &fetchResponse{r: &r}
			return nil
		}))
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		msg := args[0].Get("message").String()
		done <- struct{}{}
		ch <- &fetchResponse{e: errors.New(msg)}
		return nil
	}))

	r := <-ch

	return r.r, r.e
}
