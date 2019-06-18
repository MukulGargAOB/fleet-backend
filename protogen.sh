#!/usr/bin/env bash
protoc -I ./fleet-backend/proto  --go_out=. --micro_out=. ./fleet-backend/proto/common.proto
protoc -I ./fleet-backend/proto  --go_out=. --micro_out=. ./fleet-backend/proto/customer.proto
protoc -I ./fleet-backend/proto  --go_out=. --micro_out=. ./fleet-backend/proto/truck.proto