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

	"github.com/chclaus/dt/cmd"
	"github.com/chclaus/dt/utils"
	"github.com/spf13/cobra"
	"log"
)

// uriCmd represents the uri command
var uriCmd = &cobra.Command{
	Use:   "uri",
	Short: "Encodes or decodes an URI",
	Long:  "Encodes or decodes an URI",
	Run: func(cmd *cobra.Command, args []string) {
		encode := cmd.Flag("encode").Value.String()
		if encode != "" {
			fmt.Println(utils.EncodeUri(encode))
		}

		decode := cmd.Flag("decode").Value.String()
		if decode != "" {
			result, err := utils.DecodeUri(decode)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		}
	},
	Example: `dt uri -e http://www.github.com'
dt uri -d http%3A%2F%2Fwww.github.com`,
}

func init() {
	cmd.RootCmd.AddCommand(uriCmd)

	uriCmd.Flags().StringP("encode", "e", "", "encodes an URI to a safe representation")
	uriCmd.Flags().StringP("decode", "d", "", "decodes an already encoded URI")
}
