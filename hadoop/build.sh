#!/bin/bash

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -ldflags "-s -w"  -o ./bin/namenode_exporter cmd/namenode_exporter/namenode_exporter.go
go build -ldflags "-s -w"  -o ./bin/datanode_exporter cmd/datanode_exporter/datanode_exporter.go
go build -ldflags "-s -w"  -o ./bin/journalnode_exporter cmd/journalnode_exporter/journalnode_exporter.go
go build -ldflags "-s -w"  -o ./bin/nodemanager_exporter cmd/nodemanager_exporter/nodemanager_exporter.go
go build -ldflags "-s -w"  -o ./bin/resourcemanager_exporter cmd/resourcemanager_exporter/resourcemanager_exporter.go
go build -ldflags "-s -w"  -o ./bin/regionserver_exporter cmd/regionserver_exporter/regionserver_exporter.go
go build -ldflags "-s -w"  -o ./bin/hmaster_exporter cmd/hmaster_exporter/hmaster_exporter.go
