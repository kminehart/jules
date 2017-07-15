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

package config

import (
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strings"
)

const (
	defaultConfigPath = "jules.yaml"
)

const LogFormat = "%12s | %12s | %s\n"

// ProjectList is a list of projects pulled from the config.
type ProjectList map[string]Project

// StageList is a list of stages pulled from the config.
type StageList map[string]Stage

// ServiceList is a simple list of services to run along side the docker run for that stage
type ServiceList map[string]Service

// Stage is a single stage defined in the config.  A stage basically runs a command on your projects.
type Stage struct {
	Command []string `yaml:"command"`
}

// A service is a docker container that runs alongside the specified stages
type Service struct {
	Image   string   `yaml:"image"`
	Env     []string `yaml:"env"`
	Only    []string `yaml:"only"`
	Command []string `yaml:"command"`
}

// Project is essentially a filepath where a stage command is run.
type Project struct {
	Image    string      `yaml:"image"`
	Path     string      `yaml:"path"`
	Stages   StageList   `yaml:"stages"`
	Services ServiceList `yaml:"services"`
	Env      []string    `yaml:"env"`
}

// The Config type defines the structure of the yaml configuration file.
type Config struct {
	Stages   StageList   `yaml:"stages"`
	Projects ProjectList `yaml:"projects"`
}

// Arguments are arguments passed to the binary.
type Arguments struct {
	ConfigPath string
	Stage      string
	Projects   []string
}

// GetArguments uses the "flag" package to parse the command line arguments.
func GetArguments() *Arguments {
	// Command-line Arguments
	var (
		configPath string // -config
		stage      string
		projects   []string
	)

	flag.StringVar(&stage, "stage", "", "Runs a stage.")
	flag.StringVar(&configPath, "config", defaultConfigPath, "-config will specify a custom config file.")
	var p string
	flag.StringVar(&p, "projects", "", "Run the specified stage on a comma deliminated list of projects. (ex: -project project1,project2,project3)")
	flag.Parse()

	if p != "" {
		projects = strings.Split(p, ",")
	}

	return &Arguments{
		ConfigPath: configPath,
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
