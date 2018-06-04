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
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

// Thanks to the kubectl authors for a great template
// https://github.com/kubernetes/kubernetes/blob/0207a090743398608b357264ab4c27aeea864abc/pkg/kubectl/cmd/completion.go

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash and zsh completion scripts",
	Long: `Installing shell completion on Linux:

  bash:
  =====
  Load the dt completion code for bash into the current shell:

    source <(dt completion --shell bash)

  Write bash completion code to a file and source if from .bash_profile

    dt completion --shell bash > ~/.dt_completion.sh
    printf "\n# dt bash completion\nsource '$HOME/.dt_completion.sh'" >> $HOME/.bash_profile
    source $HOME/.bash_profile

  zsh:
  =====
  Load the dt completion code for zsh into the current shell
  
    source <(dt completion --shell zsh)

  Set the dt completion code for zsh to autoload on startup
  
    dt completion --shell zsh > "${fpath[1]}/_dt"`,
	Run: func(cmd *cobra.Command, args []string) {

		shell, err := cmd.Flags().GetString("shell")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch shell {
		case "bash":
			RootCmd.GenBashCompletion(os.Stdout)
			break
		case "zsh":
			RootCmd.GenZshCompletion(os.Stdout)
			break
		default:
			cmd.Help()
			os.Exit(1)
			break
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)

	completionCmd.Flags().StringP("shell", "s", "",
		"the shell a completion should be generated for. Valid values are 'bash' and 'zsh'.")
}
