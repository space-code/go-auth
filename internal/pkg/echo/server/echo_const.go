// Copyright (c) 2024 space-code
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package echoserver

import "time"

const (
	// MaxHeaderBytes controls the maximum number of bytes the server will read parsing the
	// request header's keys and values, including the request line. It does not limit the
	// size of the request body. If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes = 1 << 20
	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	// A zero or negative value means there will be no timeout.
	ReadTimeout = 15 * time.Second
	// WriteTimeout is the maximum duration before timing out writes of the response.
	// It is reset whenever a new request's header is read. Like ReadTimeout, it does not let
	// Handlers make decisions on a per-request basis. A zero or negative value means there
	// will be no timeout.
	WriteTimeout = 15 * time.Second
)
