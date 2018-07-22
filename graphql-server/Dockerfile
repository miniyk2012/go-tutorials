# Do not run or build this Dockerfile directly. See service dev in docker-compose.yml
# todo change to alphine
FROM golang:1.10

ARG DEV_EXPOSE_PORT
ARG DEV_EXPOSE_PORT_FE
ARG DEV_WORKSPACE

ENV LANG C.UTF-8

RUN go get -u github.com/golang/dep/cmd/dep
RUN curl https://getcaddy.com | bash -s personal

WORKDIR $GOPATH/src/$DEV_WORKSPACE

EXPOSE $DEV_EXPOSE_PORT
EXPOSE $DEV_EXPOSE_PORT_FE
