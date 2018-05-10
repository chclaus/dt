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
	"github.com/chclaus/dt/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/chclaus/dt/config"
	"github.com/satori/go.uuid"
	"fmt"
	"os"
	"errors"
)

var version int

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates a UUID",
	Long:  "Generates a UUID.",
	Args: func(cmd *cobra.Command, args []string) error {
		if (config.Cfg.UUID.Version == 3 || config.Cfg.UUID.Version == 5) {
			if len(args) != 1 {
				return errors.New("version 3 and 5 UUIDs needing a namespace and value")
			}

			ns := config.Cfg.UUID.Namespace
			errMsg := "Error: You have to specify a namespace and value. The namespace MUST be a valid UUID"
			if ns == "" {
				return errors.New(errMsg)
			}
			if _, err := uuid.FromString(ns); err != nil {
				return errors.New(errMsg)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch config.Cfg.UUID.Version {
		case 1:
			v1 := uuid.NewV1()
			fmt.Println(v1.String())
			break
		case 2:
			domain := cmd.Flag("domain").Value.String()
			switch domain {
			case "user":
				v2 := uuid.NewV2(uuid.DomainPerson)
				fmt.Println(v2.String())
				break
			case "group":
				v2 := uuid.NewV2(uuid.DomainGroup)
				fmt.Println(v2.String())
				break
			}
			break
		case 3:
			ns := config.Cfg.UUID.Namespace
			nsUUID, _ := uuid.FromString(ns)
			fmt.Println(uuid.NewV3(nsUUID, args[0]).String())
		case 4:
			v4 := uuid.NewV4()
			fmt.Println(v4.String())
			break
		case 5:
			ns := config.Cfg.UUID.Namespace
			nsUUID, _ := uuid.FromString(ns)
			fmt.Println(uuid.NewV5(nsUUID, args[0]).String())
		default:
			fmt.Println("UUID version must be between 1 and 5")
			os.Exit(1)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(uuidCmd)
	cobra.OnInitialize(config.InitConfig)

	uuidCmd.Flags().StringP("namespace", "n", "", "the namespace that should be hashed in UUID v3 or UUID v5. It should be a domain or application specific UUID")
	uuidCmd.Flags().IntVarP(&version, "version", "v", 4, "defines the version of the generated uuid")
	uuidCmd.Flags().StringP("domain", "d", "user", "the used domain for UID/GID in UUID v2. Valid values are: user, group")
	viper.BindPFlag("uuid.namespace", uuidCmd.Flags().Lookup("namespace"))
	viper.BindPFlag("uuid.version", uuidCmd.Flags().Lookup("version"))
}
