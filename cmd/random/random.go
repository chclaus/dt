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
	"github.com/chclaus/dt/cmd"
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/chclaus/dt/utils"
	"github.com/spf13/viper"
	"github.com/chclaus/dt/config"
)

var length int

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generates random numbers and strings",
	Long:  `Generates random numbers and strings. Possible variations are:
  alpha       alphabet
  alphanum    alphanumeric letters
  complex     alphanumeric and special characters
  number      random number`,
	Run: func(cmd *cobra.Command, args []string) {
		r := config.Cfg.Random
		switch r.Algorithm {
		case "alpha":
			fmt.Println(utils.Random(r.Length, utils.Alphabet{}))
			break
		case "alphanum":
			fmt.Println(utils.Random(r.Length, utils.AlphaNumeric{}))
			break
		case "complex":
			fmt.Println(utils.Random(r.Length, utils.Complex{}))
			break
		case "number":
			fmt.Println(utils.RandomNumber(r.Length))
			break
		default:
			fmt.Println(fmt.Errorf("the given algorithm '%s' is unknown.", r.Algorithm))
			os.Exit(1)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(randomCmd)

	randomCmd.Flags().StringP("algo", "a", "complex",
		`the used random algorithm. Possible values are: alpha, alphanum, complex, number`)
	randomCmd.Flags().IntVarP(&length, "length", "l", 20, "defines the length of the generated result")
	viper.BindPFlag("random.algorithm", randomCmd.Flags().Lookup("algo"))
	viper.BindPFlag("random.length", randomCmd.Flags().Lookup("length"))
}
