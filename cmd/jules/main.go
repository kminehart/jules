//  This file is part of "jules".
//
//  "jules" is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  "jules" is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with "jules".  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var action string
	if len(os.Args) == 1 {
		action = "all"
	} else {
		action = os.Args[1]
	}

	// These actions require that they be run without checking the config
	switch action {
	case "lint":
		lint()
		return
	case "help":
		help()
		return
	default:
		initConfig()
	}

	// These actions require that they be run with a valid config
	run(action)
}
