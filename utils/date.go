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

import (
	"errors"
	"strconv"
	"time"
	"unicode"
)

func isDigitOnly(date string) bool {
	for _, c := range date {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

// ParseTimestamp tries to parses a string into a time.Time format. The following cascade is used:
// 1. It tries to parse the string as current time millis.
// 2. It tries to parse the string as current time nanos.
// 3. It tries to parse the string as RFC 3339 time.
func ParseTimestamp(date string) (time.Time, error) {
	if isDigitOnly(date) {
		intDate, err := strconv.ParseInt(date, 10, 64)
		if err != nil {
			return time.Time{}, err
		}

		if len(date) == 10 {
			return time.Unix(intDate, 0), nil
		}

		if len(date) == 13 {
			nanos := intDate % 1000
			millis := intDate / 1000

			return time.Unix(millis, nanos), nil
		}

		return time.Time{}, errors.New("Could not parse time :(")
	}

	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
