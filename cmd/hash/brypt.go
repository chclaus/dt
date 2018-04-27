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

package hash

import (
	"errors"
	"fmt"
	"github.com/chclaus/dt/utils"
	"github.com/spf13/cobra"
)

var cost int

// bcryptCmd represents the bcrypt command
var bcryptCmd = &cobra.Command{
	Use:   "bcrypt",
	Short: "Returns the bcrypt hash representation of the input",
	Long:  "Returns the bcrypt hash representation of the input.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You have to specify the text that should be hashed.")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.BcryptHash(args[0], cost))
	},
}

func init() {
	hashCmd.AddCommand(bcryptCmd)

	bcryptCmd.Flags().IntVarP(&cost, "cost", "c", 10, "defines the costs of the bcrypt algorithm")
}
