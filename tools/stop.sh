#! /bin/sh

if test $# -ge 1
then
    single=$1
else
    single=-2
fi

echo "kill $single all"
for pid in `cat pids/app.pid`; do kill ${single} ${pid}; done

sleep 1s
ps aux|grep ./builder | grep -v grep