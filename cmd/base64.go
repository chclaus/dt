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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/chclaus/dt/utils"
	"log"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Encodes or decodes a string to base64 representation",
	Long: "Encodes or decodes a string to base64 representation.",

	Run: func(cmd *cobra.Command, args []string) {
		encode := cmd.Flag("encode").Value.String()
		if encode != "" {
			fmt.Println(utils.EncodeBase64(encode))
		}

		decode := cmd.Flag("decode").Value.String()
		if decode != "" {
			result, err := utils.DecodeBase64(decode)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		}
	},
	Example: `dt base64 -e foo
dt base64 -d Zm9v`,
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.Flags().StringP("encode", "e", "", "encodes a string to it's base64 representation")
	base64Cmd.Flags().StringP("decode", "d", "", "decodes a base64 string to it's plain representation")
}
