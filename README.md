# jules
A pretty basic build system for repositories with multiple projects.

_for best results, use with Docker_

# Install

# Usage

### Before you begin

By default, there's 4 actions that `jules` will do:

_Note that commands ran in these stages are at the working directory specified in your `jules` config._

To run them, run `jules [COMMAND]`:

1. configure
2. build
3. test
4. deploy

For a list of commands, see [#commands](#commands).

### Step 1:  Configure your project

In the root of your repository:

`jules.toml`

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
image: jules:jessie-slim

stages:
  - configure
  - build
  - test
  - deploy

configure:
  stage: configure
  script:
    - jules configure
    
build:
  stage: build
  script:
    - jules build
    
test:
  stage: test
  script:
    - jules test

# You can also specify a custom config file!
deploy_staging:
  stage: deploy
  script:
    - jules deploy --config=jules.staging.toml
  only:
    - development

# Or you can run your custom command.
deploy_production:
  stage: deploy
  script:
    - jules deploy
    - jules -c deploy_docker
  only:
    - master
```

### Step 3: Start committing!

# Commands

```
jules [configure|build|test|deploy] [PROJECT]
```

Runs one of the 4 stages.  If no config is specified, then `jules` will look for a `jules.toml`. 

If it exists, then it will run the specified stage on all of the projects listed.

If `[PROJECT]` is provided, then `jules` will run on the specified project.

```
jules lint
```

If no config is specified, `jules` will look for a `jules.toml`, and it will output any problems that it finds with it.

```
jules [COMMAND] -c [CONFIG], jules [COMMAND] --config=[CONFIG]
```

`jules` will run the command with the specified configuration.

```
jules [COMMAND] -l [debug|info|warning|error]`, `jules [COMMAND] --log-level=[debug|info|warning|error]
```

`jules` will provide output at a specific level.  The defailt level is `info`.

```
jules -s [STAGE]`, `jules --stage=[STAGE]
```

`jules` will run your custom command.

```
jules [COMMAND] -d`, `jules --diffs
```

If ran in a valid `git` repository, `jules` will only run the specified stage on projects that were modified in the last commit. 
