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
	"github.com/chclaus/dt/cmd"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/chclaus/dt/utils"
	"os"
	"crypto/md5"
	"crypto/sha1"
	"golang.org/x/crypto/sha3"
	"crypto/sha256"
	"crypto/sha512"
)

var cost int

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hashes an arbitrary input with different hash algorithms",
	Long: `Hashes an arbitrary input with different hash algorithms. These are:
  - bcrypt
  - md5 (default)
  - sha1
  - sha256
  - sha3_256
  - sha3_512
  - sha512`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You have to specify the text which should be hashed.")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		algo := cmd.Flag("algo").Value.String()
		switch algo {
		case "bcrypt":
			fmt.Println(utils.BcryptHash(args[0], cost))
			break
		case "md5":
			fmt.Println(utils.Hash(md5.New(), args[0]))
			break
		case "sha1":
			fmt.Println(utils.Hash(sha1.New(), args[0]))
			break
		case "sha3_256":
			fmt.Println(utils.Hash(sha3.New256(), args[0]))
			break
		case "sha3_512":
			fmt.Println(utils.Hash(sha3.New512(), args[0]))
			break
		case "sha256":
			fmt.Println(utils.Hash(sha256.New(), args[0]))
			break
		case "sha512":
			fmt.Println(utils.Hash(sha512.New(), args[0]))
			break
		default:
			fmt.Println(fmt.Errorf("the given algorithm '%s' is unknown.", algo))
			os.Exit(1)
		}
	},
	Example: `dt hash -a bcrypt foo
dt hash -a bcrypt -c 15 foo
dt hash -a md5 foo
dt hash -a sha1 foo
dt hash -a sha3_256 foo
dt hash -a sha3_512 foo
dt hash -a sha256 foo
dt hash -a sha512 foo`,
}

func init() {
	cmd.RootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringP("algo", "a", "md5",
		`the used hash algorithm. Possible values are: bcrypt, md5, sha1, sha256, sha3_256, sha3_512, sha512`)
	hashCmd.Flags().IntVarP(&cost, "cost", "c", 10, "defines the costs of the bcrypt algorithm")
}
