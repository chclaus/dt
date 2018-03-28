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

package random

import (
	"github.com/spf13/cobra"
	"github.com/chclaus/dt/utils"
	"fmt"
)

// alphaCmd represents the alpha command
var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Generates a random string, based on an alphabet with a specific length",
	Long:  "Generates a random string, based on an alphabet with a specific length.",
	Run: func(cmd *cobra.Command, args []string) {
		result := utils.Random(length, utils.Alphabet{})
		fmt.Println(result)
	},
}

func init() {
	randomCmd.AddCommand(alphaCmd)

	alphaCmd.Flags().IntVarP(&length, "length", "l", 10, "defines the length (number of letters) of the generated string")
}
