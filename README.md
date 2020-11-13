# gofsud
The smallest possible HTTP upload and download file server. 
GoFSUD stands for Golang File Server Upload & Download.

## API
* [OpenAPI specification](https://app.swaggerhub.com/apis/AgentKarbon/gofsud)

## Quick start

### Docker
```shell script
$ docker run -it --rm --init \
    --user ${UID} \
    --publish "8080:8080" \
    --env GOFSUD_DIRECTORY='/web' \
    --volume /tmp/gofsud:/web \
    theagentk/gofsud:latest
```

### Golang
Follow the development instructions, and then run:
```shell script
$ mage bin:run
```

## Development
### Requirements
* Install `Go >= 1.15.3`
* Install [`mage`](https://magefile.org/):
  ```shell script
  $ git clone https://github.com/magefile/mage /tmp/mage
  $ cd /tmp/mage
  $ go run bootstrap.go
  ```
* For image building, you can define the following environment variables:

  |        Name        |             Default             | Description                                          |
  |--------------------|---------------------------------|------------------------------------------------------|
  | GO_IMAGE_VERSION   | 1.15.3-alpine                   | Base Golang docker image version to use for building |
  | DISTROLESS_IMAGE   | gcr.io/distroless/base-debian10 | Base image to use for running the app                |
  | DISTROLESS_VERSION | nonroot                         | Version of the image for running the app             |
  | IMAGE_DESTINATION  | theagentk/gofsud                | Name of the resulting image to build                 |

### Bootstrap
* Install tools
  ```shell script
  $ mage install
  ```
* Install dependencies and run sanity check
  ```shell script
  $ mage
  ```
