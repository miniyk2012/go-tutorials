# Do not run or build this Dockerfile directly. See service dev in docker-compose.yml
FROM golang:1.10-alpine3.8

ARG WORKSPACE_DIR

RUN apk add make
RUN apk add git
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/${WORKSPACE_DIR}
