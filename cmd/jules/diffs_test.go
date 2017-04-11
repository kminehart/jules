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
	"fmt"
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	path := "/test/path"

	// Compare with origin/master
	args := &Arguments{
		Diffs: "origin/master",
	}

	// The git command to run to check for diffs.
	command := []string{
		"git",
		"--no-pager",
		"diff",
		"--name-status",
		fmt.Sprintf("%s...%s", args.Diffs, "HEAD"),
		path,
	}

	cmd := GetDiffCommand(path, args.Diffs)

	if reflect.DeepEqual(cmd.Args, command) != true {
		t.Errorf("Error:  Expected %+v\n, Got %+v\n", command, cmd.Args)
	}
}
