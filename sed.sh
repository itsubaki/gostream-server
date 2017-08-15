#!/bin/bash

if [ $# -ne 1 ]; then
  echo "ImageID required."
  exit 1
fi

ImageID=$1

sed -e s@"IMAGES"@${ImageID}@g gostream.yml > gostream.yml.tmp
