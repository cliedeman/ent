#!/usr/bin/env sh
echo Wait for servers to be up

HOSTPARAMS="--host cockroachdb --insecure"
SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

$SQL -e "CREATE DATABASE test;"
$SQL -d test -e "CREATE USER test WITH LOGIN PASSWORD 'pass';"