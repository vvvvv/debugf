#!/usr/bin/env bash

go build -tags delog;
export DELOG_STACKTRACE=ERROR; 
./example02;