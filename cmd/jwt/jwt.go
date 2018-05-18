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

package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chclaus/dt/cmd"
	"github.com/chclaus/dt/config"
	"github.com/chclaus/dt/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Allows decoding of a jwt",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You have to specify a base64 encoded jwt")
		}

		return nil
	},
	Long: "Allows decoding of a jwt.",
	Run: func(cmd *cobra.Command, args []string) {
		jwtString := args[0]
		parts := strings.Split(jwtString, ".")
		if len(parts) != 3 {
			log.Fatal("Invalid JWT. It must has a JOSE Header, JWS Payload and JWS Signature")
		}

		printJWT(parts)

		secret, err := determineSecret()
		if err != nil {
			log.Fatal(err)
		}

		if secret != nil {
			_, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				red := color.New(color.FgRed)
				red.Printf("\nOh no! %s.\n", err)
				fmt.Printf("Your chosen secret is: '%s'\n", secret)
			} else {
				green := color.New(color.FgGreen)
				green.Printf("\ntoken signature is valid.\n")
			}
		}
	},
	// secret: foobar
	Example: `dt jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmb28iLCJzdWIiOiJiYXIifQ.UxyRHFY_BpuDQ1Qp9MVvbn5uAlaoWCUKUIeq1qQIcCw
dt jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmb28iLCJzdWIiOiJiYXIifQ.UxyRHFY_BpuDQ1Qp9MVvbn5uAlaoWCUKUIeq1qQIcCw -s foobar`,
}

func determineSecret() ([]byte, error) {
	jwtCfg := config.Cfg.JWT

	if jwtCfg.Secret != "" {
		if jwtCfg.Base64Secret {
			return []byte(utils.DecodeBase64(base64.StdEncoding, jwtCfg.Secret)), nil
		}

		return []byte(jwtCfg.Secret), nil
	}

	if jwtCfg.SecretFile != "" {
		if _, err := os.Stat(jwtCfg.SecretFile); err != nil {
			return nil, fmt.Errorf("secret file doesn't exist: %s", err)
		}

		b, err := ioutil.ReadFile(jwtCfg.SecretFile)
		if err != nil {
			return nil, fmt.Errorf("error reading secret file: %s", err)
		}

		b = bytes.TrimSpace(b)
		if jwtCfg.Base64Secret {
			return []byte(utils.DecodeBase64(base64.StdEncoding, string(b))), nil
		}

		return b, nil
	}

	return nil, nil
}

func printJWT(parts []string) {
	fmt.Println("JOSE Header:")
	fmt.Println(prettifyPart(parts[0]))
	fmt.Printf("\nJWS Payload:\n")
	fmt.Println(prettifyPart(parts[1]))
}

// prettifyPart decodes the base64 string and generates a pretty, colorful representation of the resulting json
func prettifyPart(part string) string {
	decodedPart := utils.DecodeBase64(base64.RawURLEncoding, part)
	var tmp interface{}
	json.Unmarshal([]byte(decodedPart), &tmp)
	prettyBytes, _ := prettyjson.Marshal(tmp)

	return string(prettyBytes)
}

func init() {
	cmd.RootCmd.AddCommand(jwtCmd)

	jwtCmd.Flags().StringP("secret", "s", "", "the secret to validate the token signature")
	jwtCmd.Flags().StringP("secretFile", "f", "", "a file of the secret to validate the token signature")
	jwtCmd.Flags().BoolP("base64Secret", "b", false, "set to true if the secret is base64 encoded")
	viper.BindPFlag("jwt.secret", jwtCmd.Flags().Lookup("secret"))
	viper.BindPFlag("jwt.secretFile", jwtCmd.Flags().Lookup("secretFile"))
	viper.BindPFlag("jwt.base64Secret", jwtCmd.Flags().Lookup("base64Secret"))
}
