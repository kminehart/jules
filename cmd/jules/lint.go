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
	"strings"
)

func lintStage(k string, v *Stage) {
	if strings.TrimSpace(k) == "" {
		log.Printf("On stage named %s, the name is empty or is only whitespace.\n", k)
	}

	if len(v.Command) == 0 {
		log.Printf("On stage named %s, the command array has no members.\n", k)
	} else {
		if strings.Contains(v.Command[0], " ") {
			log.Printf("On stage named %s, the first command element has a space. (It definitely won't work). Separate each argument as different members in the array.  \"make test --name=test\" becomes [\"make\", \"test\", \"--name\", \"test\"]\n", k)
		}
	}
}

func lintProject(k string, p *Project) {
	if strings.TrimSpace(k) == "" {
		log.Printf("On project named %s, the name is empty or is only whitespace.\n", k)
	}

	if strings.TrimSpace(p.Path) == "" {
		log.Printf("On project named %s, the name is empty or is only whitespace.\n", k)
	}

	if p.Path[0] == '/' {
		log.Printf("On project named %s, the path is absolute.\n", k)
	}

	for ii, v := range p.Env {
		if strings.Contains(v, "=") != true {
			log.Printf("On project named %s, environment variable index %d, the environment variable does not contain a '='.\n", k, ii)
		}
	}

	for k, v := range p.Stages {
		lintStage(k, &v)
	}
}

// The lint function will print warnings if the desired configuration has any possible issues.
func lint(conf *Config) {
	log.Printf("The following problems were found:\n")
	for k, v := range conf.Stages {
		lintStage(k, &v)
	}

	for k, v := range conf.Projects {
		lintProject(k, &v)
	}
}
