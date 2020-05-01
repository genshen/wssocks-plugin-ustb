#!/bin/sh
set -e

# this script build go api into an archive file (static library)

go build --buildmode=c-archive -o libwssocks_go_api.a
