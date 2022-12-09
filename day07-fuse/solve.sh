#!/bin/sh

mkdir -p mount
go run fuse.go day07.txt mount &
sleep 1

echo "Part 1"
find mount/ -type d -print0 | du --files0-from - -b --apparent-size | awk '$1 <= 100000 {s+=$1} END {print s}'
echo

echo "Part 2"
DISKSIZE=70000000
NEEDED=30000000
find mount/ -type d -print0 | du --files0-from - -b --apparent-size | sort -n | awk '$1 >= '$(echo "$NEEDED-($DISKSIZE-$(du -b --apparent-size -s mount | cut -f1)) " | bc)'{print $1; exit}'

umount mount
