#!/usr/bin/env bash

rm -f ./example03_no_tag.bin
rm -f ./example03_delog.bin
rm -f ./example03_no_tag_objdump 
rm -f ./example03_delog_objdump
go build -o example03_no_tag.bin
go build -tags delog -o example03_delog.bin

go tool objdump -S -s delog example03_no_tag.bin > example03_no_tag_objdump
go tool objdump -S -s delog example03_delog.bin > example03_delog_objdump
