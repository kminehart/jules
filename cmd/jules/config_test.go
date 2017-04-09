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
	"reflect"
	"testing"
)

var testConfig = `
stages:
  - name: configure
    # The 'command' value can be configured with an array (like a Dockerfile)
    # Or with standard yaml syntax (below)
    command: ["make", "configure"]
  - name: build
    command: ["make", "build"]
  - name: test
    command: ["make", "test"]
  - name: benchmark
    command: ["make", "benchmark"]
  - name: deploy_staging
    command: ["make", "deploy_staging"]
  - name: deploy_docker
    # Or you can just use normal yaml syntax
    command: 
      - make
      - deploy_docker
  - name: deploy
    command: ["make", "deploy"]

# Each project will have these stages ran on it.
projects:
  - name: test1
    # Prefer relative paths to absolute paths.
    # I won't stop you from using absolute paths if you want to do that though.
    path: "path/to/project1"
    stages:
      - name: configure
        command: ["npm", "configure"]
    env:
      # This is technically a []string it just looks like a map.
      - ENV_PROJECT1=value
  - name: test2
    path: "./path/to/project2"
    # Or JSON syntax.
    env: ["ENV_PROJECT2=value"]`

func TestConfig(t *testing.T) {
	c, err := ReadConfigString(testConfig)
	if err != nil {
		t.Error(err)
	}

	validConfig := &Config{
		Stages: []Stage{
			{
				Name: "configure",
				Command: []string{
					"make",
					"configure",
				},
			},
			{
				Name: "build",
				Command: []string{
					"make",
					"build",
				},
			},
			{
				Name: "test",
				Command: []string{
					"make",
					"test",
				},
			},
			{
				Name: "benchmark",
				Command: []string{
					"make",
					"benchmark",
				},
			},
			{
				Name: "deploy_staging",
				Command: []string{
					"make",
					"deploy_staging",
				},
			},
			{
				Name: "deploy_docker",
				Command: []string{
					"make",
					"deploy_docker",
				},
			},
			{
				Name: "deploy",
				Command: []string{
					"make",
					"deploy",
				},
			},
		},
		Projects: []Project{
			{
				Name: "test1",
				Path: "path/to/project1",
				Stages: []Stage{
					{
						Name: "configure",
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
			{
				Name: "test2",
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
