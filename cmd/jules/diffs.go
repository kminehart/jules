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
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Executes the provided command, and returns true if it has output.
func ExecuteDiff(cmd *exec.Cmd) (bool, error) {
	var buff bytes.Buffer
	cmd.Stdout = &buff
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	return (buff.Len() > 0), nil
}

// GetDiffCommand returns an os/exec.Cmd struct.  Run it, and it will tell you what files have changed for the specified project.
func GetDiffCommand(path string, branch string) *exec.Cmd {
	var command = []string{
		"git",
		"--no-pager",
		"diff",
		"--name-status",
		fmt.Sprintf("%s...%s", branch, "HEAD"),
		path,
	}

	log.Printf(LogFormat, path, branch, strings.Join(command, " "))
	return exec.Command(command[0], command[1:]...)
}
