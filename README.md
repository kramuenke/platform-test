# platform-test

# Getting started

Go >= 1.16 is required

This project uses Go modules. For best results don't check out this project in $GOPATH or use go get

## Running tests

The tests can be invoked locally from the project root with

```
make local-test
```

## Running the app

The web app can be started locally from the project root with

```
make start
```

The app will run on `localhost:8080` and has three endpoints:

 - `/` will output 'Hello World'
 - `/status` can be used for health checks
 - `/metadata` can be use to check git commit sha and version of app.

# Build pipeline

The build pipeline is managed through Circle CI. The setup of the pipeline is managed via `.circleci/config.yml`. We are using CircleCi's standard linux build agents.

It has two sequential steps, `test` and `publish`. `test` will run the test suite in a dockerized environment to avoid managing any build agent configuration. `publish` is building a docker image and pushing the image to [docker hub](https://hub.docker.com/r/kramuenke/platform-test/tags). Images are tagged with CircleCi's build number.

We don't need any dependencies to run our go app so we are building a very thin image based on alpine. The `Dockerfile` has two parts. The first part is building the go application binary and the second part is using the binary in within a bare bones alpine installation. Thus the resulting image is very small (currently less than 6MB).

During building the image we use CircleCI's build-in env variables to pass through the git sha.

# Using the published artifacts

Pull any published artifact

```
docker pull kramuenke/platform-test:<tag>
```

Run the docker image via

```
docker run -p 8080:8080 kramuenke/platform-test:<tag>
```

# Potential improvements

We currently don't know whether the docker image is actually working after it has been build. We could add a step between docker build and push where we run the image and hit the health check endpoint.