#!/bin/bash

if ! git add .; then
  echo "fail"
  return 1
fi

if ! git commit -m updates; then
  echo "fail"
  return 1
fi

if ! git push origin master; then
  echo "fail"
  return 1
fi

echo "done"
