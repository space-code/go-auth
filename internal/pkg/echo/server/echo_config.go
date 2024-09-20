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

// EchoConfig defines the configuration options for the Echo server.
type EchoConfig struct {
	// The port on which the server will listen. This field is required.
	Port string `mapstructure:"port" validate:"required"`
	// Development indicates if the server is in development mode.
	Development bool `mapstructure:"development"`
	// Host specifies the hostname or IP address that the server will bind to.
	Host string `mapstructure:"host"`
	// Timeout sets the maximum duration, in seconds, that the server should wait
	// for a request to complete before timing out.
	Timeout int `mapstructure:"timeout"`
	// BasePath defines the base URL path for all routes served by the Echo server.
	BasePath string `mapstructure:"basePath" validate:"required"`
}
