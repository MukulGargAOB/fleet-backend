#!/bin/bash

pids=$(ps ax | grep 'go run' | awk '{ print $1}')
for p in $pids; do
    kill -s TERM $p
done

netPids=$(netstat -npl | grep 'main' | awk '{ print $7 }' | grep -oP '\d*(?=/main)')
for p in $netPids; do
    kill -s TERM $p
done

kill -9 `lsof -t -i:20022`
