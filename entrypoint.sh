#!/bin/sh
while [ -z "$(wget -O - ${API_HOST}/v1/health 2> /dev/null)" ]; do
  echo "API not ready... Trying again in 5s"
  sleep 5
done
main
