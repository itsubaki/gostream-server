#!/bin/bash

if [ $# -ne 1 ]; then
  echo "ImageID required."
  exit 1
fi

ImageID=$1

sed -e s@"IMAGES"@${ImageID}@g kube/gostream.yml > kube/gostream-deployment.yml
