// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package output

import (
	"fmt"

	"github.com/yarpc/crossdock/execute"
	"github.com/yarpc/crossdock/plan"
)

// Reporter is responsible for outputting test results
type Reporter interface {
	Stream(config *plan.Config, tests <-chan execute.TestResponse) Summary
}

// ReporterFunc streams test results to the console
type ReporterFunc func(config *plan.Config, tests <-chan execute.TestResponse) Summary

// Stream allows ReporterFunc to fulfill the Reporter interface
func (f ReporterFunc) Stream(config *plan.Config, tests <-chan execute.TestResponse) Summary {
	return f(config, tests)
}

// GetReporter returns a ReporterFunc for the given name
func GetReporter(name string) (Reporter, error) {
	// default reporter
	if name == "" {
		return List, nil
	}
	// else a specific value was provided
	switch name {
	case "list":
		return List, nil
	case "json":
		return JSON, nil
	default:
		return nil, fmt.Errorf("%v is not a valid reporter", name)
	}
}