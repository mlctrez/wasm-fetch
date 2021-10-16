package fetch

import (
	"context"
	"io"
	"io/ioutil"
)

// Opts are the options you can pass to the fetch call.
type Opts struct {
	// Method is the http verb (constants are copied from net/http to avoid import)
	Method string

	// Headers is a map of http headers to send.
	Headers map[string]string

	// Body is the body request
	Body io.Reader

	// Mode docs https://developer.mozilla.org/en-US/docs/Web/API/Request/mode
	Mode string

	// Credentials docs https://developer.mozilla.org/en-US/docs/Web/API/Request/credentials
	Credentials string

	// Cache docs https://developer.mozilla.org/en-US/docs/Web/API/Request/cache
	Cache string

	// Redirect docs https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch
	Redirect string

	// Referrer docs https://developer.mozilla.org/en-US/docs/Web/API/Request/referrer
	Referrer string

	// ReferrerPolicy docs https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch
	ReferrerPolicy string

	// Integrity docs https://developer.mozilla.org/en-US/docs/Web/Security/Subresource_Integrity
	Integrity string

	// KeepAlive docs https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch
	KeepAlive *bool

	// Signal docs https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal
	Signal context.Context
}

func mapOpts(opts *Opts) (map[string]interface{}, error) {
	mp := map[string]interface{}{}

	if opts.Method != "" {
		mp["method"] = opts.Method
	}
	if opts.Headers != nil {
		mp["headers"] = mapHeaders(opts.Headers)
	}
	if opts.Mode != "" {
		mp["mode"] = opts.Mode
	}
	if opts.Credentials != "" {
		mp["credentials"] = opts.Credentials
	}
	if opts.Cache != "" {
		mp["cache"] = opts.Cache
	}
	if opts.Redirect != "" {
		mp["redirect"] = opts.Redirect
	}
	if opts.Referrer != "" {
		mp["referrer"] = opts.Referrer
	}
	if opts.ReferrerPolicy != "" {
		mp["referrerPolicy"] = opts.ReferrerPolicy
	}
	if opts.Integrity != "" {
		mp["integrity"] = opts.Integrity
	}
	if opts.KeepAlive != nil {
		mp["keepalive"] = *opts.KeepAlive
	}

	if opts.Body != nil {
		bts, err := ioutil.ReadAll(opts.Body)
		if err != nil {
			return nil, err
		}

		mp["body"] = string(bts)
	}

	return mp, nil
}

func mapHeaders(mp map[string]string) map[string]interface{} {
	newMap := map[string]interface{}{}
	for k, v := range mp {
		newMap[k] = v
	}
	return newMap
}
