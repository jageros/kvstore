#! /bin/sh

mkdir -p logs
touch logs/strsrv.log
mkdir -p pids

nohup ./builder/strsrv --id 1 2 >> logs/strsrv.log & echo $! > pids/strsrv.pid

sleep 1s
ps aux|grep ./builder | grep -v grep
