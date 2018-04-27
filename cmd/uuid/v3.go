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
	"github.com/spf13/cobra"
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/pkg/errors"
)

// uuidV3Cmd represents the uuidV3 command
var uuidV3Cmd = &cobra.Command{
	Use:   "v3",
	Short: "Generates a UUID Version 3",
	Long: "Generates a v3 UUID, based on MD5 hashing of (namespace(UUID), value) (RFC 4122)",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ns := cmd.Flag("namespace").Value.String()
		errMsg := "You have to specify a namespace and value. The namespace MUST be a valid UUID"
		if ns == "" {
			return errors.New(errMsg)
		}
		if _, err := uuid.FromString(ns); err != nil {
			return errors.New(errMsg)
		}

		val := cmd.Flag("value").Value.String()
		if val == "" {
			return errors.New(errMsg)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ns := cmd.Flag("namespace").Value.String()
		nsUUID, _:= uuid.FromString(ns);
		val := cmd.Flag("value").Value.String()

		fmt.Println(uuid.NewV3(nsUUID, val).String())
	},
	Example: "dt uuid v5 -n cacae610-c76a-4736-90ef-0271126b4345 -v foo",
}

func init() {
	uuidCmd.AddCommand(uuidV3Cmd)

	uuidV3Cmd.Flags().StringP("namespace", "n", "", "The namespace that should be hashed. It should be a domain or application specific UUID")
	uuidV3Cmd.Flags().StringP("value", "v", "", "The value that should be hashed")
}
