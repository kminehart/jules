# jules

[![Go Report Card](https://goreportcard.com/badge/github.com/kminehart/jules)](https://goreportcard.com/report/github.com/kminehart/jules)

A pretty basic build system for repositories with multiple projects.

_for best results, use with Docker_

Check out our [cookbook](#cookbook) for some detailed examples and how you can leverage Docker to speed up your build process!

# Documentation

For detailed documentation on Jules, click [here](http://gojules.io/)

# Progress
* [x] -config
* [x] -stage
  * [x] Custom stages in project file.
  * [x] Execute the script specified.
  * [x] Test environment variables
* [x] -lint
* [x] -help
* [ ] -diff

# Install

```
go get github.com/kminehart/jules/cmd/jules
```

# Usage

### Before you begin

_Note that commands ran in these stages are at the working directory specified in your `jules` config._

To run a stage, run `jules -stage=[COMMAND]`:

For a list of commands, see [#commands](#commands).

### Step 1:  Configure your project

In the root of your repository:

`jules.yaml`

```yaml
# Each stage can be ran with 'jules -stage [STAGE]'
stages:
  configure:
    # The 'command' value can be configured with an array (like a Dockerfile)
    # Or with standard yaml syntax (below)
    command: ["make", "configure"]
  build:
    command: ["make", "build"]
  test:
    command: ["make", "test"]
  benchmark:
    command: ["make", "benchmark"]
  deploy_staging:
    command: ["make", "deploy_staging"]
  deploy_docker:
    # Or you can just use normal yaml syntax
    command: 
      - make
      - deploy_docker
  deploy:
    command: ["make", "deploy"]

# Each project will have these stages ran on it.
projects:
  test1:
    # Prefer relative paths to absolute paths.
    # I won't stop you from using absolute paths if you want to do that though.
    path: "path/to/project1"
    stages:
      configure:
        command: ["npm"]
    env:
      # This is technically a []string it just looks like a map.
      - ENV_PROJECT1=value
  test2:
    path: "./path/to/project2"
    # Or JSON syntax.
    env: ["ENV_PROJECT2=value"]
```

### Step 2:  Configure your CI

#### Travis CI

```yml
language: go
```

#### Gitlab CI

```yml
# Use the Debian Jessie image for that package manager
# Ideally though you should use your own docker image so that npm, go, cmake, cargo, etc. don't have to be installed every time.
image: jules:jessie-slim

stages:
  - configure
  - build
  - test
  - deploy

configure:
  stage: configure
  script:
    - jules -stage=configure
    
build:
  stage: build
  script:
    - jules -stage=build
    
test:
  stage: test
  script:
    - jules -stage=test

# You can also specify a custom config file!
deploy_staging:
  stage: deploy
  script:
    - jules -stage=deploy_staging -config=jules.staging.toml
    - jules -stage=deploy_docker -config=jules.staging.toml
  only:
    - development

# Or you can run your custom command.
deploy_production:
  stage: deploy
  script:
    - jules -stage=deploy
    - jules -stage=deploy_docker
  only:
    - master
```

### Step 3: Start committing!

```
jules lint
```

If no config is specified, `jules` will look for a `jules.toml`, and it will output any problems that it finds with it.

```
jules -stage=[STAGE]  -config=[CONFIG]
```

`jules` will run the stage on the specified configuration.

```
jules -stage=[STAGE] -project=[PROJECT1,PROJECT2...]
```

`jules` will run the command on the specified project(s).

```
jules -stage=[STAGE] -diffs
```

If ran in a valid `git` repository, `jules` will only run the specified stage on projects that were modified in the last commit. 

# License

```
    This file is part of "jules".

    "jules" is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    "jules" is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with "jules".  If not, see <http://www.gnu.org/licenses/>.
```

The full GPLv3 can be read [here](LICENSE).
