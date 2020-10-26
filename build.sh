#!/bin/sh

curdir=`pwd`

echo '###### building All ... ...'
go build -v -ldflags "-w -s" -o ./bin/WebsocketServer github.com/luckyweiwei/websocketserver/executable/WebsocketServer
