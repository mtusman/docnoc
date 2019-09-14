FROM golang:1.11.1-alpine
RUN apk update && apk add --no-cache git gcc musl-dev
RUN mkdir -p /go/src/github.com/mtusman/docnoc
WORKDIR /go/src/github.com/mtusman/docnoc/
COPY . .
RUN go get -d ./...
RUN go install -v
#ENTRYPOINT docnoc -f /tmp/docnoc_config.yaml