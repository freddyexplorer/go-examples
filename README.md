# go sample
A sample go microservice block for [kintohub](http://kintohub.com)

It is using go `1.14.0`

# First time setup
* `go build main.go`

# Run
- `./main` will launch the server on port `8063`

# Run in docker
- `docker build -t go-test .`
- `docker run -d --name=go-test -p 8063:8063 go-test` 
