#!/bin/sh

/root/zebra_linux -f /root/global.yml >> /root/zebra.log 2>&1 &

/usr/local/bin/filebeat -e -c /root/filebeat.yml >> /root/filebeat.log 2>&1 &

while true; do
    echo 'sleep'
    sleep 5;
done