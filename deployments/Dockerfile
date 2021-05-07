from golang:alpine

run go get github.com/cosmtrek/air

env GO111MODULE=on
env GOFLAGS=-mod=vendor

user root
workdir /go/src/app

expose 8080
cmd ["air"]