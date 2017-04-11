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

func ExecuteCommand(stage string, project string, cmd *exec.Cmd) error {
	var buff bytes.Buffer

	cmd.Stdout = &buff
	err := cmd.Run()
	log.Printf(LogFormat, stage, project, buff.String())
	return err
}

// GetCommand will return an "os/exec" command based on the stage, project, and configuration provided.
func GetCommand(stage string, project string, conf *Config) (*exec.Cmd, error) {
	if conf == nil {
		return nil, fmt.Errorf("You can not provide a nil config.\n")
	}

	var (
		command []string
	)

	if _, ok := conf.Projects[project]; ok != true {
		return nil, fmt.Errorf("Project %s not found in the config.", project)
	}

	// Step 1:  See if the project has the stage.
	if val, ok := conf.Projects[project].Stages[stage]; ok {
		command = val.Command
	}

	// The project didn't have the command
	if len(command) == 0 {
		if val, ok := conf.Stages[stage]; ok {
			command = val.Command
		}
	}

	if len(command) == 0 {
		return nil, fmt.Errorf("Could not find a suitable command. Please check your config file.\n")
	}

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = conf.Projects[project].Env
	cmd.Dir = conf.Projects[project].Path

	log.Printf(LogFormat, project, stage, strings.Join(cmd.Args, " "))
	return cmd, nil
}
