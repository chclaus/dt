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

package uri

import (
	"fmt"

	"errors"
	"github.com/chclaus/dt/utils"
	"github.com/spf13/cobra"
	"os"
)

// decodeCmd represents the uri decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes an URI-encoded string",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("you have to specify a string which should be decoded")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		result, err := utils.DecodeUri(args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(result)
		return
	},
	Example: "dt uri decode http%3A%2F%2Fwww.github.com",
}

func init() {
	uriCmd.AddCommand(decodeCmd)
}
