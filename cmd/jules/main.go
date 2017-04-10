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
	"log"
	"os"
	"os/exec"
	"strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func execute(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func runAll(conf *Config, args *Arguments) error {
	projects := args.Projects

	if len(args.Projects) == 0 {
		projects = make([]string, len(conf.Projects))
		for i, v := range conf.Projects {
			projects[i] = v.Name
		}
	}

	// For each stage defined in the configuration
	for _, v := range conf.Stages {
		log.Printf("Running stage %s.\n", v.Name)
		for _, p := range projects {
			log.Printf("Running stage %s on project %s.\n", v.Name, p)
			cmd, err := GetCommand(v.Name, p, conf)
			if err != nil {
				return err
			}

			err = execute(cmd)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	args := GetArguments()

	conf, err := ReadConfig(args.ConfigPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Lint?
	for _, v := range os.Args {
		if strings.ToLower(v) == "lint" {
			lint(conf)
			return
		}
		if strings.ToLower(v) == "help" {
			help()
			return
		}
	}

	if args.Stage == "all" {
		err = runAll(conf, args)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if len(args.Projects) != 0 {
		for _, v := range args.Projects {
			log.Printf("Running stage %s on project %s.\n", args.Stage, v)
			cmd, err := GetCommand(args.Stage, v, conf)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = execute(cmd)

			if err != nil {
				log.Fatal(err.Error())
			}
		}
		return
	}

	for _, v := range conf.Projects {
		log.Printf("Running stage %s on project %s.\n", args.Stage, v.Name)
		cmd, err := GetCommand(args.Stage, v.Name, conf)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = execute(cmd)

		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
