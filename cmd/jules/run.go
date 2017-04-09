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
	"os/exec"
	"strings"
)

func GetCommand(stage *Stage, project *Project, conf *Config) (*exec.Cmd, error) {
	if stage == nil {
		return nil, fmt.Errorf("You can not provide a nil stage.\n")
	}
	if project == nil {
		return nil, fmt.Errorf("You can not provide a nil project.\n")
	}
	if conf == nil {
		return nil, fmt.Errorf("You can not provide a nil config.\n")
	}

	var (
		command []string
	)
	n := strings.ToLower(stage.Name)
	for _, v := range project.Stages {
		if strings.ToLower(v.Name) == n {
			command = v.Command
			break
		}
	}

	if len(command) == 0 {
		command = stage.Command
	}

	if len(command) == 0 {
		return nil, fmt.Errorf("Could not find a suitable command. Please check your config file.\n")
	}

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = project.Env
	cmd.Dir = project.Path
	return cmd, nil
}

func GetCommandFromStrings(stage string, project string, conf *Config) (*exec.Cmd, error) {
	if conf == nil {
		return nil, fmt.Errorf("You can not provide a nil config.\n")
	}

	var (
		command []string
		p       *Project
	)

	// Step 1:  See if the project has the stage.
	project = strings.ToLower(project)
	for i, v := range conf.Projects {
		if strings.ToLower(v.Name) == project {
			p = &conf.Projects[i]
			for _, s := range v.Stages {
				if s.Name == stage {
					command = s.Command
					break
				}
			}
		}
	}

	if p == nil {
		return nil, fmt.Errorf("%s is not a project found in the provided config.", project)
	}

	// The project didn't have the command
	if len(command) == 0 {
		for _, s := range conf.Stages {
			if s.Name == stage {
				command = s.Command
			}
		}
	}

	if len(command) == 0 {
		return nil, fmt.Errorf("Could not find a suitable command. Please check your config file.\n")
	}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = p.Env
	cmd.Dir = p.Path
	return cmd, nil
}
