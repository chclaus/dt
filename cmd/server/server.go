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
	"path/filepath"
	"os"
	"github.com/spf13/viper"
	"github.com/chclaus/dt/config"
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

		srv := config.Cfg.Server
		addr := fmt.Sprintf("%s:%s", srv.Address, srv.Port)

		path, err := filepath.Abs(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Serving %s on %s\n", path, addr)

		http.ListenAndServe(addr, nil)
	},
	Example: ``,
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	cmd.RootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("address", "a", "0.0.0.0", "the hostname or ip address")
	serverCmd.Flags().StringP("port", "p", "3000", "the listening port")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.address", serverCmd.Flags().Lookup("address"))
}
