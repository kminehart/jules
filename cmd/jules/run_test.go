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

var runConfig = `
# Each stage can be ran with 'jules -stage [STAGE]'
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

func TestRun(t *testing.T) {
	// Run a stage

	conf, err := ReadConfigString(runConfig)

	if err != nil {
		t.Error(err)
	}

	// stage, project, config
	cmd, err := GetCommandFromStrings("test", "test1", conf)

	if err != nil {
		t.Error(err)
	}

	path, _ := exec.LookPath("make")

	if cmd.Path != path {
		t.Errorf("%s != %s", cmd.Path, path)
	}

	if cmd.Env[0] != "ENV_PROJECT1=value" {
		t.Errorf("cmd for %s on %s did not have the correct environment variable. Got %s, wanted %s.", "test", "test1", cmd.Env[0], "ENV_PROJECT1=value")
	}

	// This project overrides a stage
	cmd, err = GetCommandFromStrings("configure", "test1", conf)
	path, _ = exec.LookPath("npm")
	if cmd.Path != path {
		t.Errorf("%s != %s", cmd.Path, path)
	}
	if err != nil {
		t.Error(err)
	}

	// A project that doesn't exist
	cmd, err = GetCommandFromStrings("test", "non_existent", conf)
	if err == nil {
		t.Fail()
	}

	// A stage that doesn't exist
	cmd, err = GetCommandFromStrings("non_existent", "test1", conf)
	if err == nil {
		t.Fail()
	}

	// A nil config
	cmd, err = GetCommandFromStrings("build", "test1", nil)
	if err == nil {
		t.Fail()
	}
}
