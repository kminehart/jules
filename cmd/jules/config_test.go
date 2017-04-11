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
	"os"
	"reflect"
	"testing"
)

var testConfig = `
# Each stage can be ran with 'jules -stage [STAGE]'
stages:
  configure:
    # The 'command' value can be configured with an array (like a Dockerfile)
    # Or with standard yaml syntax (below)
    command: ["make", "configure"]
  build:
    command: ["make", "build"]
  test:
    command: ["make", "test"]
  benchmark:
    command: ["make", "benchmark"]
  deploy_staging:
    command: ["make", "deploy_staging"]
  deploy_docker:
    # Or you can just use normal yaml syntax
    command: 
      - make
      - deploy_docker
  deploy:
    command: ["make", "deploy"]

# Each project will have these stages ran on it.
projects:
  test1:
    # Prefer relative paths to absolute paths.
    # I won't stop you from using absolute paths if you want to do that though.
    path: "path/to/project1"
    stages:
      configure:
        command: ["npm", "configure"]
    env:
      # This is technically a []string it just looks like a map.
      - ENV_PROJECT1=value
  test2:
    path: "./path/to/project2"
    # Or JSON syntax.
    env: ["ENV_PROJECT2=value"]
`

func TestConfig(t *testing.T) {
	c, err := ReadConfigString(testConfig)
	if err != nil {
		t.Error(err)
	}

	validConfig := &Config{
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

	if reflect.DeepEqual(*validConfig, *c) != true {
		t.Errorf("%+v\n != %+v\n", *validConfig, *c)
	}
}

func TestArgs(t *testing.T) {
	os.Args = append(os.Args, []string{
		"-projects",
		"test1,test2",
	}...)

	args := GetArguments()
	expect := &Arguments{
		ConfigPath: "jules.yaml",
		Projects: []string{
			"test1",
			"test2",
		},
	}

	if reflect.DeepEqual(*args, *expect) != true {
		t.Errorf("%+v\n != %+v\n", *args, *expect)
	}
}
