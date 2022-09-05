#!/bin/bash

if [ X$1 = X ]; then
        read -p "Enter the image version number(default version:latest)：" version
else
        version=$1
fi

if [ X$version = X ]; then
version=latest
fi

echo -e "\n"
echo "------------------------"
echo "image version is：$version"
echo "------------------------"


docker build . -t patient-go

docker tag winterchen/patient-go:latest winterchen/patient-go:$version &&

docker push winterchen/patient-go:$version

echo "[upload to dockerhub success!!]"