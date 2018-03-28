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
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

// NUM are numbers [0:10)
const NUM = "0123456789"

// ALPH are alphabet letters in lower and upper case
const ALPH = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// SPECIAL are special characters
const SPECIAL = "!#$%&()*+,-./:;><?^_"

// PARALLEL_LIMIT is the limit for concurrent execution of random char generations
const PARALLEL_LIMIT = 100

// Source defines a string of letters used as source for a random string
type Source interface {
	letters() string
}

// Numeric is a default source of numbers in a range of [0,10).
type Numeric struct{}

func (n Numeric) letters() string {
	return NUM
}

// Alphabet is a source of letters: abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
type Alphabet struct{}

func (a Alphabet) letters() string {
	return ALPH
}

// AlphaNumeric is a source of letters and numbers: abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
type AlphaNumeric struct{}

func (a AlphaNumeric) letters() string {
	return ALPH + NUM
}

// Complex is the same as the AlphaNumeric source but adds also special characters
type Complex struct{}

func (c Complex) letters() string {
	return ALPH + NUM + SPECIAL
}

// Random generates a random string with a length of n. Each letter of the
// random string has it's source in the given source a
func Random(n int, a Source) string {
	letters := a.letters()
	l := len(letters)
	r := make(chan string)
	result := make([]string, n)
	sem := make(Semaphore, PARALLEL_LIMIT)

	for i := 0; i < n; i++ {
		go func(numLetters int, letters string) {
			// Wait if there are more concurrent executions than the PARALLEL_LIMIT allows
			sem.Acquire(1)
			defer sem.Release(1)

			val, _ := rand.Int(rand.Reader, big.NewInt(int64(numLetters)))
			iVal := val.Int64()
			r <- letters[iVal : iVal+1]
		}(l, letters)
	}

	for i := 0; i < n; i++ {
		result[i] = <-r
	}

	return strings.Join(result, "")
}

// RandomNumber generates a random number with a length of n. Each letter of
// the random string has it's source in the Numeric source, but the first
// number of the string is in a range of [1,10)
func RandomNumber(n int) string {
	result := Random(n, Numeric{})
	rArr := strings.Split(result, "")

	if rArr[0] == "0" {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(9)))
		iVal := val.Int64() + 1
		rArr[0] = strconv.FormatInt(iVal, 10)
	}

	return strings.Join(rArr, "")
}
