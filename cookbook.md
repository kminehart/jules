# Cookbook

* Multiple projects in many different languages:

*Solution*: Makefiles!

jules.yaml:

```yml
stages:
  - name: configure
    command: ["make", "configure"]
  - name: build
    command: ["make", "build"]
  - name: test
    command: ["make", "test"]
  - name: deploy
    command: ["make", "deploy"]
  - name: deploy_docker
    command: ["make", "deploy_docker"]
  - name: deploy_staging
    command: ["make", "deploy_staging"]

projects:
  - name: project1
    path: "projects/project1"
  - name: project2
    path: "projects/project2"
  - name: project3
    path: "projects/project3"
```

And each `Makefile` looks like this (this one is for Go, but each `Makefile` can be fine-tuned to build each project individually)

```Makefile
configure:
  go get -d v

build:
  go build -o "name"

test:
  go test ./...

deploy:
  scp -R . app@app-server.local:/app

deploy_staging:
  scp -R . app@staging-server.local:/app

deploy_docker:
  docker push
```

And for your gitlab CI:

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
    - jules -stage configure
    
build:
  stage: build
  script:
    - jules -stage build
    
test:
  stage: test
  script:
    - jules -stage test

# You can also specify a custom config file!
deploy_staging:
  stage: deploy
  script:
    - jules -stage=deploy_staging -config jules.staging.toml
    - jules -stage=deploy_docker -config jules.staging.toml
  only:
    - development

# Or you can run your custom command.
deploy_production:
  stage: deploy
  script:
    - jules -stage deploy
    - jules -stage deploy_docker
  only:
    - master
```

