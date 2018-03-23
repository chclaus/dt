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

package date

import (
	"errors"
	"fmt"
	"github.com/chclaus/dt/cmd"
	"github.com/chclaus/dt/utils"
	"github.com/spf13/cobra"
	"log"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date [the date]",
	Short: "Basic date operations",
	Long:  "Basic date operations.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You have to specify a date")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dateInput := args[0]

		time, e := utils.ParseTimestamp(dateInput)
		if e != nil {
			log.Fatal(e)
		}

		fmt.Printf("Unix timestamp (millis): %d\n", time.Unix())
		fmt.Printf("Unix timestamp (nanos): %s\n", time.String())
		fmt.Printf("Day: %d\n", time.Day())
		fmt.Printf("Day of week: %d\n", time.Weekday())
		fmt.Printf("Day of year: %d\n", time.YearDay())
		fmt.Printf("Month: %d\n", time.Month())
		fmt.Printf("Year: %d\n", time.Year())
	},
	Example: `dt date 2018-03-23T16:55:15+02:00
dt date 1521816915
dt date 1521816915000`,
}

func init() {
	cmd.RootCmd.AddCommand(dateCmd)
}
