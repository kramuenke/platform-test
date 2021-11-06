#!/bin/sh

## Usage:
##   . ./export-env.sh ; $COMMAND
##   . ./export-env.sh ; echo ${MINIENTREGA_FECHALIMITE}

# set -a
# . ./app/.env.test
# set +a

cd app
go test ./... -cover