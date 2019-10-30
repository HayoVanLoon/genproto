#!/bin/bash

OLD_DIR=$(pwd)
cd ~/go/src/github.com/HayoVanLoon/genproto

if ! git add .; then
  echo "fail"
  exit 1
fi

if ! git commit -m updates; then
  echo "fail"
  exit 1
fi

if ! git push origin master; then
  echo "fail"
  exit 1
fi

cd "${OLD_DIR}"

echo "done"
