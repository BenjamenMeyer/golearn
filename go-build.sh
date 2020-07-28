#!/bin/bash
set -e

DEPENDENCIES=`cat go.mod`

CGO_ENABLED=1 go build -ldflags "-X 'main.dependency=${DEPENDENCIES}'"
