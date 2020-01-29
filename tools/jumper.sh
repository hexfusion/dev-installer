#/usr/bin/env bash

if [ "$1" == "" ] || [ "$2" == "" ]; then
  echo "script requires bootstrap and master ip/name"
fi

BOOTSTRAP=$1
MASTER=$2

echo "ssh -vvv -i ~/.ssh/libra.pem -o ProxyCommand="ssh -i ~/.ssh/libra.pem -W %h:%p core@${BOOTSTRAP}" core@${MASTER}"

ssh -vvv -i ~/.ssh/libra.pem -o ProxyCommand="ssh -i ~/.ssh/libra.pem -W %h:%p core@${BOOTSTRAP}" core@${MASTER}
