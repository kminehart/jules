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
	"strings"
)

func init() {
	log.SetOutput(os.Stdout)
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
		for _, v := range conf.Stages {
			log.Printf("Running stage %s.\n", v.Name)

			// The user did not specify any projects.
			if len(args.Projects) == 0 {
				for _, p := range conf.Projects {
					log.Printf("Running stage %s on project %s.\n", v.Name, p.Name)
					cmd, err := GetCommand(&v, &p, conf)
					if err != nil {
						log.Fatalf(err.Error())
					}
					log.Printf("%+v\n", cmd)
				}
			} else {
				// Run all stages on specified projects
				for _, p := range args.Projects {
					log.Printf("Running stage %s on project %s.\n", v.Name, p)
					cmd, err := GetCommandFromStrings(v.Name, p, conf)
					if err != nil {
						log.Fatal(err.Error())
					}
					log.Printf("%+v\n", cmd)
				}
			}
		}
	}

	if len(args.Projects) != 0 {
		for _, v := range args.Projects {
			log.Printf("Running stage %s on project %s.\n", args.Stage, v)
			cmd, err := GetCommandFromStrings(args.Stage, v, conf)
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Printf("%+v\n", cmd)
		}
		return
	}

	for _, v := range conf.Projects {
		log.Printf("Running stage %s on project %s.\n", args.Stage, v.Name)
		cmd, err := GetCommandFromStrings(args.Stage, v.Name, conf)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("%+v\n", cmd)
	}

}
