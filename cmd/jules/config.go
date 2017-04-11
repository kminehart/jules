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
	defaultDiffs      = ""
)

// ProjectList is a list of projects pulled from the config.
type ProjectList map[string]Project

// StageList is a list of stages pulled from the config.
type StageList map[string]Stage

// Stage is a single stage defined in the config.  A stage basically runs a command on your projects.
type Stage struct {
	Command []string `yaml:"command"`
}

// Project is essentially a filepath where a stage command is run.
type Project struct {
	Path   string    `yaml:"path"`
	Stages StageList `yaml:"stages"`
	Env    []string  `yaml:"env"`
}

// The Config type defines the structure of the yaml configuration file.
type Config struct {
	Order    []string    `yaml:"order"`
	Stages   StageList   `yaml:"stages"`
	Projects ProjectList `yaml:"projects"`
}

// Arguments are arguments passed to the binary.
type Arguments struct {
	ConfigPath string
	Diffs      string
	Stage      string
	Projects   []string
}

// GetArguments uses the "flag" package to parse the command line arguments.
func GetArguments() *Arguments {
	// Command-line Arguments
	var (
		configPath string // -config
		diffs      string // -diffs
		stage      string
		projects   []string
	)

	flag.StringVar(&stage, "stage", "", "Runs a stage.")
	flag.StringVar(&configPath, "config", defaultConfigPath, "-config will specify a custom config file.")
	flag.StringVar(&diffs, "diffs", defaultDiffs, "If in a valid git repository, -diffs [branch] will run a stage only on projects that have been changed when compared to the specified branch.")
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
		return nil, fmt.Errorf("There was an error reading your config file.  Please run \"jules lint\" to find possible problems with your configuration.\n Error: %s", err.Error())
	}
	return config, nil
}

// ReadConfigString will read a string and return a Config object.
func ReadConfigString(conf string) (c *Config, err error) {
	config := &Config{}
	if err = yaml.Unmarshal([]byte(conf), config); err != nil {
		return nil, fmt.Errorf("There was an error reading your config file.  Please run \"jules lint\" to find possible problems with your configuration.\n Error: %s", err.Error())
	}
	return config, nil
}
