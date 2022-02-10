#!/bin/bash

major=$1
minor=$2

echo "Building Redweb go rpm - v$major-$minor"
podman build --build-arg MAJOR=$major --build-arg MINOR=$minor -t greeninja/redweb:build . -f Dockerfile.build

echo "Running build container"
podman container create --name extract greeninja/redweb:build
echo "Copying Redweb rpm"
podman container cp extract:/build/redweb-$major-$minor.x86_64.rpm ./rpms/
echo "Removing build container"
podman container rm -f extract
