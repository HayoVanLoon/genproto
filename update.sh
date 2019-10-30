#!/bin/bash

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

echo "done"
