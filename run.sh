#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

# go install ?
bin/revel run firefly

export GOPATH="$OLDGOPATH"
