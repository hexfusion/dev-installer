#!/usr/bin/env bash

# make clean
rm -f ./bin/dev-installer

go build -o ./bin/dev-installer ./cmd/dev-installer/dev-installer.go

sudo cp ./bin/dev-installer /usr/local/bin/

echo "install complete: $(which dev-installer)"
