#! /bin/sh

mkdir -p logs
touch logs/app.log
mkdir -p pids

nohup ./builder/app --id 1 2 >> logs/app.log & echo $! > pids/app.pid

sleep 1s
ps aux|grep ./builder | grep -v grep
