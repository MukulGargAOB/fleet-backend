#!/bin/bash

rm -rf /tmp/fms-*

listVar="customer-service truck-service"
for i in $listVar; do
    go run "$i"/main.go &
done