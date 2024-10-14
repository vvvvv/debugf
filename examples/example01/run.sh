#!/usr/bin/env bash

rm -f ./example01.bin
go build -tags delog -o example01.bin 
./example01.bin
