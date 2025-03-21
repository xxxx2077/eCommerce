#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/echoservice"
exec "$CURDIR/bin/echoservice"
