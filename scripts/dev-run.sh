#!/usr/bin/env bash

set -e

make build
node_modules/.bin/sls sam export -o template.yml
sam local start-api
