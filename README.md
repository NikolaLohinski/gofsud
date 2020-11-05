# gofsud
The smallest possible HTTP upload and download file server. 
GoFSUD stands for Golang File Server Upload & Download.

## API
* [OpenAPI specification](https://app.swaggerhub.com/apis/AgentKarbon/gofsud/1.0.0)

## Development
### Requirements
* Install `Go >= 1.15.3`
* Install [`mage`](https://magefile.org/):
  ```shell script
  $ git clone https://github.com/magefile/mage /tmp/mage
  $ cd /tmp/mage
  $ go run bootstrap.go
  ```

### Bootstrap
* Install tools
  ```shell script
  $ mage install
  ```
* Install dependencies and run sanity check
  ```shell script
  $ mage
  ```