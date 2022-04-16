#!/usr/bin/env bash

set -e
DIR="$( cd "$(dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "current dir: $DIR"
PROJECT_ROOT="$DIR/.."

cd "$PROJECT_ROOT"

ls -al

version=$(node -p "require('$PROJECT_ROOT/package.json').version")
if [ "$1" == "test" ]; then
  version="test"
fi
docker_repo="ghcr.io/117s/mdm"
docker_image="$docker_repo:$version"

echo "--> Build Docker Image $docker_image"
docker build --platform=linux/amd64 -t "$docker_image" .

if [ "$1" != "test" ]; then
  docker push "$docker_image"
fi

