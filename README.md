# joules
A pretty basic build system for repositories with multiple projects.

_for best results, use with Docker_

# Install

# Usage

### Before you begin

By default, there's 4 actions that `joules` will do:

_Note that commands ran in these stages are at the working directory specified in the config below._

To run them, run `joules [COMMAND]`:

1. configure
2. build
3. test
4. deploy

For a list of commands, see [#commands](#commands).

### Step 1:  Configure your project

```toml

# The commands to run to configure, build, test, and deploy a project.
configure = "go get -d -v"
build = "go build $j"
test = "go test ./..."
deploy = "./deploy.sh"

# You can also specify custom commands.
[custom]
deploy_staging = "./deploy_staging.sh"
deploy_docker = "./deploy_docker.sh"

[[projects]]
# Note:  name must be unique.
name = "project1"
path = "pkg/core/project1-frontend"
# You can also modify the 4 stages for specifc projects
configure = "npm install"
build = "webpack"
test = "npm run test"
deploy = "npm run deploy"


[[projects]]
name = "project2"
path = "pkg/core/project2"
```

### Step 2:  Configure your CI

#### Travis CI
```yml
language: go
```

#### Gitlab CI

```yml
# Use the Debian Jessie image for that package manager
image: joules:jessie-slim

stages:
  - configure
  - build
  - test
  - deploy

configure:
  stage: configure
  script:
    - joules configure
    
build:
  stage: build
  script:
    - joules build
    
test:
  stage: test
  script:
    - joules test

# You can also specify a custom config file!
deploy_staging:
  stage: deploy
  script:
    - joules deploy --config=joules.staging.toml
  only:
    - development

# Or you can run your custom command.
deploy_production:
  stage: deploy
  script:
    - joules deploy
    - joules -c deploy_docker
  only:
    - master
```

### Step 3: Start committing!

# Commands

* `joules configure`, `joules build`, `joules test`, `joules deploy`
Runs one of the 4 stages.  If no other options are specified, then `joules` will look for a `joules.toml`. If it exists, then it will run the specified stage on all of the projects listed.

* `joules lint`
If no config is specified, `joules` will look for a `joules.toml`, and it will output any problems that it finds with it.

* `joules [COMMAND] --config=[CONFIG]`
`joules` will run the command with the specified configuration.

* `joules [COMMAND] --log-level=[debug|info|warning|error]`
`joules` will provide output at a specific level.  The defailt level is `info`.

* `joules -c [COMMAND]`
`joules` will run your custom command.
