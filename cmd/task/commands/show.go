// Copyright 2018 Mathew Robinson <chasinglogic@gmail.com>. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var show = &cobra.Command{
	Use:     "show",
	Aliases: []string{},
	Short:   "Print detailed task info",
	Run: func(cmd *cobra.Command, args []string) {
		l, err := config.list()
		if err != nil {
			fmt.Println("ERROR Unable to load list:", err)
			os.Exit(1)
		}

		if len(args) != 1 {
			fmt.Println("Not enough arguments, please pass exactly one task ID.")
			os.Exit(1)
		}

		current, err := l.FindByID(args[0])
		if err != nil {
			fmt.Println("ERROR unable to get task:", err)
			return
		}

		fmt.Println(current.IsCompleted())

		switch {
		case titleOnly:
			fmt.Println(current.Title)
		case idOnly:
			fmt.Println(current.ID)
		default:
			fmt.Println(current)
		}
	},
}
