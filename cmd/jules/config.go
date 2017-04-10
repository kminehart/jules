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
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strings"
)

const (
	defaultConfigPath = "jules.yaml"
	defaultDiffs      = false
	defaultStage      = "all"
)

// A stage is a single stage defined in the config.  A stage basically runs a command on your projects.
type Stage struct {
	Name    string   `yaml:"name"`
	Command []string `yaml:"command"`
}

// Each project is essentially a filepath where a stage command is run.
type Project struct {
	Name   string   `yaml:"name"`
	Path   string   `yaml:"path"`
	Stages []Stage  `yaml:"stages"`
	Env    []string `yaml:"env"`
}

// The Config type defines the structure of the yaml configuration file.
type Config struct {
	Stages   []Stage   `yaml:"stages"`
	Projects []Project `yaml:"projects"`
}

// Arguments are arguments passed to the binary.
type Arguments struct {
	ConfigPath string
	Diffs      bool
	Stage      string
	Projects   []string
}

// GetArguments uses the "flag" package to parse the command line arguments.
func GetArguments() *Arguments {
	// Command-line Arguments
	var (
		configPath string // -config
		diffs      bool   // -diffs
		stage      string
		projects   []string
	)

	flag.StringVar(&stage, "stage", defaultStage, "Runs a stage. If not specified, then jules will run all of the stages it can find.")
	flag.StringVar(&configPath, "config", defaultConfigPath, "-config will specify a custom config file.")
	flag.BoolVar(&diffs, "diffs", false, "If in a valid git repository, -diffs will run a stage only on projects that have been changed.")
	var p string
	flag.StringVar(&p, "projects", "", "Run the specified stage on a comma deliminated list of projects. (ex: -project project1,project2,project3)")
	flag.Parse()

	if p != "" {
		projects = strings.Split(p, ",")
	}

	return &Arguments{
		ConfigPath: configPath,
		Diffs:      diffs,
		Stage:      stage,
		Projects:   projects,
	}
}

// ReadConfig will open a filepath and return a Config object.
func ReadConfig(path string) (c *Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Invalid config file.  Could not open \"%s\"\n", path)
	}

	config := &Config{}
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("There was an error reading your config file.  Please run \"jules lint\" to find possible problems with your configuration.")
	}
	return config, nil
}

// ReadConfigString will read a string and return a Config object.
func ReadConfigString(conf string) (c *Config, err error) {
	config := &Config{}
	if err = yaml.Unmarshal([]byte(conf), config); err != nil {
		return nil, fmt.Errorf("There was an error reading your config file.  Please run \"jules lint\" to find possible problems with your configuration.")
	}
	return config, nil
}
