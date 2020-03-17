FROM fredsimon-docker.jfrog.io/golang:1.14.0

WORKDIR /go/src/app

COPY . .

ENV GOPROXY https://fredsimon.jfrog.io/fredsimon/api/go/go

RUN go mod download
RUN go build -o app -v ./...

ENTRYPOINT ["./app"]
