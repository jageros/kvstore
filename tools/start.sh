#! /bin/sh

mkdir -p logs
mkdir -p pids

nohup ./builder/kvs 2 >> logs/kvs.log & echo $! > pids/kvs.pid

sleep 1s
ps aux|grep ./builder/kvs | grep -v grep
