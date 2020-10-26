#!/bin/sh

killproc(){
  pkill -9 $1
  sleep 1

  var=`ps -aef | grep -w $1 | grep -v sh| grep -v grep| awk '{print $2}'`
  if [ -n "$var" ];then
    kill -9 $var
  fi
}

killproc WebsocketServer
