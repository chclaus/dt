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

package utils

import "testing"

func TestParseMillisUnixTimestamp(t *testing.T) {
	time, err := ParseTimestamp("1521820515")
	if err != nil {
		t.Fatal(err)
	}
	if time.Unix() != 1521820515 {
		t.Fatal("Error while parsing timestamp")
	}
}

func TestParseNanosUnixTimestamp(t *testing.T) {
	time, err := ParseTimestamp("1521820515000")
	if err != nil {
		t.Fatal(err)
	}
	if time.Unix() != 1521820515 {
		t.Fatal("Error while parsing timestamp")
	}
}

func TestParseIsoDate(t *testing.T) {
	time, err := ParseTimestamp("2018-03-23T16:55:15+01:00")
	if err != nil {
		t.Fatal(err)
	}
	if time.Unix() != 1521820515 {
		t.Fatal("Error while parsing timestamp")
	}
}
