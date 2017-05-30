NAME=fpm
VERSION=0.0.1
RELEASE=alpha
GITID=$(shell git rev-parse --short=8 HEAD)
GOVERSION=$(shell go version|cut -f3 -d' ')

dependencies:
	@go get -u github.com/kardianos/govendor
	@govendor add +external
	@govendor fetch +missing
