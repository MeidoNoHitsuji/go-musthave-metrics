#!/bin/bash

# Provided by https://superuser.com/a/1293762/622572

comm -23 <(seq 10000 65535) <(ss -tan | awk '{print $4}' | cut -d':' -f2 | grep "[0-9]\{1,5\}" | sort | uniq) | shuf | head -n 1