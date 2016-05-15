#!/bin/sh

export GOPATH=$GOPATH:`pwd`
echo $GOPATH


go test -v -run=TestAnsiText2