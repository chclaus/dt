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

package uuid

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

// uuidV1Cmd represents the uuidV1 command
var uuidV1Cmd = &cobra.Command{
	Use:   "v1",
	Short: "Generates a UUID Version 1",
	Long:  "Generates a v1 UUID, based on timestamp and MAC address (RFC 4122)",
	Run: func(cmd *cobra.Command, args []string) {
		v1 := uuid.NewV1()
		fmt.Println(v1.String())
	},
	Example: "dt uuid v1",
}

func init() {
	uuidCmd.AddCommand(uuidV1Cmd)
}
