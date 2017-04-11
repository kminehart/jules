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

func run(stage string, projects []string, conf *Config, args *Arguments) error {
	for _, p := range projects {
		cmd, err := GetCommand(stage, p, conf)

		if err != nil {
			return err
		}

		err = execute(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	args := GetArguments()

	conf, err := ReadConfig(args.ConfigPath)

	if args.Stage == "" {
		flag.Usage()
		return
	}

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

	projects := args.Projects

	// If no projects were specified, then use all of them
	if len(args.Projects) == 0 {
		i := 0
		projects = make([]string, len(conf.Projects))
		for k, _ := range conf.Projects {
			projects[i] = k
			i++
		}
	}

	// If the user provided the "-diffs" flag, find the projects with changes.
	if args.Diffs != "" {
		for i := len(projects) - 1; i >= 0; i-- {
			if val, ok := conf.Projects[projects[i]]; ok {
				isDifferent, err := ExecuteDiff(GetDiffCommand(val.Path, args.Diffs))
				if err != nil {
					log.Fatalf("Something went wrong when trying to git diff.  Do you have `git` installed? Error: %s\n", err.Error())
				}

				// If there was no diff found for this project.
				if isDifferent != true {
					if len(projects) == 1 {
						log.Fatalf("No projects were found with diffs against the branch %s\n", args.Diffs)
					}
					projects = append(projects[:i], projects[i+1:]...)
				}
			}
		}
	}

	err = run(args.Stage, projects, conf, args)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
