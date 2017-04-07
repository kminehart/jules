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
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

var Configuration Config

const (
	defaultConfigPath = "jules.toml"
	defaultLogLevel   = "info"
	defaultDiffs      = false
)

// Command-line Arguments
var (
	configPath string // -config
	logLevel   string // -logLevel
	stage      string // -stage
	diffs      bool   // -diffs
)

type Project struct {
	Name        string `toml:"name"`
	Path        string `toml:"path"`
	Configure   string `toml:"configure"`
	Build       string `toml:"build"`
	Test        string `toml:"test"`
	Deploy      string `toml:"deploy"`
	Environment string `toml:"env"`
}

type Config struct {
	Configure string            `toml:"configure"`
	Build     string            `toml:"build"`
	Test      string            `toml:"test"`
	Deploy    string            `toml:"deploy"`
	Custom    map[string]string `toml:"custom"`
	Projects  []Project         `toml:"projects"`
}

func initArguments() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "-config will specify a custom config file.")
	flag.StringVar(&logLevel, "log-level", defaultLogLevel, "-log-level controls the verbosity of the output.  Options: \"debug\", \"info\"(default), \"warn\", and \"error\".")
	flag.StringVar(&stage, "stage", "", "-stage will specify a custom stage to run.")
	flag.BoolVar(&diffs, "diffs", false, "If in a valid git repository, -diffs will run a stage only on projects that have been changed.")

	flag.Parse()
}

func initConfig() {
	initArguments()
}

func config(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("Invalid config file.  Could not open \"%s\"\n", path)
	}

	if err := toml.Unmarshal(data, &Configuration); err != nil {
		log.Panicln("There was an error reading your config file.  Please run \"jules lint\" to find possible problems with your configuration.")
	}
}
