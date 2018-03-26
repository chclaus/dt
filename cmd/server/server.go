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

package server

import (
	"errors"
	"fmt"
	"github.com/chclaus/dt/cmd"
	"github.com/spf13/cobra"
	"net/http"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a simple web server to serve static content",
	Long:  "Starts a simple web server to serve static content",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You have to specify a folder with web content")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir(args[0]))
		http.Handle("/", fs)

		hostname := cmd.Flag("address").Value.String()
		port := cmd.Flag("port").Value.String()
		addr := fmt.Sprintf("%s:%s", hostname, port)

		fmt.Printf("Serving %s on %s\n", args[0], addr)

		http.ListenAndServe(addr, nil)
	},
	Example: ``,
}

func init() {
	cmd.RootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("address", "a", "0.0.0.0", "the hostname or ip address")
	serverCmd.Flags().StringP("port", "p", "3000", "the listening port")
}
