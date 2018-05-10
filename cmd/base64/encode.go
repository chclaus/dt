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

package base64

import (
	"fmt"

	"encoding/base64"
	"github.com/chclaus/dt/utils"
	"github.com/spf13/cobra"
	"errors"
	"os"
)

// encodeCmd represents the std encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encodes a decoded base64 string",
	Long:  `Encodes a decoded base64 string. As default, the standard encoding 'std',
defined in RFC 4648 is used. All possible encodings are:

  std         Uses the standard base64 encoding, as defined in RFC 4648
  stdRaw      Uses the standard raw, unpadded base64 encoding as defined in RFC 4648 section 3.2
  url         Uses the alternate base64 encoding defined in RFC 4648. Typically used in URLs and filenames
  urlRaw      Uses the standard raw, unpadded alternate base64 encoding defined in RFC 4648
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("you have to specify a string which should be encoded")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		format := cmd.Flag("format").Value.String()
		switch format {
		case "std":
			fmt.Println(utils.EncodeBase64(base64.StdEncoding, args[0]))
			break
		case "stdRaw":
			fmt.Println(utils.EncodeBase64(base64.RawStdEncoding, args[0]))
			break
		case "url":
			fmt.Println(utils.EncodeBase64(base64.URLEncoding, args[0]))
			break
		case "urlRaw":
			fmt.Println(utils.EncodeBase64(base64.RawURLEncoding, args[0]))
			break
		default:
			fmt.Println(fmt.Errorf("the given format '%s' is unknown.", format))
			os.Exit(1)
		}
	},
	Example: `dt base64 encode -f std foo
dt base64 encode -f stdRaw foo
dt base64 encode -f url foo
dt base64 encode -f urlRaw foo`,
}

func init() {
	base64Cmd.AddCommand(encodeCmd)
}
