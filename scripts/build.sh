#!/bin/sh
. ./build-config
go build -ldflags "-X main.version=${VERSION}"