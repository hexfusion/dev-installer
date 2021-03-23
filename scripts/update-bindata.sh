#!/usr/bin/env bash

set -x

go-bindata -pkg template_assets -o pkg/template_assets/bindata.go bindata/templates/...
                                                           


