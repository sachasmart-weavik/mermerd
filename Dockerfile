FROM golang:alpine AS base

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR ${GOPATH}/bin
COPY . .

RUN go build .
