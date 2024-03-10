#!/usr/bin/env bash

set -x
unset GOFLAGS
go install github.com/go-bindata/go-bindata/go-bindata@latest
export GOFLAG=S-mod=vendor


go-bindata -pkg template_assets -o pkg/template_assets/bindata.go bindata/templates/...
                                                           


