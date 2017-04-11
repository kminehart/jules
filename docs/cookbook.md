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
