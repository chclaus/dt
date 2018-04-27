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
	"github.com/spf13/cobra"
	"github.com/satori/go.uuid"
	"fmt"
)

// uuidV2Cmd represents the uuidV2 command
var uuidV2Cmd = &cobra.Command{
	Use:   "v2",
	Short: "Generates a UUID Version 2",
	Long: "Generates a v2 UUID, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)",
	Run: func(cmd *cobra.Command, args []string) {
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
		default:
			cmd.Help()
		}
	},
	Example: "dt uuid v3 -n cacae610-c76a-4736-90ef-0271126b4345 -v foo",
}

func init() {
	uuidCmd.AddCommand(uuidV2Cmd)

	uuidV2Cmd.Flags().StringP("domain", "d", "user", "the used domain for UID/GID. Valid values are: [user, group]")
}
