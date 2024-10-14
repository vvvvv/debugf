#!/usr/bin/env bash

rm -f example02.bin
go build -tags delog -o example02.bin -ldflags "-X 'github.com/vvvvv/delog.DisableDebugWarning=1'"
export DELOG_STACKTRACE=ERROR
./example02.bin
