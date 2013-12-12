#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

# go install ?
bin/revel build firefly ./exe

export GOPATH="$OLDGOPATH"

