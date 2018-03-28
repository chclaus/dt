// Copyright © 2018 Christian Claus <ch.claus@me.com>
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

package main

import (
	"github.com/chclaus/dt/cmd"
	_ "github.com/chclaus/dt/cmd/base64"  // import for init functions
	_ "github.com/chclaus/dt/cmd/date"    // import for init functions
	_ "github.com/chclaus/dt/cmd/hash"    // import for init functions
	_ "github.com/chclaus/dt/cmd/html"    // import for init functions
	_ "github.com/chclaus/dt/cmd/jwt"     // import for init functions
	_ "github.com/chclaus/dt/cmd/random"  // import for init functions
	_ "github.com/chclaus/dt/cmd/server"  // import for init functions
	_ "github.com/chclaus/dt/cmd/uri"     // import for init functions
	_ "github.com/chclaus/dt/cmd/version" // import for init functions
)

func main() {
	cmd.Execute()
}
