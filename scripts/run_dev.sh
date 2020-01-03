#!/bin/bash

count=0
for  _  in `seq 10`; do
  if `nc -z -w1 127.0.0.1 3306`; then
    count=$(($count+1))

    if [ $count == 3 ]; then
      break
    fi
  fi
  sleep 1
done

sleep 3
fresh -c fresh.conf
