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
	"os/exec"
	"testing"
)

func TestRun(t *testing.T) {
	// Run a stage

	conf := &Config{
		Stages: StageList{
			"configure": {
				Command: []string{
					"make",
					"configure",
				},
			},
			"build": {
				Command: []string{
					"make",
					"build",
				},
			},
			"test": {
				Command: []string{
					"make",
					"test",
				},
			},
			"benchmark": {
				Command: []string{
					"make",
					"benchmark",
				},
			},
			"deploy_staging": {
				Command: []string{
					"make",
					"deploy_staging",
				},
			},
			"deploy_docker": {
				Command: []string{
					"make",
					"deploy_docker",
				},
			},
			"deploy": {
				Command: []string{
					"make",
					"deploy",
				},
			},
		},
		Projects: ProjectList{
			"test1": {
				Path: "path/to/project1",
				Stages: StageList{
					"configure": {
						Command: []string{
							"npm",
							"configure",
						},
					},
				},
				Env: []string{
					"ENV_PROJECT1=value",
				},
			},
			"test2": {
				Path: "./path/to/project2",
				Env: []string{
					"ENV_PROJECT2=value",
				},
			},
		},
	}

	// stage, project, config
	cmd, err := GetCommand("test", "test1", conf)

	if err != nil {
		t.Error(err)
	}

	path, _ := exec.LookPath("make")

	if cmd.Path != path {
		t.Errorf("%s != %s", cmd.Path, path)
	}

	contains := false
	for _, v := range cmd.Env {
		if v == "ENV_PROJECT1=value" {
			contains = true
		}
	}

	if contains != true {
		t.Errorf("cmd for %s on %s did not have the correct environment variable. Environment variables %+v should contain  %s.", "test", "test1", cmd.Env, "ENV_PROJECT1=value")
	}

	// This project overrides a stage
	cmd, err = GetCommand("configure", "test1", conf)
	path, _ = exec.LookPath("npm")
	if cmd.Path != path {
		t.Errorf("%s != %s", cmd.Path, path)
	}
	if err != nil {
		t.Error(err)
	}

	// A project that doesn't exist
	cmd, err = GetCommand("test", "non_existent", conf)
	if err == nil {
		t.Fail()
	}

	// A stage that doesn't exist
	cmd, err = GetCommand("non_existent", "test1", conf)
	if err == nil {
		t.Fail()
	}

	// A nil config
	cmd, err = GetCommand("build", "test1", nil)
	if err == nil {
		t.Fail()
	}
}
