#!/usr/bin/env bash

set -e

make build
sls sam export -o template.yml
sam local start-api
